package producer

import (
	"log"

	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func Publish(message string) {
	// Here we connect to RabbitMQ or send a message if there are any errors connecting.
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// We create a Queue to send the message to.
	err = ch.ExchangeDeclare(
		"br.payment.topic", // name
		"fanout",           // durable
		true,               // delete when unused
		false,              // exclusive
		false,              // no-wait
		false,              // no-wait
		nil,                // arguments
	)
	failOnError(err, "Failed to declare exchange")

	// We set the payload for the message.
	body := "Golang is awesome - Keep Moving Forward!"
	// body := bodyFrom(message)
	err = ch.Publish(
		"br.payment.topic", // exchange
		"",                 // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	// If there is an error publishing the message, a log will be displayed in the terminal.
	failOnError(err, "Failed to publish a message")
	log.Printf(" [x] Congrats, sending message: %s", body)
}
