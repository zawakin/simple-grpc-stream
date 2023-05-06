package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/zawakin/simple-grpc-stream/api"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUpperCaseServiceClient(conn)

	ctx := context.Background()
	stream, err := client.ToUpperCase(ctx)
	if err != nil {
		log.Fatalf("failed to call ToUpperCase: %v", err)
	}

	go func() {
		for _, text := range []string{"hello", "world", "grpc", "streaming"} {
			if err := stream.Send(&pb.InputMessage{Text: text}); err != nil {
				log.Fatalf("failed to send message: %v", err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Fatalf("failed to close send: %v", err)
		}
	}()

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("failed to receive message: %v", err)
		}
		log.Printf("UpperCase: %s", res.Text)
	}
}
