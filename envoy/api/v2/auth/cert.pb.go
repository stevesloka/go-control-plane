// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/auth/cert.proto

package envoy_api_v2_auth

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	proto "github.com/golang/protobuf/proto"
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

func init() { proto.RegisterFile("envoy/api/v2/auth/cert.proto", fileDescriptor_0c3851c0865b2745) }

var fileDescriptor_0c3851c0865b2745 = []byte{
	// 205 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x8f, 0x4d, 0x4e, 0xc4, 0x30,
	0x0c, 0x85, 0xc5, 0xaf, 0xc4, 0xec, 0x98, 0x65, 0x81, 0x61, 0x8d, 0x90, 0x1c, 0x69, 0x7a, 0x83,
	0x72, 0x81, 0xdc, 0xa0, 0x32, 0xc5, 0x82, 0x08, 0x1a, 0x47, 0xb6, 0x1b, 0xb5, 0xa7, 0x60, 0xcb,
	0x39, 0x39, 0x00, 0x42, 0x69, 0xbb, 0x2b, 0xac, 0x3f, 0xbf, 0xf7, 0x3e, 0xef, 0x6e, 0x29, 0x66,
	0x9e, 0x1c, 0xa6, 0xe0, 0xf2, 0xd1, 0xe1, 0x60, 0x6f, 0xae, 0x23, 0x31, 0x48, 0xc2, 0xc6, 0xfb,
	0xeb, 0x99, 0x02, 0xa6, 0x00, 0xf9, 0x08, 0x85, 0x56, 0x87, 0xe1, 0x25, 0xa1, 0xc3, 0x18, 0xd9,
	0xd0, 0x02, 0x47, 0x75, 0x7d, 0x78, 0x15, 0x34, 0x5a, 0x22, 0xd5, 0xdd, 0x86, 0xab, 0xa1, 0x0d,
	0xba, 0xe2, 0xc3, 0x1f, 0x7b, 0xdc, 0xf7, 0x1c, 0xff, 0xe7, 0x4a, 0x9d, 0xd0, 0x6a, 0x54, 0xdd,
	0x6c, 0xb9, 0x7d, 0xac, 0xe5, 0x4d, 0xfb, 0xfd, 0xf5, 0xf3, 0x79, 0xf1, 0xb8, 0x7f, 0x58, 0xb4,
	0x69, 0x34, 0x8a, 0x5a, 0x1c, 0xc0, 0x04, 0xa3, 0x26, 0x16, 0x6b, 0x95, 0xbb, 0x77, 0x32, 0x85,
	0x92, 0xca, 0xf5, 0xee, 0x3e, 0x30, 0xcc, 0xd7, 0x49, 0x78, 0x9c, 0x60, 0xf3, 0x6f, 0x73, 0xf5,
	0x44, 0x62, 0xbe, 0xd4, 0xfb, 0x13, 0x7f, 0xea, 0xcf, 0xfc, 0xf9, 0xf3, 0xe5, 0xbc, 0x56, 0xff,
	0x06, 0x00, 0x00, 0xff, 0xff, 0x24, 0xf1, 0x42, 0x7e, 0x3c, 0x01, 0x00, 0x00,
}
