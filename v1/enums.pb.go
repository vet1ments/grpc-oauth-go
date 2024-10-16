// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        (unknown)
// source: grpcoauth/v1/enums.proto

package grpcoauthv1

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

type GrantType int32

const (
	GrantType_GRANT_TYPE_UNSPECIFIED        GrantType = 0 // 기본값wef
	GrantType_GRANT_TYPE_AUTHORIZATION_CODE GrantType = 1
	GrantType_GRANT_TYPE_REFRESH_TOKEN      GrantType = 2
)

// Enum value maps for GrantType.
var (
	GrantType_name = map[int32]string{
		0: "GRANT_TYPE_UNSPECIFIED",
		1: "GRANT_TYPE_AUTHORIZATION_CODE",
		2: "GRANT_TYPE_REFRESH_TOKEN",
	}
	GrantType_value = map[string]int32{
		"GRANT_TYPE_UNSPECIFIED":        0,
		"GRANT_TYPE_AUTHORIZATION_CODE": 1,
		"GRANT_TYPE_REFRESH_TOKEN":      2,
	}
)

func (x GrantType) Enum() *GrantType {
	p := new(GrantType)
	*p = x
	return p
}

func (x GrantType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GrantType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpcoauth_v1_enums_proto_enumTypes[0].Descriptor()
}

func (GrantType) Type() protoreflect.EnumType {
	return &file_grpcoauth_v1_enums_proto_enumTypes[0]
}

func (x GrantType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GrantType.Descriptor instead.
func (GrantType) EnumDescriptor() ([]byte, []int) {
	return file_grpcoauth_v1_enums_proto_rawDescGZIP(), []int{0}
}

type RoleType int32

const (
	RoleType_ROLE_TYPE_UNSPECIFIED RoleType = 0
	RoleType_ROLE_TYPE_ADMIN       RoleType = 1
	RoleType_ROLE_TYPE_USER        RoleType = 2
)

// Enum value maps for RoleType.
var (
	RoleType_name = map[int32]string{
		0: "ROLE_TYPE_UNSPECIFIED",
		1: "ROLE_TYPE_ADMIN",
		2: "ROLE_TYPE_USER",
	}
	RoleType_value = map[string]int32{
		"ROLE_TYPE_UNSPECIFIED": 0,
		"ROLE_TYPE_ADMIN":       1,
		"ROLE_TYPE_USER":        2,
	}
)

func (x RoleType) Enum() *RoleType {
	p := new(RoleType)
	*p = x
	return p
}

func (x RoleType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RoleType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpcoauth_v1_enums_proto_enumTypes[1].Descriptor()
}

func (RoleType) Type() protoreflect.EnumType {
	return &file_grpcoauth_v1_enums_proto_enumTypes[1]
}

func (x RoleType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RoleType.Descriptor instead.
func (RoleType) EnumDescriptor() ([]byte, []int) {
	return file_grpcoauth_v1_enums_proto_rawDescGZIP(), []int{1}
}

type ResponseType int32

const (
	ResponseType_RESPONSE_TYPE_UNSPECIFIED ResponseType = 0
	ResponseType_RESPONSE_TYPE_SUCCESS     ResponseType = 1
	ResponseType_RESPONSE_TYPE_FAIL        ResponseType = 2
)

// Enum value maps for ResponseType.
var (
	ResponseType_name = map[int32]string{
		0: "RESPONSE_TYPE_UNSPECIFIED",
		1: "RESPONSE_TYPE_SUCCESS",
		2: "RESPONSE_TYPE_FAIL",
	}
	ResponseType_value = map[string]int32{
		"RESPONSE_TYPE_UNSPECIFIED": 0,
		"RESPONSE_TYPE_SUCCESS":     1,
		"RESPONSE_TYPE_FAIL":        2,
	}
)

func (x ResponseType) Enum() *ResponseType {
	p := new(ResponseType)
	*p = x
	return p
}

func (x ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpcoauth_v1_enums_proto_enumTypes[2].Descriptor()
}

func (ResponseType) Type() protoreflect.EnumType {
	return &file_grpcoauth_v1_enums_proto_enumTypes[2]
}

func (x ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResponseType.Descriptor instead.
func (ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_grpcoauth_v1_enums_proto_rawDescGZIP(), []int{2}
}

type LoginType int32

const (
	LoginType_LOGIN_TYPE_UNSPECIFIED LoginType = 0
	LoginType_LOGIN_TYPE_NATIVE      LoginType = 1
	LoginType_LOGIN_TYPE_KAKAO       LoginType = 2
	LoginType_LOGIN_TYPE_NAVER       LoginType = 3
)

// Enum value maps for LoginType.
var (
	LoginType_name = map[int32]string{
		0: "LOGIN_TYPE_UNSPECIFIED",
		1: "LOGIN_TYPE_NATIVE",
		2: "LOGIN_TYPE_KAKAO",
		3: "LOGIN_TYPE_NAVER",
	}
	LoginType_value = map[string]int32{
		"LOGIN_TYPE_UNSPECIFIED": 0,
		"LOGIN_TYPE_NATIVE":      1,
		"LOGIN_TYPE_KAKAO":       2,
		"LOGIN_TYPE_NAVER":       3,
	}
)

func (x LoginType) Enum() *LoginType {
	p := new(LoginType)
	*p = x
	return p
}

func (x LoginType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LoginType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpcoauth_v1_enums_proto_enumTypes[3].Descriptor()
}

func (LoginType) Type() protoreflect.EnumType {
	return &file_grpcoauth_v1_enums_proto_enumTypes[3]
}

func (x LoginType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LoginType.Descriptor instead.
func (LoginType) EnumDescriptor() ([]byte, []int) {
	return file_grpcoauth_v1_enums_proto_rawDescGZIP(), []int{3}
}

type TokenType int32

const (
	TokenType_TOKEN_TYPE_UNSPECIFIED TokenType = 0
	TokenType_TOKEN_TYPE_BEARER      TokenType = 1
)

// Enum value maps for TokenType.
var (
	TokenType_name = map[int32]string{
		0: "TOKEN_TYPE_UNSPECIFIED",
		1: "TOKEN_TYPE_BEARER",
	}
	TokenType_value = map[string]int32{
		"TOKEN_TYPE_UNSPECIFIED": 0,
		"TOKEN_TYPE_BEARER":      1,
	}
)

func (x TokenType) Enum() *TokenType {
	p := new(TokenType)
	*p = x
	return p
}

func (x TokenType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TokenType) Descriptor() protoreflect.EnumDescriptor {
	return file_grpcoauth_v1_enums_proto_enumTypes[4].Descriptor()
}

func (TokenType) Type() protoreflect.EnumType {
	return &file_grpcoauth_v1_enums_proto_enumTypes[4]
}

func (x TokenType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TokenType.Descriptor instead.
func (TokenType) EnumDescriptor() ([]byte, []int) {
	return file_grpcoauth_v1_enums_proto_rawDescGZIP(), []int{4}
}

var File_grpcoauth_v1_enums_proto protoreflect.FileDescriptor

var file_grpcoauth_v1_enums_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x76, 0x31, 0x2f, 0x65,
	0x6e, 0x75, 0x6d, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x67, 0x72, 0x70, 0x63,
	0x6f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x76, 0x31, 0x2a, 0x68, 0x0a, 0x09, 0x47, 0x72, 0x61, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x16, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x21, 0x0a, 0x1d, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x41, 0x55, 0x54, 0x48, 0x4f, 0x52, 0x49, 0x5a, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x5f, 0x43, 0x4f,
	0x44, 0x45, 0x10, 0x01, 0x12, 0x1c, 0x0a, 0x18, 0x47, 0x52, 0x41, 0x4e, 0x54, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x52, 0x45, 0x46, 0x52, 0x45, 0x53, 0x48, 0x5f, 0x54, 0x4f, 0x4b, 0x45, 0x4e,
	0x10, 0x02, 0x2a, 0x4e, 0x0a, 0x08, 0x52, 0x6f, 0x6c, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x19,
	0x0a, 0x15, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x13, 0x0a, 0x0f, 0x52, 0x4f, 0x4c,
	0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x41, 0x44, 0x4d, 0x49, 0x4e, 0x10, 0x01, 0x12, 0x12,
	0x0a, 0x0e, 0x52, 0x4f, 0x4c, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x53, 0x45, 0x52,
	0x10, 0x02, 0x2a, 0x60, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x1d, 0x0a, 0x19, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54,
	0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10,
	0x00, 0x12, 0x19, 0x0a, 0x15, 0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12,
	0x52, 0x45, 0x53, 0x50, 0x4f, 0x4e, 0x53, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x10, 0x02, 0x2a, 0x6a, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x1a, 0x0a, 0x16, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f,
	0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a,
	0x11, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x54, 0x49,
	0x56, 0x45, 0x10, 0x01, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x4f, 0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59,
	0x50, 0x45, 0x5f, 0x4b, 0x41, 0x4b, 0x41, 0x4f, 0x10, 0x02, 0x12, 0x14, 0x0a, 0x10, 0x4c, 0x4f,
	0x47, 0x49, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x4e, 0x41, 0x56, 0x45, 0x52, 0x10, 0x03,
	0x2a, 0x3e, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a,
	0x16, 0x54, 0x4f, 0x4b, 0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x54, 0x4f, 0x4b,
	0x45, 0x4e, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x5f, 0x42, 0x45, 0x41, 0x52, 0x45, 0x52, 0x10, 0x01,
	0x42, 0xab, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x76, 0x31, 0x42, 0x0a, 0x45, 0x6e, 0x75, 0x6d, 0x73, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x76, 0x65, 0x74, 0x31, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61,
	0x75, 0x74, 0x68, 0x2f, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x76, 0x31, 0x3b, 0x67, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x76, 0x31, 0xa2,
	0x02, 0x03, 0x47, 0x58, 0x58, 0xaa, 0x02, 0x0c, 0x47, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x0c, 0x47, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68,
	0x5c, 0x56, 0x31, 0xe2, 0x02, 0x18, 0x47, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x5c,
	0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0d, 0x47, 0x72, 0x70, 0x63, 0x6f, 0x61, 0x75, 0x74, 0x68, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpcoauth_v1_enums_proto_rawDescOnce sync.Once
	file_grpcoauth_v1_enums_proto_rawDescData = file_grpcoauth_v1_enums_proto_rawDesc
)

func file_grpcoauth_v1_enums_proto_rawDescGZIP() []byte {
	file_grpcoauth_v1_enums_proto_rawDescOnce.Do(func() {
		file_grpcoauth_v1_enums_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpcoauth_v1_enums_proto_rawDescData)
	})
	return file_grpcoauth_v1_enums_proto_rawDescData
}

var file_grpcoauth_v1_enums_proto_enumTypes = make([]protoimpl.EnumInfo, 5)
var file_grpcoauth_v1_enums_proto_goTypes = []any{
	(GrantType)(0),    // 0: grpcoauth.v1.GrantType
	(RoleType)(0),     // 1: grpcoauth.v1.RoleType
	(ResponseType)(0), // 2: grpcoauth.v1.ResponseType
	(LoginType)(0),    // 3: grpcoauth.v1.LoginType
	(TokenType)(0),    // 4: grpcoauth.v1.TokenType
}
var file_grpcoauth_v1_enums_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_grpcoauth_v1_enums_proto_init() }
func file_grpcoauth_v1_enums_proto_init() {
	if File_grpcoauth_v1_enums_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_grpcoauth_v1_enums_proto_rawDesc,
			NumEnums:      5,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_grpcoauth_v1_enums_proto_goTypes,
		DependencyIndexes: file_grpcoauth_v1_enums_proto_depIdxs,
		EnumInfos:         file_grpcoauth_v1_enums_proto_enumTypes,
	}.Build()
	File_grpcoauth_v1_enums_proto = out.File
	file_grpcoauth_v1_enums_proto_rawDesc = nil
	file_grpcoauth_v1_enums_proto_goTypes = nil
	file_grpcoauth_v1_enums_proto_depIdxs = nil
}
