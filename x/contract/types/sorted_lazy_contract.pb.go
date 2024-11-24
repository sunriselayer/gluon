// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/contract/sorted_lazy_contract.proto

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

type SortedLazyContract struct {
	Expiry uint64 `protobuf:"varint,1,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Id     uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *SortedLazyContract) Reset()         { *m = SortedLazyContract{} }
func (m *SortedLazyContract) String() string { return proto.CompactTextString(m) }
func (*SortedLazyContract) ProtoMessage()    {}
func (*SortedLazyContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f86f01373e97949, []int{0}
}
func (m *SortedLazyContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *SortedLazyContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_SortedLazyContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *SortedLazyContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SortedLazyContract.Merge(m, src)
}
func (m *SortedLazyContract) XXX_Size() int {
	return m.Size()
}
func (m *SortedLazyContract) XXX_DiscardUnknown() {
	xxx_messageInfo_SortedLazyContract.DiscardUnknown(m)
}

var xxx_messageInfo_SortedLazyContract proto.InternalMessageInfo

func (m *SortedLazyContract) GetExpiry() uint64 {
	if m != nil {
		return m.Expiry
	}
	return 0
}

func (m *SortedLazyContract) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func init() {
	proto.RegisterType((*SortedLazyContract)(nil), "gluon.contract.SortedLazyContract")
}

func init() {
	proto.RegisterFile("gluon/contract/sorted_lazy_contract.proto", fileDescriptor_1f86f01373e97949)
}

var fileDescriptor_1f86f01373e97949 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x4c, 0xcf, 0x29, 0xcd,
	0xcf, 0xd3, 0x4f, 0xce, 0xcf, 0x2b, 0x29, 0x4a, 0x4c, 0x2e, 0xd1, 0x2f, 0xce, 0x2f, 0x2a, 0x49,
	0x4d, 0x89, 0xcf, 0x49, 0xac, 0xaa, 0x8c, 0x87, 0x09, 0xea, 0x15, 0x14, 0xe5, 0x97, 0xe4, 0x0b,
	0xf1, 0x81, 0x95, 0xea, 0xc1, 0x44, 0x95, 0x6c, 0xb8, 0x84, 0x82, 0xc1, 0xaa, 0x7d, 0x12, 0xab,
	0x2a, 0x9d, 0xa1, 0xa2, 0x42, 0x62, 0x5c, 0x6c, 0xa9, 0x15, 0x05, 0x99, 0x45, 0x95, 0x12, 0x8c,
	0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0x13, 0x58,
	0x8c, 0x29, 0x33, 0xc5, 0xc9, 0xe0, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0xc4, 0x20, 0x4e, 0xaa, 0x40, 0x38, 0xaa, 0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x0c,
	0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x92, 0xf8, 0xd0, 0xb3, 0x00, 0x00, 0x00,
}

func (m *SortedLazyContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *SortedLazyContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *SortedLazyContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintSortedLazyContract(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if m.Expiry != 0 {
		i = encodeVarintSortedLazyContract(dAtA, i, uint64(m.Expiry))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintSortedLazyContract(dAtA []byte, offset int, v uint64) int {
	offset -= sovSortedLazyContract(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *SortedLazyContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Expiry != 0 {
		n += 1 + sovSortedLazyContract(uint64(m.Expiry))
	}
	if m.Id != 0 {
		n += 1 + sovSortedLazyContract(uint64(m.Id))
	}
	return n
}

func sovSortedLazyContract(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozSortedLazyContract(x uint64) (n int) {
	return sovSortedLazyContract(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *SortedLazyContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSortedLazyContract
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
			return fmt.Errorf("proto: SortedLazyContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: SortedLazyContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			m.Expiry = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedLazyContract
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
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSortedLazyContract
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
		default:
			iNdEx = preIndex
			skippy, err := skipSortedLazyContract(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthSortedLazyContract
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
func skipSortedLazyContract(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSortedLazyContract
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
					return 0, ErrIntOverflowSortedLazyContract
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
					return 0, ErrIntOverflowSortedLazyContract
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
				return 0, ErrInvalidLengthSortedLazyContract
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupSortedLazyContract
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthSortedLazyContract
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthSortedLazyContract        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSortedLazyContract          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupSortedLazyContract = fmt.Errorf("proto: unexpected end of group")
)
