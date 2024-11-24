package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	pb "example.com/Replication/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	pb.UnimplementedAuctionServer
	servers       map[int32]pb.AuctionClient
	id            int32
	responsecount uint32
	servercrash   uint32
}

var clientport = flag.Int("id", 0, "The ID of the client (used to calculate the port)")

func main() {
	flag.Parse()

	ownport := int(*clientport) + 6000

	logFile, err := os.OpenFile("../log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", err)
	}
	defer logFile.Close() // Ensure the file is closed when the program ends

	// Set the output of the log package to the log file
	log.SetOutput(logFile)

	client := &Client{
		id:          int32(ownport),
		servers:     make(map[int32]pb.AuctionClient),
		servercrash: 0,
	}

	for i := 0; i < 3; i++ {
		port := int32(5001) + int32(i)

		var conn *grpc.ClientConn
		conn, err := grpc.Dial(fmt.Sprintf(":%v", port), grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Could not connect: %s\n", err)
		}
		defer conn.Close()
		c := pb.NewAuctionClient(conn)
		client.servers[port] = c
	}

	fmt.Printf("\n_____________________________________________________________\nWelcome to the AUCTION. You have the ID: %v.\nWrite 'result' and press ENTER to get info about the auction.\nWrite a number and press ENTER to send a BID.\n_____________________________________________________________\n", client.id)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() != "result" {
			bid, _ := strconv.Atoi(scanner.Text())
			ClientBids(context.Background(), int32(bid), client)
		} else {
			AskForState(context.Background(), client)
		}

	}
}
