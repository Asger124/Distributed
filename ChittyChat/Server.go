package __

import (
	"context"
	"fmt"
	"sync"

	pb "ChittyChat/proto"

	//"google.golang.org/protobuf/proto"
)

type Connection struct {
	pb.UnimplementedChittyChatServiceServer
	stream pb.ChittyChatService_ChatStreamServer
	id     string
	active bool
	error  chan error
}

type Pool struct {
 	pb.UnimplementedChittyChatServiceServer
 	Connection []*Connection
}

func (p *Pool) CreateStream(pconn *pb.JoinRequestCtS, stream pb.ChittyChatService_ChatStreamServer) error {
	conn := &Connection{
	 stream: stream,
	 id:     pconn.UserName,
	 active: true,
	 error:  make(chan error),
	}
   
	p.Connection = append(p.Connection, conn)
   
	return <-conn.error
}

func (s *Pool) BroadcastMessage(ctx context.Context, msg *pb.MessageRequestCtS) (*pb.Close, error) {
	wait := sync.WaitGroup{}
	done := make(chan int)
   
	for _, conn := range s.Connection {
	 wait.Add(1)
   
	 go func(msg *pb.MessageRequestCtS, conn *Connection) {
	  defer wait.Done()
   
	  if conn.active {
	   err := conn.stream.Send(msg)
	   fmt.Printf("Sending message to: %v from %v", conn.id, msg.UserName)
   
	   if err != nil {
		fmt.Printf("Error with Stream: %v - Error: %v\n", conn.stream, err)
		conn.active = false
		conn.error <- err
	   }
	  }
	 }(msg, conn)
   
	}
   
	go func() {
	 wait.Wait()
	 close(done)
	}()
   
	<-done
	return &pb.Close{}, nil

}



/*func server() {
	
	fmt.Println("i am main")

	
}*/