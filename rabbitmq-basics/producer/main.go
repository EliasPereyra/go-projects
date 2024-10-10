package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	// we declare a queue
	channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// we begin a tcp server using gin
	router := gin.Default()

	// declare a route for receiving messages
	router.GET("/send", func(c *gin.Context) {
		msg := c.Query("msg")
		if msg == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "you need to write a message"})
			return
		}

		// we create a message to publish
		message := amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		}

		// then we publish the message created to the queue
		err = channel.Publish("", queueName, false, false, message)
		if err != nil {
			log.Printf("Failed to publish the message: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to publish the message"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": msg, "status": "success"})
	})

	log.Fatal(router.Run(":8080"))
}
