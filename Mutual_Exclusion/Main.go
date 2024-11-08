package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "example.com/Mutual_Exclusion/proto"
	"google.golang.org/grpc"
)

type peer struct {
	id         int
	address    string
	p_next     *peer
	token_chan chan bool
	server     *grpc.Server
	serve      *pb.UnimplementedTokenRingServer
}

func createPeer(id int, addr string) *peer {
	p := peer{}
	p.id = id
	p.address = addr
	p.token_chan = make(chan bool)
	return &p
}

func (p *peer) StartServer() {
	lis, err := net.Listen("tcp", p.address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", p.address, err)
	}

	p.server = grpc.NewServer()
	pb.RegisterTokenRingServer(p.server, p.serve) // Register this node as a PeerServer

	log.Printf("Node %d listening on %s\n", p.id, p.address)
	if err := p.server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func main() {

	number := flag.Int("UserID", -1, "portnumber")
	flag.Parse()
	const baseport = 5000
	num_peers := make([]*peer, 3)

	for i := range num_peers {
		address := fmt.Sprintf("localhost:%d", baseport+i)

		num_peers[i] = createPeer(i, address)

	}

	go num_peers[*number].StartServer()

	for i := range num_peers {

		num_peers[i].p_next = num_peers[(i+1)%len(num_peers)]

	}

	select {}

	// conn1, err := grpc.Dial(addr1, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }
	// defer conn1.Close()

	// client1 := pb.NewTimeServiceClient(conn1)

	// addr2 := fmt.Sprintf("localhost:%s", *port2)

	// conn2, err := grpc.Dial(addr2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("Did not connect: %v", err)
	// }
}
