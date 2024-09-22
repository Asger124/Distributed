package main

import (
	"fmt"
)

type client struct {
	name string
}

func makeClient(name string) *client {
	c := client{}
	c.name = name
	return &c

}

func (c *client) SendPacket(msg string, syn int, ack int, channel chan *packet) {
	defer wg.Done()

	header := makeHeader(syn, ack)
	packet := makePacket(msg, header)

	//Sends a packet through the channel (first step of Three-way-handshake)
	//When the packet have been send, the client will wait until another go routine receives the packet.
	channel <- packet

	fmt.Printf("%s sent SYN=%d and ACK=%d\n", c.name, packet.header.syn_number, packet.header.ack_number)

	// Receives a packet with ack
	syn_ack_pack := <-channel

	fmt.Printf("%s received SYN=%d and ACK=%d\n", c.name, syn_ack_pack.header.syn_number, syn_ack_pack.header.ack_number)

	header2 := makeHeader(syn+1, syn_ack_pack.header.syn_number+1)
	packet2 := makePacket(msg, header2)
	fmt.Printf("%s sent final ACK=%d\n", c.name, packet2.header.ack_number)

	//Sends final packet throught the channel
	channel <- packet2

}

func (c *client) receive_packet(channel chan *packet) {
	defer wg.Done()

	// receive the first packet which is send through a channel
	// The host will wait until they receive a packet from the channel from another go routine
	received_pack := <-channel

	fmt.Printf("%s received SYN=%d and ACK=%d\n", c.name, received_pack.header.syn_number, received_pack.header.ack_number)

	//increments the received syn number
	received_pack.header.syn_number++

	//Set Ack number to be sent
	received_pack.header.ack_number = received_pack.header.syn_number

	//make new syn number to be sent
	received_pack.header.syn_number = 100

	// Send the updated packet with new syn number and an ack number
	// Host will halt until a go routine receives the pack from a channel
	channel <- received_pack

	fmt.Printf("%s Sent SYN=%d and ACK=%d\n", c.name, received_pack.header.syn_number, received_pack.header.ack_number)

	// Final step of three-way handshake
	// The host will halt until it receives a final packet from another go routine - effectively establishing connection.
	finalAck := <-channel

	fmt.Printf("%s received final ACK=%d\n", c.name, finalAck.header.ack_number)
	fmt.Println("A message has been sent: ", finalAck.message)

}
