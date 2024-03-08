package rabbitmq

import (
	"fmt"
	"go.uber.org/zap"
	"go_bbs/setting"

	"github.com/streadway/amqp"
)

var (
	Conn           *amqp.Connection
	PrivateChannel *amqp.Channel
)

func Init(cfg *setting.RabbitConfig) (err error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	Conn, err = amqp.Dial(dsn)
	if err != nil {
		return
	}
	PrivateChannel, err = Conn.Channel()
	if err != nil {
		zap.L().Fatal("Failed to create RabbitMQ channel", zap.Error(err))
	}
	// 声明交换机
	err = PrivateChannel.ExchangeDeclare(
		"private_chat", // 交换机名称
		"topic",        // 交换机类型
		true,           // 是否持久化
		false,          // 是否自动删除
		false,          // 是否内部
		false,          // 是否等待
		nil,            // 附加参数
	)
	if err != nil {
		zap.L().Fatal("Failed to declare RabbitMQ exchange", zap.Error(err))
	}
	return
}

func Close() {
	_ = Conn.Close()
}
