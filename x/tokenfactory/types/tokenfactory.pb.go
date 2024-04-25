// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: osmosis/tokenfactory/v1beta1/tokenfactory.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
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

// DenomAuthorityMetadata specifies metadata for addresses that have specific
// capabilities over a token factory denom. Right now there is only one Admin
// permission, but is planned to be extended to the future.
type DenomAuthorityMetadata struct {
	// Can be empty for no admin, or a valid stargaze address
	Admin string `protobuf:"bytes,1,opt,name=admin,proto3" json:"admin,omitempty" yaml:"admin"`
}

func (m *DenomAuthorityMetadata) Reset()         { *m = DenomAuthorityMetadata{} }
func (m *DenomAuthorityMetadata) String() string { return proto.CompactTextString(m) }
func (*DenomAuthorityMetadata) ProtoMessage()    {}
func (*DenomAuthorityMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce345928cf8ce182, []int{0}
}
func (m *DenomAuthorityMetadata) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomAuthorityMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomAuthorityMetadata.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomAuthorityMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomAuthorityMetadata.Merge(m, src)
}
func (m *DenomAuthorityMetadata) XXX_Size() int {
	return m.Size()
}
func (m *DenomAuthorityMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomAuthorityMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_DenomAuthorityMetadata proto.InternalMessageInfo

func (m *DenomAuthorityMetadata) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

// Params defines the parameters for the tokenfactory module.
type Params struct {
	// DenomCreationFee defines the fee to be charged on the creation of a new
	// denom. The fee is drawn from the MsgCreateDenom's sender account, and
	// transferred to the community pool.
	DenomCreationFee github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,1,rep,name=denom_creation_fee,json=denomCreationFee,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"denom_creation_fee" yaml:"denom_creation_fee"`
	// DenomCreationGasConsume defines the gas cost for creating a new denom.
	// This is intended as a spam deterrence mechanism.
	//
	// See: https://github.com/CosmWasm/token-factory/issues/11
	DenomCreationGasConsume uint64 `protobuf:"varint,2,opt,name=denom_creation_gas_consume,json=denomCreationGasConsume,proto3" json:"denom_creation_gas_consume,omitempty" yaml:"denom_creation_gas_consume"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce345928cf8ce182, []int{1}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetDenomCreationFee() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.DenomCreationFee
	}
	return nil
}

func (m *Params) GetDenomCreationGasConsume() uint64 {
	if m != nil {
		return m.DenomCreationGasConsume
	}
	return 0
}

func init() {
	proto.RegisterType((*DenomAuthorityMetadata)(nil), "osmosis.tokenfactory.v1beta1.DenomAuthorityMetadata")
	proto.RegisterType((*Params)(nil), "osmosis.tokenfactory.v1beta1.Params")
}

func init() {
	proto.RegisterFile("osmosis/tokenfactory/v1beta1/tokenfactory.proto", fileDescriptor_ce345928cf8ce182)
}

var fileDescriptor_ce345928cf8ce182 = []byte{
	// 396 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x52, 0xb1, 0xae, 0xd3, 0x30,
	0x14, 0x8d, 0x1f, 0x8f, 0x27, 0x11, 0x18, 0x9e, 0x22, 0x04, 0x6d, 0x85, 0x92, 0x92, 0x01, 0x85,
	0xe1, 0xc5, 0x2a, 0x30, 0x95, 0x89, 0x14, 0x95, 0xa9, 0x12, 0xaa, 0xc4, 0xc2, 0x12, 0xdd, 0x24,
	0x6e, 0x6a, 0xb5, 0xce, 0xad, 0x62, 0x07, 0x08, 0x7f, 0xc0, 0xc6, 0xc4, 0xcc, 0xcc, 0x97, 0x74,
	0xec, 0xc8, 0x14, 0x50, 0xbb, 0x30, 0xf7, 0x0b, 0x50, 0xe2, 0x80, 0x1a, 0x78, 0x93, 0x7d, 0xcf,
	0xf1, 0x39, 0xf7, 0x1e, 0x5d, 0x9b, 0x14, 0xa5, 0x40, 0xc9, 0x25, 0x55, 0xb8, 0x62, 0xd9, 0x02,
	0x62, 0x85, 0x79, 0x49, 0xdf, 0x8d, 0x22, 0xa6, 0x60, 0xd4, 0x01, 0xfd, 0x4d, 0x8e, 0x0a, 0xad,
	0x07, 0xad, 0xc0, 0xef, 0x70, 0xad, 0x60, 0x60, 0xc7, 0x0d, 0x4d, 0x23, 0x90, 0xec, 0xaf, 0x4b,
	0x8c, 0x3c, 0xd3, 0xea, 0x41, 0x5f, 0xf3, 0x61, 0x53, 0x51, 0x5d, 0xb4, 0xd4, 0xdd, 0x14, 0x53,
	0xd4, 0x78, 0x7d, 0xd3, 0xa8, 0x3b, 0x35, 0xef, 0xbd, 0x64, 0x19, 0x8a, 0x17, 0x85, 0x5a, 0x62,
	0xce, 0x55, 0x39, 0x63, 0x0a, 0x12, 0x50, 0x60, 0x3d, 0x32, 0x6f, 0x42, 0x22, 0x78, 0xd6, 0x23,
	0x43, 0xe2, 0xdd, 0x0a, 0x2e, 0x8f, 0x95, 0x73, 0xa7, 0x04, 0xb1, 0x1e, 0xbb, 0x0d, 0xec, 0xce,
	0x35, 0x3d, 0x3e, 0xff, 0xf5, 0xd5, 0x21, 0xee, 0xa7, 0x33, 0xf3, 0xe2, 0x35, 0xe4, 0x20, 0xa4,
	0xf5, 0x85, 0x98, 0x56, 0x52, 0x7b, 0x86, 0x71, 0xce, 0x40, 0x71, 0xcc, 0xc2, 0x05, 0x63, 0x3d,
	0x32, 0xbc, 0xe1, 0xdd, 0x7e, 0xd2, 0xf7, 0xdb, 0xa1, 0xea, 0x04, 0x7f, 0x62, 0xf9, 0x13, 0xe4,
	0x59, 0x30, 0xdb, 0x56, 0x8e, 0x71, 0xac, 0x9c, 0xbe, 0xee, 0xf2, 0xbf, 0x85, 0xfb, 0xed, 0x87,
	0xe3, 0xa5, 0x5c, 0x2d, 0x8b, 0xc8, 0x8f, 0x51, 0xb4, 0xf1, 0xda, 0xe3, 0x4a, 0x26, 0x2b, 0xaa,
	0xca, 0x0d, 0x93, 0x8d, 0x9b, 0x9c, 0x5f, 0x36, 0x06, 0x93, 0x56, 0x3f, 0x65, 0xcc, 0x5a, 0x98,
	0x83, 0x7f, 0x4c, 0x53, 0x90, 0x61, 0x8c, 0x99, 0x2c, 0x04, 0xeb, 0x9d, 0x0d, 0x89, 0x77, 0x1e,
	0x3c, 0xde, 0x56, 0x0e, 0x39, 0x56, 0xce, 0xc3, 0x6b, 0x87, 0x38, 0x79, 0xef, 0xce, 0xef, 0x77,
	0x1a, 0xbc, 0x02, 0x39, 0xd1, 0x4c, 0xf0, 0x66, 0xbb, 0xb7, 0xc9, 0x6e, 0x6f, 0x93, 0x9f, 0x7b,
	0x9b, 0x7c, 0x3e, 0xd8, 0xc6, 0xee, 0x60, 0x1b, 0xdf, 0x0f, 0xb6, 0xf1, 0xf6, 0xf9, 0xc9, 0xf4,
	0x9b, 0x22, 0x5a, 0xf3, 0xf8, 0x0a, 0xde, 0x33, 0x89, 0x82, 0x51, 0xa9, 0x20, 0x4f, 0xe1, 0x63,
	0xbd, 0xd5, 0x67, 0xf4, 0x43, 0xf7, 0xbb, 0x34, 0xb1, 0xa2, 0x8b, 0x66, 0x63, 0x4f, 0x7f, 0x07,
	0x00, 0x00, 0xff, 0xff, 0x2e, 0xc9, 0x92, 0xe0, 0x53, 0x02, 0x00, 0x00,
}

func (this *DenomAuthorityMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*DenomAuthorityMetadata)
	if !ok {
		that2, ok := that.(DenomAuthorityMetadata)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Admin != that1.Admin {
		return false
	}
	return true
}
func (m *DenomAuthorityMetadata) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomAuthorityMetadata) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomAuthorityMetadata) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintTokenfactory(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DenomCreationGasConsume != 0 {
		i = encodeVarintTokenfactory(dAtA, i, uint64(m.DenomCreationGasConsume))
		i--
		dAtA[i] = 0x10
	}
	if len(m.DenomCreationFee) > 0 {
		for iNdEx := len(m.DenomCreationFee) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomCreationFee[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTokenfactory(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintTokenfactory(dAtA []byte, offset int, v uint64) int {
	offset -= sovTokenfactory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomAuthorityMetadata) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovTokenfactory(uint64(l))
	}
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DenomCreationFee) > 0 {
		for _, e := range m.DenomCreationFee {
			l = e.Size()
			n += 1 + l + sovTokenfactory(uint64(l))
		}
	}
	if m.DenomCreationGasConsume != 0 {
		n += 1 + sovTokenfactory(uint64(m.DenomCreationGasConsume))
	}
	return n
}

func sovTokenfactory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTokenfactory(x uint64) (n int) {
	return sovTokenfactory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomAuthorityMetadata) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenfactory
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
			return fmt.Errorf("proto: DenomAuthorityMetadata: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomAuthorityMetadata: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenfactory
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
				return ErrInvalidLengthTokenfactory
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTokenfactory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTokenfactory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenfactory
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTokenfactory
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationFee", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenfactory
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
				return ErrInvalidLengthTokenfactory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTokenfactory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomCreationFee = append(m.DenomCreationFee, types.Coin{})
			if err := m.DenomCreationFee[len(m.DenomCreationFee)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomCreationGasConsume", wireType)
			}
			m.DenomCreationGasConsume = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTokenfactory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.DenomCreationGasConsume |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTokenfactory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTokenfactory
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
func skipTokenfactory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTokenfactory
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
					return 0, ErrIntOverflowTokenfactory
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
					return 0, ErrIntOverflowTokenfactory
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
				return 0, ErrInvalidLengthTokenfactory
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTokenfactory
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTokenfactory
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTokenfactory        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTokenfactory          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTokenfactory = fmt.Errorf("proto: unexpected end of group")
)
