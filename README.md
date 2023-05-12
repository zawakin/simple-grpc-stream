# Simple Chat gRPC Streaming Example

## Description

This is a simple chat gRPC stream example. It is a simple client/server application that sends a stream of messages from the client to the server and the server responds with a stream of messages back to the client.

## Files

```
.
├── README.md
├── api
│   └── chat_service.pb.go
├── chat_service.proto
├── client
│   └── main.go
├── go.mod
├── go.sum
└── server
    └── main.go

3 directories, 7 files
```

## Usage

### Server

To run the server, run the following command:

```bash
go run server/main.go
```

### Client

To run the client, run the following command:

```bash
go run client/main.go
```

### Generate gRPC code

Install the gRPC code generator(protoc):

```bash
# macOS
brew install protobuf
```

To generate the gRPC code, run the following command:

```bash
protoc --go_out=plugins=grpc:./api ./*.proto
```
