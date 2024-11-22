package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/Replication/proto"
)

var response string

func (s *Server) Bid(ctx context.Context, amount *pb.Amount) (*pb.Ack, error) {
	if s.lamportstamp == 0 && !s.auctionisactive {

		log.Printf("AUCTION HAS BEGUN DUE TO CLIENT WITH ID %d MAKING THE FIRST BID AT %d", amount.ClientID, amount.Amount)
		log.Printf("THE AUCTION WILL BE ACTIVE FOR 1 MINUTE")
		s.auctionisactive = true
		s.time = time.NewTimer(1 * time.Minute)
		go func() {
			<-s.time.C
			s.auctionisactive = false
			log.Printf("TIME HAS EXPIRED AND AUCTION HAS ENDED %d", s.serverid)
			log.Printf("Client %d HAS WON WITH HIGHEST BID %d", s.ClientWithHighestbid, s.Highestbid)

		}()
		s.Highestbid = amount.Amount
		s.ClientWithHighestbid = amount.ClientID

		log.Printf("%d CURRENTLY HAS HIGHEST BID AT: %d", s.ClientWithHighestbid, s.Highestbid)
	}

	s.lamportstamp++
	s.clientID = amount.ClientID
	bid := amount.Amount

	if s.auctionisactive {

		if bid > s.Highestbid {
			s.Highestbid = bid
			s.ClientWithHighestbid = s.clientID
			log.Printf("Client %d has made a bid %d, at Lamport stamp %d, and it is now the highest bid", s.ClientWithHighestbid, s.lamportstamp, bid)
			response = fmt.Sprintf("Your bid %d has been received and you are now the highest bidder!.\n", bid)

		}

		if bid == s.Highestbid {
			log.Printf("CLIENT %d HAS MADE A BID %d , at Lamport stamp %d, WHICH IS EQUAL TO THE HIGHEST BID, SO IT DOES NOT COUNT", s.clientID, s.lamportstamp, bid)
			log.Printf("%d CURRENTLY HOLDS HIGHEST BID AT: %d", s.ClientWithHighestbid, s.Highestbid)
			response = fmt.Sprintf("YOUR BID %d HAS BEEN RECEIVED, YOUR BID WAS EQUAL TO THE HIGHEST, SO IT DOES NOT COUNT", bid)

		}

		if bid < s.Highestbid {
			log.Printf("CLIENT %d HAS MADE A BID %d, at Lamport stamp%d, THE BID IS NOT HIGH ENOUGH", s.clientID, s.lamportstamp, bid)
			log.Printf("%d CURRENTLY HOLDS HIGHEST BID AT: %d", s.ClientWithHighestbid, s.Highestbid)
			response = fmt.Sprintf("YOUR BID %d WAS NOT HIGH ENOUGH PLEASE TRY AGAIN", bid)

		}
	} else {

		log.Printf("CLIENT %d MADE A BID AT LAMPORT STAMP %d, ERROR: AUCTION HAS ENDED, BID DOES NOT COUNT", s.clientID, s.lamportstamp)
		response = fmt.Sprintf("AUCTION HAS ENDED. YOUR BID %d WILL NOT BE ACCEPTED", bid)

	}

	return &pb.Ack{Ack: response}, nil

}

func (s *Server) Result(ctx context.Context, void *pb.Void) (*pb.Outcome, error) {

	s.clientID = void.ClientID
	state_c := "CLIENT"

	if s.clientID == s.ClientWithHighestbid {
		state_c = "YOU"

	}

	if s.auctionisactive {
		s.lamportstamp++
		log.Printf("CLIENT %d HAS REQUESTED STATE OF AUCTION, AT LAMPORT STAMP %d", s.clientID, s.lamportstamp)
		log.Printf("CLIENT %d CURRENTLY HOLDS HIGHEST BID AT %d", s.ClientWithHighestbid, s.Highestbid)

		response = fmt.Sprintf("YOU HAVE REQUESTED STATE OF AUCTION:\n AUCTION IS RUNNING. %s %d CURRENTLY HOLDS HIGHEST BID AT %d", state_c, s.ClientWithHighestbid, s.Highestbid)

	} else if s.lamportstamp == 0 {

		log.Printf("CLIENT %d HAS REQUESTED STATE OF AUCTION BEFORE AUCTION HAS STARTED", s.clientID)
		response = fmt.Sprintln("THE AUCTION HAS NOT STARTED YET. MAKE A BID TO START THE AUCTION")
	} else {
		s.lamportstamp++
		log.Printf("CLIENT %d HAS REQUESTED STATE OF AUCTION, AT LAMPORT STAMP %d, AFTER AUCTION HAS ENDED", s.clientID, s.lamportstamp)
		log.Printf("AUCTION HAS ENDED AND CLIENT %d WON THE AUCTION WITH HIGHEST BID AT %d", s.ClientWithHighestbid, s.Highestbid)
		response = fmt.Sprintf("YOU HAVE REQUESTED STATE OF AUCTION:\n AUCTION HAS ENDED. %s ,%d WON THE AUCTION WITH HIGHEST BID AT %d", state_c, s.ClientWithHighestbid, s.Highestbid)

	}

	return &pb.Outcome{Outcome: response}, nil

}
