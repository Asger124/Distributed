package main

import (
	"fmt"
)

type client struct {
	name string
	//packet  *packet
	//channel chan packet
}

func makeClient(name string) *client {
	c := client{}
	c.name = name
	//c.packet = pack
	//c.channel = make(chan packet)
	return &c

}

func (c *client) SendPacket(msg string, syn int, ack int, channel chan *packet) {
	defer wg.Done()

	header := makeHeader(syn, ack)
	packet := makePacket(msg, header)

	channel <- packet

	fmt.Println("Send packet_1 :", packet.header.syn_number, packet.header.ack_number)

	//time.Sleep(time.Millisecond * time.Duration(3000))

	syn_ack_pack := <-channel

	header2 := makeHeader(syn+1, syn_ack_pack.header.syn_number+1)
	packet2 := makePacket(msg, header2)
	fmt.Println("Send final packet", packet2.header.syn_number, packet2.header.ack_number)

	channel <- packet2

}

func (c *client) receive_packet(channel chan *packet) {
	defer wg.Done()

	received_pack := <-channel

	fmt.Println("Received pack before modifying", received_pack.header.syn_number, received_pack.header.ack_number)

	received_pack.header.syn_number++

	received_pack.header.ack_number = received_pack.header.syn_number

	//make new syn number to be send
	received_pack.header.syn_number = 100

	channel <- received_pack

	fmt.Println("Sending back a new syn number, alongside and ack", received_pack.header.syn_number, received_pack.header.ack_number)

	finalAck := <-channel

	fmt.Printf("%s received final ACK: SYN=%d, ACK=%d\n", c.name, finalAck.header.syn_number, finalAck.header.ack_number)

}
