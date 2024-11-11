package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"

	pb "example.com/Mutual_Exclusion/Proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//type Token struct{ message string }

type peer struct {
	id         int
	address    string
	p_next     *peer
	token_chan chan *pb.Token
	server     *grpc.Server
	pb.TokenRingServer
	cient   pb.TokenRingClient
	request bool
}

func createPeer(id int, addr string) *peer {
	p := peer{}
	p.id = id
	p.address = addr
	p.token_chan = make(chan *pb.Token)
	return &p
}

func (p *peer) StartServer() {
	lis, err := net.Listen("tcp", p.address)
	if err != nil {
		log.Fatalf("Failed to listen on %s: %v", p.address, err)
	}

	p.server = grpc.NewServer()
	pb.RegisterTokenRingServer(p.server, p)

	//log.Printf("Node %d listening on %s\n", p.id, p.address)
	if err := p.server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
func (p *peer) StartClient() {

	conn, err := grpc.Dial(p.p_next.address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Printf("Failed to connect to node %s: %v", p.p_next.address, err)
		return
	}
	p.cient = pb.NewTokenRingClient(conn)
	log.Printf("Node %d has connected to node %d", p.id, p.p_next.id)

}

func (p *peer) ReceiveToken(ctx context.Context, msg *pb.Token) (*pb.Empty, error) {
	// This is called when a peer receives a token
	// Simulate critical section work
	time.Sleep(time.Second)

	// Token is passed into peers channel

	p.token_chan <- msg
	return &pb.Empty{}, nil
}

func (p *peer) simulateRequest() {
	time.Sleep(time.Second)
	for {
		if !p.request {
			// Wait for a random interval before making a request
			time.Sleep(time.Duration(rand.Intn(10)+8) * time.Second)
			p.request = true
			log.Printf("Node %d is requesting access critical section", p.id)
		} else {
			time.Sleep(time.Second)
		}
	}
}

func (p *peer) EnterCs_andPasstoken() {
	for {
		//This will block until a token is received from another go routine. receiveToken() is what makes this section unblock
		token := <-p.token_chan
		time.Sleep(time.Second)
		log.Printf("Node %d got %s\n", p.id, token.Message)

		if p.request {
			log.Printf("Node %d is entering critical section", p.id)
			// Simulate critical section work
			time.Sleep(3 * time.Second)
			log.Printf("Node %d is leaving critical section", p.id)
			p.request = false
		}

		time.Sleep(time.Second)

		//This is effectively what passes on the token. p.cient is a GRPC client that communicates with the next peer in the ring
		//This means that when p.cientReceiveToken() is called, it is actually the next Peer in the ring that receives the token
		if p.cient != nil {
			_, err := p.cient.ReceiveToken(context.Background(), token)
			if err != nil {
				log.Printf("Node %d failed to pass token to node %d: %v", p.id, p.p_next.id, err)
			}

		}
		time.Sleep(1 * time.Second)
	}

}

func main() {

	log.SetFlags(0)
	const baseport = 5000
	num_peers := make([]*peer, 3)

	for i := range num_peers {
		address := fmt.Sprintf("localhost:%d", baseport+i)

		num_peers[i] = createPeer(i, address)

		go num_peers[i].StartServer()

	}

	for i := range num_peers {
		num_peers[i].p_next = num_peers[(i+1)%len(num_peers)]
	}

	for i := range num_peers {
		go num_peers[i].StartClient()
		go num_peers[i].EnterCs_andPasstoken()
	}

	initialToken := pb.Token{Message: "encrypted token"}
	time.Sleep(time.Second)
	log.Printf("\nHello and welcome to a token ring simulation that demonstrates distributed mutual exclusion for three Nodes\nTo start the simulation write : 'START' and hit enter. To quit the program hit ctrl+c")

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()

		if input == "START" {

			num_peers[0].ReceiveToken(context.Background(), &initialToken)

			go num_peers[0].simulateRequest()
			go num_peers[1].simulateRequest()
			go num_peers[2].simulateRequest()

		}

	}

	select {}
}
