package main

import (
	"fmt"
	"sync"
)

func main() {

serverCh := make(chan int)

clientCh := make(chan int)
var wg sync.WaitGroup

wg.Add(2) 



//clientCh goroutine 
go func() {

//wating for the syn msg from the client-channel
	defer wg.Done()
	serverCh <- 100
	synAck1 := <- clientCh
	synAck2 := <- clientCh
	fmt.Println(synAck1, synAck2)
	
	serverCh <- 301
	serverCh <- 101

	
}()
	 
//serverCh goroutine 
go func() {

	defer wg.Done() 
    syn := <-serverCh 
	fmt.Println( syn )

	
	clientCh <- 101
	clientCh <- 300
	ack1:= <-serverCh 
	ack2:= <- serverCh
	fmt.Println( ack1, ack2 )

	}() 
		wg.Wait()

}


