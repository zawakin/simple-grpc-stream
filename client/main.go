package main

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"

	"github.com/zawakin/simple-grpc-stream/api"
)

func main() {
	conn, err := grpc.Dial(":50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                45 * time.Second, // Send keepalive pings every 45 seconds
			Timeout:             10 * time.Second,
			PermitWithoutStream: false, // Don't allow pings when there are no active streams

		}),
	)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := api.NewChatServiceClient(conn)
	ctx := context.Background()

	stream, err := client.ChatStream(ctx)
	if err != nil {
		log.Fatalf("Failed to open chat stream: %v", err)
	}

	go receiveMessage(stream)

	reader := bufio.NewReader(os.Stdin)
	readLine := func() string {
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Failed to read username: %v", err)
		}
		return strings.TrimRight(text, "\n")
	}

	fmt.Print("Enter your username: ")
	username := readLine()

	fmt.Println("Start chatting! Type your messages and press Enter to send.")
	for {
		text := readLine()

		msg := &api.ChatMessage{
			User:      username,
			Message:   text,
			Timestamp: uint64(time.Now().Unix()),
		}

		err = stream.Send(msg)
		if err != nil {
			log.Printf("Failed to send message: %v", err)
			return
		}
	}
}

func receiveMessage(stream api.ChatService_ChatStreamClient) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Server closed the connection")
				return
			}
			log.Printf("Failed to receive message: %v", err)
			return
		}
		fmt.Printf("%s: %s\n", msg.GetUser(), msg.GetMessage())
	}
}
