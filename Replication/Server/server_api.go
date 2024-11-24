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

		log.Printf("SERVER %d --> AUCTION HAS BEGUN DUE TO CLIENT WITH ID %d MAKING THE FIRST BID AT %d", s.serverid, amount.ClientID, amount.Amount)
		fmt.Print("THE AUCTION HAS BEGUN IT WILL BE ACTIVE FOR 1 MINUTE\n")
		s.auctionisactive = true
		s.time = time.NewTimer(1 * time.Minute)
		go func() {
			<-s.time.C
			s.auctionisactive = false
			log.Printf("SERVER %d --> AUCTION HAS ENDED - CLIENT %d WON WITH HIGHEST BID %d", s.serverid, s.ClientWithHighestbid, s.Highestbid)
			fmt.Printf("AUCTION HAS ENDED - CLIENT %d WON WITH HIGHEST BID %d", s.ClientWithHighestbid, s.Highestbid)
			//log.Printf("SERVER %d --> CLIENT %d HAS WON WITH HIGHEST BID %d", s.serverid, s.ClientWithHighestbid, s.Highestbid)

		}()

	}

	s.lamportstamp++
	s.clientID = amount.ClientID
	bid := amount.Amount

	if s.auctionisactive {

		if bid > s.Highestbid {
			s.Highestbid = bid
			s.ClientWithHighestbid = s.clientID
			log.Printf("SERVER %d --> CLIENT %d HAS MADE A BID %d, AT LAMPORTSTAMP %d, AND IT IS NOW THE HIGHEST BID", s.serverid, s.ClientWithHighestbid, bid, s.lamportstamp)
			response = fmt.Sprintf("YOUR BID %d HAS BEEN RECEIVED AND YOU ARE NOW THE HIGHEST BIDDER.\n", bid)

		} else if bid == s.Highestbid {
			log.Printf("SERVER %d --> CLIENT %d HAS MADE A BID %d, AT LAMPORTSTAMP %d, WHICH IS EQUAL TO THE HIGHEST BID, SO IT DOES NOT COUNT", s.serverid, s.clientID, bid, s.lamportstamp)
			//log.Printf("%d CURRENTLY HOLDS HIGHEST BID AT: %d", s.ClientWithHighestbid, s.Highestbid)
			response = fmt.Sprintf("YOUR BID %d HAS BEEN RECEIVED, YOUR BID WAS EQUAL TO THE HIGHEST, SO IT DOES NOT COUNT", bid)

		} else {
			log.Printf("SERVER %d --> CLIENT %d HAS MADE A BID %d, AT LAMPORTSTAMP %d, THE BID IS NOT HIGH ENOUGH", s.serverid, s.clientID, bid, s.lamportstamp)
			//log.Printf("%d CURRENTLY HOLDS HIGHEST BID AT: %d", s.ClientWithHighestbid, s.Highestbid)
			response = fmt.Sprintf("YOUR BID %d WAS NOT HIGH ENOUGH PLEASE TRY AGAIN", bid)

		}
	} else {

		log.Printf("SERVER %d --> CLIENT %d MADE A BID AT LAMPORTSTAMP %d, ERROR: AUCTION HAS ENDED, BID DOES NOT COUNT", s.serverid, s.clientID, s.lamportstamp)
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
		log.Printf("SERVER %d --> CLIENT %d HAS REQUESTED STATE OF AUCTION, AT LAMPORTSTAMP %d", s.serverid, s.clientID, s.lamportstamp)
		log.Printf("SERVER %d --> CLIENT %d CURRENTLY HOLDS HIGHEST BID AT %d", s.serverid, s.ClientWithHighestbid, s.Highestbid)

		response = fmt.Sprintf("AUCTION IS RUNNING. %s %d CURRENTLY HOLDS HIGHEST BID AT %d\n", state_c, s.ClientWithHighestbid, s.Highestbid)

	} else if s.lamportstamp == 0 {

		log.Printf("SERVER %d --> CLIENT %d HAS REQUESTED STATE OF AUCTION BEFORE AUCTION HAS STARTED", s.serverid, s.clientID)
		response = fmt.Sprintln("THE AUCTION HAS NOT STARTED YET. MAKE A BID TO START THE AUCTION")
	} else {
		s.lamportstamp++
		log.Printf("SERVER %d --> CLIENT %d HAS REQUESTED STATE OF AUCTION, AT LAMPORTSTAMP %d, AFTER AUCTION HAS ENDED", s.serverid, s.clientID, s.lamportstamp)
		log.Printf("SERVER %d --> AUCTION HAS ENDED AND CLIENT %d WON THE AUCTION WITH HIGHEST BID AT %d", s.serverid, s.ClientWithHighestbid, s.Highestbid)
		response = fmt.Sprintf("AUCTION HAS ENDED. %s %d WON THE AUCTION WITH HIGHEST BID AT %d", state_c, s.ClientWithHighestbid, s.Highestbid)

	}

	return &pb.Outcome{Outcome: response}, nil

}
