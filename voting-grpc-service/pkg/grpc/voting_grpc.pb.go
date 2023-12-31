// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.25.1
// source: voting.proto

package grpc

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

// VotingClient is the client API for Voting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VotingClient interface {
	CreatePoll(ctx context.Context, in *NewPoll, opts ...grpc.CallOption) (*Poll, error)
	AddPollAlbums(ctx context.Context, opts ...grpc.CallOption) (Voting_AddPollAlbumsClient, error)
	GetPollDetails(ctx context.Context, in *PollQuery, opts ...grpc.CallOption) (*PollDetails, error)
	VoteAlbum(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*Vote, error)
	GetAlbumVote(ctx context.Context, in *AlbumVotesQuery, opts ...grpc.CallOption) (Voting_GetAlbumVoteClient, error)
}

type votingClient struct {
	cc grpc.ClientConnInterface
}

func NewVotingClient(cc grpc.ClientConnInterface) VotingClient {
	return &votingClient{cc}
}

func (c *votingClient) CreatePoll(ctx context.Context, in *NewPoll, opts ...grpc.CallOption) (*Poll, error) {
	out := new(Poll)
	err := c.cc.Invoke(ctx, "/grpc.Voting/CreatePoll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingClient) AddPollAlbums(ctx context.Context, opts ...grpc.CallOption) (Voting_AddPollAlbumsClient, error) {
	stream, err := c.cc.NewStream(ctx, &Voting_ServiceDesc.Streams[0], "/grpc.Voting/AddPollAlbums", opts...)
	if err != nil {
		return nil, err
	}
	x := &votingAddPollAlbumsClient{stream}
	return x, nil
}

type Voting_AddPollAlbumsClient interface {
	Send(*NewPollAlbum) error
	Recv() (*NewPollAlbumResponse, error)
	grpc.ClientStream
}

type votingAddPollAlbumsClient struct {
	grpc.ClientStream
}

func (x *votingAddPollAlbumsClient) Send(m *NewPollAlbum) error {
	return x.ClientStream.SendMsg(m)
}

func (x *votingAddPollAlbumsClient) Recv() (*NewPollAlbumResponse, error) {
	m := new(NewPollAlbumResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *votingClient) GetPollDetails(ctx context.Context, in *PollQuery, opts ...grpc.CallOption) (*PollDetails, error) {
	out := new(PollDetails)
	err := c.cc.Invoke(ctx, "/grpc.Voting/GetPollDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingClient) VoteAlbum(ctx context.Context, in *Vote, opts ...grpc.CallOption) (*Vote, error) {
	out := new(Vote)
	err := c.cc.Invoke(ctx, "/grpc.Voting/VoteAlbum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *votingClient) GetAlbumVote(ctx context.Context, in *AlbumVotesQuery, opts ...grpc.CallOption) (Voting_GetAlbumVoteClient, error) {
	stream, err := c.cc.NewStream(ctx, &Voting_ServiceDesc.Streams[1], "/grpc.Voting/GetAlbumVote", opts...)
	if err != nil {
		return nil, err
	}
	x := &votingGetAlbumVoteClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Voting_GetAlbumVoteClient interface {
	Recv() (*Vote, error)
	grpc.ClientStream
}

type votingGetAlbumVoteClient struct {
	grpc.ClientStream
}

func (x *votingGetAlbumVoteClient) Recv() (*Vote, error) {
	m := new(Vote)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// VotingServer is the server API for Voting service.
// All implementations must embed UnimplementedVotingServer
// for forward compatibility
type VotingServer interface {
	CreatePoll(context.Context, *NewPoll) (*Poll, error)
	AddPollAlbums(Voting_AddPollAlbumsServer) error
	GetPollDetails(context.Context, *PollQuery) (*PollDetails, error)
	VoteAlbum(context.Context, *Vote) (*Vote, error)
	GetAlbumVote(*AlbumVotesQuery, Voting_GetAlbumVoteServer) error
	mustEmbedUnimplementedVotingServer()
}

// UnimplementedVotingServer must be embedded to have forward compatible implementations.
type UnimplementedVotingServer struct {
}

func (UnimplementedVotingServer) CreatePoll(context.Context, *NewPoll) (*Poll, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePoll not implemented")
}
func (UnimplementedVotingServer) AddPollAlbums(Voting_AddPollAlbumsServer) error {
	return status.Errorf(codes.Unimplemented, "method AddPollAlbums not implemented")
}
func (UnimplementedVotingServer) GetPollDetails(context.Context, *PollQuery) (*PollDetails, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPollDetails not implemented")
}
func (UnimplementedVotingServer) VoteAlbum(context.Context, *Vote) (*Vote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteAlbum not implemented")
}
func (UnimplementedVotingServer) GetAlbumVote(*AlbumVotesQuery, Voting_GetAlbumVoteServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAlbumVote not implemented")
}
func (UnimplementedVotingServer) mustEmbedUnimplementedVotingServer() {}

// UnsafeVotingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VotingServer will
// result in compilation errors.
type UnsafeVotingServer interface {
	mustEmbedUnimplementedVotingServer()
}

func RegisterVotingServer(s grpc.ServiceRegistrar, srv VotingServer) {
	s.RegisterService(&Voting_ServiceDesc, srv)
}

func _Voting_CreatePoll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewPoll)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServer).CreatePoll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Voting/CreatePoll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServer).CreatePoll(ctx, req.(*NewPoll))
	}
	return interceptor(ctx, in, info, handler)
}

func _Voting_AddPollAlbums_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(VotingServer).AddPollAlbums(&votingAddPollAlbumsServer{stream})
}

type Voting_AddPollAlbumsServer interface {
	Send(*NewPollAlbumResponse) error
	Recv() (*NewPollAlbum, error)
	grpc.ServerStream
}

type votingAddPollAlbumsServer struct {
	grpc.ServerStream
}

func (x *votingAddPollAlbumsServer) Send(m *NewPollAlbumResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *votingAddPollAlbumsServer) Recv() (*NewPollAlbum, error) {
	m := new(NewPollAlbum)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _Voting_GetPollDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PollQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServer).GetPollDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Voting/GetPollDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServer).GetPollDetails(ctx, req.(*PollQuery))
	}
	return interceptor(ctx, in, info, handler)
}

func _Voting_VoteAlbum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Vote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VotingServer).VoteAlbum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.Voting/VoteAlbum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VotingServer).VoteAlbum(ctx, req.(*Vote))
	}
	return interceptor(ctx, in, info, handler)
}

func _Voting_GetAlbumVote_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(AlbumVotesQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(VotingServer).GetAlbumVote(m, &votingGetAlbumVoteServer{stream})
}

type Voting_GetAlbumVoteServer interface {
	Send(*Vote) error
	grpc.ServerStream
}

type votingGetAlbumVoteServer struct {
	grpc.ServerStream
}

func (x *votingGetAlbumVoteServer) Send(m *Vote) error {
	return x.ServerStream.SendMsg(m)
}

// Voting_ServiceDesc is the grpc.ServiceDesc for Voting service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Voting_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.Voting",
	HandlerType: (*VotingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePoll",
			Handler:    _Voting_CreatePoll_Handler,
		},
		{
			MethodName: "GetPollDetails",
			Handler:    _Voting_GetPollDetails_Handler,
		},
		{
			MethodName: "VoteAlbum",
			Handler:    _Voting_VoteAlbum_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AddPollAlbums",
			Handler:       _Voting_AddPollAlbums_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "GetAlbumVote",
			Handler:       _Voting_GetAlbumVote_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "voting.proto",
}
