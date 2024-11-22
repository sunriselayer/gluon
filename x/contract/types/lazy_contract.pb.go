// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/contract/lazy_contract.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	github_com_cosmos_gogoproto_types "github.com/cosmos/gogoproto/types"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	math "math"
	math_bits "math/bits"
	time "time"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf
var _ = time.Kitchen

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type LazyContract struct {
	Id                  uint64                `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Buyer               string                `protobuf:"bytes,2,opt,name=buyer,proto3" json:"buyer,omitempty"`
	Seller              string                `protobuf:"bytes,3,opt,name=seller,proto3" json:"seller,omitempty"`
	DenomBase           string                `protobuf:"bytes,4,opt,name=denom_base,json=denomBase,proto3" json:"denom_base,omitempty"`
	DenomQuote          string                `protobuf:"bytes,5,opt,name=denom_quote,json=denomQuote,proto3" json:"denom_quote,omitempty"`
	AmountEscrowBuyer   cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=amount_escrow_buyer,json=amountEscrowBuyer,proto3,customtype=cosmossdk.io/math.Int" json:"amount_escrow_buyer"`
	AmountEscrowSeller  cosmossdk_io_math.Int `protobuf:"bytes,7,opt,name=amount_escrow_seller,json=amountEscrowSeller,proto3,customtype=cosmossdk.io/math.Int" json:"amount_escrow_seller"`
	AmountPendingBuyer  cosmossdk_io_math.Int `protobuf:"bytes,8,opt,name=amount_pending_buyer,json=amountPendingBuyer,proto3,customtype=cosmossdk.io/math.Int" json:"amount_pending_buyer"`
	AmountPendingSeller cosmossdk_io_math.Int `protobuf:"bytes,9,opt,name=amount_pending_seller,json=amountPendingSeller,proto3,customtype=cosmossdk.io/math.Int" json:"amount_pending_seller"`
	Expiry              time.Time             `protobuf:"bytes,10,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *LazyContract) Reset()         { *m = LazyContract{} }
func (m *LazyContract) String() string { return proto.CompactTextString(m) }
func (*LazyContract) ProtoMessage()    {}
func (*LazyContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_df50d0b8489692fb, []int{0}
}
func (m *LazyContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *LazyContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_LazyContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *LazyContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LazyContract.Merge(m, src)
}
func (m *LazyContract) XXX_Size() int {
	return m.Size()
}
func (m *LazyContract) XXX_DiscardUnknown() {
	xxx_messageInfo_LazyContract.DiscardUnknown(m)
}

var xxx_messageInfo_LazyContract proto.InternalMessageInfo

func (m *LazyContract) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LazyContract) GetBuyer() string {
	if m != nil {
		return m.Buyer
	}
	return ""
}

func (m *LazyContract) GetSeller() string {
	if m != nil {
		return m.Seller
	}
	return ""
}

func (m *LazyContract) GetDenomBase() string {
	if m != nil {
		return m.DenomBase
	}
	return ""
}

func (m *LazyContract) GetDenomQuote() string {
	if m != nil {
		return m.DenomQuote
	}
	return ""
}

func (m *LazyContract) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*LazyContract)(nil), "gluon.contract.LazyContract")
}

func init() {
	proto.RegisterFile("gluon/contract/lazy_contract.proto", fileDescriptor_df50d0b8489692fb)
}

var fileDescriptor_df50d0b8489692fb = []byte{
	// 434 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x93, 0x31, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x7d, 0x21, 0x35, 0xcd, 0x15, 0x55, 0xea, 0x35, 0xad, 0x4c, 0x24, 0xec, 0xa8, 0x53,
	0x84, 0x84, 0x5d, 0xc1, 0xca, 0x64, 0xc4, 0x50, 0x89, 0x01, 0x0c, 0x13, 0x8b, 0x39, 0xdb, 0x87,
	0xb1, 0xb0, 0xef, 0x19, 0xdf, 0x59, 0xd4, 0xfd, 0x14, 0xfd, 0x18, 0x8c, 0x0c, 0x7c, 0x88, 0x8e,
	0x15, 0x13, 0xea, 0x50, 0x50, 0x32, 0xf0, 0x35, 0x90, 0xef, 0xce, 0x0a, 0x61, 0x0c, 0x4b, 0x74,
	0xef, 0xff, 0x7f, 0xf7, 0x7e, 0xf9, 0x5b, 0xef, 0xf0, 0x49, 0x5e, 0xb6, 0xc0, 0x83, 0x14, 0xb8,
	0x6c, 0x68, 0x2a, 0x83, 0x92, 0x5e, 0x74, 0xf1, 0x50, 0xf9, 0x75, 0x03, 0x12, 0xc8, 0xbe, 0xea,
	0xf1, 0x07, 0x75, 0x76, 0x40, 0xab, 0x82, 0x43, 0xa0, 0x7e, 0x75, 0xcb, 0x6c, 0x9a, 0x43, 0x0e,
	0xea, 0x18, 0xf4, 0x27, 0xa3, 0xde, 0x4f, 0x41, 0x54, 0x20, 0x62, 0x6d, 0xe8, 0xc2, 0x58, 0x5e,
	0x0e, 0x90, 0x97, 0x2c, 0x50, 0x55, 0xd2, 0xbe, 0x0f, 0x64, 0x51, 0x31, 0x21, 0x69, 0x55, 0xeb,
	0x86, 0x93, 0x9b, 0x31, 0xbe, 0xf7, 0x82, 0x5e, 0x74, 0xcf, 0x0c, 0x95, 0xec, 0xe3, 0x51, 0x91,
	0x39, 0x68, 0x8e, 0x16, 0xe3, 0x68, 0x54, 0x64, 0x64, 0x8a, 0x77, 0x92, 0xb6, 0x63, 0x8d, 0x33,
	0x9a, 0xa3, 0xc5, 0x24, 0xd2, 0x05, 0x39, 0xc6, 0xb6, 0x60, 0x65, 0xc9, 0x1a, 0xe7, 0x8e, 0x92,
	0x4d, 0x45, 0x1e, 0x60, 0x9c, 0x31, 0x0e, 0x55, 0x9c, 0x50, 0xc1, 0x9c, 0xb1, 0xf2, 0x26, 0x4a,
	0x09, 0xa9, 0x60, 0xc4, 0xc3, 0x7b, 0xda, 0xfe, 0xd4, 0x82, 0x64, 0xce, 0x8e, 0xf2, 0xf5, 0x8d,
	0x57, 0xbd, 0x42, 0xde, 0xe1, 0x43, 0x5a, 0x41, 0xcb, 0x65, 0xcc, 0x44, 0xda, 0xc0, 0xe7, 0x58,
	0xb3, 0xed, 0xbe, 0x31, 0x3c, 0xbd, 0xba, 0xf5, 0xac, 0x9b, 0x5b, 0xef, 0x48, 0x47, 0x14, 0xd9,
	0x47, 0xbf, 0x80, 0xa0, 0xa2, 0xf2, 0x83, 0x7f, 0xc6, 0xe5, 0xf7, 0x6f, 0x8f, 0xb0, 0xc9, 0x7e,
	0xc6, 0xe5, 0x97, 0xdf, 0x5f, 0x1f, 0xa2, 0xe8, 0x40, 0x0f, 0x7b, 0xae, 0x66, 0x85, 0xea, 0x9f,
	0x27, 0x78, 0xba, 0x49, 0x30, 0x39, 0xee, 0x6e, 0x89, 0x20, 0x7f, 0x23, 0x5e, 0xeb, 0xaf, 0xb0,
	0x66, 0xd4, 0x8c, 0x67, 0x05, 0xcf, 0x4d, 0x8c, 0xdd, 0xff, 0x63, 0xbc, 0xd4, 0xc3, 0x74, 0x8e,
	0x0c, 0x1f, 0xfd, 0xc3, 0x30, 0x41, 0x26, 0x5b, 0x42, 0x0e, 0x37, 0x20, 0x26, 0xc9, 0x53, 0x6c,
	0xb3, 0xf3, 0xba, 0x68, 0x3a, 0x07, 0xcf, 0xd1, 0x62, 0xef, 0xf1, 0xcc, 0xd7, 0x0b, 0xe5, 0x0f,
	0x0b, 0xe5, 0xbf, 0x19, 0x16, 0x2a, 0xdc, 0xed, 0x91, 0x97, 0x3f, 0x3d, 0x14, 0x99, 0x3b, 0xe1,
	0xe9, 0xd5, 0xd2, 0x45, 0xd7, 0x4b, 0x17, 0xfd, 0x5a, 0xba, 0xe8, 0x72, 0xe5, 0x5a, 0xd7, 0x2b,
	0xd7, 0xfa, 0xb1, 0x72, 0xad, 0xb7, 0xc7, 0xfa, 0x3d, 0x9c, 0xaf, 0x5f, 0x84, 0xec, 0x6a, 0x26,
	0x12, 0x5b, 0xcd, 0x7d, 0xf2, 0x27, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x7e, 0xae, 0xf6, 0x30, 0x03,
	0x00, 0x00,
}

func (m *LazyContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *LazyContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *LazyContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintLazyContract(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x52
	{
		size := m.AmountPendingSeller.Size()
		i -= size
		if _, err := m.AmountPendingSeller.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLazyContract(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x4a
	{
		size := m.AmountPendingBuyer.Size()
		i -= size
		if _, err := m.AmountPendingBuyer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLazyContract(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x42
	{
		size := m.AmountEscrowSeller.Size()
		i -= size
		if _, err := m.AmountEscrowSeller.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLazyContract(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x3a
	{
		size := m.AmountEscrowBuyer.Size()
		i -= size
		if _, err := m.AmountEscrowBuyer.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintLazyContract(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if len(m.DenomQuote) > 0 {
		i -= len(m.DenomQuote)
		copy(dAtA[i:], m.DenomQuote)
		i = encodeVarintLazyContract(dAtA, i, uint64(len(m.DenomQuote)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.DenomBase) > 0 {
		i -= len(m.DenomBase)
		copy(dAtA[i:], m.DenomBase)
		i = encodeVarintLazyContract(dAtA, i, uint64(len(m.DenomBase)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Seller) > 0 {
		i -= len(m.Seller)
		copy(dAtA[i:], m.Seller)
		i = encodeVarintLazyContract(dAtA, i, uint64(len(m.Seller)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Buyer) > 0 {
		i -= len(m.Buyer)
		copy(dAtA[i:], m.Buyer)
		i = encodeVarintLazyContract(dAtA, i, uint64(len(m.Buyer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintLazyContract(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintLazyContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovLazyContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *LazyContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovLazyContract(uint64(m.Id))
	}
	l = len(m.Buyer)
	if l > 0 {
		n += 1 + l + sovLazyContract(uint64(l))
	}
	l = len(m.Seller)
	if l > 0 {
		n += 1 + l + sovLazyContract(uint64(l))
	}
	l = len(m.DenomBase)
	if l > 0 {
		n += 1 + l + sovLazyContract(uint64(l))
	}
	l = len(m.DenomQuote)
	if l > 0 {
		n += 1 + l + sovLazyContract(uint64(l))
	}
	l = m.AmountEscrowBuyer.Size()
	n += 1 + l + sovLazyContract(uint64(l))
	l = m.AmountEscrowSeller.Size()
	n += 1 + l + sovLazyContract(uint64(l))
	l = m.AmountPendingBuyer.Size()
	n += 1 + l + sovLazyContract(uint64(l))
	l = m.AmountPendingSeller.Size()
	n += 1 + l + sovLazyContract(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovLazyContract(uint64(l))
	return n
}

func sovLazyContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozLazyContract(x uint64) (n int) {
	return sovLazyContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *LazyContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowLazyContract
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
			return fmt.Errorf("proto: LazyContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: LazyContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Buyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Buyer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Seller = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomQuote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomQuote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountEscrowBuyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountEscrowBuyer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountEscrowSeller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountEscrowSeller.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountPendingBuyer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountPendingBuyer.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AmountPendingSeller", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.AmountPendingSeller.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowLazyContract
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
				return ErrInvalidLengthLazyContract
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthLazyContract
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipLazyContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthLazyContract
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
func skipLazyContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowLazyContract
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
					return 0, ErrIntOverflowLazyContract
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
					return 0, ErrIntOverflowLazyContract
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
				return 0, ErrInvalidLengthLazyContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupLazyContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthLazyContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthLazyContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowLazyContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupLazyContract = fmt.Errorf("proto: unexpected end of group")
)
