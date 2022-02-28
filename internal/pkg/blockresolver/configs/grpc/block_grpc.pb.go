// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package blockresolver

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlockResolverClient is the client API for BlockResolver service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockResolverClient interface {
	SendBlocks(ctx context.Context, opts ...grpc.CallOption) (BlockResolver_SendBlocksClient, error)
}

type blockResolverClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockResolverClient(cc grpc.ClientConnInterface) BlockResolverClient {
	return &blockResolverClient{cc}
}

func (c *blockResolverClient) SendBlocks(ctx context.Context, opts ...grpc.CallOption) (BlockResolver_SendBlocksClient, error) {
	stream, err := c.cc.NewStream(ctx, &BlockResolver_ServiceDesc.Streams[0], "/BlockResolver/SendBlocks", opts...)
	if err != nil {
		return nil, err
	}
	x := &blockResolverSendBlocksClient{stream}
	return x, nil
}

type BlockResolver_SendBlocksClient interface {
	Send(*Block) error
	Recv() (*BlockId, error)
	grpc.ClientStream
}

type blockResolverSendBlocksClient struct {
	grpc.ClientStream
}

func (x *blockResolverSendBlocksClient) Send(m *Block) error {
	return x.ClientStream.SendMsg(m)
}

func (x *blockResolverSendBlocksClient) Recv() (*BlockId, error) {
	m := new(BlockId)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockResolverServer is the server API for BlockResolver service.
// All implementations must embed UnimplementedBlockResolverServer
// for forward compatibility
type BlockResolverServer interface {
	SendBlocks(BlockResolver_SendBlocksServer) error
	mustEmbedUnimplementedBlockResolverServer()
}

// UnimplementedBlockResolverServer must be embedded to have forward compatible implementations.
type UnimplementedBlockResolverServer struct {
}

func (UnimplementedBlockResolverServer) SendBlocks(BlockResolver_SendBlocksServer) error {
	return status.Errorf(codes.Unimplemented, "method SendBlocks not implemented")
}
func (UnimplementedBlockResolverServer) mustEmbedUnimplementedBlockResolverServer() {}

// UnsafeBlockResolverServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockResolverServer will
// result in compilation errors.
type UnsafeBlockResolverServer interface {
	mustEmbedUnimplementedBlockResolverServer()
}

func RegisterBlockResolverServer(s grpc.ServiceRegistrar, srv BlockResolverServer) {
	s.RegisterService(&BlockResolver_ServiceDesc, srv)
}

func _BlockResolver_SendBlocks_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(BlockResolverServer).SendBlocks(&blockResolverSendBlocksServer{stream})
}

type BlockResolver_SendBlocksServer interface {
	Send(*BlockId) error
	Recv() (*Block, error)
	grpc.ServerStream
}

type blockResolverSendBlocksServer struct {
	grpc.ServerStream
}

func (x *blockResolverSendBlocksServer) Send(m *BlockId) error {
	return x.ServerStream.SendMsg(m)
}

func (x *blockResolverSendBlocksServer) Recv() (*Block, error) {
	m := new(Block)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// BlockResolver_ServiceDesc is the grpc.ServiceDesc for BlockResolver service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockResolver_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "BlockResolver",
	HandlerType: (*BlockResolverServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "SendBlocks",
			Handler:       _BlockResolver_SendBlocks_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "configs/grpc/block.proto",
}