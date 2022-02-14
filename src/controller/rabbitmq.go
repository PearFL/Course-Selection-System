package controller

import (
	global "course_select/src/global"
	"course_select/src/rabbitmq"
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

func InitProducer(request global.BookCourseRequest) error {

	q, err := rabbitmq.RabbitMQChannel.QueueDeclare(
		// 队列名称
		"simple:queue", // name
		false,          // durable
		false,          // delete when unused
		false,          // exclusive
		false,          // no-wait
		nil,            // arguments
	)
	if err != nil {
		log.Println(err)
		return err
	}

	dataBytes, err := json.Marshal(request)
	if err != nil {
		log.Println(err)
		return err
	}
	err = rabbitmq.RabbitMQChannel.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        dataBytes,
		})
	// log.Printf(" [x] Sent %s", dataBytes)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
