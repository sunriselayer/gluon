// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/perp/order.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	types "gluon/x/order/types"
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

// PerpPositionCreateOrder
type PerpPositionCreateOrder struct {
	types.BaseOrder `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base"`
	IsolatedMargin  bool `protobuf:"varint,2,opt,name=isolated_margin,json=isolatedMargin,proto3" json:"isolated_margin,omitempty"`
	// Zero margin is only accepted for cross margin
	MarginAmount cosmossdk_io_math.Int `protobuf:"bytes,3,opt,name=margin_amount,json=marginAmount,proto3,customtype=cosmossdk.io/math.Int" json:"margin_amount"`
}

func (m *PerpPositionCreateOrder) Reset()         { *m = PerpPositionCreateOrder{} }
func (m *PerpPositionCreateOrder) String() string { return proto.CompactTextString(m) }
func (*PerpPositionCreateOrder) ProtoMessage()    {}
func (*PerpPositionCreateOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e20fa245d287f6, []int{0}
}
func (m *PerpPositionCreateOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpPositionCreateOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpPositionCreateOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpPositionCreateOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpPositionCreateOrder.Merge(m, src)
}
func (m *PerpPositionCreateOrder) XXX_Size() int {
	return m.Size()
}
func (m *PerpPositionCreateOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpPositionCreateOrder.DiscardUnknown(m)
}

var xxx_messageInfo_PerpPositionCreateOrder proto.InternalMessageInfo

func (m *PerpPositionCreateOrder) GetIsolatedMargin() bool {
	if m != nil {
		return m.IsolatedMargin
	}
	return false
}

// PerpPositionCancelOrder
type PerpPositionCancelOrder struct {
	types.BaseOrder   `protobuf:"bytes,1,opt,name=base,proto3,embedded=base" json:"base"`
	PositionOrderHash string `protobuf:"bytes,2,opt,name=position_order_hash,json=positionOrderHash,proto3" json:"position_order_hash,omitempty"`
}

func (m *PerpPositionCancelOrder) Reset()         { *m = PerpPositionCancelOrder{} }
func (m *PerpPositionCancelOrder) String() string { return proto.CompactTextString(m) }
func (*PerpPositionCancelOrder) ProtoMessage()    {}
func (*PerpPositionCancelOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_63e20fa245d287f6, []int{1}
}
func (m *PerpPositionCancelOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PerpPositionCancelOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PerpPositionCancelOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PerpPositionCancelOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PerpPositionCancelOrder.Merge(m, src)
}
func (m *PerpPositionCancelOrder) XXX_Size() int {
	return m.Size()
}
func (m *PerpPositionCancelOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_PerpPositionCancelOrder.DiscardUnknown(m)
}

var xxx_messageInfo_PerpPositionCancelOrder proto.InternalMessageInfo

func (m *PerpPositionCancelOrder) GetPositionOrderHash() string {
	if m != nil {
		return m.PositionOrderHash
	}
	return ""
}

func init() {
	proto.RegisterType((*PerpPositionCreateOrder)(nil), "gluon.perp.PerpPositionCreateOrder")
	proto.RegisterType((*PerpPositionCancelOrder)(nil), "gluon.perp.PerpPositionCancelOrder")
}

func init() { proto.RegisterFile("gluon/perp/order.proto", fileDescriptor_63e20fa245d287f6) }

var fileDescriptor_63e20fa245d287f6 = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x14, 0x85, 0x33, 0x2a, 0xd2, 0x8e, 0x7f, 0x18, 0xb5, 0xd6, 0x22, 0x69, 0xe9, 0xc6, 0x82, 0x9a,
	0x80, 0xfa, 0x02, 0xc6, 0x8d, 0x5d, 0x88, 0x25, 0x4b, 0x37, 0x61, 0xda, 0x0e, 0x49, 0xb0, 0x99,
	0x3b, 0xcc, 0x4c, 0x41, 0x57, 0xbe, 0x82, 0x0f, 0xe3, 0x43, 0x14, 0x57, 0xc5, 0x95, 0xb8, 0x28,
	0xd2, 0xbc, 0x88, 0x64, 0x6e, 0x0b, 0xe2, 0xd6, 0xdd, 0xcc, 0xb9, 0xf7, 0x7c, 0xcc, 0x99, 0x43,
	0x6b, 0xc9, 0x68, 0x0c, 0x22, 0x90, 0x5c, 0xc9, 0x00, 0xd4, 0x90, 0x2b, 0x5f, 0x2a, 0x30, 0xe0,
	0x52, 0xab, 0xfb, 0xa5, 0xde, 0x38, 0x1a, 0x80, 0xce, 0x41, 0xc7, 0x76, 0x12, 0xe0, 0x05, 0xd7,
	0x1a, 0xc7, 0x68, 0xb7, 0xce, 0xa0, 0xcf, 0x34, 0x8f, 0x7f, 0x41, 0x1a, 0xfb, 0x09, 0x24, 0x80,
	0xae, 0xf2, 0x84, 0x6a, 0xfb, 0x9d, 0xd0, 0xc3, 0x1e, 0x57, 0xb2, 0x07, 0x3a, 0x33, 0x19, 0x88,
	0x1b, 0xc5, 0x99, 0xe1, 0xf7, 0xa5, 0xcf, 0xbd, 0xa2, 0x6b, 0x25, 0xa5, 0x4e, 0x5a, 0xa4, 0xb3,
	0x71, 0x51, 0xf3, 0xf1, 0x15, 0xc8, 0x0c, 0x99, 0xc6, 0xad, 0xb0, 0x32, 0x99, 0x35, 0x9d, 0xe9,
	0xac, 0x49, 0x22, 0xbb, 0xed, 0x9e, 0xd0, 0x9d, 0x4c, 0xc3, 0x88, 0x19, 0x3e, 0x8c, 0x73, 0xa6,
	0x92, 0x4c, 0xd4, 0x57, 0x5a, 0xa4, 0x53, 0x89, 0xb6, 0x97, 0xf2, 0x9d, 0x55, 0xdd, 0x1e, 0xdd,
	0xc2, 0x79, 0xcc, 0x72, 0x18, 0x0b, 0x53, 0x5f, 0x6d, 0x91, 0x4e, 0x35, 0x3c, 0x2d, 0x79, 0x5f,
	0xb3, 0xe6, 0x01, 0x66, 0xd3, 0xc3, 0x47, 0x3f, 0x83, 0x20, 0x67, 0x26, 0xf5, 0xbb, 0xc2, 0x7c,
	0xbc, 0x9d, 0xd3, 0x45, 0xe8, 0xae, 0x30, 0xd1, 0x26, 0x12, 0xae, 0x2d, 0xa0, 0xfd, 0xf2, 0x27,
	0x0b, 0x13, 0x03, 0x3e, 0xfa, 0x4f, 0x16, 0x9f, 0xee, 0xc9, 0x05, 0x0c, 0xff, 0x32, 0x4e, 0x99,
	0x4e, 0x6d, 0x9e, 0x6a, 0xb4, 0xbb, 0x1c, 0x59, 0xef, 0x2d, 0xd3, 0x69, 0x78, 0x36, 0x99, 0x7b,
	0x64, 0x3a, 0xf7, 0xc8, 0xf7, 0xdc, 0x23, 0xaf, 0x85, 0xe7, 0x4c, 0x0b, 0xcf, 0xf9, 0x2c, 0x3c,
	0xe7, 0xc1, 0xc5, 0x6e, 0x9e, 0xb0, 0x5c, 0xf3, 0x2c, 0xb9, 0xee, 0xaf, 0xdb, 0x0a, 0x2e, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xe8, 0x08, 0x09, 0xd9, 0xf7, 0x01, 0x00, 0x00,
}

func (m *PerpPositionCreateOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpPositionCreateOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpPositionCreateOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.MarginAmount.Size()
		i -= size
		if _, err := m.MarginAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if m.IsolatedMargin {
		i--
		if m.IsolatedMargin {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	{
		size, err := m.BaseOrder.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *PerpPositionCancelOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PerpPositionCancelOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PerpPositionCancelOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PositionOrderHash) > 0 {
		i -= len(m.PositionOrderHash)
		copy(dAtA[i:], m.PositionOrderHash)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.PositionOrderHash)))
		i--
		dAtA[i] = 0x12
	}
	{
		size, err := m.BaseOrder.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
func (m *PerpPositionCreateOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseOrder.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.IsolatedMargin {
		n += 2
	}
	l = m.MarginAmount.Size()
	n += 1 + l + sovOrder(uint64(l))
	return n
}

func (m *PerpPositionCancelOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.BaseOrder.Size()
	n += 1 + l + sovOrder(uint64(l))
	l = len(m.PositionOrderHash)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	return n
}

func sovOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozOrder(x uint64) (n int) {
	return sovOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PerpPositionCreateOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PerpPositionCreateOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpPositionCreateOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseOrder", wireType)
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
			if err := m.BaseOrder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsolatedMargin", wireType)
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
			m.IsolatedMargin = bool(v != 0)
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MarginAmount", wireType)
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
			if err := m.MarginAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *PerpPositionCancelOrder) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: PerpPositionCancelOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PerpPositionCancelOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BaseOrder", wireType)
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
			if err := m.BaseOrder.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionOrderHash", wireType)
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
			m.PositionOrderHash = string(dAtA[iNdEx:postIndex])
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
