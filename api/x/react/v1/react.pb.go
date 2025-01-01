// Copyright 2024 The Convērs Authors.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.21.12
// source: react.proto

package react

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SayHelloRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	StrVal    *wrapperspb.StringValue `protobuf:"bytes,2,opt,name=str_val,json=strVal,proto3" json:"str_val,omitempty"`
	FloatVal  *wrapperspb.FloatValue  `protobuf:"bytes,3,opt,name=float_val,json=floatVal,proto3" json:"float_val,omitempty"`
	DoubleVal *wrapperspb.DoubleValue `protobuf:"bytes,4,opt,name=double_val,json=doubleVal,proto3" json:"double_val,omitempty"`
	BoolVal   *wrapperspb.BoolValue   `protobuf:"bytes,5,opt,name=bool_val,json=boolVal,proto3" json:"bool_val,omitempty"`
	BytesVal  *wrapperspb.BytesValue  `protobuf:"bytes,6,opt,name=bytes_val,json=bytesVal,proto3" json:"bytes_val,omitempty"`
	Int32Val  *wrapperspb.Int32Value  `protobuf:"bytes,7,opt,name=int32_val,json=int32Val,proto3" json:"int32_val,omitempty"`
	Uint32Val *wrapperspb.UInt32Value `protobuf:"bytes,8,opt,name=uint32_val,json=uint32Val,proto3" json:"uint32_val,omitempty"`
	Int64Val  *wrapperspb.Int64Value  `protobuf:"bytes,9,opt,name=int64_val,json=int64Val,proto3" json:"int64_val,omitempty"`
	Uint64Val *wrapperspb.UInt64Value `protobuf:"bytes,10,opt,name=uint64_val,json=uint64Val,proto3" json:"uint64_val,omitempty"`
}

func (x *SayHelloRequest) Reset() {
	*x = SayHelloRequest{}
	mi := &file_react_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloRequest) ProtoMessage() {}

func (x *SayHelloRequest) ProtoReflect() protoreflect.Message {
	mi := &file_react_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloRequest.ProtoReflect.Descriptor instead.
func (*SayHelloRequest) Descriptor() ([]byte, []int) {
	return file_react_proto_rawDescGZIP(), []int{0}
}

func (x *SayHelloRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SayHelloRequest) GetStrVal() *wrapperspb.StringValue {
	if x != nil {
		return x.StrVal
	}
	return nil
}

func (x *SayHelloRequest) GetFloatVal() *wrapperspb.FloatValue {
	if x != nil {
		return x.FloatVal
	}
	return nil
}

func (x *SayHelloRequest) GetDoubleVal() *wrapperspb.DoubleValue {
	if x != nil {
		return x.DoubleVal
	}
	return nil
}

func (x *SayHelloRequest) GetBoolVal() *wrapperspb.BoolValue {
	if x != nil {
		return x.BoolVal
	}
	return nil
}

func (x *SayHelloRequest) GetBytesVal() *wrapperspb.BytesValue {
	if x != nil {
		return x.BytesVal
	}
	return nil
}

func (x *SayHelloRequest) GetInt32Val() *wrapperspb.Int32Value {
	if x != nil {
		return x.Int32Val
	}
	return nil
}

func (x *SayHelloRequest) GetUint32Val() *wrapperspb.UInt32Value {
	if x != nil {
		return x.Uint32Val
	}
	return nil
}

func (x *SayHelloRequest) GetInt64Val() *wrapperspb.Int64Value {
	if x != nil {
		return x.Int64Val
	}
	return nil
}

func (x *SayHelloRequest) GetUint64Val() *wrapperspb.UInt64Value {
	if x != nil {
		return x.Uint64Val
	}
	return nil
}

type SayHelloResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SayHelloResponse) Reset() {
	*x = SayHelloResponse{}
	mi := &file_react_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SayHelloResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SayHelloResponse) ProtoMessage() {}

func (x *SayHelloResponse) ProtoReflect() protoreflect.Message {
	mi := &file_react_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SayHelloResponse.ProtoReflect.Descriptor instead.
func (*SayHelloResponse) Descriptor() ([]byte, []int) {
	return file_react_proto_rawDescGZIP(), []int{1}
}

func (x *SayHelloResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_react_proto protoreflect.FileDescriptor

var file_react_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63,
	0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x1a,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61,
	0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1c, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x77, 0x72, 0x61, 0x70,
	0x70, 0x65, 0x72, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbb, 0x04, 0x0a, 0x0f, 0x53,
	0x61, 0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xba, 0x48,
	0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x07, 0x73,
	0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x53,
	0x74, 0x72, 0x69, 0x6e, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x06, 0x73, 0x74, 0x72, 0x56,
	0x61, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x08, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x56, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x0a,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x44, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09,
	0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x56, 0x61, 0x6c, 0x12, 0x35, 0x0a, 0x08, 0x62, 0x6f, 0x6f,
	0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x6f,
	0x6f, 0x6c, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x07, 0x62, 0x6f, 0x6f, 0x6c, 0x56, 0x61, 0x6c,
	0x12, 0x38, 0x0a, 0x09, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x42, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x52, 0x08, 0x62, 0x79, 0x74, 0x65, 0x73, 0x56, 0x61, 0x6c, 0x12, 0x38, 0x0a, 0x09, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x49, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x0a, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76,
	0x61, 0x6c, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x33,
	0x32, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x56, 0x61,
	0x6c, 0x12, 0x38, 0x0a, 0x09, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x09,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x0a, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x1c, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x55, 0x49, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x52, 0x09, 0x75,
	0x69, 0x6e, 0x74, 0x36, 0x34, 0x56, 0x61, 0x6c, 0x22, 0x2c, 0x0a, 0x10, 0x53, 0x61, 0x79, 0x48,
	0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x80, 0x03, 0x0a, 0x0e, 0x47, 0x72, 0x65, 0x65, 0x74,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0xed, 0x02, 0x0a, 0x08, 0x53, 0x61,
	0x79, 0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x12, 0x21, 0x2e, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73,
	0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79, 0x48, 0x65, 0x6c,
	0x6c, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x63, 0x6f, 0x6e, 0x76,
	0x65, 0x72, 0x73, 0x2e, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x61, 0x79,
	0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x99, 0x02,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x92, 0x02, 0x5a, 0x17, 0x12, 0x15, 0x2f, 0x73, 0x61, 0x79, 0x2f,
	0x73, 0x74, 0x72, 0x76, 0x61, 0x6c, 0x2f, 0x7b, 0x73, 0x74, 0x72, 0x5f, 0x76, 0x61, 0x6c, 0x7d,
	0x5a, 0x1b, 0x12, 0x19, 0x2f, 0x73, 0x61, 0x79, 0x2f, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x76, 0x61,
	0x6c, 0x2f, 0x7b, 0x66, 0x6c, 0x6f, 0x61, 0x74, 0x5f, 0x76, 0x61, 0x6c, 0x7d, 0x5a, 0x1d, 0x12,
	0x1b, 0x2f, 0x73, 0x61, 0x79, 0x2f, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x76, 0x61, 0x6c, 0x2f,
	0x7b, 0x64, 0x6f, 0x75, 0x62, 0x6c, 0x65, 0x5f, 0x76, 0x61, 0x6c, 0x7d, 0x5a, 0x19, 0x12, 0x17,
	0x2f, 0x73, 0x61, 0x79, 0x2f, 0x62, 0x6f, 0x6f, 0x6c, 0x76, 0x61, 0x6c, 0x2f, 0x7b, 0x62, 0x6f,
	0x6f, 0x6c, 0x5f, 0x76, 0x61, 0x6c, 0x7d, 0x5a, 0x1b, 0x12, 0x19, 0x2f, 0x73, 0x61, 0x79, 0x2f,
	0x62, 0x79, 0x74, 0x65, 0x73, 0x76, 0x61, 0x6c, 0x2f, 0x7b, 0x62, 0x79, 0x74, 0x65, 0x73, 0x5f,
	0x76, 0x61, 0x6c, 0x7d, 0x5a, 0x1b, 0x12, 0x19, 0x2f, 0x73, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74,
	0x33, 0x32, 0x76, 0x61, 0x6c, 0x2f, 0x7b, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c,
	0x7d, 0x5a, 0x1d, 0x12, 0x1b, 0x2f, 0x73, 0x61, 0x79, 0x2f, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32,
	0x76, 0x61, 0x6c, 0x2f, 0x7b, 0x75, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x5f, 0x76, 0x61, 0x6c, 0x7d,
	0x5a, 0x1b, 0x12, 0x19, 0x2f, 0x73, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x76, 0x61,
	0x6c, 0x2f, 0x7b, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x7d, 0x5a, 0x1d, 0x12,
	0x1b, 0x2f, 0x73, 0x61, 0x79, 0x2f, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x76, 0x61, 0x6c, 0x2f,
	0x7b, 0x75, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x5f, 0x76, 0x61, 0x6c, 0x7d, 0x12, 0x0b, 0x2f, 0x73,
	0x61, 0x79, 0x2f, 0x7b, 0x6e, 0x61, 0x6d, 0x65, 0x7d, 0x42, 0x94, 0x01, 0x92, 0x41, 0x63, 0x12,
	0x14, 0x0a, 0x0d, 0x52, 0x65, 0x61, 0x63, 0x74, 0x20, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x02, 0x02, 0x01, 0x32, 0x10, 0x61, 0x70, 0x70, 0x6c, 0x69,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6a, 0x73, 0x6f, 0x6e, 0x5a, 0x23, 0x0a, 0x21, 0x0a,
	0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68, 0x12, 0x13, 0x08, 0x02, 0x1a,
	0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x20, 0x02,
	0x62, 0x10, 0x0a, 0x0e, 0x0a, 0x0a, 0x41, 0x70, 0x69, 0x4b, 0x65, 0x79, 0x41, 0x75, 0x74, 0x68,
	0x12, 0x00, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x6e, 0x76, 0x72, 0x7a, 0x2f, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x72, 0x65, 0x61, 0x63, 0x74, 0x2f, 0x76, 0x31, 0x3b, 0x72, 0x65, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_react_proto_rawDescOnce sync.Once
	file_react_proto_rawDescData = file_react_proto_rawDesc
)

func file_react_proto_rawDescGZIP() []byte {
	file_react_proto_rawDescOnce.Do(func() {
		file_react_proto_rawDescData = protoimpl.X.CompressGZIP(file_react_proto_rawDescData)
	})
	return file_react_proto_rawDescData
}

var file_react_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_react_proto_goTypes = []any{
	(*SayHelloRequest)(nil),        // 0: convers.react.v1.SayHelloRequest
	(*SayHelloResponse)(nil),       // 1: convers.react.v1.SayHelloResponse
	(*wrapperspb.StringValue)(nil), // 2: google.protobuf.StringValue
	(*wrapperspb.FloatValue)(nil),  // 3: google.protobuf.FloatValue
	(*wrapperspb.DoubleValue)(nil), // 4: google.protobuf.DoubleValue
	(*wrapperspb.BoolValue)(nil),   // 5: google.protobuf.BoolValue
	(*wrapperspb.BytesValue)(nil),  // 6: google.protobuf.BytesValue
	(*wrapperspb.Int32Value)(nil),  // 7: google.protobuf.Int32Value
	(*wrapperspb.UInt32Value)(nil), // 8: google.protobuf.UInt32Value
	(*wrapperspb.Int64Value)(nil),  // 9: google.protobuf.Int64Value
	(*wrapperspb.UInt64Value)(nil), // 10: google.protobuf.UInt64Value
}
var file_react_proto_depIdxs = []int32{
	2,  // 0: convers.react.v1.SayHelloRequest.str_val:type_name -> google.protobuf.StringValue
	3,  // 1: convers.react.v1.SayHelloRequest.float_val:type_name -> google.protobuf.FloatValue
	4,  // 2: convers.react.v1.SayHelloRequest.double_val:type_name -> google.protobuf.DoubleValue
	5,  // 3: convers.react.v1.SayHelloRequest.bool_val:type_name -> google.protobuf.BoolValue
	6,  // 4: convers.react.v1.SayHelloRequest.bytes_val:type_name -> google.protobuf.BytesValue
	7,  // 5: convers.react.v1.SayHelloRequest.int32_val:type_name -> google.protobuf.Int32Value
	8,  // 6: convers.react.v1.SayHelloRequest.uint32_val:type_name -> google.protobuf.UInt32Value
	9,  // 7: convers.react.v1.SayHelloRequest.int64_val:type_name -> google.protobuf.Int64Value
	10, // 8: convers.react.v1.SayHelloRequest.uint64_val:type_name -> google.protobuf.UInt64Value
	0,  // 9: convers.react.v1.GreeterService.SayHello:input_type -> convers.react.v1.SayHelloRequest
	1,  // 10: convers.react.v1.GreeterService.SayHello:output_type -> convers.react.v1.SayHelloResponse
	10, // [10:11] is the sub-list for method output_type
	9,  // [9:10] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_react_proto_init() }
func file_react_proto_init() {
	if File_react_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_react_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_react_proto_goTypes,
		DependencyIndexes: file_react_proto_depIdxs,
		MessageInfos:      file_react_proto_msgTypes,
	}.Build()
	File_react_proto = out.File
	file_react_proto_rawDesc = nil
	file_react_proto_goTypes = nil
	file_react_proto_depIdxs = nil
}
