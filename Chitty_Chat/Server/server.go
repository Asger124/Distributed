package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/Chitty_Chat/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedChitty_ChatServer
	Messages      []*pb.Message
	participants  []*Participant
	ServerLamport uint32
}

type Participant struct {
	Participant *pb.Participant
	channel     chan *pb.Message
}

func (s *Server) Broadcast(par *pb.Participant, stream pb.Chitty_Chat_BroadcastServer) error {

	s.ServerLamport++
	log.Printf("-> Participant %v joined Chitty-Chat at Lamport time %v", par.Name, s.ServerLamport)

	//lamportString := strconv.FormatInt(int64(s.ServerLamport), 10)

	NewParticipant := &Participant{
		Participant: par,
		channel:     make(chan *pb.Message, 1),
	}

	s.participants = append(s.participants, NewParticipant)

	Join := &pb.Message{
		Message:     "Joined Chitty-Chat",
		Par:         par,
		LamportTime: s.ServerLamport,
	}

	for _, p := range s.participants {

		p.channel <- Join

	}

	if len(s.Messages) != 0 {
		for _, msg := range s.Messages {
			stream.Send(msg)
		}
	}

	s.Messages = append(s.Messages, Join)

	for {
		select {
		case msg := <-NewParticipant.channel:
			stream.Send(msg)
		}
	}
}

func (s *Server) Publish(ctx context.Context, msg *pb.Message) (*pb.Empty, error) {

	s.ServerLamport++
	s.Messages = append(s.Messages, msg)
	for _, p := range s.participants {

		p.channel <- msg
	}

	return &pb.Empty{}, nil
}

func (s *Server) Leave(ctx context.Context, par *pb.Participant) (*pb.Empty, error) {

	s.ServerLamport++

	log.Printf("-> Participant %v Left Chitty-Chat at Lamport time %v", par.Name, s.ServerLamport)

	//lamportString := strconv.FormatInt(int64(s.ServerLamport), 10)

	Leave := &pb.Message{
		Message:     "Left Chitty-Chat",
		Par:         par,
		LamportTime: s.ServerLamport,
	}

	for _, p := range s.participants {

		p.channel <- Leave

	}

	s.Messages = append(s.Messages, Leave)

	return &pb.Empty{}, nil

}

var ServerLamport = uint(0)

func main() {

	port := flag.String("port", "50051", "The server port")
	flag.Parse()
	log.SetFlags(0)

	// Build the address using the provided port
	addr := fmt.Sprintf("0.0.0.0:%s", *port)

	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Server failed to listen")
	}

	log.Printf("Listening on: %s", addr)

	S := grpc.NewServer()

	pb.RegisterChitty_ChatServer(S, &Server{})

	if err = S.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
