package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "example.com/session_05_time/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	port1 := flag.String("port1", "5001", "The server port")
	port2 := flag.String("port2", "5002", "The server port")
	flag.Parse()

	// Build the address using the provided port
	addr1 := fmt.Sprintf("localhost:%s", *port1)

	conn1, err := grpc.Dial(addr1, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn1.Close()

	client1 := pb.NewTimeServiceClient(conn1)

	addr2 := fmt.Sprintf("localhost:%s", *port2)

	conn2, err := grpc.Dial(addr2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn2.Close()

	client2 := pb.NewTimeServiceClient(conn2)

	log.Println("Calling the first service...")
	res1, err := client1.GetTime(context.Background(), &pb.Timerequest{})
	if err != nil {
		log.Fatalf("Error calling first service: %v", err)
	}
	log.Printf("First service responded with time: %v", res1.Time)

	log.Println("Calling the second service...")
	res2, err := client2.GetTime(context.Background(), &pb.Timerequest{})
	if err != nil {
		log.Fatalf("Error calling second service: %v", err)
	}
	log.Printf("Second service responded with time: %v", res2.Time)
}
