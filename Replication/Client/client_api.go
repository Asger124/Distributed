package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/Replication/proto"
)

func ClientBids(ctx context.Context, bid uint32, c *Client) {

	request := pb.Amount{
		ClientID: c.id,
		Amount:   bid,
	}

	response := &pb.Ack{}

	c.responsecount = 0

	for _, connect := range c.servers {

		res, err := connect.Bid(context.Background(), &request)
		c.responsecount++
		if err != nil {
			c.responsecount--
			continue
		}

		response = res
	}

	if c.responsecount == 0 {

		log.Printf("No responses where provided servers have crashed")
	}

	length := uint32(len(c.servers))
	ServersDown := length - c.responsecount

	if ServersDown > c.servercrash {

		c.servercrash = ServersDown

		log.Printf("Warning: %d/%d Servers did not respond", c.servercrash, length)

	}

	if response.Ack == "" {

		fmt.Printf("YOU RECEIVED NO RESPONSE FROM THE SERVICE. STANDBY WHILE WE CHECK THE SERVERS..\n")
		log.Printf("Client %d GOT NO RESPONSE AT ALL. CHECKING SERVERS..", c.id)
		time.Sleep(2 * time.Second)
		if c.servercrash == length {
			fmt.Printf("ALL SERVERS ARE DOWN. AUCTION HAS BEEN CLOSED, NO WINNERS WAS FOUND\n")
			log.Printf("CLIENT %d RECEIVED MESSAGE: ALL SERVERS ARE DOWN. AUCTION IS CLOSED AND NO WINNER WAS FOUND DUE TO CRASH", c.id)
		}
	} else {
		fmt.Printf("RESPONSE FROM AUCTION: %v\n", response.Ack)
		log.Printf("CLIENT %d REVEIVED RESPONSE FROM AUCTON: %v\n", c.id, response.Ack)
	}

}

func AskForState(ctx context.Context, c *Client) {

	request := pb.Void{
		ClientID: c.id}

	response := &pb.Outcome{}

	c.responsecount = 0

	for _, connect := range c.servers {

		res, err := connect.Result(context.Background(), &request)
		c.responsecount++
		if err != nil {
			c.responsecount--
			continue
		}

		response = res
	}

	if c.responsecount == 0 {

		log.Printf("No responses where provided servers have crashed")
	}

	length := uint32(len(c.servers))
	ServersDown := length - c.responsecount

	if ServersDown > c.servercrash {

		c.servercrash = ServersDown

		log.Printf("Warning: %d/%d Servers did not respond", c.servercrash, length)

	}

	if response.Outcome == "" {

		fmt.Printf("YOU RECEIVED NO RESPONSE FROM THE SERVICE. STANDBY WHILE WE CHECK THE SERVERS..\n")
		log.Printf("Client %d GOT NO RESPONSE AT ALL. CHECKING SERVERS..", c.id)
		time.Sleep(2 * time.Second)
		if c.servercrash == length {
			fmt.Printf("ALL SERVERS ARE DOWN. AUCTION HAS BEEN CLOSED, NO WINNERS WAS FOUND\n")
			log.Printf("CLIENT %d RECEIVED MESSAGE: ALL SERVERS ARE DOWN. AUCTION IS CLOSED AND NO WINNER WAS FOUND DUE TO CRASH", c.id)
		}
	} else {
		fmt.Printf("RESPONSE FROM AUCTION: %v\n", response.Outcome)
		log.Printf("CLIENT %d REVEIVED RESPONSE FROM AUCTON: %v\n", c.id, response.Outcome)
	}

}
