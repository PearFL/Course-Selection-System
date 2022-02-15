package rabbitmq

import (
	"course_select/src/config"
	global "course_select/src/global"
	"course_select/src/model"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection
var RabbitMQChannel *amqp.Channel

func init() {

	mqConfig := config.GetRabbitMQConfig()

	// RabbitMQ分配的用户名称
	var user string = mqConfig.USER
	// RabbitMQ用户的密码
	var pwd string = mqConfig.PASSWORD
	// RabbitMQ Broker 的ip地址
	var host string = mqConfig.HOST
	// RabbitMQ Broker 监听的端口
	var port string = mqConfig.PORT
	url := "amqp://" + user + ":" + pwd + "@" + host + ":" + port + "/"
	// 新建一个连接
	var err error
	RabbitMQConn, err = amqp.Dial(url)
	if err != nil {
		panic("rabbitmq connection error! " + err.Error())
	}

	RabbitMQChannel, err = RabbitMQConn.Channel()
	if err != nil {
		panic("rabbitmq connection error! " + err.Error())
	}

}

func InitConsumer() {
	q, _ := RabbitMQChannel.QueueDeclare(
		"simple:queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)

	msgs, _ := RabbitMQChannel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
			HandleMessage(d)
		}
	}()

	select {}
}

func HandleMessage(d amqp.Delivery) {
	request := global.BookCourseRequest{}
	if err := json.Unmarshal(d.Body, &request); err != nil {

	}
	err := model.SaveChoice(request.StudentID, request.CourseID)
	if err != nil {

	}
}
