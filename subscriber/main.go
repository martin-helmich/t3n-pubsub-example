package main

import (
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/streadway/amqp"
)

type UserCreatedMessage struct {
	Username     string `json:"username"`
	EmailAddress string `json:"emailAddress"`
}

func main() {
	var conn *amqp.Connection
	var err error

	broker := flag.String("broker", "amqp://localhost:5672", "address of AMQP broker")
	flag.Parse()

	for {
		conn, err = amqp.Dial(*broker)
		if err == nil {
			break
		}

		log.Printf("connection to AMQP failed: %s", err.Error())
		time.Sleep(1 * time.Second)
	}

	defer conn.Close()

	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	queue, err := channel.QueueDeclare("", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	log.Printf("created queue: %s", queue.Name)

	if err := channel.QueueBind(queue.Name, "userCreated", "events", false, nil); err != nil {
		panic(err)
	}

	msgs, err := channel.Consume(queue.Name, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		created := UserCreatedMessage{}
		if err := json.Unmarshal(msg.Body, &created); err != nil {
			panic(err)
		}

		log.Printf("user created: %s <%s>", created.Username, created.EmailAddress)
	}
}
