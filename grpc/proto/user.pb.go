// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.6.1
// source: grpc/proto/user.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type UserRegisterStatus int32

const (
	UserRegisterStatus_UNKNOWN UserRegisterStatus = 0
	UserRegisterStatus_SUCCESS UserRegisterStatus = 1
	UserRegisterStatus_ERROR   UserRegisterStatus = 2
)

// Enum value maps for UserRegisterStatus.
var (
	UserRegisterStatus_name = map[int32]string{
		0: "UNKNOWN",
		1: "SUCCESS",
		2: "ERROR",
	}
	UserRegisterStatus_value = map[string]int32{
		"UNKNOWN": 0,
		"SUCCESS": 1,
		"ERROR":   2,
	}
)

func (x UserRegisterStatus) Enum() *UserRegisterStatus {
	p := new(UserRegisterStatus)
	*p = x
	return p
}

func (x UserRegisterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserRegisterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_grpc_proto_user_proto_enumTypes[0].Descriptor()
}

func (UserRegisterStatus) Type() protoreflect.EnumType {
	return &file_grpc_proto_user_proto_enumTypes[0]
}

func (x UserRegisterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserRegisterStatus.Descriptor instead.
func (UserRegisterStatus) EnumDescriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{0}
}

type UserRegisterRequesut struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AttendanceGroupId string `protobuf:"bytes,1,opt,name=attendance_group_id,json=attendanceGroupId,proto3" json:"attendance_group_id,omitempty"`
	LoginId           string `protobuf:"bytes,2,opt,name=login_id,json=loginId,proto3" json:"login_id,omitempty"`
	Password          string `protobuf:"bytes,3,opt,name=password,proto3" json:"password,omitempty"`
	Name              string `protobuf:"bytes,4,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *UserRegisterRequesut) Reset() {
	*x = UserRegisterRequesut{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterRequesut) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterRequesut) ProtoMessage() {}

func (x *UserRegisterRequesut) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterRequesut.ProtoReflect.Descriptor instead.
func (*UserRegisterRequesut) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserRegisterRequesut) GetAttendanceGroupId() string {
	if x != nil {
		return x.AttendanceGroupId
	}
	return ""
}

func (x *UserRegisterRequesut) GetLoginId() string {
	if x != nil {
		return x.LoginId
	}
	return ""
}

func (x *UserRegisterRequesut) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserRegisterRequesut) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type UserRegisterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status  UserRegisterStatus `protobuf:"varint,1,opt,name=status,proto3,enum=proto.UserRegisterStatus" json:"status,omitempty"`
	Message string             `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *UserRegisterResponse) Reset() {
	*x = UserRegisterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRegisterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRegisterResponse) ProtoMessage() {}

func (x *UserRegisterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRegisterResponse.ProtoReflect.Descriptor instead.
func (*UserRegisterResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{1}
}

func (x *UserRegisterResponse) GetStatus() UserRegisterStatus {
	if x != nil {
		return x.Status
	}
	return UserRegisterStatus_UNKNOWN
}

func (x *UserRegisterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_grpc_proto_user_proto protoreflect.FileDescriptor

var file_grpc_proto_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x91,
	0x01, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x75, 0x74, 0x12, 0x2e, 0x0a, 0x13, 0x61, 0x74, 0x74, 0x65, 0x6e,
	0x64, 0x61, 0x6e, 0x63, 0x65, 0x5f, 0x67, 0x72, 0x6f, 0x75, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x61, 0x74, 0x74, 0x65, 0x6e, 0x64, 0x61, 0x6e, 0x63, 0x65,
	0x47, 0x72, 0x6f, 0x75, 0x70, 0x49, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c, 0x6f, 0x67, 0x69, 0x6e,
	0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x63, 0x0a, 0x14, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x18, 0x0a,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2a, 0x39, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0b, 0x0a,
	0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x55,
	0x43, 0x43, 0x45, 0x53, 0x53, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x52, 0x52, 0x4f, 0x52,
	0x10, 0x02, 0x32, 0x4e, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x75, 0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_user_proto_rawDescOnce sync.Once
	file_grpc_proto_user_proto_rawDescData = file_grpc_proto_user_proto_rawDesc
)

func file_grpc_proto_user_proto_rawDescGZIP() []byte {
	file_grpc_proto_user_proto_rawDescOnce.Do(func() {
		file_grpc_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_user_proto_rawDescData)
	})
	return file_grpc_proto_user_proto_rawDescData
}

var file_grpc_proto_user_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_grpc_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_grpc_proto_user_proto_goTypes = []interface{}{
	(UserRegisterStatus)(0),      // 0: proto.UserRegisterStatus
	(*UserRegisterRequesut)(nil), // 1: proto.UserRegisterRequesut
	(*UserRegisterResponse)(nil), // 2: proto.UserRegisterResponse
}
var file_grpc_proto_user_proto_depIdxs = []int32{
	0, // 0: proto.UserRegisterResponse.status:type_name -> proto.UserRegisterStatus
	1, // 1: proto.User.Register:input_type -> proto.UserRegisterRequesut
	2, // 2: proto.User.Register:output_type -> proto.UserRegisterResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_user_proto_init() }
func file_grpc_proto_user_proto_init() {
	if File_grpc_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterRequesut); i {
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
		file_grpc_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRegisterResponse); i {
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
			RawDescriptor: file_grpc_proto_user_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_user_proto_goTypes,
		DependencyIndexes: file_grpc_proto_user_proto_depIdxs,
		EnumInfos:         file_grpc_proto_user_proto_enumTypes,
		MessageInfos:      file_grpc_proto_user_proto_msgTypes,
	}.Build()
	File_grpc_proto_user_proto = out.File
	file_grpc_proto_user_proto_rawDesc = nil
	file_grpc_proto_user_proto_goTypes = nil
	file_grpc_proto_user_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// UserClient is the client API for User service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type UserClient interface {
	Register(ctx context.Context, in *UserRegisterRequesut, opts ...grpc.CallOption) (*UserRegisterResponse, error)
}

type userClient struct {
	cc grpc.ClientConnInterface
}

func NewUserClient(cc grpc.ClientConnInterface) UserClient {
	return &userClient{cc}
}

func (c *userClient) Register(ctx context.Context, in *UserRegisterRequesut, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	out := new(UserRegisterResponse)
	err := c.cc.Invoke(ctx, "/proto.User/Register", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServer is the server API for User service.
type UserServer interface {
	Register(context.Context, *UserRegisterRequesut) (*UserRegisterResponse, error)
}

// UnimplementedUserServer can be embedded to have forward compatible implementations.
type UnimplementedUserServer struct {
}

func (*UnimplementedUserServer) Register(context.Context, *UserRegisterRequesut) (*UserRegisterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Register not implemented")
}

func RegisterUserServer(s *grpc.Server, srv UserServer) {
	s.RegisterService(&_User_serviceDesc, srv)
}

func _User_Register_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRegisterRequesut)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServer).Register(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.User/Register",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServer).Register(ctx, req.(*UserRegisterRequesut))
	}
	return interceptor(ctx, in, info, handler)
}

var _User_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.User",
	HandlerType: (*UserServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Register",
			Handler:    _User_Register_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "grpc/proto/user.proto",
}
