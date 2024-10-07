package main

import (
	"fmt"
	"github.com/dimastephen/chatServer/grpc/pkg/chatServerV1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const addr = "50052"

type server struct {
	chatServerV1.UnimplementedChatServerServer
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", addr))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	chatServerV1.RegisterChatServerServer(s, &server{})
	log.Println("Started ChatServer")
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to start chatServer", err)
	}

}
