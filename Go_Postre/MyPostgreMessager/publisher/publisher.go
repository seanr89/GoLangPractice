package publisher

import (
	"flag"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@localhost:5672/", "AMQP URI")
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func Init() {
	fmt.Println("Publisher Init!")
	flag.Parse()
	initAmqp()
}

var conn *amqp.Connection
var ch *amqp.Channel
var q *amqp.Queue

func initAmqp() {
	// var err error
	// conn, err = amqp.Dial(*amqpURI)
	// //conn, err = amqp.Dial("amqp://guest:guest@localhost:5672/")
	// failOnError(err, "Failed to connect to RabbitMQ")
	// ch, err = conn.Channel()
	// failOnError(err, "Failed to open a channel")
	// err = ch.ExchangeDeclare(
	// 	"test-exchange", // name
	// 	"direct",        // type
	// 	true,            // durable
	// 	false,           // auto-deleted
	// 	false,           // internal
	// 	false,           // noWait
	// 	nil,             // arguments
	// )
	// failOnError(err, "Failed to declare the Exchange")

	conn, err := amqp.Dial(*amqpURI)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError(err, "Failed to publish a message")
}

func PublishAMessage(message string) {
	log.Printf("Publish a Message")

	err := ch.Publish(
		"",      // exchange
		"hello", // routing key
		false,   // mandatory
		false,   // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	log.Printf(" [x] Sent %s", message)

	failOnError(err, "Failed to Publish on RabbitMQ")
}
