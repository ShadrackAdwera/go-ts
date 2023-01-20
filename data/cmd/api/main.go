package main

import (
	"context"
	"data/repo"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	webPort = "5001"
)

type Config struct {
	Models repo.Models
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
		Models: repo.New(client),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.Routes(),
	}

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
