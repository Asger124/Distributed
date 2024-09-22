package main

import "sync"

var wg sync.WaitGroup

func main() {

	Client := makeClient("Client")
	Server := makeClient("Server")
	Channel := make(chan *packet)

	wg.Add(2)

	go Client.SendPacket("Hello i would like to establish a connection", 0, 0, Channel)
	go Server.receive_packet(Channel)

	wg.Wait()

}
