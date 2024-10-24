package main

import (
	"bufio"
	"context"
	"flag"
	"io"
	"log"
	"os"
	"unicode/utf8"

	pb "example.com/Chitty_Chat/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

var ParName = flag.String("name", "Unknown", "Participant")

var clientLamport = uint32(0)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	log.SetFlags(0)

	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	// Create a client using the established connection
	client := pb.NewChitty_ChatClient(conn)

	stream, err := client.Broadcast(context.Background(), &pb.Participant{Name: *ParName})

	if err != nil {
		log.Fatalf("Error calling broadcast %v", err)
	} else {
		log.Printf("\nWelcome to Chitty-chat %s! To send a message type it in below and hit enter! To exit the chat type in 'Quit', and hit enter.\nFollowing messages have already been sent:\n\n", *ParName)

	}

	go func() {

		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				break // End of stream
			}
			if err != nil {
				log.Fatalf("Error receiving message: %v", err)
			}

			if msg.LamportTime > clientLamport {
				clientLamport = msg.LamportTime
				clientLamport++
			} else {
				clientLamport++
			}

			if msg.Par.Name == *ParName {

				log.Printf("(%d) %s: %s", msg.LamportTime, "You", msg.Message)

			} else {

				log.Printf("(%d) %s: %s", msg.LamportTime, msg.Par.Name, msg.Message)

			}
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "Quit" {
			client.Leave(context.Background(), &pb.Participant{Name: *ParName})
			break
		}
		if utf8.RuneCountInString(scanner.Text()) <= 128 {
			client.Publish(context.Background(), &pb.Message{
				Message:     scanner.Text(),
				Par:         &pb.Participant{Name: *ParName},
				LamportTime: clientLamport,
			})

		} else {

			log.Println("The message is too long. Message must at max contain 128 characters, please try again.")

		}

	}
}
