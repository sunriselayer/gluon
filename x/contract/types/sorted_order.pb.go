// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/contract/sorted_order.proto

package types

import (
	fmt "fmt"
	io "io"
	math "math"
	math_bits "math/bits"

	cosmossdk_io_math "cosmossdk.io/math"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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
type SortedOrder struct {
	Expiry           uint64                `protobuf:"varint,1,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Id               string                `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cancelled        bool                  `protobuf:"varint,3,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	ContractedAmount cosmossdk_io_math.Int `protobuf:"bytes,4,opt,name=contracted_amount,json=contractedAmount,proto3,customtype=cosmossdk.io/math.Int" json:"contracted_amount"`
}

func (m *SortedOrder) Reset()         { *m = SortedOrder{} }
func (m *SortedOrder) String() string { return proto.CompactTextString(m) }
func (*SortedOrder) ProtoMessage()    {}
func (*SortedOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_f4a98d023afd6b1e, []int{0}
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

func (m *SortedOrder) GetExpiry() uint64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *SortedOrder) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *SortedOrder) GetCancelled() bool {
	if m != nil {
		return m.Cancelled
	}
	return false
}

func init() {
	proto.RegisterType((*SortedOrder)(nil), "gluon.contract.SortedOrder")
}

func init() { proto.RegisterFile("gluon/contract/sorted_order.proto", fileDescriptor_f4a98d023afd6b1e) }

var fileDescriptor_f4a98d023afd6b1e = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0xcf, 0x29, 0xcd,
	0xcf, 0xd3, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2f, 0xce, 0x2f, 0x2a, 0x49,
	0x4d, 0x89, 0xcf, 0x2f, 0x4a, 0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x03,
	0x2b, 0xd1, 0x83, 0x29, 0x91, 0x92, 0x4c, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0x8e, 0x07, 0xcb, 0xea,
	0x43, 0x38, 0x10, 0xa5, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x10, 0x71, 0x10, 0x0b, 0x22, 0xaa,
	0xb4, 0x96, 0x91, 0x8b, 0x3b, 0x18, 0x6c, 0xae, 0x3f, 0xc8, 0x58, 0x21, 0x31, 0x2e, 0xb6, 0xd4,
	0x8a, 0x82, 0xcc, 0xa2, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d, 0x96, 0x20, 0x28, 0x4f, 0x88, 0x8f,
	0x8b, 0x29, 0x33, 0x45, 0x82, 0x49, 0x81, 0x51, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45, 0x48, 0x86,
	0x8b, 0x33, 0x39, 0x31, 0x2f, 0x39, 0x35, 0x27, 0x27, 0x35, 0x45, 0x82, 0x59, 0x81, 0x51, 0x83,
	0x23, 0x08, 0x21, 0x20, 0x14, 0xc1, 0x25, 0x08, 0x73, 0x52, 0x6a, 0x4a, 0x7c, 0x62, 0x6e, 0x7e,
	0x69, 0x5e, 0x89, 0x04, 0x0b, 0x48, 0xb3, 0x93, 0xf6, 0x89, 0x7b, 0xf2, 0x0c, 0xb7, 0xee, 0xc9,
	0x8b, 0x42, 0x1c, 0x57, 0x9c, 0x92, 0xad, 0x97, 0x99, 0xaf, 0x9f, 0x9b, 0x58, 0x92, 0xa1, 0xe7,
	0x99, 0x57, 0x72, 0x69, 0x8b, 0x2e, 0x17, 0xd4, 0xd5, 0x9e, 0x79, 0x25, 0x41, 0x02, 0x08, 0x53,
	0x1c, 0xc1, 0x86, 0x38, 0x19, 0x9c, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91, 0x1c, 0xe3, 0x83, 0x47,
	0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3, 0xb1, 0x1c, 0x43, 0x94,
	0x18, 0x24, 0xb4, 0x2a, 0x10, 0xe1, 0x55, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xf6, 0xa8,
	0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x6c, 0x54, 0x7a, 0x06, 0x4e, 0x01, 0x00, 0x00,
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
	{
		size := m.ContractedAmount.Size()
		i -= size
		if _, err := m.ContractedAmount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintSortedOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.Cancelled {
		i--
		if m.Cancelled {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintSortedOrder(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if m.Expiry != 0 {
		i = encodeVarintSortedOrder(dAtA, i, uint64(m.Expiry))
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
	if m.Expiry != 0 {
		n += 1 + sovSortedOrder(uint64(m.Expiry))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovSortedOrder(uint64(l))
	}
	if m.Cancelled {
		n += 2
	}
	l = m.ContractedAmount.Size()
	n += 1 + l + sovSortedOrder(uint64(l))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Expiry |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cancelled", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedOrder
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
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractedAmount", wireType)
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
			if err := m.ContractedAmount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
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
