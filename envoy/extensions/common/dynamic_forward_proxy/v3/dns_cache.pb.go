// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto

package envoy_extensions_common_dynamic_forward_proxy_v3

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3 "github.com/envoyproxy/go-control-plane/envoy/config/cluster/v3"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type DnsCacheConfig struct {
	Name                  string                     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	DnsLookupFamily       v3.Cluster_DnsLookupFamily `protobuf:"varint,2,opt,name=dns_lookup_family,json=dnsLookupFamily,proto3,enum=envoy.config.cluster.v3.Cluster_DnsLookupFamily" json:"dns_lookup_family,omitempty"`
	DnsRefreshRate        *duration.Duration         `protobuf:"bytes,3,opt,name=dns_refresh_rate,json=dnsRefreshRate,proto3" json:"dns_refresh_rate,omitempty"`
	HostTtl               *duration.Duration         `protobuf:"bytes,4,opt,name=host_ttl,json=hostTtl,proto3" json:"host_ttl,omitempty"`
	MaxHosts              *wrappers.UInt32Value      `protobuf:"bytes,5,opt,name=max_hosts,json=maxHosts,proto3" json:"max_hosts,omitempty"`
	DnsFailureRefreshRate *v3.Cluster_RefreshRate    `protobuf:"bytes,6,opt,name=dns_failure_refresh_rate,json=dnsFailureRefreshRate,proto3" json:"dns_failure_refresh_rate,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                   `json:"-"`
	XXX_unrecognized      []byte                     `json:"-"`
	XXX_sizecache         int32                      `json:"-"`
}

func (m *DnsCacheConfig) Reset()         { *m = DnsCacheConfig{} }
func (m *DnsCacheConfig) String() string { return proto.CompactTextString(m) }
func (*DnsCacheConfig) ProtoMessage()    {}
func (*DnsCacheConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_f57893b6dd868364, []int{0}
}

func (m *DnsCacheConfig) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DnsCacheConfig.Unmarshal(m, b)
}
func (m *DnsCacheConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DnsCacheConfig.Marshal(b, m, deterministic)
}
func (m *DnsCacheConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DnsCacheConfig.Merge(m, src)
}
func (m *DnsCacheConfig) XXX_Size() int {
	return xxx_messageInfo_DnsCacheConfig.Size(m)
}
func (m *DnsCacheConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_DnsCacheConfig.DiscardUnknown(m)
}

var xxx_messageInfo_DnsCacheConfig proto.InternalMessageInfo

func (m *DnsCacheConfig) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DnsCacheConfig) GetDnsLookupFamily() v3.Cluster_DnsLookupFamily {
	if m != nil {
		return m.DnsLookupFamily
	}
	return v3.Cluster_AUTO
}

func (m *DnsCacheConfig) GetDnsRefreshRate() *duration.Duration {
	if m != nil {
		return m.DnsRefreshRate
	}
	return nil
}

func (m *DnsCacheConfig) GetHostTtl() *duration.Duration {
	if m != nil {
		return m.HostTtl
	}
	return nil
}

func (m *DnsCacheConfig) GetMaxHosts() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxHosts
	}
	return nil
}

func (m *DnsCacheConfig) GetDnsFailureRefreshRate() *v3.Cluster_RefreshRate {
	if m != nil {
		return m.DnsFailureRefreshRate
	}
	return nil
}

func init() {
	proto.RegisterType((*DnsCacheConfig)(nil), "envoy.extensions.common.dynamic_forward_proxy.v3.DnsCacheConfig")
}

func init() {
	proto.RegisterFile("envoy/extensions/common/dynamic_forward_proxy/v3/dns_cache.proto", fileDescriptor_f57893b6dd868364)
}

var fileDescriptor_f57893b6dd868364 = []byte{
	// 523 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0x97, 0xd0, 0x75, 0x6d, 0x80, 0x52, 0x22, 0x21, 0xc2, 0x80, 0xa9, 0x20, 0x21, 0x55,
	0x13, 0xb2, 0xa7, 0xf6, 0x86, 0xc4, 0x84, 0xd2, 0x6a, 0x80, 0xc4, 0x61, 0x44, 0x80, 0xc4, 0x29,
	0xfa, 0x96, 0x38, 0x6d, 0x44, 0x62, 0x47, 0xb6, 0x93, 0xb5, 0x37, 0x84, 0x38, 0xf0, 0x0c, 0x88,
	0x27, 0xd8, 0x23, 0x70, 0xe2, 0x82, 0xc4, 0x95, 0xd7, 0xe9, 0x09, 0xd9, 0x4e, 0x60, 0xd5, 0x80,
	0x69, 0x37, 0xdb, 0x9f, 0xff, 0x3f, 0xfb, 0xfb, 0xc9, 0x76, 0x9e, 0x10, 0x5a, 0xb1, 0x25, 0x26,
	0x0b, 0x49, 0xa8, 0x48, 0x19, 0x15, 0x38, 0x62, 0x79, 0xce, 0x28, 0x8e, 0x97, 0x14, 0xf2, 0x34,
	0x0a, 0x13, 0xc6, 0x8f, 0x81, 0xc7, 0x61, 0xc1, 0xd9, 0x62, 0x89, 0xab, 0x31, 0x8e, 0xa9, 0x08,
	0x23, 0x88, 0xe6, 0x04, 0x15, 0x9c, 0x49, 0xe6, 0xee, 0x69, 0x02, 0xfa, 0x43, 0x40, 0x86, 0x80,
	0xfe, 0x4a, 0x40, 0xd5, 0x78, 0xfb, 0x81, 0x39, 0x33, 0x62, 0x34, 0x49, 0x67, 0x38, 0xca, 0x4a,
	0x21, 0x09, 0x57, 0xe8, 0x7a, 0x68, 0xc0, 0xdb, 0x3b, 0x33, 0xc6, 0x66, 0x19, 0xc1, 0x7a, 0x76,
	0x54, 0x26, 0x38, 0x2e, 0x39, 0xc8, 0x94, 0xd1, 0x7f, 0xd5, 0x8f, 0x39, 0x14, 0x05, 0xe1, 0xa2,
	0xae, 0xdf, 0x2d, 0xe3, 0x02, 0x30, 0x50, 0xca, 0xa4, 0x8e, 0x09, 0x2c, 0x24, 0xc8, 0xb2, 0x29,
	0xdf, 0x3b, 0x53, 0xae, 0x08, 0x57, 0x0d, 0xa4, 0x74, 0x56, 0x6f, 0xb9, 0x59, 0x41, 0x96, 0xc6,
	0x20, 0x09, 0x6e, 0x06, 0xa6, 0x70, 0xff, 0x4b, 0xcb, 0xe9, 0x4d, 0xa9, 0x98, 0x28, 0x0d, 0x13,
	0xdd, 0x86, 0x7b, 0xdb, 0x69, 0x51, 0xc8, 0x89, 0x67, 0x0d, 0xac, 0x61, 0xd7, 0xdf, 0x5a, 0xf9,
	0x2d, 0x6e, 0x0f, 0xac, 0x40, 0x2f, 0xba, 0x89, 0x73, 0x5d, 0x69, 0xcb, 0x18, 0x7b, 0x57, 0x16,
	0x61, 0x02, 0x79, 0x9a, 0x2d, 0x3d, 0x7b, 0x60, 0x0d, 0x7b, 0xa3, 0x3d, 0x64, 0xfc, 0x19, 0x1b,
	0xa8, 0x51, 0x50, 0x8d, 0xd1, 0xa4, 0x1e, 0x4e, 0xa9, 0x78, 0xa1, 0x83, 0x07, 0x3a, 0xe7, 0x77,
	0x56, 0xfe, 0xe6, 0x07, 0xcb, 0xee, 0x5b, 0xc1, 0xb5, 0x78, 0xbd, 0xe4, 0xbe, 0x74, 0xfa, 0xea,
	0x1c, 0x4e, 0x12, 0x4e, 0xc4, 0x3c, 0xe4, 0x20, 0x89, 0x77, 0x69, 0x60, 0x0d, 0x2f, 0x8f, 0x6e,
	0x21, 0x63, 0x0b, 0x35, 0xb6, 0xd0, 0xb4, 0xb6, 0xe9, 0x5f, 0x59, 0xf9, 0xdd, 0x13, 0xab, 0x3d,
	0x6a, 0xf5, 0xbf, 0x7d, 0x7c, 0x1c, 0xf4, 0x62, 0x2a, 0x02, 0x93, 0x0f, 0x40, 0x12, 0x77, 0xdf,
	0xe9, 0xcc, 0x99, 0x90, 0xa1, 0x94, 0x99, 0xd7, 0x3a, 0x0f, 0xa5, 0xae, 0x76, 0x62, 0xd9, 0xbb,
	0x1b, 0xc1, 0x96, 0x0a, 0xbd, 0x92, 0x99, 0xeb, 0x3b, 0xdd, 0x1c, 0x16, 0xa1, 0x9a, 0x0a, 0x6f,
	0x53, 0x03, 0xee, 0x9c, 0x01, 0xbc, 0x7e, 0x4e, 0xe5, 0x78, 0xf4, 0x06, 0xb2, 0x92, 0x68, 0x75,
	0xbb, 0xf6, 0x60, 0x23, 0xe8, 0xe4, 0xb0, 0x78, 0xa6, 0x62, 0x2e, 0x71, 0x3c, 0xd5, 0x56, 0x02,
	0x69, 0x56, 0x72, 0xb2, 0xde, 0x5e, 0x5b, 0x23, 0x1f, 0x9e, 0x6b, 0xf1, 0x54, 0x4f, 0xc1, 0x8d,
	0x98, 0x8a, 0x03, 0x03, 0x3b, 0xb5, 0xfc, 0xe8, 0xe9, 0xe7, 0xef, 0x9f, 0x76, 0xfc, 0xfa, 0x4b,
	0xfc, 0x46, 0xfd, 0xf7, 0x31, 0x8f, 0x20, 0x2b, 0xe6, 0x80, 0xd6, 0xdf, 0x82, 0xff, 0xf6, 0xeb,
	0xfb, 0x1f, 0x3f, 0xdb, 0x76, 0xdf, 0x76, 0xf6, 0x53, 0x66, 0x6e, 0x66, 0xf6, 0x5f, 0xf4, 0xab,
	0xf8, 0x57, 0x1b, 0xf2, 0xa1, 0x52, 0x75, 0x68, 0x1d, 0xb5, 0xb5, 0xb3, 0xf1, 0xaf, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x9e, 0x1e, 0x72, 0xc8, 0xb8, 0x03, 0x00, 0x00,
}
