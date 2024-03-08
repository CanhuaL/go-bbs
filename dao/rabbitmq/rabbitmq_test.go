package rabbitmq

import (
	"go_bbs/setting"
	"testing"
)

func TestRabbitMQInit(t *testing.T) {
	dbCfg := setting.RabbitConfig{
		Host:     "127.0.0.1",
		User:     "guest",
		Password: "guest",
		Port:     5672,
	}

	err := Init(&dbCfg)
	if err != nil {
		t.Fatalf("Failed to initialize RabbitMQ: %v", err)
	}

	defer Close()

	p, err := Conn.Channel()
	if err != nil {
		t.Fatalf("Failed to initialize Conn: %v", err)
	}

	// 声明交换机
	err = p.ExchangeDeclare(
		"private_chat", // 交换机名称
		"topic",        // 交换机类型
		true,           // 是否持久化
		false,          // 是否自动删除
		false,          // 是否内部
		false,          // 是否等待
		nil,            // 附加参数
	)

	if err != nil {
		t.Fatalf("Failed to initialize Exchange: %v", err)
	}
}
