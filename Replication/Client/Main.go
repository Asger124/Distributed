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
