package events

import (
	"context"
	"encoding/json"
	"fmt"
	"listener/protobufs"
	"log"
	"strconv"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Consumer struct {
	conn      *amqp.Connection
	queueName string
}

type Payload struct {
	UserId string `json:"userId"`
	Expiry int    `json:"expiry"`
}

// receiving events from RabbitMQ
func NewConnection(conn *amqp.Connection) (Consumer, error) {
	consumer := Consumer{
		conn: conn,
	}

	err := consumer.setUp()

	if err != nil {
		return Consumer{}, err
	}
	return consumer, nil
}

func (consumer *Consumer) setUp() error {
	channel, err := consumer.conn.Channel()

	if err != nil {
		return err
	}

	return declareExchange(channel)
}

func (consumer *Consumer) Listen(topics []string) error {
	ch, err := consumer.conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	q, err := declareRandomQueue(ch)
	if err != nil {
		return err
	}

	for _, s := range topics {
		ch.QueueBind(
			q.Name,
			s,
			"auth_topic",
			false,
			nil,
		)

		if err != nil {
			return err
		}
	}

	messages, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		return err
	}
	// consume messages forever
	forever := make(chan bool)
	go func() {
		for d := range messages {
			var payload Payload
			_ = json.Unmarshal(d.Body, &payload)

			//go handlePayload(payload)
		}
	}()

	fmt.Printf("Waiting for message [Exchange, Queue] [auth_topic, %s]\n", q.Name)
	<-forever

	return nil
}

func handlePayload(payload Payload) {
	// send auth data to data service
	cc, err := grpc.Dial("data-service:5001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())

	defer cc.Close()

	if err != nil {
		return
	}

	c := protobufs.NewAuthServiceClient(cc)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)

	defer cancel()

	_, err = c.WriteAuth(ctx, &protobufs.AuthRequest{
		AuthEntry: &protobufs.Auth{
			UserId: payload.UserId,
			Expiry: strconv.Itoa(payload.Expiry),
		},
	})

	if err != nil {
		log.Println("Error sending Auth Data via gRPC. . . ", err)
		return
	}

	log.Println("Auth Data sent to Data Service via gRPC. . . ")
}
