package main

import (
	"context"
	"log"
	"net"

	pb "example.com/session_04_grpc/proto"

	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type server struct {
	pb.UnimplementedCourseServiceServer
}

func (s *server) GetCourse(ctx context.Context, req *pb.CourseRequest) (*pb.CourseResponse, error) {

	return &pb.CourseResponse{
		Id:          req.Id,
		Name:        "Beginner in grpc",
		Description: "Trying to understand grpc",
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Server failed to listen")
	}

	log.Printf("Listening on %s addr", addr)

	S := grpc.NewServer()

	pb.RegisterCourseServiceServer(S, &server{})

	if err = S.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
