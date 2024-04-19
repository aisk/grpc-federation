// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        (unknown)
// source: federation/federation.proto

package federation

import (
	_ "example/comment"
	_ "example/favorite"
	_ "github.com/mercari/grpc-federation/grpc/federation"
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

type MyFavoriteType int32

const (
	MyFavoriteType_UNKNOWN MyFavoriteType = 0
	MyFavoriteType_TYPE1   MyFavoriteType = 5000
)

// Enum value maps for MyFavoriteType.
var (
	MyFavoriteType_name = map[int32]string{
		0:    "UNKNOWN",
		5000: "TYPE1",
	}
	MyFavoriteType_value = map[string]int32{
		"UNKNOWN": 0,
		"TYPE1":   5000,
	}
)

func (x MyFavoriteType) Enum() *MyFavoriteType {
	p := new(MyFavoriteType)
	*p = x
	return p
}

func (x MyFavoriteType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MyFavoriteType) Descriptor() protoreflect.EnumDescriptor {
	return file_federation_federation_proto_enumTypes[0].Descriptor()
}

func (MyFavoriteType) Type() protoreflect.EnumType {
	return &file_federation_federation_proto_enumTypes[0]
}

func (x MyFavoriteType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MyFavoriteType.Descriptor instead.
func (MyFavoriteType) EnumDescriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{0}
}

type GetPostRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPostRequest) Reset() {
	*x = GetPostRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostRequest) ProtoMessage() {}

func (x *GetPostRequest) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostRequest.ProtoReflect.Descriptor instead.
func (*GetPostRequest) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{0}
}

func (x *GetPostRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetPostResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *Post `protobuf:"bytes,1,opt,name=post,proto3" json:"post,omitempty"`
}

func (x *GetPostResponse) Reset() {
	*x = GetPostResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPostResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPostResponse) ProtoMessage() {}

func (x *GetPostResponse) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPostResponse.ProtoReflect.Descriptor instead.
func (*GetPostResponse) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{1}
}

func (x *GetPostResponse) GetPost() *Post {
	if x != nil {
		return x.Post
	}
	return nil
}

type Post struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string         `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title         string         `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	Content       string         `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	User          *User          `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	Reaction      *Reaction      `protobuf:"bytes,5,opt,name=reaction,proto3" json:"reaction,omitempty"`
	FavoriteValue MyFavoriteType `protobuf:"varint,6,opt,name=favorite_value,json=favoriteValue,proto3,enum=federation.MyFavoriteType" json:"favorite_value,omitempty"`
	Cmp           bool           `protobuf:"varint,7,opt,name=cmp,proto3" json:"cmp,omitempty"`
}

func (x *Post) Reset() {
	*x = Post{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Post) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Post) ProtoMessage() {}

func (x *Post) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Post.ProtoReflect.Descriptor instead.
func (*Post) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{2}
}

func (x *Post) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Post) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Post) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Post) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Post) GetReaction() *Reaction {
	if x != nil {
		return x.Reaction
	}
	return nil
}

func (x *Post) GetFavoriteValue() MyFavoriteType {
	if x != nil {
		return x.FavoriteValue
	}
	return MyFavoriteType_UNKNOWN
}

func (x *Post) GetCmp() bool {
	if x != nil {
		return x.Cmp
	}
	return false
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{3}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type GetStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetStatusRequest) Reset() {
	*x = GetStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusRequest) ProtoMessage() {}

func (x *GetStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusRequest.ProtoReflect.Descriptor instead.
func (*GetStatusRequest) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{4}
}

type GetStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetStatusResponse) Reset() {
	*x = GetStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_federation_federation_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetStatusResponse) ProtoMessage() {}

func (x *GetStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_federation_federation_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetStatusResponse.ProtoReflect.Descriptor instead.
func (*GetStatusResponse) Descriptor() ([]byte, []int) {
	return file_federation_federation_proto_rawDescGZIP(), []int{5}
}

func (x *GetStatusResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_federation_federation_proto protoreflect.FileDescriptor

var file_federation_federation_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64,
	0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x19, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x72,
	0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x66,
	0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x51, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50,
	0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x64, 0x65,
	0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x42, 0x06, 0x9a, 0x4a, 0x03,
	0x12, 0x01, 0x70, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x3a, 0x10, 0x9a, 0x4a, 0x0d, 0x0a, 0x0b,
	0x0a, 0x01, 0x70, 0x6a, 0x06, 0x0a, 0x04, 0x50, 0x6f, 0x73, 0x74, 0x22, 0x9d, 0x04, 0x0a, 0x04,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0d, 0x9a, 0x4a, 0x0a, 0xf2, 0x01, 0x07, 0x70, 0x6f, 0x73, 0x74, 0x2d, 0x69, 0x64, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x21, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x42, 0x0b, 0x9a, 0x4a, 0x08, 0xf2, 0x01, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x52,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x27, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x9a, 0x4a, 0x0a, 0xf2, 0x01, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x2c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42,
	0x06, 0x9a, 0x4a, 0x03, 0x12, 0x01, 0x75, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x3f, 0x0a,
	0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0d, 0x9a, 0x4a, 0x0a, 0x12, 0x08, 0x72, 0x65, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x56,
	0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1a, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x79, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x42, 0x13, 0x9a, 0x4a, 0x10, 0x12, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x03, 0x63, 0x6d, 0x70, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x08, 0x42, 0x08, 0x9a, 0x4a, 0x05, 0x12, 0x03, 0x63, 0x6d, 0x70, 0x52, 0x03, 0x63,
	0x6d, 0x70, 0x3a, 0xc6, 0x01, 0x9a, 0x4a, 0xc2, 0x01, 0x0a, 0x25, 0x0a, 0x01, 0x75, 0x6a, 0x20,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0a, 0x0a, 0x02, 0x69, 0x64, 0xf2, 0x01, 0x03, 0x66,
	0x6f, 0x6f, 0x12, 0x0c, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xf2, 0x01, 0x03, 0x62, 0x61, 0x72,
	0x0a, 0x36, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x5a, 0x24, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x28,
	0x27, 0x54, 0x59, 0x50, 0x45, 0x31, 0x27, 0x29, 0x0a, 0x34, 0x0a, 0x03, 0x63, 0x6d, 0x70, 0x5a,
	0x2d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x20,
	0x3d, 0x3d, 0x20, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x2e, 0x54, 0x59, 0x50, 0x45, 0x31, 0x0a, 0x2b,
	0x0a, 0x08, 0x72, 0x65, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x6a, 0x1f, 0x0a, 0x08, 0x52, 0x65,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x13, 0x0a, 0x01, 0x76, 0x12, 0x0e, 0x66, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x42, 0x0a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x12, 0x19, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x09, 0x9a, 0x4a, 0x06, 0x12, 0x04, 0x24, 0x2e, 0x69, 0x64, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1f,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0b, 0x9a, 0x4a,
	0x08, 0x12, 0x06, 0x24, 0x2e, 0x6e, 0x61, 0x6d, 0x65, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22,
	0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x6f, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x42, 0x06, 0x9a, 0x4a, 0x03, 0x12, 0x01, 0x75,
	0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x3a, 0x2c, 0x9a, 0x4a, 0x29, 0x0a, 0x27, 0x0a, 0x01, 0x75,
	0x6a, 0x22, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x0a, 0x02, 0x69, 0x64, 0xf2, 0x01,
	0x04, 0x78, 0x78, 0x78, 0x78, 0x12, 0x0d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xf2, 0x01, 0x04,
	0x79, 0x79, 0x79, 0x79, 0x2a, 0x5f, 0x0a, 0x0e, 0x4d, 0x79, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x1a, 0x0c, 0x9a, 0x4a, 0x09, 0x12, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x12, 0x16, 0x0a, 0x05, 0x54, 0x59, 0x50, 0x45, 0x31, 0x10, 0x88, 0x27, 0x1a, 0x0a, 0x9a,
	0x4a, 0x07, 0x12, 0x05, 0x54, 0x59, 0x50, 0x45, 0x31, 0x1a, 0x1a, 0x9a, 0x4a, 0x17, 0x0a, 0x15,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x2e, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x32, 0x5e, 0x0a, 0x11, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x44, 0x0a, 0x07, 0x47, 0x65,
	0x74, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1a, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x1a, 0x03, 0x9a, 0x4a, 0x00, 0x32, 0x5f, 0x0a, 0x0c, 0x44, 0x65, 0x62, 0x75, 0x67, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x1c, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x47, 0x65, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x1a, 0x03, 0x9a, 0x4a, 0x00, 0x42, 0x88, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x66,
	0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0f, 0x46, 0x65, 0x64, 0x65, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1d, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x3b, 0x66, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xa2, 0x02, 0x03, 0x46, 0x58,
	0x58, 0xaa, 0x02, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xca, 0x02,
	0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x16, 0x46, 0x65,
	0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0a, 0x46, 0x65, 0x64, 0x65, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_federation_federation_proto_rawDescOnce sync.Once
	file_federation_federation_proto_rawDescData = file_federation_federation_proto_rawDesc
)

func file_federation_federation_proto_rawDescGZIP() []byte {
	file_federation_federation_proto_rawDescOnce.Do(func() {
		file_federation_federation_proto_rawDescData = protoimpl.X.CompressGZIP(file_federation_federation_proto_rawDescData)
	})
	return file_federation_federation_proto_rawDescData
}

var file_federation_federation_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_federation_federation_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_federation_federation_proto_goTypes = []interface{}{
	(MyFavoriteType)(0),       // 0: federation.MyFavoriteType
	(*GetPostRequest)(nil),    // 1: federation.GetPostRequest
	(*GetPostResponse)(nil),   // 2: federation.GetPostResponse
	(*Post)(nil),              // 3: federation.Post
	(*User)(nil),              // 4: federation.User
	(*GetStatusRequest)(nil),  // 5: federation.GetStatusRequest
	(*GetStatusResponse)(nil), // 6: federation.GetStatusResponse
	(*Reaction)(nil),          // 7: federation.Reaction
}
var file_federation_federation_proto_depIdxs = []int32{
	3, // 0: federation.GetPostResponse.post:type_name -> federation.Post
	4, // 1: federation.Post.user:type_name -> federation.User
	7, // 2: federation.Post.reaction:type_name -> federation.Reaction
	0, // 3: federation.Post.favorite_value:type_name -> federation.MyFavoriteType
	4, // 4: federation.GetStatusResponse.user:type_name -> federation.User
	1, // 5: federation.FederationService.GetPost:input_type -> federation.GetPostRequest
	5, // 6: federation.DebugService.GetStatus:input_type -> federation.GetStatusRequest
	2, // 7: federation.FederationService.GetPost:output_type -> federation.GetPostResponse
	6, // 8: federation.DebugService.GetStatus:output_type -> federation.GetStatusResponse
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_federation_federation_proto_init() }
func file_federation_federation_proto_init() {
	if File_federation_federation_proto != nil {
		return
	}
	file_federation_reaction_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_federation_federation_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostRequest); i {
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
		file_federation_federation_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetPostResponse); i {
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
		file_federation_federation_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Post); i {
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
		file_federation_federation_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_federation_federation_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusRequest); i {
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
		file_federation_federation_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetStatusResponse); i {
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
			RawDescriptor: file_federation_federation_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_federation_federation_proto_goTypes,
		DependencyIndexes: file_federation_federation_proto_depIdxs,
		EnumInfos:         file_federation_federation_proto_enumTypes,
		MessageInfos:      file_federation_federation_proto_msgTypes,
	}.Build()
	File_federation_federation_proto = out.File
	file_federation_federation_proto_rawDesc = nil
	file_federation_federation_proto_goTypes = nil
	file_federation_federation_proto_depIdxs = nil
}
