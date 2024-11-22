package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	pb "example.com/Replication/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAuctionServer
	clientID             uint32
	port                 uint32
	lamportstamp         uint32
	ClientWithHighestbid uint32
	auctionisactive      bool
	Highestbid           uint32
	time                 *time.Timer
	serverid             uint32
}

var serverport = flag.Int("id", 0, "The ID of the client (used to calculate the port)")

func main() {

	flag.Parse()

	ownport := uint32(*serverport) + 5000

	server := Server{
		port:                 ownport,
		lamportstamp:         0,
		ClientWithHighestbid: 0,
		auctionisactive:      false,
		Highestbid:           0,
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", ownport))

	if err != nil {
		log.Fatalf("Server failed to listen")
	}

	log.Printf("Listening on: %d\n", ownport)

	S := grpc.NewServer()

	pb.RegisterAuctionServer(S, &server)

	if err = S.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
