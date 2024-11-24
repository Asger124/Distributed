package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"

	pb "example.com/Replication/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedAuctionServer
	clientID             uint32
	port                 int32
	lamportstamp         uint32
	ClientWithHighestbid uint32
	auctionisactive      bool
	Highestbid           uint32
	time                 *time.Timer
	serverid             int32
}

var serverport = flag.Int("id", 0, "The ID of the client (used to calculate the port)")

func main() {

	flag.Parse()

	ownport := int(*serverport) + 5000

	logFile, err := os.OpenFile("../log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close() // Ensure the file is closed when the program ends

	// Set the output of the log package to the log file
	log.SetOutput(logFile)

	server := Server{
		serverid:             int32(ownport),
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
