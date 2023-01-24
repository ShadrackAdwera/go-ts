package main

import (
	"context"
	"data/repo"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-redis/redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort   = "5001"
	gRPCPort  = "5005"
	redisAddr = "auth-redis-service:6379"
)

type Config struct {
	Models   repo.Models
	AuthData AuthData
	Redis    *redis.Client
}

func main() {

	client, err := connectToMongo()

	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)

	defer cancel()

	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	app := Config{
		Models:   repo.New(client),
		AuthData: AuthData{},
		Redis:    InitRedis(),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

	go app.gRPCListen()

	log.Printf("Starting DATA service on port %s\n", webPort)

	err = srv.ListenAndServe()

	if err != nil {
		log.Panicln(err)
		return
	}
}

func connectToMongo() (*mongo.Client, error) {
	// create connection options
	mongoURL := os.Getenv("DATA_MONGO_URI")
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "password",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Println("Error connecting:", err)
		return nil, err
	}

	log.Println("Connected to mongo!")

	return c, nil
}

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "",
		DB:       0,
	})
	return rdb
}
