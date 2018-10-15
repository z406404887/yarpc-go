// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: src/keyvalue/key_value.proto

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package keyvaluepb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "go.uber.org/yarpc/v2/yarpcprotobuf/protoc-gen-yarpc-go/internal/tests/gen/proto/src/common"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func init() {
	proto.RegisterFile("src/keyvalue/key_value.proto", fileDescriptor_key_value_89e7f2910f5faa8a)
}

var fileDescriptor_key_value_89e7f2910f5faa8a = []byte{
	// 177 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x29, 0x2e, 0x4a, 0xd6,
	0xcf, 0x4e, 0xad, 0x2c, 0x4b, 0xcc, 0x29, 0x4d, 0x05, 0x31, 0xe2, 0xc1, 0x2c, 0xbd, 0x82, 0xa2,
	0xfc, 0x92, 0x7c, 0x21, 0x0e, 0x98, 0x8c, 0x94, 0x38, 0x48, 0x5d, 0x72, 0x7e, 0x6e, 0x6e, 0x7e,
	0x1e, 0x94, 0x82, 0x28, 0x31, 0x4a, 0xe7, 0x62, 0x0d, 0x2e, 0xc9, 0x2f, 0x4a, 0x15, 0xd2, 0xe3,
	0x62, 0x76, 0x4f, 0x2d, 0x11, 0x12, 0xd2, 0x83, 0x4a, 0xbb, 0xa7, 0x96, 0x04, 0xa5, 0x16, 0x96,
	0xa6, 0x16, 0x97, 0x48, 0x09, 0xa3, 0x88, 0x15, 0x17, 0xe4, 0xe7, 0x15, 0x83, 0xd5, 0x07, 0x23,
	0xab, 0x0f, 0xc6, 0xa2, 0x3e, 0x18, 0xa1, 0xde, 0xc9, 0xe2, 0xc2, 0x43, 0x39, 0x86, 0x1b, 0x0f,
	0xe5, 0x18, 0x3e, 0x3c, 0x94, 0x63, 0x6c, 0x78, 0x24, 0xc7, 0xb8, 0xe2, 0x91, 0x1c, 0xe3, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0xf8, 0xe2, 0x91, 0x1c, 0xc3,
	0x87, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x44, 0x71, 0xc1, 0x5c, 0x5e, 0x90, 0x94, 0xc4,
	0x06, 0x76, 0xa9, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x0b, 0x1c, 0xdb, 0x62, 0xec, 0x00, 0x00,
	0x00,
}