// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/order/order.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// Order
type Order struct {
	Hash             string                `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`
	Owner            string                `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Body             types.Any             `protobuf:"bytes,3,opt,name=body,proto3" json:"body"`
	Cancelled        bool                  `protobuf:"varint,4,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	ContractedAmount cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=contracted_amount,json=contractedAmount,proto3,customtype=cosmossdk.io/math.Int" json:"contracted_amount"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_12d6294cf91d7886, []int{0}
}
func (m *Order) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Order) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Order.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Order) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Order.Merge(m, src)
}
func (m *Order) XXX_Size() int {
	return m.Size()
}
func (m *Order) XXX_DiscardUnknown() {
	xxx_messageInfo_Order.DiscardUnknown(m)
}

var xxx_messageInfo_Order proto.InternalMessageInfo

func (m *Order) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func (m *Order) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Order) GetBody() types.Any {
	if m != nil {
		return m.Body
	}
	return types.Any{}
}

func (m *Order) GetCancelled() bool {
	if m != nil {
		return m.Cancelled
	}
	return false
}

func init() {
	proto.RegisterType((*Order)(nil), "gluon.order.Order")
}

func init() { proto.RegisterFile("gluon/order/order.proto", fileDescriptor_12d6294cf91d7886) }

var fileDescriptor_12d6294cf91d7886 = []byte{
	// 321 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0x51, 0x41, 0x4e, 0x32, 0x31,
	0x18, 0x9d, 0xfe, 0xff, 0x60, 0xa4, 0x6c, 0xb4, 0x62, 0x1c, 0x88, 0x19, 0x88, 0x2b, 0x12, 0x43,
	0x27, 0xd1, 0x13, 0xc0, 0x8e, 0x95, 0xc9, 0xb8, 0x31, 0x6e, 0x48, 0x99, 0xd6, 0x42, 0x1c, 0xfa,
	0x91, 0xb6, 0x44, 0xe7, 0x16, 0x1e, 0x86, 0x43, 0xb0, 0x24, 0xac, 0x8c, 0x0b, 0x62, 0xe0, 0x08,
	0x5e, 0xc0, 0x4c, 0x3b, 0x86, 0x4d, 0xf3, 0x7d, 0xef, 0xbd, 0xbc, 0xf7, 0x9a, 0x0f, 0x5f, 0xc9,
	0x7c, 0x09, 0x2a, 0x01, 0xcd, 0x85, 0xf6, 0x2f, 0x5d, 0x68, 0xb0, 0x40, 0x1a, 0x8e, 0xa0, 0x0e,
	0x6a, 0xb7, 0x32, 0x30, 0x73, 0x30, 0x63, 0x47, 0x25, 0x7e, 0xf1, 0xba, 0x76, 0x53, 0x82, 0x04,
	0x8f, 0x97, 0x53, 0x85, 0xb6, 0x24, 0x80, 0xcc, 0x45, 0xe2, 0xb6, 0xc9, 0xf2, 0x25, 0x61, 0xaa,
	0xf0, 0xd4, 0xcd, 0x0f, 0xc2, 0xb5, 0x87, 0xd2, 0x95, 0x10, 0x1c, 0x4e, 0x99, 0x99, 0x46, 0xa8,
	0x8b, 0x7a, 0xf5, 0xd4, 0xcd, 0x84, 0xe2, 0x1a, 0xbc, 0x29, 0xa1, 0xa3, 0x7f, 0x25, 0x38, 0x8c,
	0xb6, 0xab, 0x7e, 0xb3, 0xca, 0x1b, 0x70, 0xae, 0x85, 0x31, 0x8f, 0x56, 0xcf, 0x94, 0x4c, 0xbd,
	0x8c, 0x50, 0x1c, 0x4e, 0x80, 0x17, 0xd1, 0xff, 0x2e, 0xea, 0x35, 0xee, 0x9a, 0xd4, 0xe7, 0xd2,
	0xbf, 0x5c, 0x3a, 0x50, 0xc5, 0x30, 0x5c, 0xef, 0x3a, 0x41, 0xea, 0x74, 0xe4, 0x1a, 0xd7, 0x33,
	0xa6, 0x32, 0x91, 0xe7, 0x82, 0x47, 0x61, 0x17, 0xf5, 0x4e, 0xd3, 0x23, 0x40, 0x9e, 0xf0, 0x79,
	0x06, 0xca, 0x6a, 0x96, 0x59, 0xc1, 0xc7, 0x6c, 0x0e, 0x4b, 0x65, 0xa3, 0x9a, 0x6b, 0x72, 0x5b,
	0x9a, 0x7c, 0xed, 0x3a, 0x97, 0xbe, 0x8d, 0xe1, 0xaf, 0x74, 0x06, 0xc9, 0x9c, 0xd9, 0x29, 0x1d,
	0x29, 0xbb, 0x5d, 0xf5, 0x71, 0x55, 0x73, 0xa4, 0x6c, 0x7a, 0x76, 0x74, 0x19, 0x38, 0x93, 0x61,
	0x7f, 0xbd, 0x8f, 0xd1, 0x66, 0x1f, 0xa3, 0xef, 0x7d, 0x8c, 0x3e, 0x0e, 0x71, 0xb0, 0x39, 0xc4,
	0xc1, 0xe7, 0x21, 0x0e, 0x9e, 0x2f, 0xfc, 0x05, 0xde, 0xab, 0x1b, 0xd8, 0x62, 0x21, 0xcc, 0xe4,
	0xc4, 0x7d, 0xe0, 0xfe, 0x37, 0x00, 0x00, 0xff, 0xff, 0xa5, 0xab, 0xd7, 0x72, 0x9f, 0x01, 0x00,
	0x00,
}

func (m *Order) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Order) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Order) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.ContractedAmount.Size()
		i -= size
		if _, err := m.ContractedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Cancelled {
		i--
		if m.Cancelled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	{
		size, err := m.Body.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Order) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Hash)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = m.Body.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.Cancelled {
		n += 2
	}
	l = m.ContractedAmount.Size()
	n += 1 + l + sovOrder(uint64(l))
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Order) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowOrder
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
			return fmt.Errorf("proto: Order: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Order: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Body", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Body.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancelled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
			m.Cancelled = bool(v != 0)
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractedAmount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
				return ErrInvalidLengthOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ContractedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthOrder
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
func skipOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
					return 0, ErrIntOverflowOrder
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
				return 0, ErrInvalidLengthOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupOrder = fmt.Errorf("proto: unexpected end of group")
)
