package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	amqp "github.com/rabbitmq/amqp091-go"
)

const queueName = "Service1Queue"

func main() {
	// starting a new connection
	conn, err := amqp.Dial("amqp://localhost:5672")
	if err != nil {
		panic(err)
	}
	// we make sure the connection ends properly
	defer conn.Close()

	// We open a channel
	channel, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// we also make sure the chanel closes properly
	defer channel.Close()

	// In the consumer side, we consume messages from the queue, using the producer's name
	messages, err := channel.Consume(queueName, "", false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// Besides of processing the messages, we need to handle System interrupts from the user
	// if the user wants to stop the process, we need to handle it
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	// Now with a loop we wait for incoming messages or interrupt signals
	for {
		select {
		case message := <-messages:
			log.Printf("The message is: %s\n", message.Body)
		case <-sigchan:
			log.Println("Process interrupted!")
			os.Exit(0)
		}
	}
}
