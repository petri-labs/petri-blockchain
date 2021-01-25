// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: did/did.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type DidCredential struct {
	Credtype []string `protobuf:"bytes,1,rep,name=credtype,proto3" json:"credtype,omitempty" json:"type" yaml:"type"`
	Issuer   string   `protobuf:"bytes,2,opt,name=issuer,proto3" json:"issuer,omitempty" json:"issuer" yaml:"issuer"`
	Issued   string   `protobuf:"bytes,3,opt,name=issued,proto3" json:"issued,omitempty" json:"issued" yaml:"issued"`
	Claim    *Claim   `protobuf:"bytes,4,opt,name=claim,proto3" json:"claim,omitempty" json:"claim" yaml:"claim"`
}

func (m *DidCredential) Reset()         { *m = DidCredential{} }
func (m *DidCredential) String() string { return proto.CompactTextString(m) }
func (*DidCredential) ProtoMessage()    {}
func (*DidCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95066e7350e77d5, []int{0}
}
func (m *DidCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DidCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DidCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DidCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DidCredential.Merge(m, src)
}
func (m *DidCredential) XXX_Size() int {
	return m.Size()
}
func (m *DidCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_DidCredential.DiscardUnknown(m)
}

var xxx_messageInfo_DidCredential proto.InternalMessageInfo

func (m *DidCredential) GetCredtype() []string {
	if m != nil {
		return m.Credtype
	}
	return nil
}

func (m *DidCredential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *DidCredential) GetIssued() string {
	if m != nil {
		return m.Issued
	}
	return ""
}

func (m *DidCredential) GetClaim() *Claim {
	if m != nil {
		return m.Claim
	}
	return nil
}

type Claim struct {
	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" json:"id" yaml:"id"`
	KYCvalidated bool   `protobuf:"varint,2,opt,name=KYCvalidated,proto3" json:"KYCvalidated,omitempty" json:"KYCValidated" yaml:"KYCValidated"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95066e7350e77d5, []int{1}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Claim) GetKYCvalidated() bool {
	if m != nil {
		return m.KYCvalidated
	}
	return false
}

type IxoDid struct {
	Did                 string  `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty" json:"did" yaml:"did"`
	VerifyKey           string  `protobuf:"bytes,2,opt,name=verifyKey,proto3" json:"verifyKey,omitempty" json:"verifyKey" yaml:"verifyKey"`
	EncryptionPublicKey string  `protobuf:"bytes,3,opt,name=encryptionPublicKey,proto3" json:"encryptionPublicKey,omitempty" json:"encryptionPublicKey" yaml:"encryptionPublicKey"`
	Secret              *Secret `protobuf:"bytes,4,opt,name=secret,proto3" json:"secret,omitempty" json:"secret" yaml:"secret"`
}

func (m *IxoDid) Reset()         { *m = IxoDid{} }
func (m *IxoDid) String() string { return proto.CompactTextString(m) }
func (*IxoDid) ProtoMessage()    {}
func (*IxoDid) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95066e7350e77d5, []int{2}
}
func (m *IxoDid) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IxoDid) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IxoDid.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IxoDid) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IxoDid.Merge(m, src)
}
func (m *IxoDid) XXX_Size() int {
	return m.Size()
}
func (m *IxoDid) XXX_DiscardUnknown() {
	xxx_messageInfo_IxoDid.DiscardUnknown(m)
}

var xxx_messageInfo_IxoDid proto.InternalMessageInfo

func (m *IxoDid) GetDid() string {
	if m != nil {
		return m.Did
	}
	return ""
}

func (m *IxoDid) GetVerifyKey() string {
	if m != nil {
		return m.VerifyKey
	}
	return ""
}

func (m *IxoDid) GetEncryptionPublicKey() string {
	if m != nil {
		return m.EncryptionPublicKey
	}
	return ""
}

func (m *IxoDid) GetSecret() *Secret {
	if m != nil {
		return m.Secret
	}
	return nil
}

type Secret struct {
	Seed                 string `protobuf:"bytes,1,opt,name=seed,proto3" json:"seed,omitempty" json:"seed" yaml:"seed"`
	SignKey              string `protobuf:"bytes,2,opt,name=signKey,proto3" json:"signKey,omitempty" json:"signKey" yaml:"signKey"`
	EncryptionPrivateKey string `protobuf:"bytes,3,opt,name=encryptionPrivateKey,proto3" json:"encryptionPrivateKey,omitempty" json:"encryptionPrivateKey" yaml:"encryptionPrivateKey"`
}

func (m *Secret) Reset()         { *m = Secret{} }
func (m *Secret) String() string { return proto.CompactTextString(m) }
func (*Secret) ProtoMessage()    {}
func (*Secret) Descriptor() ([]byte, []int) {
	return fileDescriptor_a95066e7350e77d5, []int{3}
}
func (m *Secret) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Secret) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Secret.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Secret) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Secret.Merge(m, src)
}
func (m *Secret) XXX_Size() int {
	return m.Size()
}
func (m *Secret) XXX_DiscardUnknown() {
	xxx_messageInfo_Secret.DiscardUnknown(m)
}

var xxx_messageInfo_Secret proto.InternalMessageInfo

func (m *Secret) GetSeed() string {
	if m != nil {
		return m.Seed
	}
	return ""
}

func (m *Secret) GetSignKey() string {
	if m != nil {
		return m.SignKey
	}
	return ""
}

func (m *Secret) GetEncryptionPrivateKey() string {
	if m != nil {
		return m.EncryptionPrivateKey
	}
	return ""
}

func init() {
	proto.RegisterType((*DidCredential)(nil), "did.DidCredential")
	proto.RegisterType((*Claim)(nil), "did.Claim")
	proto.RegisterType((*IxoDid)(nil), "did.IxoDid")
	proto.RegisterType((*Secret)(nil), "did.Secret")
}

func init() { proto.RegisterFile("did/did.proto", fileDescriptor_a95066e7350e77d5) }

var fileDescriptor_a95066e7350e77d5 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x94, 0x41, 0x6b, 0xdb, 0x30,
	0x14, 0xc7, 0xe3, 0xa4, 0xcd, 0x1a, 0x75, 0xbd, 0xa8, 0x1d, 0x4d, 0x57, 0x6a, 0xa5, 0x86, 0xd1,
	0x40, 0x59, 0x0c, 0x1b, 0x23, 0x6c, 0x1d, 0x0c, 0x92, 0x5c, 0x46, 0x0e, 0x1b, 0x1a, 0x14, 0xba,
	0x9b, 0xe3, 0xa7, 0xa6, 0x5a, 0x1d, 0x2b, 0xd8, 0x4e, 0x48, 0x8e, 0xfb, 0x06, 0xfb, 0x22, 0xfb,
	0x1e, 0x3b, 0xf6, 0xb8, 0x93, 0x19, 0xc9, 0x61, 0x77, 0x5f, 0x77, 0x19, 0xb2, 0x64, 0x3b, 0xa1,
	0xbe, 0xe9, 0xbd, 0xf7, 0xff, 0xfd, 0x25, 0xbd, 0x67, 0x19, 0x1d, 0x00, 0x07, 0x1b, 0x38, 0x74,
	0xa6, 0x81, 0x88, 0x04, 0xae, 0x01, 0x87, 0xe7, 0x47, 0x63, 0x31, 0x16, 0x69, 0x6c, 0xcb, 0x95,
	0x2a, 0x59, 0xff, 0x0c, 0x74, 0x30, 0xe0, 0xd0, 0x0f, 0x18, 0x30, 0x3f, 0xe2, 0x8e, 0x87, 0xbb,
	0x68, 0xcf, 0x0d, 0x18, 0x44, 0xcb, 0x29, 0x6b, 0x1a, 0xad, 0x5a, 0xbb, 0xd1, 0x3b, 0x4d, 0x62,
	0x72, 0xfc, 0x2d, 0x14, 0xfe, 0x3b, 0x4b, 0x66, 0xad, 0xd6, 0xd2, 0x99, 0x78, 0x7a, 0x4d, 0x73,
	0x31, 0xee, 0xa2, 0x3a, 0x0f, 0xc3, 0x19, 0x0b, 0x9a, 0xd5, 0x96, 0xd1, 0x6e, 0xf4, 0x48, 0x12,
	0x93, 0x53, 0x85, 0xa9, 0x7c, 0x06, 0xea, 0x88, 0x6a, 0x79, 0x0e, 0x42, 0xb3, 0x56, 0x0a, 0xc2,
	0x16, 0x08, 0x19, 0x08, 0xf8, 0x03, 0xda, 0x75, 0x3d, 0x87, 0x4f, 0x9a, 0x3b, 0x2d, 0xa3, 0xbd,
	0xff, 0x0a, 0x75, 0xe4, 0x95, 0xfb, 0x32, 0xd3, 0x3b, 0x4b, 0x62, 0x72, 0xa2, 0x3c, 0x52, 0x49,
	0x66, 0xa1, 0x02, 0xaa, 0x38, 0xeb, 0xbb, 0x81, 0x76, 0x53, 0x3d, 0xbe, 0x40, 0x55, 0x0e, 0x4d,
	0x23, 0xdd, 0xff, 0x38, 0x89, 0xc9, 0xa1, 0xde, 0xbf, 0xd8, 0x1b, 0x2c, 0x5a, 0xe5, 0x80, 0x3f,
	0xa1, 0xa7, 0xc3, 0x9b, 0xfe, 0xdc, 0xf1, 0x38, 0x38, 0x11, 0x83, 0xf4, 0xae, 0x7b, 0xbd, 0xcb,
	0x24, 0x26, 0x17, 0x0a, 0x19, 0xde, 0xf4, 0xaf, 0xb3, 0x6a, 0x06, 0x6f, 0xe5, 0xe8, 0x96, 0x81,
	0xf5, 0xb3, 0x8a, 0xea, 0x1f, 0x17, 0x62, 0xc0, 0x01, 0x5f, 0x22, 0x39, 0x29, 0x7d, 0x8a, 0x93,
	0x24, 0x26, 0xcf, 0x94, 0x25, 0x14, 0xc7, 0x90, 0x4b, 0x2a, 0x55, 0xb8, 0x8f, 0x1a, 0x73, 0x16,
	0xf0, 0xdb, 0xe5, 0x90, 0x2d, 0x75, 0xc7, 0x5f, 0x24, 0x31, 0x39, 0x57, 0x48, 0x5e, 0xca, 0xc0,
	0x22, 0x41, 0x0b, 0x0e, 0xdf, 0xa3, 0x43, 0xe6, 0xbb, 0xc1, 0x72, 0x1a, 0x71, 0xe1, 0x7f, 0x9e,
	0x8d, 0x3c, 0xee, 0x4a, 0x3b, 0x35, 0x87, 0xb7, 0x49, 0x4c, 0xde, 0x28, 0xbb, 0x12, 0x51, 0x66,
	0x5c, 0x56, 0xa2, 0x65, 0xae, 0x78, 0x80, 0xea, 0x21, 0x73, 0x03, 0x16, 0xe9, 0x79, 0xed, 0xa7,
	0xf3, 0xfa, 0x92, 0xa6, 0x36, 0x87, 0xae, 0x44, 0x99, 0xbf, 0x8e, 0xa8, 0x66, 0xad, 0xbf, 0x06,
	0xaa, 0x2b, 0x06, 0xdb, 0x68, 0x27, 0x64, 0x2c, 0x6b, 0xd8, 0xc6, 0x67, 0x2a, 0xb3, 0x05, 0x2f,
	0x7b, 0x9e, 0x0a, 0xf1, 0x15, 0x7a, 0x12, 0xf2, 0xb1, 0x5f, 0x74, 0xec, 0x3c, 0x89, 0xc9, 0x99,
	0x66, 0x54, 0x21, 0xc7, 0x74, 0x48, 0x33, 0x02, 0x0b, 0x74, 0xb4, 0x71, 0xab, 0x80, 0xcf, 0x9d,
	0x88, 0x15, 0xcd, 0xba, 0x4a, 0x62, 0xd2, 0x7d, 0xd4, 0xac, 0x5c, 0x55, 0xd2, 0xad, 0xa2, 0x46,
	0x4b, 0x8d, 0x7b, 0xd7, 0xbf, 0x56, 0xa6, 0xf1, 0xb0, 0x32, 0x8d, 0x3f, 0x2b, 0xd3, 0xf8, 0xb1,
	0x36, 0x2b, 0x0f, 0x6b, 0xb3, 0xf2, 0x7b, 0x6d, 0x56, 0xbe, 0xbe, 0x1f, 0xf3, 0xe8, 0x6e, 0x36,
	0xea, 0xb8, 0x62, 0x62, 0xf3, 0x85, 0xb8, 0x15, 0x33, 0x1f, 0x1c, 0x49, 0xcb, 0xe8, 0xe5, 0xc8,
	0x13, 0xee, 0xbd, 0x7b, 0xe7, 0x70, 0xdf, 0x5e, 0xc8, 0xbf, 0x80, 0xcd, 0xfd, 0x88, 0x05, 0xbe,
	0xe3, 0xd9, 0xf2, 0x9d, 0x86, 0xa3, 0x7a, 0xfa, 0xf4, 0x5f, 0xff, 0x0f, 0x00, 0x00, 0xff, 0xff,
	0x18, 0x0b, 0x3d, 0x02, 0x26, 0x04, 0x00, 0x00,
}

func (m *DidCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DidCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DidCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Claim != nil {
		{
			size, err := m.Claim.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.Issued) > 0 {
		i -= len(m.Issued)
		copy(dAtA[i:], m.Issued)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Issued)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Credtype) > 0 {
		for iNdEx := len(m.Credtype) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Credtype[iNdEx])
			copy(dAtA[i:], m.Credtype[iNdEx])
			i = encodeVarintDid(dAtA, i, uint64(len(m.Credtype[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.KYCvalidated {
		i--
		if m.KYCvalidated {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IxoDid) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IxoDid) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IxoDid) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Secret != nil {
		{
			size, err := m.Secret.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDid(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if len(m.EncryptionPublicKey) > 0 {
		i -= len(m.EncryptionPublicKey)
		copy(dAtA[i:], m.EncryptionPublicKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.EncryptionPublicKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.VerifyKey) > 0 {
		i -= len(m.VerifyKey)
		copy(dAtA[i:], m.VerifyKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.VerifyKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Secret) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Secret) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Secret) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.EncryptionPrivateKey) > 0 {
		i -= len(m.EncryptionPrivateKey)
		copy(dAtA[i:], m.EncryptionPrivateKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.EncryptionPrivateKey)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SignKey) > 0 {
		i -= len(m.SignKey)
		copy(dAtA[i:], m.SignKey)
		i = encodeVarintDid(dAtA, i, uint64(len(m.SignKey)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Seed) > 0 {
		i -= len(m.Seed)
		copy(dAtA[i:], m.Seed)
		i = encodeVarintDid(dAtA, i, uint64(len(m.Seed)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDid(dAtA []byte, offset int, v uint64) int {
	offset -= sovDid(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DidCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Credtype) > 0 {
		for _, s := range m.Credtype {
			l = len(s)
			n += 1 + l + sovDid(uint64(l))
		}
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.Issued)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.Claim != nil {
		l = m.Claim.Size()
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.KYCvalidated {
		n += 2
	}
	return n
}

func (m *IxoDid) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.VerifyKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.EncryptionPublicKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	if m.Secret != nil {
		l = m.Secret.Size()
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func (m *Secret) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Seed)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.SignKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	l = len(m.EncryptionPrivateKey)
	if l > 0 {
		n += 1 + l + sovDid(uint64(l))
	}
	return n
}

func sovDid(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDid(x uint64) (n int) {
	return sovDid(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DidCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DidCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DidCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Credtype", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Credtype = append(m.Credtype, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issued", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issued = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claim", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Claim == nil {
				m.Claim = &Claim{}
			}
			if err := m.Claim.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field KYCvalidated", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.KYCvalidated = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IxoDid) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IxoDid: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IxoDid: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Did = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifyKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerifyKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionPublicKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionPublicKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secret", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Secret == nil {
				m.Secret = &Secret{}
			}
			if err := m.Secret.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Secret) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDid
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Secret: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Secret: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seed", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seed = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SignKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SignKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EncryptionPrivateKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDid
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthDid
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDid
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EncryptionPrivateKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDid(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthDid
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipDid(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDid
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowDid
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthDid
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDid
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDid
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDid        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDid          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDid = fmt.Errorf("proto: unexpected end of group")
)
