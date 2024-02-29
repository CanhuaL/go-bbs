package logic

import (
	"fmt"
	"go.uber.org/zap"
	"go_bbs/dao/mysql"
	"go_bbs/dao/redis"
	"go_bbs/models"
	"go_bbs/pkg/jwt"
	"go_bbs/pkg/snowflake"
	"math/rand"
	"strconv"
	"time"
)

// 存放业务逻辑的代码

func SignUp(p *models.ParamSignUp) (err error) {
	// 1.判断用户存不存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		return err
	}
	// 2.生成UID
	userID := snowflake.GenID()
	// 构造一个User实例
	user := &models.User{
		UserID:   userID,
		Username: p.Username,
		Password: p.Password,
		Phone:    p.Phone,
		Email:    p.Email,
	}
	// 3.保存进数据库
	return mysql.InsertUser(user)
}

func Login(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Username: p.Username,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.Login(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

func PhoneLogin(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Phone:    p.Phone,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.PhoneLogin(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

func GetUserFromPhone(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Phone: p.Phone,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.GetUserFromPhone(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

func EmailLogin(p *models.ParamLogin) (user *models.User, err error) {
	user = &models.User{
		Email:    p.Email,
		Password: p.Password,
	}
	// 传递的是指针，就能拿到user.UserID
	if err := mysql.EmailLogin(user); err != nil {
		return nil, err
	}
	// 生成JWT
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return
	}
	user.Token = token
	return
}

func SaveCode(phone string) (err error) {
	//  如果发过短信
	if redis.GetPhone(phone) {
		s, err := mysql.GetSmsMessages(phone)
		if err != nil {
			zap.L().Error("mysql.GetSmsMessages failed", zap.Error(err))
			return err
		}
		now := time.Now().Unix()
		t := s.SendTime.Unix()
		//  如果验证码间隔小于60s则提示报错
		if now-t < 60 {
			fmt.Println("请在：", 60-(now-t), "s后再试")
			zap.L().Error("验证码间隔小于60s，请稍后再试！", zap.Error(err))
			return err
		}
	}

	// 生成随机验证码
	code := generateCode()
	sms := &models.SMS{
		SMSType:    "验证码登录",
		SMSContent: code,
		Phone:      phone,
		SendTime:   time.Now(),
	}
	if err = mysql.SaveSmsMessages(sms); err != nil {
		return err
	}
	return redis.SaveCode(code, phone)
}

// generateCode 生成随机验证码
func generateCode() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(1000000)) // 生成一个六位数的随机数
}

func GetCode(code, phone string) (err error) {
	return redis.GetCode(code, phone)
}
