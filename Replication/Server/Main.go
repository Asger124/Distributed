package main

import (
	"time"

	pb "example.com/Replication/proto"
)

type Server struct {
	pb.UnimplementedAuctionServer
	clientID             uint32
	port                 string
	lamportstamp         uint32
	ClientWithHighestbid uint32
	auctionisactive      bool
	Highestbid           uint32
	time                 *time.Timer
	serverid             uint32
}
