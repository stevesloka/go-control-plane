// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v2/auth/tls.proto

package envoy_api_v2_auth

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
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

type UpstreamTlsContext struct {
	CommonTlsContext     *CommonTlsContext     `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext,proto3" json:"common_tls_context,omitempty"`
	Sni                  string                `protobuf:"bytes,2,opt,name=sni,proto3" json:"sni,omitempty"`
	AllowRenegotiation   bool                  `protobuf:"varint,3,opt,name=allow_renegotiation,json=allowRenegotiation,proto3" json:"allow_renegotiation,omitempty"`
	MaxSessionKeys       *wrappers.UInt32Value `protobuf:"bytes,4,opt,name=max_session_keys,json=maxSessionKeys,proto3" json:"max_session_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *UpstreamTlsContext) Reset()         { *m = UpstreamTlsContext{} }
func (m *UpstreamTlsContext) String() string { return proto.CompactTextString(m) }
func (*UpstreamTlsContext) ProtoMessage()    {}
func (*UpstreamTlsContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_463a52fb5d41af5b, []int{0}
}

func (m *UpstreamTlsContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpstreamTlsContext.Unmarshal(m, b)
}
func (m *UpstreamTlsContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpstreamTlsContext.Marshal(b, m, deterministic)
}
func (m *UpstreamTlsContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpstreamTlsContext.Merge(m, src)
}
func (m *UpstreamTlsContext) XXX_Size() int {
	return xxx_messageInfo_UpstreamTlsContext.Size(m)
}
func (m *UpstreamTlsContext) XXX_DiscardUnknown() {
	xxx_messageInfo_UpstreamTlsContext.DiscardUnknown(m)
}

var xxx_messageInfo_UpstreamTlsContext proto.InternalMessageInfo

func (m *UpstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *UpstreamTlsContext) GetSni() string {
	if m != nil {
		return m.Sni
	}
	return ""
}

func (m *UpstreamTlsContext) GetAllowRenegotiation() bool {
	if m != nil {
		return m.AllowRenegotiation
	}
	return false
}

func (m *UpstreamTlsContext) GetMaxSessionKeys() *wrappers.UInt32Value {
	if m != nil {
		return m.MaxSessionKeys
	}
	return nil
}

type DownstreamTlsContext struct {
	CommonTlsContext         *CommonTlsContext   `protobuf:"bytes,1,opt,name=common_tls_context,json=commonTlsContext,proto3" json:"common_tls_context,omitempty"`
	RequireClientCertificate *wrappers.BoolValue `protobuf:"bytes,2,opt,name=require_client_certificate,json=requireClientCertificate,proto3" json:"require_client_certificate,omitempty"`
	RequireSni               *wrappers.BoolValue `protobuf:"bytes,3,opt,name=require_sni,json=requireSni,proto3" json:"require_sni,omitempty"`
	// Types that are valid to be assigned to SessionTicketKeysType:
	//	*DownstreamTlsContext_SessionTicketKeys
	//	*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig
	//	*DownstreamTlsContext_DisableStatelessSessionResumption
	SessionTicketKeysType isDownstreamTlsContext_SessionTicketKeysType `protobuf_oneof:"session_ticket_keys_type"`
	SessionTimeout        *duration.Duration                           `protobuf:"bytes,6,opt,name=session_timeout,json=sessionTimeout,proto3" json:"session_timeout,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                     `json:"-"`
	XXX_unrecognized      []byte                                       `json:"-"`
	XXX_sizecache         int32                                        `json:"-"`
}

func (m *DownstreamTlsContext) Reset()         { *m = DownstreamTlsContext{} }
func (m *DownstreamTlsContext) String() string { return proto.CompactTextString(m) }
func (*DownstreamTlsContext) ProtoMessage()    {}
func (*DownstreamTlsContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_463a52fb5d41af5b, []int{1}
}

func (m *DownstreamTlsContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DownstreamTlsContext.Unmarshal(m, b)
}
func (m *DownstreamTlsContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DownstreamTlsContext.Marshal(b, m, deterministic)
}
func (m *DownstreamTlsContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DownstreamTlsContext.Merge(m, src)
}
func (m *DownstreamTlsContext) XXX_Size() int {
	return xxx_messageInfo_DownstreamTlsContext.Size(m)
}
func (m *DownstreamTlsContext) XXX_DiscardUnknown() {
	xxx_messageInfo_DownstreamTlsContext.DiscardUnknown(m)
}

var xxx_messageInfo_DownstreamTlsContext proto.InternalMessageInfo

func (m *DownstreamTlsContext) GetCommonTlsContext() *CommonTlsContext {
	if m != nil {
		return m.CommonTlsContext
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireClientCertificate() *wrappers.BoolValue {
	if m != nil {
		return m.RequireClientCertificate
	}
	return nil
}

func (m *DownstreamTlsContext) GetRequireSni() *wrappers.BoolValue {
	if m != nil {
		return m.RequireSni
	}
	return nil
}

type isDownstreamTlsContext_SessionTicketKeysType interface {
	isDownstreamTlsContext_SessionTicketKeysType()
}

type DownstreamTlsContext_SessionTicketKeys struct {
	SessionTicketKeys *TlsSessionTicketKeys `protobuf:"bytes,4,opt,name=session_ticket_keys,json=sessionTicketKeys,proto3,oneof"`
}

type DownstreamTlsContext_SessionTicketKeysSdsSecretConfig struct {
	SessionTicketKeysSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,5,opt,name=session_ticket_keys_sds_secret_config,json=sessionTicketKeysSdsSecretConfig,proto3,oneof"`
}

type DownstreamTlsContext_DisableStatelessSessionResumption struct {
	DisableStatelessSessionResumption bool `protobuf:"varint,7,opt,name=disable_stateless_session_resumption,json=disableStatelessSessionResumption,proto3,oneof"`
}

func (*DownstreamTlsContext_SessionTicketKeys) isDownstreamTlsContext_SessionTicketKeysType() {}

func (*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig) isDownstreamTlsContext_SessionTicketKeysType() {
}

func (*DownstreamTlsContext_DisableStatelessSessionResumption) isDownstreamTlsContext_SessionTicketKeysType() {
}

func (m *DownstreamTlsContext) GetSessionTicketKeysType() isDownstreamTlsContext_SessionTicketKeysType {
	if m != nil {
		return m.SessionTicketKeysType
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeys() *TlsSessionTicketKeys {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeys); ok {
		return x.SessionTicketKeys
	}
	return nil
}

func (m *DownstreamTlsContext) GetSessionTicketKeysSdsSecretConfig() *SdsSecretConfig {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig); ok {
		return x.SessionTicketKeysSdsSecretConfig
	}
	return nil
}

func (m *DownstreamTlsContext) GetDisableStatelessSessionResumption() bool {
	if x, ok := m.GetSessionTicketKeysType().(*DownstreamTlsContext_DisableStatelessSessionResumption); ok {
		return x.DisableStatelessSessionResumption
	}
	return false
}

func (m *DownstreamTlsContext) GetSessionTimeout() *duration.Duration {
	if m != nil {
		return m.SessionTimeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DownstreamTlsContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DownstreamTlsContext_SessionTicketKeys)(nil),
		(*DownstreamTlsContext_SessionTicketKeysSdsSecretConfig)(nil),
		(*DownstreamTlsContext_DisableStatelessSessionResumption)(nil),
	}
}

type CommonTlsContext struct {
	TlsParams                      *TlsParameters     `protobuf:"bytes,1,opt,name=tls_params,json=tlsParams,proto3" json:"tls_params,omitempty"`
	TlsCertificates                []*TlsCertificate  `protobuf:"bytes,2,rep,name=tls_certificates,json=tlsCertificates,proto3" json:"tls_certificates,omitempty"`
	TlsCertificateSdsSecretConfigs []*SdsSecretConfig `protobuf:"bytes,6,rep,name=tls_certificate_sds_secret_configs,json=tlsCertificateSdsSecretConfigs,proto3" json:"tls_certificate_sds_secret_configs,omitempty"`
	// Types that are valid to be assigned to ValidationContextType:
	//	*CommonTlsContext_ValidationContext
	//	*CommonTlsContext_ValidationContextSdsSecretConfig
	//	*CommonTlsContext_CombinedValidationContext
	ValidationContextType isCommonTlsContext_ValidationContextType `protobuf_oneof:"validation_context_type"`
	AlpnProtocols         []string                                 `protobuf:"bytes,4,rep,name=alpn_protocols,json=alpnProtocols,proto3" json:"alpn_protocols,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}                                 `json:"-"`
	XXX_unrecognized      []byte                                   `json:"-"`
	XXX_sizecache         int32                                    `json:"-"`
}

func (m *CommonTlsContext) Reset()         { *m = CommonTlsContext{} }
func (m *CommonTlsContext) String() string { return proto.CompactTextString(m) }
func (*CommonTlsContext) ProtoMessage()    {}
func (*CommonTlsContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_463a52fb5d41af5b, []int{2}
}

func (m *CommonTlsContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonTlsContext.Unmarshal(m, b)
}
func (m *CommonTlsContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonTlsContext.Marshal(b, m, deterministic)
}
func (m *CommonTlsContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonTlsContext.Merge(m, src)
}
func (m *CommonTlsContext) XXX_Size() int {
	return xxx_messageInfo_CommonTlsContext.Size(m)
}
func (m *CommonTlsContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonTlsContext.DiscardUnknown(m)
}

var xxx_messageInfo_CommonTlsContext proto.InternalMessageInfo

func (m *CommonTlsContext) GetTlsParams() *TlsParameters {
	if m != nil {
		return m.TlsParams
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificates() []*TlsCertificate {
	if m != nil {
		return m.TlsCertificates
	}
	return nil
}

func (m *CommonTlsContext) GetTlsCertificateSdsSecretConfigs() []*SdsSecretConfig {
	if m != nil {
		return m.TlsCertificateSdsSecretConfigs
	}
	return nil
}

type isCommonTlsContext_ValidationContextType interface {
	isCommonTlsContext_ValidationContextType()
}

type CommonTlsContext_ValidationContext struct {
	ValidationContext *CertificateValidationContext `protobuf:"bytes,3,opt,name=validation_context,json=validationContext,proto3,oneof"`
}

type CommonTlsContext_ValidationContextSdsSecretConfig struct {
	ValidationContextSdsSecretConfig *SdsSecretConfig `protobuf:"bytes,7,opt,name=validation_context_sds_secret_config,json=validationContextSdsSecretConfig,proto3,oneof"`
}

type CommonTlsContext_CombinedValidationContext struct {
	CombinedValidationContext *CommonTlsContext_CombinedCertificateValidationContext `protobuf:"bytes,8,opt,name=combined_validation_context,json=combinedValidationContext,proto3,oneof"`
}

func (*CommonTlsContext_ValidationContext) isCommonTlsContext_ValidationContextType() {}

func (*CommonTlsContext_ValidationContextSdsSecretConfig) isCommonTlsContext_ValidationContextType() {}

func (*CommonTlsContext_CombinedValidationContext) isCommonTlsContext_ValidationContextType() {}

func (m *CommonTlsContext) GetValidationContextType() isCommonTlsContext_ValidationContextType {
	if m != nil {
		return m.ValidationContextType
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContext() *CertificateValidationContext {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_ValidationContext); ok {
		return x.ValidationContext
	}
	return nil
}

func (m *CommonTlsContext) GetValidationContextSdsSecretConfig() *SdsSecretConfig {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_ValidationContextSdsSecretConfig); ok {
		return x.ValidationContextSdsSecretConfig
	}
	return nil
}

func (m *CommonTlsContext) GetCombinedValidationContext() *CommonTlsContext_CombinedCertificateValidationContext {
	if x, ok := m.GetValidationContextType().(*CommonTlsContext_CombinedValidationContext); ok {
		return x.CombinedValidationContext
	}
	return nil
}

func (m *CommonTlsContext) GetAlpnProtocols() []string {
	if m != nil {
		return m.AlpnProtocols
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CommonTlsContext) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CommonTlsContext_ValidationContext)(nil),
		(*CommonTlsContext_ValidationContextSdsSecretConfig)(nil),
		(*CommonTlsContext_CombinedValidationContext)(nil),
	}
}

type CommonTlsContext_CombinedCertificateValidationContext struct {
	DefaultValidationContext         *CertificateValidationContext `protobuf:"bytes,1,opt,name=default_validation_context,json=defaultValidationContext,proto3" json:"default_validation_context,omitempty"`
	ValidationContextSdsSecretConfig *SdsSecretConfig              `protobuf:"bytes,2,opt,name=validation_context_sds_secret_config,json=validationContextSdsSecretConfig,proto3" json:"validation_context_sds_secret_config,omitempty"`
	XXX_NoUnkeyedLiteral             struct{}                      `json:"-"`
	XXX_unrecognized                 []byte                        `json:"-"`
	XXX_sizecache                    int32                         `json:"-"`
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) Reset() {
	*m = CommonTlsContext_CombinedCertificateValidationContext{}
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) String() string {
	return proto.CompactTextString(m)
}
func (*CommonTlsContext_CombinedCertificateValidationContext) ProtoMessage() {}
func (*CommonTlsContext_CombinedCertificateValidationContext) Descriptor() ([]byte, []int) {
	return fileDescriptor_463a52fb5d41af5b, []int{2, 0}
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Unmarshal(m, b)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Marshal(b, m, deterministic)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Merge(m, src)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_Size() int {
	return xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.Size(m)
}
func (m *CommonTlsContext_CombinedCertificateValidationContext) XXX_DiscardUnknown() {
	xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext.DiscardUnknown(m)
}

var xxx_messageInfo_CommonTlsContext_CombinedCertificateValidationContext proto.InternalMessageInfo

func (m *CommonTlsContext_CombinedCertificateValidationContext) GetDefaultValidationContext() *CertificateValidationContext {
	if m != nil {
		return m.DefaultValidationContext
	}
	return nil
}

func (m *CommonTlsContext_CombinedCertificateValidationContext) GetValidationContextSdsSecretConfig() *SdsSecretConfig {
	if m != nil {
		return m.ValidationContextSdsSecretConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*UpstreamTlsContext)(nil), "envoy.api.v2.auth.UpstreamTlsContext")
	proto.RegisterType((*DownstreamTlsContext)(nil), "envoy.api.v2.auth.DownstreamTlsContext")
	proto.RegisterType((*CommonTlsContext)(nil), "envoy.api.v2.auth.CommonTlsContext")
	proto.RegisterType((*CommonTlsContext_CombinedCertificateValidationContext)(nil), "envoy.api.v2.auth.CommonTlsContext.CombinedCertificateValidationContext")
}

func init() { proto.RegisterFile("envoy/api/v2/auth/tls.proto", fileDescriptor_463a52fb5d41af5b) }

var fileDescriptor_463a52fb5d41af5b = []byte{
	// 881 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0x4f, 0x8f, 0xdb, 0x44,
	0x14, 0x5f, 0x27, 0x4d, 0x36, 0x3b, 0xab, 0x6e, 0xdd, 0x29, 0x52, 0xbd, 0x29, 0x94, 0x34, 0x6c,
	0x45, 0x10, 0x92, 0x2d, 0x65, 0x8f, 0x1c, 0x90, 0x9c, 0x0a, 0x2d, 0x7f, 0x0e, 0x5b, 0x67, 0x5b,
	0xc1, 0xc9, 0xcc, 0xda, 0x93, 0x30, 0xaa, 0xed, 0x31, 0xf3, 0xc6, 0xd9, 0x44, 0xe2, 0x90, 0x73,
	0x2f, 0x48, 0x5c, 0xe0, 0xc2, 0x17, 0xe0, 0x23, 0xf0, 0x09, 0xb8, 0xf2, 0x2d, 0x38, 0x73, 0xac,
	0x10, 0xa0, 0x19, 0xdb, 0x9b, 0xd4, 0xb6, 0xb4, 0x81, 0x43, 0x6f, 0xf1, 0xfb, 0xf7, 0x7b, 0xef,
	0xf7, 0x7e, 0xf3, 0x14, 0xf4, 0x80, 0x26, 0x0b, 0xbe, 0x72, 0x48, 0xca, 0x9c, 0xc5, 0xd8, 0x21,
	0x99, 0xfc, 0xc6, 0x91, 0x11, 0xd8, 0xa9, 0xe0, 0x92, 0xe3, 0xbb, 0xda, 0x69, 0x93, 0x94, 0xd9,
	0x8b, 0xb1, 0xad, 0x9c, 0xfd, 0x87, 0xf5, 0xf8, 0x80, 0xc7, 0x31, 0x4f, 0xf2, 0x94, 0x26, 0x3f,
	0xd0, 0x40, 0x50, 0x59, 0xfa, 0xe7, 0x9c, 0xcf, 0x23, 0xea, 0xe8, 0xaf, 0xcb, 0x6c, 0xe6, 0x84,
	0x99, 0x20, 0x92, 0x6d, 0xf2, 0xab, 0xfe, 0x2b, 0x41, 0xd2, 0x94, 0x0a, 0x28, 0xfd, 0x59, 0x98,
	0x12, 0x87, 0x24, 0x09, 0x97, 0x3a, 0x0d, 0x9c, 0x98, 0xcd, 0x05, 0x91, 0xb4, 0xf0, 0xbf, 0x53,
	0xf3, 0x83, 0x24, 0x32, 0x2b, 0xd3, 0xef, 0x2f, 0x48, 0xc4, 0x42, 0x22, 0xa9, 0x53, 0xfe, 0xc8,
	0x1d, 0xc3, 0xbf, 0x0c, 0x84, 0x9f, 0xa5, 0x20, 0x05, 0x25, 0xf1, 0x45, 0x04, 0x13, 0x9e, 0x48,
	0xba, 0x94, 0xf8, 0x29, 0xc2, 0xf9, 0x78, 0xbe, 0x8c, 0xc0, 0x0f, 0x72, 0xab, 0x65, 0x0c, 0x8c,
	0xd1, 0xe1, 0xf8, 0x3d, 0xbb, 0x46, 0x8f, 0x3d, 0xd1, 0xc1, 0x9b, 0x02, 0x9e, 0x19, 0x54, 0x2c,
	0xb8, 0x8f, 0xda, 0x90, 0x30, 0xab, 0x35, 0x30, 0x46, 0x07, 0x6e, 0xef, 0x95, 0xdb, 0x11, 0xed,
	0xd1, 0x3f, 0x86, 0xa7, 0x8c, 0xd8, 0x41, 0xf7, 0x48, 0x14, 0xf1, 0x2b, 0x5f, 0xd0, 0x84, 0xce,
	0xb9, 0x64, 0x7a, 0x06, 0xab, 0x3d, 0x30, 0x46, 0x3d, 0x0f, 0x6b, 0x97, 0xb7, 0xed, 0xc1, 0x9f,
	0x20, 0x33, 0x26, 0x4b, 0x1f, 0x28, 0x00, 0xe3, 0x89, 0xff, 0x82, 0xae, 0xc0, 0xba, 0xa5, 0xbb,
	0x7b, 0xdb, 0xce, 0x99, 0xb4, 0x4b, 0x26, 0xed, 0x67, 0x9f, 0x26, 0xf2, 0x74, 0xfc, 0x9c, 0x44,
	0x19, 0xf5, 0x8e, 0x62, 0xb2, 0x9c, 0xe6, 0x49, 0x9f, 0xd3, 0x15, 0x0c, 0x7f, 0xec, 0xa0, 0xb7,
	0x9e, 0xf0, 0xab, 0xe4, 0x4d, 0x10, 0xf0, 0x25, 0xea, 0x0b, 0xfa, 0x6d, 0xc6, 0x04, 0xf5, 0x83,
	0x88, 0xd1, 0x44, 0xfa, 0x01, 0x15, 0x92, 0xcd, 0x58, 0x40, 0x24, 0xd5, 0xbc, 0x1c, 0x8e, 0xfb,
	0xb5, 0xee, 0x5d, 0xce, 0xa3, 0xbc, 0x77, 0xab, 0xc8, 0x9e, 0xe8, 0xe4, 0xc9, 0x26, 0x17, 0x7f,
	0x84, 0x0e, 0xcb, 0xca, 0x8a, 0xe2, 0xf6, 0x8d, 0xa5, 0x50, 0x11, 0x3e, 0x4d, 0x18, 0xfe, 0x0a,
	0xdd, 0x2b, 0x69, 0x94, 0x2c, 0x78, 0x41, 0xe5, 0x36, 0x9b, 0xef, 0x37, 0x8c, 0x7a, 0x11, 0x41,
	0x41, 0xe1, 0x85, 0x8e, 0x57, 0x44, 0x9e, 0xed, 0x79, 0x77, 0xa1, 0x6a, 0xc4, 0x19, 0x7a, 0xdc,
	0x50, 0xda, 0x87, 0x10, 0xfc, 0xfc, 0x71, 0x28, 0x62, 0x67, 0x6c, 0x6e, 0x75, 0x34, 0xd8, 0xb0,
	0x01, 0x6c, 0x1a, 0xc2, 0x54, 0x87, 0x4e, 0x74, 0xe4, 0xd9, 0x9e, 0x37, 0xa8, 0xe1, 0x54, 0x62,
	0xf0, 0x53, 0x74, 0x12, 0x32, 0x20, 0x97, 0x11, 0xf5, 0xd5, 0x23, 0xa0, 0x11, 0x05, 0xb8, 0x96,
	0x8a, 0xa0, 0x90, 0xc5, 0xa9, 0x96, 0xd7, 0xbe, 0x92, 0xd7, 0xd9, 0x9e, 0xf7, 0xa8, 0x88, 0x9e,
	0x96, 0xc1, 0xc5, 0x78, 0xde, 0x75, 0x28, 0xbe, 0x40, 0x77, 0x36, 0x93, 0xc4, 0x94, 0x67, 0xd2,
	0xea, 0xea, 0x9e, 0x8f, 0x6b, 0x2c, 0x3f, 0x29, 0x1e, 0xb6, 0x6b, 0xbe, 0x72, 0x6f, 0xff, 0x62,
	0xa0, 0x7e, 0xb7, 0xb7, 0x5e, 0xaf, 0xd7, 0xe6, 0x78, 0xcf, 0x3b, 0xba, 0x6e, 0x5d, 0x97, 0x70,
	0xfb, 0xc8, 0x6a, 0xe2, 0x47, 0xae, 0x52, 0x3a, 0xfc, 0x63, 0x1f, 0x99, 0x55, 0x51, 0xe1, 0x8f,
	0x11, 0x52, 0x72, 0x4c, 0x89, 0x20, 0x31, 0x14, 0x6a, 0x1c, 0x34, 0xaf, 0xe8, 0x5c, 0xc5, 0x50,
	0x49, 0x05, 0x78, 0x07, 0xb2, 0xf8, 0x04, 0xfc, 0x05, 0x32, 0xb5, 0x9e, 0x37, 0xe2, 0x01, 0xab,
	0x35, 0x68, 0x8f, 0x0e, 0xc7, 0x8f, 0x9a, 0xcb, 0x6c, 0xc9, 0xcc, 0xbb, 0x23, 0x5f, 0xfb, 0x06,
	0xbc, 0x44, 0xc3, 0x4a, 0xb5, 0xfa, 0x6e, 0xc1, 0xea, 0xea, 0xfa, 0x3b, 0x2c, 0x57, 0x5f, 0x85,
	0x1f, 0x8c, 0x96, 0x69, 0x78, 0x0f, 0x5f, 0x87, 0xaa, 0x04, 0x02, 0xfe, 0x1a, 0xe1, 0xe2, 0x90,
	0x29, 0xf2, 0xca, 0xe7, 0x99, 0x0b, 0xdf, 0x69, 0x7a, 0x9e, 0x9b, 0x5a, 0xcf, 0xaf, 0xf3, 0x0a,
	0x56, 0x95, 0x76, 0x17, 0x55, 0x23, 0x96, 0xe8, 0xa4, 0x8e, 0xd0, 0x20, 0xdd, 0xfd, 0xff, 0x22,
	0xdd, 0x1a, 0x4c, 0x55, 0xba, 0x2f, 0x0d, 0xf4, 0x20, 0xe0, 0xf1, 0x25, 0x4b, 0x68, 0xe8, 0x37,
	0x4c, 0xd8, 0xd3, 0x68, 0x67, 0x3b, 0x1c, 0x20, 0x65, 0xd0, 0x65, 0x6e, 0x18, 0xfd, 0xb8, 0x84,
	0xab, 0x39, 0xf1, 0x63, 0x74, 0x44, 0xa2, 0x34, 0xf1, 0xb5, 0xb4, 0x03, 0x1e, 0xa9, 0xa3, 0xd0,
	0x1e, 0x1d, 0x78, 0xb7, 0x95, 0xf5, 0xbc, 0x34, 0xf6, 0x7f, 0x6e, 0xa1, 0x93, 0x5d, 0xc0, 0xf0,
	0x15, 0xea, 0x87, 0x74, 0x46, 0xb2, 0x48, 0x36, 0x8d, 0x66, 0xfc, 0xaf, 0xe5, 0x69, 0xcd, 0xbc,
	0xd4, 0x9a, 0xb1, 0x8a, 0xe2, 0x75, 0xe0, 0xef, 0x76, 0xdc, 0x65, 0x6b, 0xd7, 0x5d, 0x6e, 0xa1,
	0xde, 0xb8, 0x53, 0xf7, 0x18, 0xdd, 0x6f, 0x40, 0x57, 0x8f, 0xfc, 0xb3, 0x5b, 0xbd, 0x8e, 0xd9,
	0x75, 0x67, 0x7f, 0xfe, 0xf4, 0xf7, 0xf7, 0x9d, 0x0f, 0xf1, 0x07, 0x39, 0x2e, 0x5d, 0x4a, 0x9a,
	0xa8, 0xa3, 0x00, 0xb6, 0x14, 0x24, 0x81, 0x94, 0x0b, 0xe9, 0x03, 0x57, 0xf7, 0x01, 0x6c, 0xf5,
	0x27, 0x65, 0x71, 0xfa, 0xeb, 0xfa, 0xb7, 0xdf, 0xbb, 0x2d, 0xd3, 0x40, 0xef, 0x32, 0x9e, 0x77,
	0x9b, 0x0a, 0xbe, 0x5c, 0xd5, 0x1b, 0x77, 0x7b, 0xea, 0x14, 0xa8, 0x55, 0x9d, 0x1b, 0x97, 0x5d,
	0xbd, 0xc8, 0xd3, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x8f, 0x61, 0xf9, 0x44, 0xfc, 0x08, 0x00,
	0x00,
}
