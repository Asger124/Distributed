package main

import (
	"context"
	"log"
	"time"

	pb "example.com/session_04_grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCourseServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.GetCourse(ctx, &pb.CourseRequest{Id: 1})
	if err != nil {
		log.Fatalf("Could not get course: %v", err)
	}

	log.Printf("Course: %v", res)

}
