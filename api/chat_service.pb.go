// Code generated by protoc-gen-go. DO NOT EDIT.
// source: chat_service.proto

package api

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ChatMessage struct {
	User                 string   `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	Message              string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Timestamp            uint64   `protobuf:"varint,3,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ChatMessage) Reset()         { *m = ChatMessage{} }
func (m *ChatMessage) String() string { return proto.CompactTextString(m) }
func (*ChatMessage) ProtoMessage()    {}
func (*ChatMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_e881f93b317ee75c, []int{0}
}

func (m *ChatMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ChatMessage.Unmarshal(m, b)
}
func (m *ChatMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ChatMessage.Marshal(b, m, deterministic)
}
func (m *ChatMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChatMessage.Merge(m, src)
}
func (m *ChatMessage) XXX_Size() int {
	return xxx_messageInfo_ChatMessage.Size(m)
}
func (m *ChatMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_ChatMessage.DiscardUnknown(m)
}

var xxx_messageInfo_ChatMessage proto.InternalMessageInfo

func (m *ChatMessage) GetUser() string {
	if m != nil {
		return m.User
	}
	return ""
}

func (m *ChatMessage) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *ChatMessage) GetTimestamp() uint64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func init() {
	proto.RegisterType((*ChatMessage)(nil), "api.ChatMessage")
}

func init() { proto.RegisterFile("chat_service.proto", fileDescriptor_e881f93b317ee75c) }

var fileDescriptor_e881f93b317ee75c = []byte{
	// 155 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4a, 0xce, 0x48, 0x2c,
	0x89, 0x2f, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62,
	0x4e, 0x2c, 0xc8, 0x54, 0x8a, 0xe4, 0xe2, 0x76, 0xce, 0x48, 0x2c, 0xf1, 0x4d, 0x2d, 0x2e, 0x4e,
	0x4c, 0x4f, 0x15, 0x12, 0xe2, 0x62, 0x29, 0x2d, 0x4e, 0x2d, 0x92, 0x60, 0x54, 0x60, 0xd4, 0xe0,
	0x0c, 0x02, 0xb3, 0x85, 0x24, 0xb8, 0xd8, 0x73, 0x21, 0xd2, 0x12, 0x4c, 0x60, 0x61, 0x18, 0x57,
	0x48, 0x86, 0x8b, 0xb3, 0x24, 0x33, 0x37, 0xb5, 0xb8, 0x24, 0x31, 0xb7, 0x40, 0x82, 0x59, 0x81,
	0x51, 0x83, 0x25, 0x08, 0x21, 0x60, 0xe4, 0x0c, 0x31, 0x3a, 0x18, 0x62, 0xa9, 0x90, 0x09, 0x17,
	0x17, 0x98, 0x5b, 0x52, 0x94, 0x9a, 0x98, 0x2b, 0x24, 0xa0, 0x97, 0x58, 0x90, 0xa9, 0x87, 0x64,
	0xb5, 0x14, 0x86, 0x88, 0x06, 0xa3, 0x01, 0x63, 0x12, 0x1b, 0xd8, 0xad, 0xc6, 0x80, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x3d, 0xc3, 0x49, 0xda, 0xc1, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ChatServiceClient is the client API for ChatService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ChatServiceClient interface {
	ChatStream(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatStreamClient, error)
}

type chatServiceClient struct {
	cc *grpc.ClientConn
}

func NewChatServiceClient(cc *grpc.ClientConn) ChatServiceClient {
	return &chatServiceClient{cc}
}

func (c *chatServiceClient) ChatStream(ctx context.Context, opts ...grpc.CallOption) (ChatService_ChatStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ChatService_serviceDesc.Streams[0], "/api.ChatService/ChatStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &chatServiceChatStreamClient{stream}
	return x, nil
}

type ChatService_ChatStreamClient interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ClientStream
}

type chatServiceChatStreamClient struct {
	grpc.ClientStream
}

func (x *chatServiceChatStreamClient) Send(m *ChatMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chatServiceChatStreamClient) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ChatServiceServer is the server API for ChatService service.
type ChatServiceServer interface {
	ChatStream(ChatService_ChatStreamServer) error
}

// UnimplementedChatServiceServer can be embedded to have forward compatible implementations.
type UnimplementedChatServiceServer struct {
}

func (*UnimplementedChatServiceServer) ChatStream(srv ChatService_ChatStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method ChatStream not implemented")
}

func RegisterChatServiceServer(s *grpc.Server, srv ChatServiceServer) {
	s.RegisterService(&_ChatService_serviceDesc, srv)
}

func _ChatService_ChatStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChatServiceServer).ChatStream(&chatServiceChatStreamServer{stream})
}

type ChatService_ChatStreamServer interface {
	Send(*ChatMessage) error
	Recv() (*ChatMessage, error)
	grpc.ServerStream
}

type chatServiceChatStreamServer struct {
	grpc.ServerStream
}

func (x *chatServiceChatStreamServer) Send(m *ChatMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chatServiceChatStreamServer) Recv() (*ChatMessage, error) {
	m := new(ChatMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChatService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.ChatService",
	HandlerType: (*ChatServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ChatStream",
			Handler:       _ChatService_ChatStream_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "chat_service.proto",
}
