package logic

import (
	"github.com/gorilla/websocket"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"go_bbs/dao/mysql"
	"go_bbs/dao/rabbitmq"
	"go_bbs/models"
	"net/http"
	"sync"
)

// Node 本核心在于形成userid和Node的映射关系
type Node struct {
	Conn *websocket.Conn //  保存websocket连接，conn是io型的资源
}

var (
	// 映射关系表(map的键userid值是Node，全局的map，所有协程共享)
	clientMap map[int64]*Node = make(map[int64]*Node, 0)
	// 读写锁
	rwlocker sync.RWMutex
)

func PrivateChat(p *models.ParamPrivateChat, w http.ResponseWriter, r *http.Request) (err error) {
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(w, r, nil)
	if err != nil {
		zap.L().Error("websocket.Upgrader has err")
	}

	node := &Node{
		Conn: conn,
	}

	rwlocker.Lock()
	clientMap[p.ReceiverId] = node
	rwlocker.Unlock()
	//  发送conn
	go SendProc(p, node)
	//  接收conn
	go ReceiveProc(node)
	return err
}

// SendProc ws发送协程
func SendProc(p *models.ParamPrivateChat, node *Node) {
	ch, err := rabbitmq.Conn.Channel()
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	defer ch.Close()

	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			zap.L().Error(err.Error())
			return
		}
		//err = node.Conn.WriteMessage(websocket.TextMessage, data)
		err = ch.ExchangeDeclare(
			"private_chat", // name
			"topic",        // type
			true,           // durable
			false,          // auto-deletedZ
			false,          // internal
			false,          // no-wait
			nil,            // arguments
		)
		// 发送消息到 RabbitMQ
		err = ch.Publish(
			"private_chat", // 交换机名称
			"private.P",    // 路由键，使用默认交换机，直接发送到队列中
			false,          // 是否等待服务器确认
			true,           // 是否持久化消息
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        data,
			},
		)
		if err != nil {
			zap.L().Error(err.Error())
			break
		}
		// 将data持久化到mysql中，data再转成sting
		s := string(data)
		m := &models.PrivateChat{
			SenderId:   p.SenderId,
			ReceiverId: p.SenderId,
			Content:    s,
		}
		if err = mysql.SavePrivateChat(m); err != nil {
			zap.L().Error(err.Error())
			return
		}
	}
}

// ReceiveProc ws接收协程
func ReceiveProc(node *Node) {
	ch, err := rabbitmq.Conn.Channel()
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	defer ch.Close()
	err = ch.ExchangeDeclare(
		"private_chat", // name
		"topic",        // type
		true,           // durable
		false,          // auto-deletedZ
		false,          // internal
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	// 声明队列
	q, err := ch.QueueDeclare(
		"private_queue", // 队列名称
		false,           // 是否持久化
		false,           // 是否自动删除
		true,            // 是否排他
		false,           // 是否阻塞
		nil,             // 额外参数
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	err = ch.QueueBind(
		q.Name,         // queue name
		"private.P",    // routing key
		"private_chat", // exchange
		false,
		nil)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	// 将接收到的数据直接发送给接收方的客户端
	for msg := range msgs {
		err = node.Conn.WriteMessage(websocket.TextMessage, msg.Body)
		if err != nil {
			zap.L().Error(err.Error())
			break
		}
	}
}
