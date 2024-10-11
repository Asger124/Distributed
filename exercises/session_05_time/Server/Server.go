package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "example.com/session_05_time/proto"
	"google.golang.org/grpc"
)

//var addr string = "0.0.0.0:50051"

type server struct {
	pb.UnimplementedTimeServiceServer
}

func (s *server) GetTime(ctx context.Context, timereq *pb.Timerequest) (*pb.Timeresponse, error) {

	log.Printf("Timerequest was invoked %v:\n,", timereq)

	return &pb.Timeresponse{
		Time: time.Now().String(),
	}, nil

}

func main() {

	port := flag.String("port", "50051", "The server port")
	flag.Parse()

	// Build the address using the provided port
	addr := fmt.Sprintf("0.0.0.0:%s", *port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Server failed to listen")
	}

	log.Printf("Listening on %s addr", addr)

	S := grpc.NewServer()

	pb.RegisterTimeServiceServer(S, &server{})

	if err = S.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
