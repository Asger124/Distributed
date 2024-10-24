package __

import (
	pb "ChittyChat/proto"
	"fmt"
	"log"
	"google.golang.org/protobuf/proto"
)

func main() {
	//fmt.Println("i am main")

	chittyChat := &pb.MessageRequestCtS{
       Message: "hello", 
	   UserName: "jnojo", 

	}

	data, err := proto.Marshal(chittyChat)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(data)

	
}

