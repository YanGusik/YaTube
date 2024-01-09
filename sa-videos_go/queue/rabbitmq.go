package queue

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

var conn *amqp.Connection

func Init(c string) {
	var err error
	conn, err = amqp.Dial(c)
	if err != nil {
		log.Fatalf("could not connect to rabbitmq: %v", err)
		panic(err)
	}
}

func Publish(q string, msg []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()
	
	payload := amqp.Publishing{
		DeliveryMode: amqp.Persistent,
		ContentType:  "application/json",
		Body:         msg,
	}

	if err := ch.Publish("", q, false, false, payload); err != nil {
		return fmt.Errorf("[Publish] failed to publish to queue %v", err)
	}

	return nil
}

func Subscribe(qName string) (<-chan amqp.Delivery, func(), error) {
	ch, err := conn.Channel()
	if err != nil {
		return nil, nil, err
	}
	q, err := ch.QueueDeclare(qName, false, false, false, false, nil)
	c, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	return c, func() { ch.Close() }, err
}
