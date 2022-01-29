// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: test-grpc/hello.proto

package test_grpc

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

type HelloReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *HelloReq) Reset() {
	*x = HelloReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_hello_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloReq) ProtoMessage() {}

func (x *HelloReq) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_hello_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloReq.ProtoReflect.Descriptor instead.
func (*HelloReq) Descriptor() ([]byte, []int) {
	return file_test_grpc_hello_proto_rawDescGZIP(), []int{0}
}

func (x *HelloReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type HelloRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age int64 `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *HelloRsp) Reset() {
	*x = HelloRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_hello_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HelloRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HelloRsp) ProtoMessage() {}

func (x *HelloRsp) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_hello_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HelloRsp.ProtoReflect.Descriptor instead.
func (*HelloRsp) Descriptor() ([]byte, []int) {
	return file_test_grpc_hello_proto_rawDescGZIP(), []int{1}
}

func (x *HelloRsp) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

type WorldReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
}

func (x *WorldReq) Reset() {
	*x = WorldReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_hello_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldReq) ProtoMessage() {}

func (x *WorldReq) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_hello_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldReq.ProtoReflect.Descriptor instead.
func (*WorldReq) Descriptor() ([]byte, []int) {
	return file_test_grpc_hello_proto_rawDescGZIP(), []int{2}
}

func (x *WorldReq) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type WorldRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Age int64 `protobuf:"varint,2,opt,name=Age,proto3" json:"Age,omitempty"`
}

func (x *WorldRsp) Reset() {
	*x = WorldRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_test_grpc_hello_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WorldRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorldRsp) ProtoMessage() {}

func (x *WorldRsp) ProtoReflect() protoreflect.Message {
	mi := &file_test_grpc_hello_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorldRsp.ProtoReflect.Descriptor instead.
func (*WorldRsp) Descriptor() ([]byte, []int) {
	return file_test_grpc_hello_proto_rawDescGZIP(), []int{3}
}

func (x *WorldRsp) GetAge() int64 {
	if x != nil {
		return x.Age
	}
	return 0
}

var File_test_grpc_hello_proto protoreflect.FileDescriptor

var file_test_grpc_hello_proto_rawDesc = []byte{
	0x0a, 0x15, 0x74, 0x65, 0x73, 0x74, 0x2d, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x68, 0x65, 0x6c, 0x6c,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x1e, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x65, 0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x48, 0x65, 0x6c, 0x6c, 0x6f,
	0x52, 0x73, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x41, 0x67, 0x65, 0x22, 0x1e, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x12, 0x12, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x1c, 0x0a, 0x08, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x73,
	0x70, 0x12, 0x10, 0x0a, 0x03, 0x41, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03,
	0x41, 0x67, 0x65, 0x32, 0x28, 0x0a, 0x07, 0x54, 0x65, 0x73, 0x74, 0x52, 0x50, 0x43, 0x12, 0x1d,
	0x0a, 0x05, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x09, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52,
	0x65, 0x71, 0x1a, 0x09, 0x2e, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x73, 0x70, 0x32, 0x2d, 0x0a,
	0x0c, 0x54, 0x65, 0x73, 0x74, 0x4d, 0x75, 0x6c, 0x74, 0x69, 0x52, 0x50, 0x43, 0x12, 0x1d, 0x0a,
	0x05, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x12, 0x09, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x65,
	0x71, 0x1a, 0x09, 0x2e, 0x57, 0x6f, 0x72, 0x6c, 0x64, 0x52, 0x73, 0x70, 0x42, 0x2f, 0x5a, 0x2d,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x64, 0x78, 0x6a, 0x2f,
	0x73, 0x69, 0x67, 0x6e, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x65, 0x73, 0x74, 0x5f, 0x67, 0x72, 0x70, 0x63, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_test_grpc_hello_proto_rawDescOnce sync.Once
	file_test_grpc_hello_proto_rawDescData = file_test_grpc_hello_proto_rawDesc
)

func file_test_grpc_hello_proto_rawDescGZIP() []byte {
	file_test_grpc_hello_proto_rawDescOnce.Do(func() {
		file_test_grpc_hello_proto_rawDescData = protoimpl.X.CompressGZIP(file_test_grpc_hello_proto_rawDescData)
	})
	return file_test_grpc_hello_proto_rawDescData
}

var file_test_grpc_hello_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_test_grpc_hello_proto_goTypes = []interface{}{
	(*HelloReq)(nil), // 0: HelloReq
	(*HelloRsp)(nil), // 1: HelloRsp
	(*WorldReq)(nil), // 2: WorldReq
	(*WorldRsp)(nil), // 3: WorldRsp
}
var file_test_grpc_hello_proto_depIdxs = []int32{
	0, // 0: TestRPC.Hello:input_type -> HelloReq
	2, // 1: TestMultiRPC.World:input_type -> WorldReq
	1, // 2: TestRPC.Hello:output_type -> HelloRsp
	3, // 3: TestMultiRPC.World:output_type -> WorldRsp
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_test_grpc_hello_proto_init() }
func file_test_grpc_hello_proto_init() {
	if File_test_grpc_hello_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_test_grpc_hello_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloReq); i {
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
		file_test_grpc_hello_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HelloRsp); i {
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
		file_test_grpc_hello_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldReq); i {
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
		file_test_grpc_hello_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WorldRsp); i {
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
			RawDescriptor: file_test_grpc_hello_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_test_grpc_hello_proto_goTypes,
		DependencyIndexes: file_test_grpc_hello_proto_depIdxs,
		MessageInfos:      file_test_grpc_hello_proto_msgTypes,
	}.Build()
	File_test_grpc_hello_proto = out.File
	file_test_grpc_hello_proto_rawDesc = nil
	file_test_grpc_hello_proto_goTypes = nil
	file_test_grpc_hello_proto_depIdxs = nil
}
