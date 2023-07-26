// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.24.0--rc1
// source: proto/query.proto

package query

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

type LogsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogsRequest) Reset() {
	*x = LogsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_query_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsRequest) ProtoMessage() {}

func (x *LogsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_query_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsRequest.ProtoReflect.Descriptor instead.
func (*LogsRequest) Descriptor() ([]byte, []int) {
	return file_proto_query_proto_rawDescGZIP(), []int{0}
}

type LogsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Logs []*Log `protobuf:"bytes,1,rep,name=logs,proto3" json:"logs,omitempty"`
}

func (x *LogsResponse) Reset() {
	*x = LogsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_query_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsResponse) ProtoMessage() {}

func (x *LogsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_query_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsResponse.ProtoReflect.Descriptor instead.
func (*LogsResponse) Descriptor() ([]byte, []int) {
	return file_proto_query_proto_rawDescGZIP(), []int{1}
}

func (x *LogsResponse) GetLogs() []*Log {
	if x != nil {
		return x.Logs
	}
	return nil
}

type Log struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TimeUnixNano uint64 `protobuf:"fixed64,1,opt,name=time_unix_nano,json=timeUnixNano,proto3" json:"time_unix_nano,omitempty"`
	Body         string `protobuf:"bytes,2,opt,name=body,proto3" json:"body,omitempty"` // TODO get rid of this
}

func (x *Log) Reset() {
	*x = Log{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_query_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Log) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Log) ProtoMessage() {}

func (x *Log) ProtoReflect() protoreflect.Message {
	mi := &file_proto_query_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Log.ProtoReflect.Descriptor instead.
func (*Log) Descriptor() ([]byte, []int) {
	return file_proto_query_proto_rawDescGZIP(), []int{2}
}

func (x *Log) GetTimeUnixNano() uint64 {
	if x != nil {
		return x.TimeUnixNano
	}
	return 0
}

func (x *Log) GetBody() string {
	if x != nil {
		return x.Body
	}
	return ""
}

var File_proto_query_proto protoreflect.FileDescriptor

var file_proto_query_proto_rawDesc = []byte{
	0x0a, 0x11, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x0d, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x28, 0x0a, 0x0c, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x18, 0x0a, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x04, 0x2e, 0x4c, 0x6f, 0x67, 0x52, 0x04, 0x6c, 0x6f, 0x67, 0x73, 0x22, 0x3f, 0x0a, 0x03,
	0x4c, 0x6f, 0x67, 0x12, 0x24, 0x0a, 0x0e, 0x74, 0x69, 0x6d, 0x65, 0x5f, 0x75, 0x6e, 0x69, 0x78,
	0x5f, 0x6e, 0x61, 0x6e, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x06, 0x52, 0x0c, 0x74, 0x69, 0x6d,
	0x65, 0x55, 0x6e, 0x69, 0x78, 0x4e, 0x61, 0x6e, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x62, 0x6f, 0x64,
	0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x62, 0x6f, 0x64, 0x79, 0x32, 0x31, 0x0a,
	0x05, 0x51, 0x75, 0x65, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67,
	0x73, 0x12, 0x0c, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x0d, 0x2e, 0x4c, 0x6f, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61,
	0x6c, 0x62, 0x65, 0x72, 0x74, 0x6c, 0x6f, 0x63, 0x6b, 0x65, 0x74, 0x74, 0x2f, 0x6f, 0x74, 0x65,
	0x6c, 0x2d, 0x6c, 0x6f, 0x67, 0x73, 0x2d, 0x66, 0x72, 0x6f, 0x73, 0x74, 0x64, 0x62, 0x2f, 0x71,
	0x75, 0x65, 0x72, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_query_proto_rawDescOnce sync.Once
	file_proto_query_proto_rawDescData = file_proto_query_proto_rawDesc
)

func file_proto_query_proto_rawDescGZIP() []byte {
	file_proto_query_proto_rawDescOnce.Do(func() {
		file_proto_query_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_query_proto_rawDescData)
	})
	return file_proto_query_proto_rawDescData
}

var file_proto_query_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_query_proto_goTypes = []interface{}{
	(*LogsRequest)(nil),  // 0: LogsRequest
	(*LogsResponse)(nil), // 1: LogsResponse
	(*Log)(nil),          // 2: Log
}
var file_proto_query_proto_depIdxs = []int32{
	2, // 0: LogsResponse.logs:type_name -> Log
	0, // 1: Query.GetLogs:input_type -> LogsRequest
	1, // 2: Query.GetLogs:output_type -> LogsResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_query_proto_init() }
func file_proto_query_proto_init() {
	if File_proto_query_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_query_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsRequest); i {
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
		file_proto_query_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsResponse); i {
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
		file_proto_query_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Log); i {
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
			RawDescriptor: file_proto_query_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_query_proto_goTypes,
		DependencyIndexes: file_proto_query_proto_depIdxs,
		MessageInfos:      file_proto_query_proto_msgTypes,
	}.Build()
	File_proto_query_proto = out.File
	file_proto_query_proto_rawDesc = nil
	file_proto_query_proto_goTypes = nil
	file_proto_query_proto_depIdxs = nil
}