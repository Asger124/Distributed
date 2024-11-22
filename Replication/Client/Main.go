package main

import (
	pb "example.com/Replication/proto"
)

type Client struct {
	pb.UnimplementedAuctionServer
	servers       []pb.AuctionClient
	id            uint32
	responsecount uint32
	servercrash   uint32
}
