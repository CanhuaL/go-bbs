package rabbitmq

import (
	"fmt"
	"go_bbs/setting"

	"github.com/streadway/amqp"
)

var Conn *amqp.Connection

func Init(cfg *setting.RabbitConfig) (err error) {
	dsn := fmt.Sprintf("amqp://%s:%s@%s:%d/", cfg.User, cfg.Password, cfg.Host, cfg.Port)
	Conn, err = amqp.Dial(dsn)
	if err != nil {
		return
	}
	return
}

func Close() {
	_ = Conn.Close()
}

//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
//failOnError(err, "Failed to connect to RabbitMQ")
//defer conn.Close()
//
//ch, err := conn.Channel()
//failOnError(err, "Failed to open a channel")
//defer ch.Close()
//
//err = ch.ExchangeDeclare(
//"logs_topic", // name
//"topic",      // type
//true,         // durable
//false,        // auto-deleted
//false,        // internal
//false,        // no-wait
//nil,          // arguments
//)
//failOnError(err, "Failed to declare an exchange")
//
//body := bodyFrom(os.Args)
//err = ch.Publish(
//"logs_topic",          // exchange
//severityFrom(os.Args), // routing key
//false,                 // mandatory
//false,                 // immediate
//amqp.Publishing{
//ContentType: "text/plain",
//Body:        []byte(body),
//})
//failOnError(err, "Failed to publish a message")
//
//log.Printf(" [x] Sent %s", body)
//}
//
//func bodyFrom(args []string) string {
//	var s string
//	if (len(args) < 3) || os.Args[2] == "" {
//		s = "hello"
//	} else {
//		s = strings.Join(args[2:], " ")
//	}
//	return s
//}
//
//func severityFrom(args []string) string {
//	var s string
//	if (len(args) < 2) || os.Args[1] == "" {
//		s = "anonymous.info"
//	} else {
//		s = os.Args[1]
//	}
//	return s
//}
