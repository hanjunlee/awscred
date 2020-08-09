// Code generated by protoc-gen-go. DO NOT EDIT.
// source: awscred.proto

package awscred

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

type SetOnRequest struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOnRequest) Reset()         { *m = SetOnRequest{} }
func (m *SetOnRequest) String() string { return proto.CompactTextString(m) }
func (*SetOnRequest) ProtoMessage()    {}
func (*SetOnRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{0}
}

func (m *SetOnRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOnRequest.Unmarshal(m, b)
}
func (m *SetOnRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOnRequest.Marshal(b, m, deterministic)
}
func (m *SetOnRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOnRequest.Merge(m, src)
}
func (m *SetOnRequest) XXX_Size() int {
	return xxx_messageInfo_SetOnRequest.Size(m)
}
func (m *SetOnRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOnRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetOnRequest proto.InternalMessageInfo

func (m *SetOnRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type SetOnResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOnResponse) Reset()         { *m = SetOnResponse{} }
func (m *SetOnResponse) String() string { return proto.CompactTextString(m) }
func (*SetOnResponse) ProtoMessage()    {}
func (*SetOnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{1}
}

func (m *SetOnResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOnResponse.Unmarshal(m, b)
}
func (m *SetOnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOnResponse.Marshal(b, m, deterministic)
}
func (m *SetOnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOnResponse.Merge(m, src)
}
func (m *SetOnResponse) XXX_Size() int {
	return xxx_messageInfo_SetOnResponse.Size(m)
}
func (m *SetOnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetOnResponse proto.InternalMessageInfo

type SetOffRequest struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOffRequest) Reset()         { *m = SetOffRequest{} }
func (m *SetOffRequest) String() string { return proto.CompactTextString(m) }
func (*SetOffRequest) ProtoMessage()    {}
func (*SetOffRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{2}
}

func (m *SetOffRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOffRequest.Unmarshal(m, b)
}
func (m *SetOffRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOffRequest.Marshal(b, m, deterministic)
}
func (m *SetOffRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOffRequest.Merge(m, src)
}
func (m *SetOffRequest) XXX_Size() int {
	return xxx_messageInfo_SetOffRequest.Size(m)
}
func (m *SetOffRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOffRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetOffRequest proto.InternalMessageInfo

func (m *SetOffRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

type SetOffResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetOffResponse) Reset()         { *m = SetOffResponse{} }
func (m *SetOffResponse) String() string { return proto.CompactTextString(m) }
func (*SetOffResponse) ProtoMessage()    {}
func (*SetOffResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{3}
}

func (m *SetOffResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetOffResponse.Unmarshal(m, b)
}
func (m *SetOffResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetOffResponse.Marshal(b, m, deterministic)
}
func (m *SetOffResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetOffResponse.Merge(m, src)
}
func (m *SetOffResponse) XXX_Size() int {
	return xxx_messageInfo_SetOffResponse.Size(m)
}
func (m *SetOffResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetOffResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetOffResponse proto.InternalMessageInfo

type SetConfigRequest struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Serial               string   `protobuf:"bytes,2,opt,name=serial,proto3" json:"serial,omitempty"`
	Duration             int64    `protobuf:"varint,3,opt,name=duration,proto3" json:"duration,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigRequest) Reset()         { *m = SetConfigRequest{} }
func (m *SetConfigRequest) String() string { return proto.CompactTextString(m) }
func (*SetConfigRequest) ProtoMessage()    {}
func (*SetConfigRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{4}
}

func (m *SetConfigRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigRequest.Unmarshal(m, b)
}
func (m *SetConfigRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigRequest.Marshal(b, m, deterministic)
}
func (m *SetConfigRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigRequest.Merge(m, src)
}
func (m *SetConfigRequest) XXX_Size() int {
	return xxx_messageInfo_SetConfigRequest.Size(m)
}
func (m *SetConfigRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigRequest proto.InternalMessageInfo

func (m *SetConfigRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *SetConfigRequest) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *SetConfigRequest) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

type SetConfigResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetConfigResponse) Reset()         { *m = SetConfigResponse{} }
func (m *SetConfigResponse) String() string { return proto.CompactTextString(m) }
func (*SetConfigResponse) ProtoMessage()    {}
func (*SetConfigResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{5}
}

func (m *SetConfigResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetConfigResponse.Unmarshal(m, b)
}
func (m *SetConfigResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetConfigResponse.Marshal(b, m, deterministic)
}
func (m *SetConfigResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetConfigResponse.Merge(m, src)
}
func (m *SetConfigResponse) XXX_Size() int {
	return xxx_messageInfo_SetConfigResponse.Size(m)
}
func (m *SetConfigResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetConfigResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetConfigResponse proto.InternalMessageInfo

type SetGenerateRequest struct {
	Profile              string   `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGenerateRequest) Reset()         { *m = SetGenerateRequest{} }
func (m *SetGenerateRequest) String() string { return proto.CompactTextString(m) }
func (*SetGenerateRequest) ProtoMessage()    {}
func (*SetGenerateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{6}
}

func (m *SetGenerateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGenerateRequest.Unmarshal(m, b)
}
func (m *SetGenerateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGenerateRequest.Marshal(b, m, deterministic)
}
func (m *SetGenerateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGenerateRequest.Merge(m, src)
}
func (m *SetGenerateRequest) XXX_Size() int {
	return xxx_messageInfo_SetGenerateRequest.Size(m)
}
func (m *SetGenerateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGenerateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SetGenerateRequest proto.InternalMessageInfo

func (m *SetGenerateRequest) GetProfile() string {
	if m != nil {
		return m.Profile
	}
	return ""
}

func (m *SetGenerateRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type SetGenerateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SetGenerateResponse) Reset()         { *m = SetGenerateResponse{} }
func (m *SetGenerateResponse) String() string { return proto.CompactTextString(m) }
func (*SetGenerateResponse) ProtoMessage()    {}
func (*SetGenerateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{7}
}

func (m *SetGenerateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SetGenerateResponse.Unmarshal(m, b)
}
func (m *SetGenerateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SetGenerateResponse.Marshal(b, m, deterministic)
}
func (m *SetGenerateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SetGenerateResponse.Merge(m, src)
}
func (m *SetGenerateResponse) XXX_Size() int {
	return xxx_messageInfo_SetGenerateResponse.Size(m)
}
func (m *SetGenerateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_SetGenerateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_SetGenerateResponse proto.InternalMessageInfo

type GetProfileListRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetProfileListRequest) Reset()         { *m = GetProfileListRequest{} }
func (m *GetProfileListRequest) String() string { return proto.CompactTextString(m) }
func (*GetProfileListRequest) ProtoMessage()    {}
func (*GetProfileListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{8}
}

func (m *GetProfileListRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileListRequest.Unmarshal(m, b)
}
func (m *GetProfileListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileListRequest.Marshal(b, m, deterministic)
}
func (m *GetProfileListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileListRequest.Merge(m, src)
}
func (m *GetProfileListRequest) XXX_Size() int {
	return xxx_messageInfo_GetProfileListRequest.Size(m)
}
func (m *GetProfileListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileListRequest proto.InternalMessageInfo

type GetProfileListResponse struct {
	Profiles             []*Profile `protobuf:"bytes,1,rep,name=profiles,proto3" json:"profiles,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetProfileListResponse) Reset()         { *m = GetProfileListResponse{} }
func (m *GetProfileListResponse) String() string { return proto.CompactTextString(m) }
func (*GetProfileListResponse) ProtoMessage()    {}
func (*GetProfileListResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{9}
}

func (m *GetProfileListResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetProfileListResponse.Unmarshal(m, b)
}
func (m *GetProfileListResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetProfileListResponse.Marshal(b, m, deterministic)
}
func (m *GetProfileListResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetProfileListResponse.Merge(m, src)
}
func (m *GetProfileListResponse) XXX_Size() int {
	return xxx_messageInfo_GetProfileListResponse.Size(m)
}
func (m *GetProfileListResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetProfileListResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetProfileListResponse proto.InternalMessageInfo

func (m *GetProfileListResponse) GetProfiles() []*Profile {
	if m != nil {
		return m.Profiles
	}
	return nil
}

type Profile struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	On                   bool     `protobuf:"varint,2,opt,name=on,proto3" json:"on,omitempty"`
	Serial               string   `protobuf:"bytes,3,opt,name=serial,proto3" json:"serial,omitempty"`
	Duration             int64    `protobuf:"varint,4,opt,name=duration,proto3" json:"duration,omitempty"`
	Expired              string   `protobuf:"bytes,5,opt,name=expired,proto3" json:"expired,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Profile) Reset()         { *m = Profile{} }
func (m *Profile) String() string { return proto.CompactTextString(m) }
func (*Profile) ProtoMessage()    {}
func (*Profile) Descriptor() ([]byte, []int) {
	return fileDescriptor_94e10cbc5c16adca, []int{10}
}

func (m *Profile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Profile.Unmarshal(m, b)
}
func (m *Profile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Profile.Marshal(b, m, deterministic)
}
func (m *Profile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Profile.Merge(m, src)
}
func (m *Profile) XXX_Size() int {
	return xxx_messageInfo_Profile.Size(m)
}
func (m *Profile) XXX_DiscardUnknown() {
	xxx_messageInfo_Profile.DiscardUnknown(m)
}

var xxx_messageInfo_Profile proto.InternalMessageInfo

func (m *Profile) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Profile) GetOn() bool {
	if m != nil {
		return m.On
	}
	return false
}

func (m *Profile) GetSerial() string {
	if m != nil {
		return m.Serial
	}
	return ""
}

func (m *Profile) GetDuration() int64 {
	if m != nil {
		return m.Duration
	}
	return 0
}

func (m *Profile) GetExpired() string {
	if m != nil {
		return m.Expired
	}
	return ""
}

func init() {
	proto.RegisterType((*SetOnRequest)(nil), "awscred.SetOnRequest")
	proto.RegisterType((*SetOnResponse)(nil), "awscred.SetOnResponse")
	proto.RegisterType((*SetOffRequest)(nil), "awscred.SetOffRequest")
	proto.RegisterType((*SetOffResponse)(nil), "awscred.SetOffResponse")
	proto.RegisterType((*SetConfigRequest)(nil), "awscred.SetConfigRequest")
	proto.RegisterType((*SetConfigResponse)(nil), "awscred.SetConfigResponse")
	proto.RegisterType((*SetGenerateRequest)(nil), "awscred.SetGenerateRequest")
	proto.RegisterType((*SetGenerateResponse)(nil), "awscred.SetGenerateResponse")
	proto.RegisterType((*GetProfileListRequest)(nil), "awscred.GetProfileListRequest")
	proto.RegisterType((*GetProfileListResponse)(nil), "awscred.GetProfileListResponse")
	proto.RegisterType((*Profile)(nil), "awscred.Profile")
}

func init() { proto.RegisterFile("awscred.proto", fileDescriptor_94e10cbc5c16adca) }

var fileDescriptor_94e10cbc5c16adca = []byte{
	// 391 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0x41, 0x4f, 0xe2, 0x50,
	0x10, 0x4e, 0x5b, 0xa0, 0x30, 0x2c, 0x2c, 0xfb, 0x58, 0xa0, 0xdb, 0xdd, 0xec, 0x92, 0x9e, 0xba,
	0xc9, 0x86, 0x03, 0x9b, 0x98, 0x78, 0x54, 0x8c, 0x78, 0x30, 0x51, 0xdb, 0x83, 0x57, 0xab, 0x9d,
	0x9a, 0x46, 0x7c, 0xaf, 0xbe, 0x3e, 0xa2, 0x89, 0xbf, 0xd5, 0xff, 0x62, 0x6c, 0xa7, 0x4d, 0x8b,
	0x10, 0x6e, 0x7c, 0x33, 0xdf, 0x37, 0xdf, 0x63, 0xe6, 0x2b, 0xf4, 0x82, 0xe7, 0xf4, 0x4e, 0x62,
	0x38, 0x4b, 0xa4, 0x50, 0x82, 0x99, 0x04, 0x1d, 0x17, 0xbe, 0xf8, 0xa8, 0x2e, 0xb8, 0x87, 0x4f,
	0x6b, 0x4c, 0x15, 0xb3, 0xc0, 0x4c, 0xa4, 0x88, 0xe2, 0x15, 0x5a, 0xda, 0x54, 0x73, 0x3b, 0x5e,
	0x01, 0x9d, 0xaf, 0xd0, 0x23, 0x66, 0x9a, 0x08, 0x9e, 0xa2, 0xf3, 0x37, 0x2f, 0x44, 0xd1, 0x7e,
	0xed, 0x00, 0xfa, 0x05, 0x95, 0xc4, 0x37, 0x30, 0xf0, 0x51, 0x2d, 0x04, 0x8f, 0xe2, 0xfb, 0xbd,
	0x7a, 0x36, 0x86, 0x56, 0x8a, 0x32, 0x0e, 0x56, 0x96, 0x9e, 0x35, 0x08, 0x31, 0x1b, 0xda, 0xe1,
	0x5a, 0x06, 0x2a, 0x16, 0xdc, 0x32, 0xa6, 0x9a, 0x6b, 0x78, 0x25, 0x76, 0x86, 0xf0, 0xad, 0xe2,
	0x40, 0xb6, 0x27, 0xc0, 0x7c, 0x54, 0x4b, 0xe4, 0x28, 0x03, 0x85, 0xfb, 0x8d, 0xbf, 0x43, 0x53,
	0x89, 0x07, 0xe4, 0xe4, 0x9b, 0x03, 0x67, 0x04, 0xc3, 0xda, 0x14, 0x1a, 0x3e, 0x81, 0xd1, 0x12,
	0xd5, 0x65, 0x2e, 0x3d, 0x8f, 0x53, 0x45, 0xf3, 0x9d, 0x53, 0x18, 0x6f, 0x36, 0x72, 0x09, 0xfb,
	0x07, 0x6d, 0xb2, 0x4a, 0x2d, 0x6d, 0x6a, 0xb8, 0xdd, 0xf9, 0x60, 0x56, 0x5c, 0x8a, 0xf8, 0x5e,
	0xc9, 0x70, 0x5e, 0xc1, 0xa4, 0x22, 0x63, 0xd0, 0xe0, 0xc1, 0x63, 0xf1, 0xde, 0xec, 0x37, 0xeb,
	0x83, 0x2e, 0xf2, 0x97, 0xb6, 0x3d, 0x5d, 0xf0, 0xca, 0xd6, 0x8c, 0x9d, 0x5b, 0x6b, 0xd4, 0xb7,
	0xf6, 0xb1, 0x0a, 0x7c, 0x49, 0x62, 0x89, 0xa1, 0xd5, 0xcc, 0x57, 0x41, 0x70, 0xfe, 0xa6, 0x83,
	0x79, 0x74, 0xed, 0x2f, 0x24, 0x86, 0xec, 0x00, 0x9a, 0x59, 0x16, 0xd8, 0xa8, 0x7c, 0x6d, 0x35,
	0x45, 0xf6, 0x78, 0xb3, 0x4c, 0x7f, 0xf7, 0x10, 0x5a, 0x79, 0x0e, 0x58, 0x9d, 0x51, 0x66, 0xc8,
	0x9e, 0x7c, 0xaa, 0x93, 0xf4, 0x18, 0x3a, 0xe5, 0x39, 0xd9, 0x8f, 0x2a, 0xab, 0x16, 0x22, 0xdb,
	0xde, 0xd6, 0xa2, 0x19, 0x67, 0xd0, 0xad, 0xdc, 0x8d, 0xfd, 0xac, 0x52, 0x37, 0x32, 0x61, 0xff,
	0xda, 0xde, 0xa4, 0x49, 0x57, 0xd0, 0xaf, 0x5f, 0x94, 0xfd, 0x29, 0xf9, 0xdb, 0x4f, 0x6d, 0xff,
	0xde, 0x49, 0xc8, 0x0c, 0x6f, 0x5b, 0xd9, 0x97, 0xf9, 0xff, 0x3d, 0x00, 0x00, 0xff, 0xff, 0x00,
	0xc7, 0x12, 0xc6, 0xaa, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AWSCredClient is the client API for AWSCred service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AWSCredClient interface {
	SetOn(ctx context.Context, in *SetOnRequest, opts ...grpc.CallOption) (*SetOnResponse, error)
	SetOff(ctx context.Context, in *SetOffRequest, opts ...grpc.CallOption) (*SetOffResponse, error)
	SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error)
	SetGenerate(ctx context.Context, in *SetGenerateRequest, opts ...grpc.CallOption) (*SetGenerateResponse, error)
	GetProfileList(ctx context.Context, in *GetProfileListResponse, opts ...grpc.CallOption) (*GetProfileListRequest, error)
}

type aWSCredClient struct {
	cc *grpc.ClientConn
}

func NewAWSCredClient(cc *grpc.ClientConn) AWSCredClient {
	return &aWSCredClient{cc}
}

func (c *aWSCredClient) SetOn(ctx context.Context, in *SetOnRequest, opts ...grpc.CallOption) (*SetOnResponse, error) {
	out := new(SetOnResponse)
	err := c.cc.Invoke(ctx, "/awscred.AWSCred/SetOn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCredClient) SetOff(ctx context.Context, in *SetOffRequest, opts ...grpc.CallOption) (*SetOffResponse, error) {
	out := new(SetOffResponse)
	err := c.cc.Invoke(ctx, "/awscred.AWSCred/SetOff", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCredClient) SetConfig(ctx context.Context, in *SetConfigRequest, opts ...grpc.CallOption) (*SetConfigResponse, error) {
	out := new(SetConfigResponse)
	err := c.cc.Invoke(ctx, "/awscred.AWSCred/SetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCredClient) SetGenerate(ctx context.Context, in *SetGenerateRequest, opts ...grpc.CallOption) (*SetGenerateResponse, error) {
	out := new(SetGenerateResponse)
	err := c.cc.Invoke(ctx, "/awscred.AWSCred/SetGenerate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aWSCredClient) GetProfileList(ctx context.Context, in *GetProfileListResponse, opts ...grpc.CallOption) (*GetProfileListRequest, error) {
	out := new(GetProfileListRequest)
	err := c.cc.Invoke(ctx, "/awscred.AWSCred/GetProfileList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AWSCredServer is the server API for AWSCred service.
type AWSCredServer interface {
	SetOn(context.Context, *SetOnRequest) (*SetOnResponse, error)
	SetOff(context.Context, *SetOffRequest) (*SetOffResponse, error)
	SetConfig(context.Context, *SetConfigRequest) (*SetConfigResponse, error)
	SetGenerate(context.Context, *SetGenerateRequest) (*SetGenerateResponse, error)
	GetProfileList(context.Context, *GetProfileListResponse) (*GetProfileListRequest, error)
}

// UnimplementedAWSCredServer can be embedded to have forward compatible implementations.
type UnimplementedAWSCredServer struct {
}

func (*UnimplementedAWSCredServer) SetOn(ctx context.Context, req *SetOnRequest) (*SetOnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOn not implemented")
}
func (*UnimplementedAWSCredServer) SetOff(ctx context.Context, req *SetOffRequest) (*SetOffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetOff not implemented")
}
func (*UnimplementedAWSCredServer) SetConfig(ctx context.Context, req *SetConfigRequest) (*SetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetConfig not implemented")
}
func (*UnimplementedAWSCredServer) SetGenerate(ctx context.Context, req *SetGenerateRequest) (*SetGenerateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SetGenerate not implemented")
}
func (*UnimplementedAWSCredServer) GetProfileList(ctx context.Context, req *GetProfileListResponse) (*GetProfileListRequest, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileList not implemented")
}

func RegisterAWSCredServer(s *grpc.Server, srv AWSCredServer) {
	s.RegisterService(&_AWSCred_serviceDesc, srv)
}

func _AWSCred_SetOn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOnRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCredServer).SetOn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awscred.AWSCred/SetOn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCredServer).SetOn(ctx, req.(*SetOnRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCred_SetOff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetOffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCredServer).SetOff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awscred.AWSCred/SetOff",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCredServer).SetOff(ctx, req.(*SetOffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCred_SetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCredServer).SetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awscred.AWSCred/SetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCredServer).SetConfig(ctx, req.(*SetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCred_SetGenerate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SetGenerateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCredServer).SetGenerate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awscred.AWSCred/SetGenerate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCredServer).SetGenerate(ctx, req.(*SetGenerateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AWSCred_GetProfileList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProfileListResponse)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AWSCredServer).GetProfileList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/awscred.AWSCred/GetProfileList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AWSCredServer).GetProfileList(ctx, req.(*GetProfileListResponse))
	}
	return interceptor(ctx, in, info, handler)
}

var _AWSCred_serviceDesc = grpc.ServiceDesc{
	ServiceName: "awscred.AWSCred",
	HandlerType: (*AWSCredServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SetOn",
			Handler:    _AWSCred_SetOn_Handler,
		},
		{
			MethodName: "SetOff",
			Handler:    _AWSCred_SetOff_Handler,
		},
		{
			MethodName: "SetConfig",
			Handler:    _AWSCred_SetConfig_Handler,
		},
		{
			MethodName: "SetGenerate",
			Handler:    _AWSCred_SetGenerate_Handler,
		},
		{
			MethodName: "GetProfileList",
			Handler:    _AWSCred_GetProfileList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "awscred.proto",
}