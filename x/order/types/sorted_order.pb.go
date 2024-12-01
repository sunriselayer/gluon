// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/order/sorted_order.proto

package types

import (
	fmt "fmt"
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

// SortedOrder
// Sorted by [seconds]/[nanos]/[order_hash]
type SortedOrder struct {
	Seconds    uint64 `protobuf:"varint,1,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos      uint32 `protobuf:"varint,2,opt,name=nanos,proto3" json:"nanos,omitempty"`
	OrderOwner string `protobuf:"bytes,3,opt,name=order_owner,json=orderOwner,proto3" json:"order_owner,omitempty"`
	OrderHash  string `protobuf:"bytes,4,opt,name=order_hash,json=orderHash,proto3" json:"order_hash,omitempty"`
}

func (m *SortedOrder) Reset()         { *m = SortedOrder{} }
func (m *SortedOrder) String() string { return proto.CompactTextString(m) }
func (*SortedOrder) ProtoMessage()    {}
func (*SortedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_43dd6eb1763d8584, []int{0}
}
func (m *SortedOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SortedOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SortedOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SortedOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortedOrder.Merge(m, src)
}
func (m *SortedOrder) XXX_Size() int {
	return m.Size()
}
func (m *SortedOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_SortedOrder.DiscardUnknown(m)
}

var xxx_messageInfo_SortedOrder proto.InternalMessageInfo

func (m *SortedOrder) GetSeconds() uint64 {
	if m != nil {
		return m.Seconds
	}
	return 0
}

func (m *SortedOrder) GetNanos() uint32 {
	if m != nil {
		return m.Nanos
	}
	return 0
}

func (m *SortedOrder) GetOrderOwner() string {
	if m != nil {
		return m.OrderOwner
	}
	return ""
}

func (m *SortedOrder) GetOrderHash() string {
	if m != nil {
		return m.OrderHash
	}
	return ""
}

func init() {
	proto.RegisterType((*SortedOrder)(nil), "gluon.order.SortedOrder")
}

func init() { proto.RegisterFile("gluon/order/sorted_order.proto", fileDescriptor_43dd6eb1763d8584) }

var fileDescriptor_43dd6eb1763d8584 = []byte{
	// 197 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4b, 0xcf, 0x29, 0xcd,
	0xcf, 0xd3, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x2f, 0xce, 0x2f, 0x2a, 0x49, 0x4d, 0x89, 0x07,
	0x73, 0xf4, 0x0a, 0x8a, 0xf2, 0x4b, 0xf2, 0x85, 0xb8, 0xc1, 0xf2, 0x7a, 0x60, 0x21, 0xa5, 0x5a,
	0x2e, 0xee, 0x60, 0xb0, 0x12, 0x7f, 0x10, 0x57, 0x48, 0x82, 0x8b, 0xbd, 0x38, 0x35, 0x39, 0x3f,
	0x2f, 0xa5, 0x58, 0x82, 0x51, 0x81, 0x51, 0x83, 0x25, 0x08, 0xc6, 0x15, 0x12, 0xe1, 0x62, 0xcd,
	0x4b, 0xcc, 0xcb, 0x2f, 0x96, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0d, 0x82, 0x70, 0x84, 0xe4, 0xb9,
	0xb8, 0xc1, 0xe6, 0xc4, 0xe7, 0x97, 0xe7, 0xa5, 0x16, 0x49, 0x30, 0x2b, 0x30, 0x6a, 0x70, 0x06,
	0x71, 0x81, 0x85, 0xfc, 0x41, 0x22, 0x42, 0xb2, 0x5c, 0x10, 0x5e, 0x7c, 0x46, 0x62, 0x71, 0x86,
	0x04, 0x0b, 0x58, 0x9e, 0x13, 0x2c, 0xe2, 0x91, 0x58, 0x9c, 0xe1, 0xa4, 0x7b, 0xe2, 0x91, 0x1c,
	0xe3, 0x85, 0x47, 0x72, 0x8c, 0x0f, 0x1e, 0xc9, 0x31, 0x4e, 0x78, 0x2c, 0xc7, 0x70, 0xe1, 0xb1,
	0x1c, 0xc3, 0x8d, 0xc7, 0x72, 0x0c, 0x51, 0xc2, 0x10, 0x5f, 0x54, 0x40, 0xfd, 0x51, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0x81, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x81, 0x6a, 0x2d,
	0x6e, 0xe3, 0x00, 0x00, 0x00,
}

func (m *SortedOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SortedOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SortedOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.OrderHash) > 0 {
		i -= len(m.OrderHash)
		copy(dAtA[i:], m.OrderHash)
		i = encodeVarintSortedOrder(dAtA, i, uint64(len(m.OrderHash)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.OrderOwner) > 0 {
		i -= len(m.OrderOwner)
		copy(dAtA[i:], m.OrderOwner)
		i = encodeVarintSortedOrder(dAtA, i, uint64(len(m.OrderOwner)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Nanos != 0 {
		i = encodeVarintSortedOrder(dAtA, i, uint64(m.Nanos))
		i--
		dAtA[i] = 0x10
	}
	if m.Seconds != 0 {
		i = encodeVarintSortedOrder(dAtA, i, uint64(m.Seconds))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSortedOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovSortedOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SortedOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Seconds != 0 {
		n += 1 + sovSortedOrder(uint64(m.Seconds))
	}
	if m.Nanos != 0 {
		n += 1 + sovSortedOrder(uint64(m.Nanos))
	}
	l = len(m.OrderOwner)
	if l > 0 {
		n += 1 + l + sovSortedOrder(uint64(l))
	}
	l = len(m.OrderHash)
	if l > 0 {
		n += 1 + l + sovSortedOrder(uint64(l))
	}
	return n
}

func sovSortedOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSortedOrder(x uint64) (n int) {
	return sovSortedOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SortedOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSortedOrder
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
			return fmt.Errorf("proto: SortedOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SortedOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
			}
			m.Seconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Seconds |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
			}
			m.Nanos = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Nanos |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderOwner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedOrder
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
				return ErrInvalidLengthSortedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSortedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderOwner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OrderHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedOrder
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
				return ErrInvalidLengthSortedOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthSortedOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OrderHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSortedOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSortedOrder
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
func skipSortedOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSortedOrder
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
					return 0, ErrIntOverflowSortedOrder
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
					return 0, ErrIntOverflowSortedOrder
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
				return 0, ErrInvalidLengthSortedOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSortedOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSortedOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSortedOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSortedOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSortedOrder = fmt.Errorf("proto: unexpected end of group")
)
