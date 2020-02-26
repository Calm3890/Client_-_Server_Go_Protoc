package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/Calm3890/hellogrpc/hello"
	"google.golang.org/grpc"
)

type helloServer struct {}

func (s *helloServer) SayHello(ctx context.Context, req *hello.HelloRequest)(*hello.HelloResponse, error) {
	n := rand.Int63n(6)
	time.Sleep(time.Duration(n)*time.Millisecond)
	return &hello.HelloResponse{Message: "Hello "+ req.Name} , nil
}

func main() {
sock, err := net.Listen("tcp",":1234")
if err != nil {
	log.Fatalf("error listening to network: %v", err)
}

	server := grpc.NewServer()
	hello.RegisterHelloServer(server, &helloServer{})
	rand.Seed(time.Now().Unix())
	err = server.Serve(sock)
	if err != nil{
		log.Fatalf("error starting grpc server: %v", err)
	}
	fmt.Println("hello server in listening at :1234")
}