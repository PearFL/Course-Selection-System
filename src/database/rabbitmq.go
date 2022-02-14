package database

import (
	"course_select/src/config"

	"github.com/streadway/amqp"
)

var RabbitMQConn *amqp.Connection

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
}
