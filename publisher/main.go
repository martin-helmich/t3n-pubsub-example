package main

import (
	"encoding/json"
	"flag"
	"log"
	"time"

	"github.com/icrowley/fake"
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

	if err := channel.ExchangeDeclare("events", "topic", true, false, false, false, nil); err != nil {
		panic(err)
	}

	for {
		msg := UserCreatedMessage{fake.FullName(), fake.EmailAddress()}
		b, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		publishing := amqp.Publishing{
			ContentType: "application/json",
			Body:        b,
		}

		if err := channel.Publish("events", "userCreated", false, false, publishing); err != nil {
			panic(err)
		}

		time.Sleep(5 * time.Second)
	}
}
