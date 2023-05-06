package main

import (
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	pb "github.com/zawakin/simple-grpc-stream/api"
)

type server struct{}

func (s *server) ToUpperCase(stream pb.UpperCaseService_ToUpperCaseServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}

		log.Printf("ToUpperCase: %s", in.Text)

		err = stream.Send(&pb.OutputMessage{Text: strings.ToUpper(in.Text)})
		if err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUpperCaseServiceServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
