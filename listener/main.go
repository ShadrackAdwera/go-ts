package main

import (
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	conn, err := connectToRabbit()

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	defer conn.Close()

	log.Println("Listening to Messages . . .")

}

func connectToRabbit() (*amqp.Connection, error) {
	//back off routine
	var backOff = time.Second * 1
	var connection *amqp.Connection
	var count int64

	for {
		c, err := amqp.Dial("amqp://guest:guest@rabbitmq")

		if err != nil {
			log.Println("RabbitMQ is not ready . . . ")
			count++
		} else {
			log.Println("Connected to RabbitMQ . . .")
			connection = c
			break
		}

		if count > 5 {
			log.Println("Error connecting to RabbitMQ . . .")
			return nil, err
		}
		backOff = time.Duration(math.Pow(float64(count), 2)) * time.Second
		log.Println("Backing off for ", backOff, " seconds")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
