// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: grpcoauth/v1/user.proto

package grpcoauthv1

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	PhoneNumber string `protobuf:"bytes,3,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	Nickname    string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Thumbnail   string `protobuf:"bytes,5,opt,name=thumbnail,proto3" json:"thumbnail,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_grpcoauth_v1_user_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpcoauth_v1_user_proto_msgTypes[0]
	if x != nil {
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
	return file_grpcoauth_v1_user_proto_rawDescGZIP(), []int{0}
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

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *User) GetThumbnail() string {
	if x != nil {
		return x.Thumbnail
	}
	return ""
}

type GetAccessTokenInfoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	ExpiresIn int64  `protobuf:"varint,2,opt,name=expires_in,json=expiresIn,proto3" json:"expires_in,omitempty"`
	AppId     string `protobuf:"bytes,3,opt,name=app_id,json=appId,proto3" json:"app_id,omitempty"`
}

func (x *GetAccessTokenInfoResponse) Reset() {
	*x = GetAccessTokenInfoResponse{}
	mi := &file_grpcoauth_v1_user_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccessTokenInfoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessTokenInfoResponse) ProtoMessage() {}

func (x *GetAccessTokenInfoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcoauth_v1_user_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessTokenInfoResponse.ProtoReflect.Descriptor instead.
func (*GetAccessTokenInfoResponse) Descriptor() ([]byte, []int) {
	return file_grpcoauth_v1_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessTokenInfoResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetAccessTokenInfoResponse) GetExpiresIn() int64 {
	if x != nil {
		return x.ExpiresIn
	}
	return 0
}

func (x *GetAccessTokenInfoResponse) GetAppId() string {
	if x != nil {
		return x.AppId
	}
	return ""
}

type GetUserMeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserMeResponse) Reset() {
	*x = GetUserMeResponse{}
	mi := &file_grpcoauth_v1_user_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserMeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserMeResponse) ProtoMessage() {}

func (x *GetUserMeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcoauth_v1_user_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserMeResponse.ProtoReflect.Descriptor instead.
func (*GetUserMeResponse) Descriptor() ([]byte, []int) {
	return file_grpcoauth_v1_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserMeResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

type UserLogoutResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserLogoutResponse) Reset() {
	*x = UserLogoutResponse{}
	mi := &file_grpcoauth_v1_user_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserLogoutResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLogoutResponse) ProtoMessage() {}

func (x *UserLogoutResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcoauth_v1_user_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLogoutResponse.ProtoReflect.Descriptor instead.
func (*UserLogoutResponse) Descriptor() ([]byte, []int) {
	return file_grpcoauth_v1_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserLogoutResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetUserListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetUserListResponse) Reset() {
	*x = GetUserListResponse{}
	mi := &file_grpcoauth_v1_user_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListResponse) ProtoMessage() {}

func (x *GetUserListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpcoauth_v1_user_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListResponse.ProtoReflect.Descriptor instead.
func (*GetUserListResponse) Descriptor() ([]byte, []int) {
	return file_grpcoauth_v1_user_proto_rawDescGZIP(), []int{4}
}

func (x *GetUserListResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []string `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"` // awe
}

func (x *GetUserListRequest) Reset() {
	*x = GetUserListRequest{}
	mi := &file_grpcoauth_v1_user_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUserListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserListRequest) ProtoMessage() {}

func (x *GetUserListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpcoauth_v1_user_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserListRequest.ProtoReflect.Descriptor instead.
func (*GetUserListRequest) Descriptor() ([]byte, []int) {
	return file_grpcoauth_v1_user_proto_rawDescGZIP(), []int{5}
}

func (x *GetUserListRequest) GetUsers() []string {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_grpcoauth_v1_user_proto protoreflect.FileDescriptor

var file_grpcoauth_v1_user_proto_rawDesc = []byte{
	0x0a, 0x17, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63, 0x6f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0xd0, 0x01, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x15, 0xba, 0x48, 0x12, 0x72, 0x10, 0x32, 0x0e, 0x5e, 0x5b, 0x61,
	0x2d, 0x7a, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x32, 0x34, 0x7d, 0x24, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1d, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x09, 0xba,
	0x48, 0x06, 0x72, 0x04, 0x10, 0x02, 0x18, 0x14, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3f,
	0x0a, 0x0c, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xba, 0x48, 0x19, 0x72, 0x17, 0x32, 0x15, 0x5e, 0x30, 0x31,
	0x30, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34, 0x7d, 0x5b, 0x30, 0x2d, 0x39, 0x5d, 0x7b, 0x34,
	0x7d, 0x24, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xba, 0x48, 0x04, 0x72, 0x02, 0x18, 0x14, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61, 0x69,
	0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x68, 0x75, 0x6d, 0x62, 0x6e, 0x61,
	0x69, 0x6c, 0x22, 0x62, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x5f, 0x69, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73, 0x49, 0x6e, 0x12,
	0x15, 0x0a, 0x06, 0x61, 0x70, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x61, 0x70, 0x70, 0x49, 0x64, 0x22, 0x3b, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x4d, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x26, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x72, 0x70, 0x63,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x28, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x2a, 0x0a, 0x12, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x14, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x32, 0xa0, 0x03, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x77, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x28, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1f,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x19, 0x12, 0x17, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x12,
	0x56, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x10, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0a, 0x12, 0x08, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x2f, 0x6d, 0x65, 0x12, 0x5c, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x4c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x20, 0x2e,
	0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x12, 0x0c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x6c,
	0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x62, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x0e, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x08, 0x12, 0x06, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x42, 0xaa, 0x01, 0x0a, 0x10, 0x63, 0x6f,
	0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x09,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x76, 0x65, 0x74, 0x31, 0x6d, 0x65, 0x6e, 0x74,
	0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x67, 0x6f, 0x2f, 0x67,
	0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x70, 0x63,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0c,
	0x47, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x47,
	0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x47, 0x72,
	0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x47, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcoauth_v1_user_proto_rawDescOnce sync.Once
	file_grpcoauth_v1_user_proto_rawDescData = file_grpcoauth_v1_user_proto_rawDesc
)

func file_grpcoauth_v1_user_proto_rawDescGZIP() []byte {
	file_grpcoauth_v1_user_proto_rawDescOnce.Do(func() {
		file_grpcoauth_v1_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcoauth_v1_user_proto_rawDescData)
	})
	return file_grpcoauth_v1_user_proto_rawDescData
}

var file_grpcoauth_v1_user_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_grpcoauth_v1_user_proto_goTypes = []any{
	(*User)(nil),                       // 0: grpcoauth.v1.User
	(*GetAccessTokenInfoResponse)(nil), // 1: grpcoauth.v1.GetAccessTokenInfoResponse
	(*GetUserMeResponse)(nil),          // 2: grpcoauth.v1.GetUserMeResponse
	(*UserLogoutResponse)(nil),         // 3: grpcoauth.v1.UserLogoutResponse
	(*GetUserListResponse)(nil),        // 4: grpcoauth.v1.GetUserListResponse
	(*GetUserListRequest)(nil),         // 5: grpcoauth.v1.GetUserListRequest
	(*emptypb.Empty)(nil),              // 6: google.protobuf.Empty
}
var file_grpcoauth_v1_user_proto_depIdxs = []int32{
	0, // 0: grpcoauth.v1.GetUserMeResponse.user:type_name -> grpcoauth.v1.User
	0, // 1: grpcoauth.v1.GetUserListResponse.users:type_name -> grpcoauth.v1.User
	6, // 2: grpcoauth.v1.UserService.GetAccessTokenInfo:input_type -> google.protobuf.Empty
	6, // 3: grpcoauth.v1.UserService.GetUserMe:input_type -> google.protobuf.Empty
	6, // 4: grpcoauth.v1.UserService.UserLogout:input_type -> google.protobuf.Empty
	5, // 5: grpcoauth.v1.UserService.GetUserList:input_type -> grpcoauth.v1.GetUserListRequest
	1, // 6: grpcoauth.v1.UserService.GetAccessTokenInfo:output_type -> grpcoauth.v1.GetAccessTokenInfoResponse
	2, // 7: grpcoauth.v1.UserService.GetUserMe:output_type -> grpcoauth.v1.GetUserMeResponse
	3, // 8: grpcoauth.v1.UserService.UserLogout:output_type -> grpcoauth.v1.UserLogoutResponse
	4, // 9: grpcoauth.v1.UserService.GetUserList:output_type -> grpcoauth.v1.GetUserListResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_grpcoauth_v1_user_proto_init() }
func file_grpcoauth_v1_user_proto_init() {
	if File_grpcoauth_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcoauth_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpcoauth_v1_user_proto_goTypes,
		DependencyIndexes: file_grpcoauth_v1_user_proto_depIdxs,
		MessageInfos:      file_grpcoauth_v1_user_proto_msgTypes,
	}.Build()
	File_grpcoauth_v1_user_proto = out.File
	file_grpcoauth_v1_user_proto_rawDesc = nil
	file_grpcoauth_v1_user_proto_goTypes = nil
	file_grpcoauth_v1_user_proto_depIdxs = nil
}
