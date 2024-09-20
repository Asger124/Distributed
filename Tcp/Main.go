package main

import "sync"

var wg sync.WaitGroup

func main() {

	Client := makeClient("Client")
	Server := makeClient("Server")
	Channel := make(chan *packet)

	wg.Add(2)

	go Client.SendPacket("Hello", 0, 0, Channel)
	go Server.receive_packet(Channel)

	wg.Wait()

}

// 	var header = makeHeader(0, 0)

// 	var pack1 = makePacket("Hello, did the handshake work?", header)
// 	var client = makeClient("client", pack1)

// 	go func() {
// 		client.channel <- *client.packet
// 	}()

// 	var pack2 = <-client.channel

// 	fmt.Println("client has send a syn number", pack2.header.syn_number)

// 	var server = makeClient("server", &pack2)

// 	server.packet.header.ack_number = (server.packet.header.syn_number + 1)

// 	fmt.Println("Server received a syn number from client, and has established an ack from it :", server.packet.header.ack_number)

// 	//set server syn number (y) before sending
// 	server.packet.header.syn_number = 100

// 	go func() {
// 		server.channel <- *server.packet

// 	}()

// 	updatedClientPacket := <-server.channel

// 	fmt.Printf("Server has send a syn: %d, and an ack : %d\n", updatedClientPacket.header.syn_number, updatedClientPacket.header.ack_number)

// 	client.packet.message = updatedClientPacket.message
// 	client.packet.header.syn_number = updatedClientPacket.header.syn_number
// 	client.packet.header.ack_number = updatedClientPacket.header.ack_number

// 	client.packet.header.syn_number++

// 	fmt.Println("Client received a syn number and has established ack from it", client.packet.header.syn_number)

// 	go func() {
// 		client.channel <- *client.packet
// 	}()

// 	updatedServerPacket := <-client.channel
// 	server.packet.message = updatedServerPacket.message
// 	server.packet.header.ack_number = updatedServerPacket.header.syn_number

// 	fmt.Printf("Server received final ack: %d, Connection has been established, and messages can be send\n %s", server.packet.header.ack_number, server.packet.message)

// }
