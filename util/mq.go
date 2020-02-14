package util

import (
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"log"
)


func MqSend(ch *amqp.Channel,qName string,body []byte) error {
	q, err := ch.QueueDeclare(
		qName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		Logzap.Error("Failed to declare a queue",zap.Error(err))
		return err
	}


	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        body,
		})
	if err != nil {
		Logzap.Error("Failed to publish a message",zap.Error(err))
		return err
	}
	return nil
}


func MqReceive(ch *amqp.Channel, qName string)  {
	q, err := ch.QueueDeclare(
		qName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if err != nil {
		Logzap.Error("Failed to declare a queue",zap.Error(err))
		panic(err)
	}
	fmt.Println("q.name:",q.Name)
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		Logzap.Error("Failed to register a consumer",zap.Error(err))
		panic(err)
	}

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			err =XXX(d.Body)
			if err != nil {
				continue
			}
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}

func XXX(b []byte) error{
	fmt.Println(string(b))
	return nil
}