// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: side/btclightclient/bitcoin.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Bitcoin Signing Status
type SigningStatus int32

const (
	// SIGNING_STATUS_UNSPECIFIED - Default value, should not be used
	SigningStatus_SIGNING_STATUS_UNSPECIFIED SigningStatus = 0
	// SIGNING_STATUS_CREATED - The signing request is created
	SigningStatus_SIGNING_STATUS_CREATED SigningStatus = 1
	// SIGNING_STATUS_SIGNED - The signing request is signed
	SigningStatus_SIGNING_STATUS_SIGNED SigningStatus = 2
	// SIGNING_STATUS_BROADCASTED - The signing request is broadcasted
	SigningStatus_SIGNING_STATUS_BROADCASTED SigningStatus = 3
	// SIGNING_STATUS_CONFIRMED - The signing request is confirmed
	SigningStatus_SIGNING_STATUS_CONFIRMED SigningStatus = 4
	// SIGNING_STATUS_REJECTED - The signing request is rejected
	SigningStatus_SIGNING_STATUS_REJECTED SigningStatus = 5
)

var SigningStatus_name = map[int32]string{
	0: "SIGNING_STATUS_UNSPECIFIED",
	1: "SIGNING_STATUS_CREATED",
	2: "SIGNING_STATUS_SIGNED",
	3: "SIGNING_STATUS_BROADCASTED",
	4: "SIGNING_STATUS_CONFIRMED",
	5: "SIGNING_STATUS_REJECTED",
}

var SigningStatus_value = map[string]int32{
	"SIGNING_STATUS_UNSPECIFIED": 0,
	"SIGNING_STATUS_CREATED":     1,
	"SIGNING_STATUS_SIGNED":      2,
	"SIGNING_STATUS_BROADCASTED": 3,
	"SIGNING_STATUS_CONFIRMED":   4,
	"SIGNING_STATUS_REJECTED":    5,
}

func (x SigningStatus) String() string {
	return proto.EnumName(SigningStatus_name, int32(x))
}

func (SigningStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_d7dee3af17ecec77, []int{0}
}

// Bitcoin Block Header
type BlockHeader struct {
	Version           uint64 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	Hash              string `protobuf:"bytes,2,opt,name=hash,proto3" json:"hash,omitempty"`
	Height            uint64 `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	PreviousBlockHash string `protobuf:"bytes,4,opt,name=previous_block_hash,json=previousBlockHash,proto3" json:"previous_block_hash,omitempty"`
	MerkleRoot        string `protobuf:"bytes,5,opt,name=merkle_root,json=merkleRoot,proto3" json:"merkle_root,omitempty"`
	Nonce             uint64 `protobuf:"varint,6,opt,name=nonce,proto3" json:"nonce,omitempty"`
	Bits              string `protobuf:"bytes,7,opt,name=bits,proto3" json:"bits,omitempty"`
	Time              uint64 `protobuf:"varint,8,opt,name=time,proto3" json:"time,omitempty"`
	Ntx               uint64 `protobuf:"varint,9,opt,name=ntx,proto3" json:"ntx,omitempty"`
}

func (m *BlockHeader) Reset()         { *m = BlockHeader{} }
func (m *BlockHeader) String() string { return proto.CompactTextString(m) }
func (*BlockHeader) ProtoMessage()    {}
func (*BlockHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7dee3af17ecec77, []int{0}
}
func (m *BlockHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BlockHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BlockHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BlockHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BlockHeader.Merge(m, src)
}
func (m *BlockHeader) XXX_Size() int {
	return m.Size()
}
func (m *BlockHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_BlockHeader.DiscardUnknown(m)
}

var xxx_messageInfo_BlockHeader proto.InternalMessageInfo

func (m *BlockHeader) GetVersion() uint64 {
	if m != nil {
		return m.Version
	}
	return 0
}

func (m *BlockHeader) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *BlockHeader) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *BlockHeader) GetPreviousBlockHash() string {
	if m != nil {
		return m.PreviousBlockHash
	}
	return ""
}

func (m *BlockHeader) GetMerkleRoot() string {
	if m != nil {
		return m.MerkleRoot
	}
	return ""
}

func (m *BlockHeader) GetNonce() uint64 {
	if m != nil {
		return m.Nonce
	}
	return 0
}

func (m *BlockHeader) GetBits() string {
	if m != nil {
		return m.Bits
	}
	return ""
}

func (m *BlockHeader) GetTime() uint64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *BlockHeader) GetNtx() uint64 {
	if m != nil {
		return m.Ntx
	}
	return 0
}

// Bitcoin Signing Request
type BitcoinSigningRequest struct {
	Address  string        `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	TxBytes  string        `protobuf:"bytes,2,opt,name=tx_bytes,json=txBytes,proto3" json:"tx_bytes,omitempty"`
	Status   SigningStatus `protobuf:"varint,3,opt,name=status,proto3,enum=side.btclightclient.SigningStatus" json:"status,omitempty"`
	Sequence uint64        `protobuf:"varint,4,opt,name=sequence,proto3" json:"sequence,omitempty"`
	// The vault address that the request is associated with
	VaultAddress string `protobuf:"bytes,5,opt,name=vault_address,json=vaultAddress,proto3" json:"vault_address,omitempty"`
}

func (m *BitcoinSigningRequest) Reset()         { *m = BitcoinSigningRequest{} }
func (m *BitcoinSigningRequest) String() string { return proto.CompactTextString(m) }
func (*BitcoinSigningRequest) ProtoMessage()    {}
func (*BitcoinSigningRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_d7dee3af17ecec77, []int{1}
}
func (m *BitcoinSigningRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BitcoinSigningRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BitcoinSigningRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BitcoinSigningRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BitcoinSigningRequest.Merge(m, src)
}
func (m *BitcoinSigningRequest) XXX_Size() int {
	return m.Size()
}
func (m *BitcoinSigningRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BitcoinSigningRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BitcoinSigningRequest proto.InternalMessageInfo

func (m *BitcoinSigningRequest) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *BitcoinSigningRequest) GetTxBytes() string {
	if m != nil {
		return m.TxBytes
	}
	return ""
}

func (m *BitcoinSigningRequest) GetStatus() SigningStatus {
	if m != nil {
		return m.Status
	}
	return SigningStatus_SIGNING_STATUS_UNSPECIFIED
}

func (m *BitcoinSigningRequest) GetSequence() uint64 {
	if m != nil {
		return m.Sequence
	}
	return 0
}

func (m *BitcoinSigningRequest) GetVaultAddress() string {
	if m != nil {
		return m.VaultAddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("side.btclightclient.SigningStatus", SigningStatus_name, SigningStatus_value)
	proto.RegisterType((*BlockHeader)(nil), "side.btclightclient.BlockHeader")
	proto.RegisterType((*BitcoinSigningRequest)(nil), "side.btclightclient.BitcoinSigningRequest")
}

func init() { proto.RegisterFile("side/btclightclient/bitcoin.proto", fileDescriptor_d7dee3af17ecec77) }

var fileDescriptor_d7dee3af17ecec77 = []byte{
	// 508 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x53, 0xcb, 0x6e, 0xd3, 0x40,
	0x14, 0xcd, 0x34, 0xef, 0x5b, 0x8a, 0xc2, 0xf4, 0xc1, 0x34, 0x20, 0x53, 0xc2, 0xa6, 0x62, 0xe1,
	0x48, 0x74, 0xc7, 0x2e, 0x0f, 0xb7, 0x04, 0xa9, 0x29, 0x1a, 0xa7, 0x1b, 0x36, 0x56, 0xec, 0x8c,
	0xec, 0x51, 0x1d, 0x4f, 0xf0, 0x8c, 0xa3, 0xf4, 0x2f, 0xf8, 0x25, 0x76, 0xb0, 0xeb, 0x92, 0x25,
	0x4a, 0x7e, 0x81, 0x0f, 0x40, 0x33, 0x4e, 0x90, 0x6a, 0x75, 0x13, 0xdd, 0x73, 0xef, 0xb9, 0xe7,
	0xdc, 0x1c, 0xdb, 0xf0, 0x56, 0xf2, 0x19, 0xeb, 0xfa, 0x2a, 0x88, 0x79, 0x18, 0xe9, 0x5f, 0x96,
	0xa8, 0xae, 0xcf, 0x55, 0x20, 0x78, 0x62, 0x2f, 0x52, 0xa1, 0x04, 0x3e, 0xd4, 0x14, 0xfb, 0x31,
	0xa5, 0x7d, 0x14, 0x8a, 0x50, 0x98, 0x79, 0x57, 0x57, 0x39, 0xb5, 0xf3, 0x17, 0xc1, 0x7e, 0x3f,
	0x16, 0xc1, 0xdd, 0x27, 0x36, 0x9d, 0xb1, 0x14, 0x13, 0xa8, 0x2f, 0x59, 0x2a, 0xb9, 0x48, 0x08,
	0x3a, 0x43, 0xe7, 0x15, 0xba, 0x83, 0x18, 0x43, 0x25, 0x9a, 0xca, 0x88, 0xec, 0x9d, 0xa1, 0xf3,
	0x26, 0x35, 0x35, 0x3e, 0x81, 0x5a, 0xc4, 0xb4, 0x07, 0x29, 0x1b, 0xf2, 0x16, 0x61, 0x1b, 0x0e,
	0x17, 0x29, 0x5b, 0x72, 0x91, 0x49, 0xcf, 0xd7, 0xea, 0x9e, 0x59, 0xad, 0x98, 0xd5, 0x17, 0xbb,
	0x51, 0xee, 0xab, 0x75, 0xde, 0xc0, 0xfe, 0x9c, 0xa5, 0x77, 0x31, 0xf3, 0x52, 0x21, 0x14, 0xa9,
	0x1a, 0x1e, 0xe4, 0x2d, 0x2a, 0x84, 0xc2, 0x47, 0x50, 0x4d, 0x44, 0x12, 0x30, 0x52, 0x33, 0x3e,
	0x39, 0xd0, 0x27, 0xf9, 0x5c, 0x49, 0x52, 0xcf, 0x4f, 0xd2, 0xb5, 0xee, 0x29, 0x3e, 0x67, 0xa4,
	0x61, 0x88, 0xa6, 0xc6, 0x2d, 0x28, 0x27, 0x6a, 0x45, 0x9a, 0xa6, 0xa5, 0xcb, 0xce, 0x2f, 0x04,
	0xc7, 0xfd, 0x3c, 0x33, 0x97, 0x87, 0x09, 0x4f, 0x42, 0xca, 0xbe, 0x65, 0x4c, 0x2a, 0x1d, 0xc0,
	0x74, 0x36, 0x4b, 0x99, 0x94, 0x26, 0x80, 0x26, 0xdd, 0x41, 0x7c, 0x0a, 0x0d, 0xb5, 0xf2, 0xfc,
	0x7b, 0xc5, 0xe4, 0x36, 0x84, 0xba, 0x5a, 0xf5, 0x35, 0xc4, 0x1f, 0xa1, 0x26, 0xd5, 0x54, 0x65,
	0xd2, 0xe4, 0xf0, 0xfc, 0x43, 0xc7, 0x7e, 0xe2, 0x09, 0xd8, 0x5b, 0x27, 0xd7, 0x30, 0xe9, 0x76,
	0x03, 0xb7, 0xa1, 0x21, 0xb5, 0xb7, 0xfe, 0x77, 0x15, 0x73, 0xe1, 0x7f, 0x8c, 0xdf, 0xc1, 0xc1,
	0x72, 0x9a, 0xc5, 0xca, 0xdb, 0x9d, 0x94, 0x27, 0xf3, 0xcc, 0x34, 0x7b, 0x79, 0xef, 0xfd, 0x0f,
	0x04, 0x07, 0x8f, 0xa4, 0xb1, 0x05, 0x6d, 0x77, 0x74, 0x35, 0x1e, 0x8d, 0xaf, 0x3c, 0x77, 0xd2,
	0x9b, 0xdc, 0xba, 0xde, 0xed, 0xd8, 0xfd, 0xe2, 0x0c, 0x46, 0x97, 0x23, 0x67, 0xd8, 0x2a, 0xe1,
	0x36, 0x9c, 0x14, 0xe6, 0x03, 0xea, 0xf4, 0x26, 0xce, 0xb0, 0x85, 0xf0, 0x29, 0x1c, 0x17, 0x66,
	0x1a, 0x3a, 0xc3, 0xd6, 0xde, 0x13, 0xb2, 0x7d, 0x7a, 0xd3, 0x1b, 0x0e, 0x7a, 0xae, 0x5e, 0x2d,
	0xe3, 0xd7, 0x40, 0x8a, 0xb2, 0x37, 0xe3, 0xcb, 0x11, 0xbd, 0x76, 0x86, 0xad, 0x0a, 0x7e, 0x05,
	0x2f, 0x0b, 0x53, 0xea, 0x7c, 0x76, 0x06, 0x7a, 0xb5, 0xda, 0xbf, 0xfe, 0xb9, 0xb6, 0xd0, 0xc3,
	0xda, 0x42, 0x7f, 0xd6, 0x16, 0xfa, 0xbe, 0xb1, 0x4a, 0x0f, 0x1b, 0xab, 0xf4, 0x7b, 0x63, 0x95,
	0xbe, 0x5e, 0x84, 0x5c, 0x45, 0x99, 0x6f, 0x07, 0x62, 0xde, 0xd5, 0xa1, 0x9a, 0xd7, 0x36, 0x10,
	0xb1, 0x01, 0xdd, 0x55, 0xf1, 0x43, 0x50, 0xf7, 0x0b, 0x26, 0xfd, 0x9a, 0x61, 0x5d, 0xfc, 0x0b,
	0x00, 0x00, 0xff, 0xff, 0x2c, 0x90, 0x12, 0x79, 0x2c, 0x03, 0x00, 0x00,
}

func (m *BlockHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BlockHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BlockHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Ntx != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Ntx))
		i--
		dAtA[i] = 0x48
	}
	if m.Time != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Time))
		i--
		dAtA[i] = 0x40
	}
	if len(m.Bits) > 0 {
		i -= len(m.Bits)
		copy(dAtA[i:], m.Bits)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.Bits)))
		i--
		dAtA[i] = 0x3a
	}
	if m.Nonce != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Nonce))
		i--
		dAtA[i] = 0x30
	}
	if len(m.MerkleRoot) > 0 {
		i -= len(m.MerkleRoot)
		copy(dAtA[i:], m.MerkleRoot)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.MerkleRoot)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.PreviousBlockHash) > 0 {
		i -= len(m.PreviousBlockHash)
		copy(dAtA[i:], m.PreviousBlockHash)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.PreviousBlockHash)))
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x12
	}
	if m.Version != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Version))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *BitcoinSigningRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BitcoinSigningRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BitcoinSigningRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.VaultAddress) > 0 {
		i -= len(m.VaultAddress)
		copy(dAtA[i:], m.VaultAddress)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.VaultAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if m.Sequence != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Sequence))
		i--
		dAtA[i] = 0x20
	}
	if m.Status != 0 {
		i = encodeVarintBitcoin(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TxBytes) > 0 {
		i -= len(m.TxBytes)
		copy(dAtA[i:], m.TxBytes)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.TxBytes)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintBitcoin(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBitcoin(dAtA []byte, offset int, v uint64) int {
	offset -= sovBitcoin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BlockHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovBitcoin(uint64(m.Version))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovBitcoin(uint64(m.Height))
	}
	l = len(m.PreviousBlockHash)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	l = len(m.MerkleRoot)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	if m.Nonce != 0 {
		n += 1 + sovBitcoin(uint64(m.Nonce))
	}
	l = len(m.Bits)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	if m.Time != 0 {
		n += 1 + sovBitcoin(uint64(m.Time))
	}
	if m.Ntx != 0 {
		n += 1 + sovBitcoin(uint64(m.Ntx))
	}
	return n
}

func (m *BitcoinSigningRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	l = len(m.TxBytes)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovBitcoin(uint64(m.Status))
	}
	if m.Sequence != 0 {
		n += 1 + sovBitcoin(uint64(m.Sequence))
	}
	l = len(m.VaultAddress)
	if l > 0 {
		n += 1 + l + sovBitcoin(uint64(l))
	}
	return n
}

func sovBitcoin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBitcoin(x uint64) (n int) {
	return sovBitcoin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BlockHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoin
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
			return fmt.Errorf("proto: BlockHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BlockHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PreviousBlockHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PreviousBlockHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MerkleRoot", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MerkleRoot = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nonce", wireType)
			}
			m.Nonce = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nonce |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Bits", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Bits = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			m.Time = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Time |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Ntx", wireType)
			}
			m.Ntx = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Ntx |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBitcoin
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
func (m *BitcoinSigningRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBitcoin
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
			return fmt.Errorf("proto: BitcoinSigningRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BitcoinSigningRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TxBytes", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TxBytes = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= SigningStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sequence", wireType)
			}
			m.Sequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Sequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VaultAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBitcoin
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
				return ErrInvalidLengthBitcoin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBitcoin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VaultAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBitcoin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBitcoin
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
func skipBitcoin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBitcoin
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
					return 0, ErrIntOverflowBitcoin
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
					return 0, ErrIntOverflowBitcoin
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
				return 0, ErrInvalidLengthBitcoin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBitcoin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBitcoin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBitcoin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBitcoin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBitcoin = fmt.Errorf("proto: unexpected end of group")
)