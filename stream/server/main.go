package main

import (
	"fmt"
	"net"

	pb "github.com/SeTriones/grpc-demo/stream/proto"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(stream pb.Greeter_SayHelloServer) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			fmt.Printf("recv err=%v\n", err)
			return err
		}
		stream.Send(&pb.HelloReply{Message: fmt.Sprintf("hello %s", req.Name)})
	}
}

func main() {
	lis, err := net.Listen("tcp", "127.0.0.1:18800")
	if err != nil {
		panic(err)
	}
	serv := &server{}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, serv)
	s.Serve(lis)
}
