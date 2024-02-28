package redis

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"time"
)

func SaveCode(code, phone string) error {
	pipe := client.TxPipeline()
	// 记录一天内手机号码发送的次数和检查发送次数是否超过限制
	incr := pipe.Incr(getRedisKey(KeyAttempts + phone))
	expire := pipe.Expire(getRedisKey(KeyAttempts+phone), Expiration)
	fmt.Println(expire)
	// 执行 Pipeline
	_, err := pipe.Exec()
	if err != nil {
		zap.L().Error("Failed to increment counter", zap.Error(err))
		return err
	}
	// 如果发送次数超过限制，则返回错误信息
	count := incr.Val()
	if count > MaxAttempts {
		return errors.New("一天内验证码发送次数超过限制了！")
	}
	// 将验证码存储到 Redis 中，并设置过期时间为 30 min

	// TODO: 调用短信服务发送验证码，这里假设直接打印验证码
	fmt.Println("Send code to", phone, ":", code)
	err = client.Set(phone, code, 30*60*time.Second).Err()
	if err != nil {
		return err
	}
	return err
}

func GetCode(code, phone string) error {
	// 从 Redis 中获取验证码
	storedCode, err := client.Get(phone).Result()
	if err != nil {
		return err
	}

	// 验证验证码
	if code != storedCode {
		zap.L().Error("Invalid verification code", zap.Error(err))
		return err
	}

	// 删除 Redis 中的验证码
	err = client.Del(phone).Err()
	if err != nil {
		zap.L().Error("Failed to delete code from Redis", zap.Error(err))
		return err
	}
	return err
}

func GetPhone(phone string) bool {
	val, err := client.Get(phone).Result()
	if err != nil && err != redis.Nil {
		fmt.Println("Error:", err)
		return false
	}
	if val != "" {
		return true
	}
	return false
}
