package main

import (
	"context"
	"io"
	"log"

	"google.golang.org/grpc"

	pb "github.com/zawakin/simple-grpc-stream/api"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
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
		stream.CloseSend()
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
