package main

import (
	"context"
	"data/protobufs"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

type AuthData struct {
	UserId string `json:"userId"`
	Expiry string `json:"expiry"`
}

type AuthServer struct {
	protobufs.UnimplementedAuthServiceServer
	AuthData AuthData
}

func (c *AuthServer) WriteCategory(ctx context.Context, req *protobufs.AuthRequest) (*protobufs.AuthResponse, error) {
	input := req.GetAuthEntry()

	auth := AuthData{
		UserId: input.UserId,
		Expiry: input.Expiry,
	}
	sec, err := strconv.Atoi(auth.Expiry)
	if err != nil {
		res := &protobufs.AuthResponse{Result: "Failed to convert expiry date to seconds"}
		return res, err
	}
	log.Println(auth)
	redisClient := InitRedis()
	err = redisClient.Set(context.Background(), auth.UserId, auth.UserId, time.Second*time.Duration(sec)).Err()

	if err != nil {
		res := &protobufs.AuthResponse{Result: "Failed to insert auth data into the REDIS cache"}
		return res, err
	}
	res := &protobufs.AuthResponse{Result: "auth data successfully added into REDIS"}
	return res, nil
}

func (app *Config) gRPCListen() {
	l, err := net.Listen("tcp", fmt.Sprintf(":%s", gRPCPort))

	if err != nil {
		log.Fatalf("Fail to listen to gRPC connection: %v", err)
	}

	s := grpc.NewServer()

	protobufs.RegisterAuthServiceServer(s, AuthServer{
		AuthData: app.AuthData,
	})
	log.Printf("gRPC server started on PORT: %s", gRPCPort)

	if err := s.Serve(l); err != nil {
		log.Fatalf("Fail to listen to gRPC connection: %v", err)
	}
}
