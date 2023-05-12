package main

import (
	"io"
	"log"
	"net"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/zawakin/simple-grpc-stream/api"
)

type ClientID uuid.UUID

func NewClientID() ClientID {
	return ClientID(uuid.New())
}

type ClientStream struct {
	clientID ClientID
	stream   api.ChatService_ChatStreamServer
}

type Server struct {
	clients map[ClientID]*ClientStream
	mu      sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients: make(map[ClientID]*ClientStream),
	}
}

func (s *Server) ChatStream(stream api.ChatService_ChatStreamServer) error {
	clientStream := &ClientStream{
		clientID: NewClientID(),
		stream:   stream,
	}

	s.addClient(clientStream)
	defer s.removeClient(clientStream)

	for {
		msg, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				log.Printf("Client %s disconnected", clientStream.clientID)
				return nil
			}
			log.Printf("Failed to receive message: %v", err)
			return err
		}

		log.Printf("%s: %s", msg.GetUser(), msg.GetMessage())

		err = s.broadcast(clientStream.clientID, msg)
		if err != nil {
			log.Printf("Failed to broadcast message: %v", err)
			return err
		}
	}
}

func (s *Server) addClient(clientStream *ClientStream) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.clients[clientStream.clientID] = clientStream
}

func (s *Server) removeClient(clientStream *ClientStream) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.clients, clientStream.clientID)
}

func (s *Server) broadcast(clientID ClientID, msg *api.ChatMessage) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, client := range s.clients {
		if client.clientID == clientID {
			continue
		}
		err := client.stream.Send(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Disable timestamps in log output
	log.SetFlags(0)

	s := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             30 * time.Second, // Minimum acceptable time between client pings
			PermitWithoutStream: false,            // Don't allow pings when there are no active streams
		}),
	)
	api.RegisterChatServiceServer(s, NewServer())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
