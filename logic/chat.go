package logic

import (
	"encoding/json"
	"fmt"
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
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	// 映射关系表(map的键userid值是Node，全局的map，所有协程共享)
	clientMap map[int64]*Node = make(map[int64]*Node, 0)
	// 读写锁
	rwLocker sync.RWMutex
	// RabbitMQ Channel
	// 已声明队列记录
	declaredQueues sync.Map
)

func PrivateChat(senderId, receiverId int64, w http.ResponseWriter, r *http.Request) (err error) {
	// 升级 HTTP 连接为 WebSocket 连接
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		zap.L().Error("websocket.Upgrader has err")
		return
	}
	defer conn.Close()

	node := &Node{
		Conn: conn,
	}

	rwLocker.Lock()
	clientMap[senderId] = node
	rwLocker.Unlock()
	//  发送conn
	go SendProc(node)
	//  接收conn
	go ReceiveProc(receiverId)
	return err
}

// SendProc ws发送协程
func SendProc(node *Node) {
	for {
		var msg models.ParamPrivateChatMsg
		err := node.Conn.ReadJSON(&msg)
		if err != nil {
			zap.L().Error(err.Error())
			return
		}
		routingKey := fmt.Sprintf("private.%d.%d", msg.SenderId, msg.ReceiverId)

		jsonBytes, err := json.Marshal(msg)
		fmt.Println(jsonBytes)
		if err != nil {
			return
		}
		err = rabbitmq.PrivateChannel.Publish(
			"private_chat", // 交换机名称
			routingKey,     // 路由键，使用默认交换机，直接发送到队列中
			false,          // 是否等待服务器确认
			true,           // 是否持久化消息
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        jsonBytes,
			},
		)
		if err != nil {
			zap.L().Error(err.Error())
			break
		}

		fmt.Printf(msg.Content, msg.SenderId, msg.ReceiverId)
		m := &models.PrivateChat{
			SenderId:   msg.SenderId,
			ReceiverId: msg.SenderId,
			Content:    msg.Content,
		}
		if err = mysql.SavePrivateChat(m); err != nil {
			zap.L().Error(err.Error())
			return
		}
	}
}

// ReceiveProc ws接收协程
func ReceiveProc(receiverId int64) {
	// 检查是否已经声明过该用户的队列
	if _, loaded := declaredQueues.LoadOrStore(receiverId, true); loaded {
		return // 如果已经声明过，直接返回
	}
	queueName := fmt.Sprintf("privateMessages_%d", receiverId)
	// 声明队列
	_, err := rabbitmq.PrivateChannel.QueueDeclare(
		queueName, // 队列名称
		true,      // 是否持久化
		false,     // 是否自动删除
		false,     // 是否排他
		false,     // 是否阻塞
		nil,       // 额外参数
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	bindingKey := fmt.Sprintf("private.*.%d", receiverId)
	err = rabbitmq.PrivateChannel.QueueBind(
		queueName,      // queue name
		bindingKey,     // routing key
		"private_chat", // exchange
		false,
		nil)
	if err != nil {
		zap.L().Error(err.Error())
		// 如果绑定队列失败，从已声明队列记录中移除
		declaredQueues.Delete(receiverId)
		return
	}
	msgs, err := rabbitmq.PrivateChannel.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto ack
		false,     // exclusive
		false,     // no local
		false,     // no wait
		nil,       // args
	)
	if err != nil {
		zap.L().Error(err.Error())
		return
	}
	receiverCon, ok := clientMap[receiverId]
	if !ok {
		// 如果找不到接收方的连接，可能需要采取措施，例如存储消息以便稍后传送
		// todo 暂时不做处理
	}
	// 将接收到的数据直接发送给接收方的客户端
	for msg := range msgs {
		err = receiverCon.Conn.WriteMessage(websocket.TextMessage, msg.Body)
		fmt.Printf(string(msg.Body))
		if err != nil {
			zap.L().Error(err.Error())
			break
		}
	}
}
