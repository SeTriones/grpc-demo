package main

import (
	"context"
	"fmt"
	"strconv"
	"time"

	pb "github.com/SeTriones/grpc-demo/stream/proto"
	"google.golang.org/grpc"
)

var (
	conn   *grpc.ClientConn
	stream pb.Greeter_SayHelloClient
	err    error
)

func connect(addr string) {
	conn, err = grpc.DialContext(context.Background(), addr, grpc.WithInsecure())
	if err != nil {
		conn.Close()
		fmt.Printf("dial fail, err=%v\n", err)
		return
	}
	client := pb.NewGreeterClient(conn)
	stream, err = client.SayHello(context.Background())
	if err != nil {
		conn.Close()
		stream = nil
		return
	}
	return
}

func main() {
	connect("127.0.0.1:18800")
	if err != nil {
		panic(err)
	}

	idx := 0
	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		if stream == nil {
			connect("127.0.0.1:18800")
			continue
		}
		stream.Send(&pb.HelloRequest{Name: strconv.Itoa(idx)})
		resp, err := stream.Recv()
		if err != nil {
			conn.Close()
			connect("127.0.0.1:18800")
			continue
		}
		fmt.Printf("%s, resp=%s\n", t.Format(time.RFC1123), resp.Message)
		idx = idx + 1
	}
}
