// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.25.1
// source: voting.proto

package grpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Album struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Artist string `protobuf:"bytes,3,opt,name=artist,proto3" json:"artist,omitempty"`
}

func (x *Album) Reset() {
	*x = Album{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Album) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Album) ProtoMessage() {}

func (x *Album) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Album.ProtoReflect.Descriptor instead.
func (*Album) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{0}
}

func (x *Album) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Album) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Album) GetArtist() string {
	if x != nil {
		return x.Artist
	}
	return ""
}

type Poll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Id        int64    `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
	Albums    []*Album `protobuf:"bytes,3,rep,name=albums,proto3" json:"albums,omitempty"`
	CreatorId int32    `protobuf:"varint,4,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *Poll) Reset() {
	*x = Poll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Poll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Poll) ProtoMessage() {}

func (x *Poll) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Poll.ProtoReflect.Descriptor instead.
func (*Poll) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{1}
}

func (x *Poll) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Poll) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Poll) GetAlbums() []*Album {
	if x != nil {
		return x.Albums
	}
	return nil
}

func (x *Poll) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

type Vote struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId  int32 `protobuf:"varint,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
	AlbumId int64 `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
	UserId  int32 `protobuf:"varint,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *Vote) Reset() {
	*x = Vote{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vote) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vote) ProtoMessage() {}

func (x *Vote) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vote.ProtoReflect.Descriptor instead.
func (*Vote) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{2}
}

func (x *Vote) GetPollId() int32 {
	if x != nil {
		return x.PollId
	}
	return 0
}

func (x *Vote) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

func (x *Vote) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

type PollQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId int32 `protobuf:"varint,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
}

func (x *PollQuery) Reset() {
	*x = PollQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PollQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PollQuery) ProtoMessage() {}

func (x *PollQuery) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PollQuery.ProtoReflect.Descriptor instead.
func (*PollQuery) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{3}
}

func (x *PollQuery) GetPollId() int32 {
	if x != nil {
		return x.PollId
	}
	return 0
}

type AlbumVotesQuery struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId  int32 `protobuf:"varint,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
	AlbumId int64 `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *AlbumVotesQuery) Reset() {
	*x = AlbumVotesQuery{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AlbumVotesQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AlbumVotesQuery) ProtoMessage() {}

func (x *AlbumVotesQuery) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AlbumVotesQuery.ProtoReflect.Descriptor instead.
func (*AlbumVotesQuery) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{4}
}

func (x *AlbumVotesQuery) GetPollId() int32 {
	if x != nil {
		return x.PollId
	}
	return 0
}

func (x *AlbumVotesQuery) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

type NewPoll struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	CreatorId int32  `protobuf:"varint,2,opt,name=creator_id,json=creatorId,proto3" json:"creator_id,omitempty"`
}

func (x *NewPoll) Reset() {
	*x = NewPoll{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPoll) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPoll) ProtoMessage() {}

func (x *NewPoll) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPoll.ProtoReflect.Descriptor instead.
func (*NewPoll) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{5}
}

func (x *NewPoll) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *NewPoll) GetCreatorId() int32 {
	if x != nil {
		return x.CreatorId
	}
	return 0
}

type NewPollAlbum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId  int64 `protobuf:"varint,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
	AlbumId int64 `protobuf:"varint,2,opt,name=album_id,json=albumId,proto3" json:"album_id,omitempty"`
}

func (x *NewPollAlbum) Reset() {
	*x = NewPollAlbum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPollAlbum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPollAlbum) ProtoMessage() {}

func (x *NewPollAlbum) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPollAlbum.ProtoReflect.Descriptor instead.
func (*NewPollAlbum) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{6}
}

func (x *NewPollAlbum) GetPollId() int64 {
	if x != nil {
		return x.PollId
	}
	return 0
}

func (x *NewPollAlbum) GetAlbumId() int64 {
	if x != nil {
		return x.AlbumId
	}
	return 0
}

type NewPollAlbumResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PollId int32  `protobuf:"varint,1,opt,name=poll_id,json=pollId,proto3" json:"poll_id,omitempty"`
	Album  *Album `protobuf:"bytes,2,opt,name=album,proto3" json:"album,omitempty"`
}

func (x *NewPollAlbumResponse) Reset() {
	*x = NewPollAlbumResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_voting_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewPollAlbumResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewPollAlbumResponse) ProtoMessage() {}

func (x *NewPollAlbumResponse) ProtoReflect() protoreflect.Message {
	mi := &file_voting_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewPollAlbumResponse.ProtoReflect.Descriptor instead.
func (*NewPollAlbumResponse) Descriptor() ([]byte, []int) {
	return file_voting_proto_rawDescGZIP(), []int{7}
}

func (x *NewPollAlbumResponse) GetPollId() int32 {
	if x != nil {
		return x.PollId
	}
	return 0
}

func (x *NewPollAlbumResponse) GetAlbum() *Album {
	if x != nil {
		return x.Album
	}
	return nil
}

var File_voting_proto protoreflect.FileDescriptor

var file_voting_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x76, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04,
	0x67, 0x72, 0x70, 0x63, 0x1a, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x43, 0x0a, 0x05, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a,
	0x06, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x22, 0x6e, 0x0a, 0x04, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x23, 0x0a, 0x06, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x06,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x53, 0x0a, 0x04, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x17, 0x0a,
	0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06,
	0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x24, 0x0a, 0x09, 0x50, 0x6f,
	0x6c, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64,
	0x22, 0x45, 0x0a, 0x0f, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x51, 0x75,
	0x65, 0x72, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x3c, 0x0a, 0x07, 0x4e, 0x65, 0x77, 0x50, 0x6f,
	0x6c, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x42, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x50, 0x6f, 0x6c, 0x6c,
	0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x49, 0x64, 0x22, 0x52, 0x0a, 0x14, 0x4e, 0x65, 0x77,
	0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x6f, 0x6c, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x70, 0x6f, 0x6c, 0x6c, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x61, 0x6c,
	0x62, 0x75, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x05, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x32, 0x89, 0x02,
	0x0a, 0x06, 0x56, 0x6f, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x29, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x6f, 0x6c, 0x6c, 0x12, 0x0d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65,
	0x77, 0x50, 0x6f, 0x6c, 0x6c, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x6c,
	0x6c, 0x22, 0x00, 0x12, 0x45, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6c,
	0x62, 0x75, 0x6d, 0x73, 0x12, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x4e, 0x65, 0x77, 0x50,
	0x6f, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x1a, 0x1a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x4e, 0x65, 0x77, 0x50, 0x6f, 0x6c, 0x6c, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x6c, 0x6c, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x0f, 0x2e, 0x67,
	0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0a, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x6f, 0x6c, 0x6c, 0x22, 0x00, 0x12, 0x25, 0x0a, 0x09, 0x56,
	0x6f, 0x74, 0x65, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x12, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x56, 0x6f, 0x74, 0x65, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x6f, 0x74, 0x65,
	0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x56, 0x6f,
	0x74, 0x65, 0x12, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x41, 0x6c, 0x62, 0x75, 0x6d, 0x56,
	0x6f, 0x74, 0x65, 0x73, 0x51, 0x75, 0x65, 0x72, 0x79, 0x1a, 0x0a, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x2e, 0x56, 0x6f, 0x74, 0x65, 0x22, 0x00, 0x30, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_voting_proto_rawDescOnce sync.Once
	file_voting_proto_rawDescData = file_voting_proto_rawDesc
)

func file_voting_proto_rawDescGZIP() []byte {
	file_voting_proto_rawDescOnce.Do(func() {
		file_voting_proto_rawDescData = protoimpl.X.CompressGZIP(file_voting_proto_rawDescData)
	})
	return file_voting_proto_rawDescData
}

var file_voting_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_voting_proto_goTypes = []interface{}{
	(*Album)(nil),                // 0: grpc.Album
	(*Poll)(nil),                 // 1: grpc.Poll
	(*Vote)(nil),                 // 2: grpc.Vote
	(*PollQuery)(nil),            // 3: grpc.PollQuery
	(*AlbumVotesQuery)(nil),      // 4: grpc.AlbumVotesQuery
	(*NewPoll)(nil),              // 5: grpc.NewPoll
	(*NewPollAlbum)(nil),         // 6: grpc.NewPollAlbum
	(*NewPollAlbumResponse)(nil), // 7: grpc.NewPollAlbumResponse
}
var file_voting_proto_depIdxs = []int32{
	0, // 0: grpc.Poll.albums:type_name -> grpc.Album
	0, // 1: grpc.NewPollAlbumResponse.album:type_name -> grpc.Album
	5, // 2: grpc.Voting.CreatePoll:input_type -> grpc.NewPoll
	6, // 3: grpc.Voting.AddPollAlbums:input_type -> grpc.NewPollAlbum
	3, // 4: grpc.Voting.GetPollDetails:input_type -> grpc.PollQuery
	2, // 5: grpc.Voting.VoteAlbum:input_type -> grpc.Vote
	4, // 6: grpc.Voting.GetAlbumVote:input_type -> grpc.AlbumVotesQuery
	1, // 7: grpc.Voting.CreatePoll:output_type -> grpc.Poll
	7, // 8: grpc.Voting.AddPollAlbums:output_type -> grpc.NewPollAlbumResponse
	1, // 9: grpc.Voting.GetPollDetails:output_type -> grpc.Poll
	2, // 10: grpc.Voting.VoteAlbum:output_type -> grpc.Vote
	2, // 11: grpc.Voting.GetAlbumVote:output_type -> grpc.Vote
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_voting_proto_init() }
func file_voting_proto_init() {
	if File_voting_proto != nil {
		return
	}
	file_user_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_voting_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Album); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Poll); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vote); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PollQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AlbumVotesQuery); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPoll); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPollAlbum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_voting_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*NewPollAlbumResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_voting_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_voting_proto_goTypes,
		DependencyIndexes: file_voting_proto_depIdxs,
		MessageInfos:      file_voting_proto_msgTypes,
	}.Build()
	File_voting_proto = out.File
	file_voting_proto_rawDesc = nil
	file_voting_proto_goTypes = nil
	file_voting_proto_depIdxs = nil
}
