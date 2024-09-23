package main

import (
	"fmt"
	"time"
)

// Define TCP flags as constants
const (
	SYN     = 1
	ACK     = 2
	SYN_ACK = SYN | ACK
)

// packet
type Packet struct {
	Flag int
}

//client
//Send a SYN
//recive a SYN-ACK
// send ACK

func client(clientToServer chan Packet, serverToClient chan Packet) {
	// Step 1: Send SYN to the server
	fmt.Println("Client: Sending SYN to server.")
	clientToServer <- Packet{Flag: SYN}

	// Step 2: Wait for SYN-ACK from the server
	synAckPacket := <-serverToClient
	if synAckPacket.Flag == SYN_ACK {
		fmt.Println("Client: Received SYN-ACK from server.")

		// Step 3: Send ACK to the server to establish the connection
		fmt.Println("Client: Sending ACK to server.")
		clientToServer <- Packet{Flag: ACK}

		fmt.Println("Client: Connection established with server.")
	}

}

// server
// Wait for a SYN, add something
// Send a SYN-ACK
// recive a ACK
func server(clientToServer chan Packet, serverToClient chan Packet) {

	// Step 1: Wait for SYN from the client
	synPacket := <-clientToServer
	if synPacket.Flag == SYN {
		fmt.Println("Server: Received SYN from client.")

		// Step 2: Send SYN-ACK to the client
		fmt.Println("Server: Sending SYN-ACK to client.")
		serverToClient <- Packet{Flag: SYN_ACK}

		// Step 3: Wait for the final ACK from the client
		ackPacket := <-clientToServer
		if ackPacket.Flag == ACK {
			fmt.Println("Server: Received ACK from client. Connection established.")
		}
	}

}

// main
func main() {
	// Create two channels to simulate the communication between client and server
	clientToServer := make(chan Packet)
	serverToClient := make(chan Packet)

	// Simulate the server in a separate goroutine
	go server(clientToServer, serverToClient)

	// Simulate the client in the main goroutine (for simplicity)
	client(clientToServer, serverToClient)

	// Wait a bit for the goroutines to complete
	time.Sleep(1 * time.Second)
}
