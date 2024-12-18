// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/perp/cross_margin.proto

package types

import (
	fmt "fmt"
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

// CrossMargin
type CrossMargin struct {
	Owner  string                                   `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	Assets github_com_cosmos_cosmos_sdk_types.Coins `protobuf:"bytes,2,rep,name=assets,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"assets"`
}

func (m *CrossMargin) Reset()         { *m = CrossMargin{} }
func (m *CrossMargin) String() string { return proto.CompactTextString(m) }
func (*CrossMargin) ProtoMessage()    {}
func (*CrossMargin) Descriptor() ([]byte, []int) {
	return fileDescriptor_4dc086b41b91396d, []int{0}
}
func (m *CrossMargin) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CrossMargin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CrossMargin.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CrossMargin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CrossMargin.Merge(m, src)
}
func (m *CrossMargin) XXX_Size() int {
	return m.Size()
}
func (m *CrossMargin) XXX_DiscardUnknown() {
	xxx_messageInfo_CrossMargin.DiscardUnknown(m)
}

var xxx_messageInfo_CrossMargin proto.InternalMessageInfo

func (m *CrossMargin) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *CrossMargin) GetAssets() github_com_cosmos_cosmos_sdk_types.Coins {
	if m != nil {
		return m.Assets
	}
	return nil
}

func init() {
	proto.RegisterType((*CrossMargin)(nil), "gluon.perp.CrossMargin")
}

func init() { proto.RegisterFile("gluon/perp/cross_margin.proto", fileDescriptor_4dc086b41b91396d) }

var fileDescriptor_4dc086b41b91396d = []byte{
	// 246 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x4d, 0xcf, 0x29, 0xcd,
	0xcf, 0xd3, 0x2f, 0x48, 0x2d, 0x2a, 0xd0, 0x4f, 0x2e, 0xca, 0x2f, 0x2e, 0x8e, 0xcf, 0x4d, 0x2c,
	0x4a, 0xcf, 0xcc, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x02, 0x4b, 0xeb, 0x81, 0xa4,
	0xa5, 0xe4, 0x92, 0xf3, 0x8b, 0x73, 0xf3, 0x8b, 0xf5, 0x93, 0x12, 0x8b, 0x53, 0xf5, 0xcb, 0x0c,
	0x93, 0x52, 0x4b, 0x12, 0x0d, 0xf5, 0x93, 0xf3, 0x61, 0x6a, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3,
	0xc1, 0x4c, 0x7d, 0x10, 0x0b, 0x22, 0xaa, 0xd4, 0xc1, 0xc8, 0xc5, 0xed, 0x0c, 0x32, 0xd8, 0x17,
	0x6c, 0xae, 0x90, 0x08, 0x17, 0x6b, 0x7e, 0x79, 0x5e, 0x6a, 0x91, 0x04, 0xa3, 0x02, 0xa3, 0x06,
	0x67, 0x10, 0x84, 0x23, 0x94, 0xcc, 0xc5, 0x96, 0x58, 0x5c, 0x9c, 0x5a, 0x52, 0x2c, 0xc1, 0xa4,
	0xc0, 0xac, 0xc1, 0x6d, 0x24, 0xa9, 0x07, 0xb1, 0x4c, 0x0f, 0x64, 0x99, 0x1e, 0xd4, 0x32, 0x3d,
	0xe7, 0xfc, 0xcc, 0x3c, 0x27, 0x83, 0x13, 0xf7, 0xe4, 0x19, 0x56, 0xdd, 0x97, 0xd7, 0x48, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0xba, 0x0c, 0x42, 0xe9, 0x16, 0xa7,
	0x64, 0xeb, 0x97, 0x54, 0x16, 0xa4, 0x16, 0x83, 0x35, 0x14, 0x07, 0x41, 0x8d, 0x76, 0xd2, 0x39,
	0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63,
	0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0x21, 0x48, 0x28, 0x54, 0x40, 0xc2,
	0x01, 0xac, 0x37, 0x89, 0x0d, 0xec, 0x7e, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x47, 0x28,
	0xe5, 0x03, 0x22, 0x01, 0x00, 0x00,
}

func (m *CrossMargin) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CrossMargin) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CrossMargin) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Assets) > 0 {
		for iNdEx := len(m.Assets) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Assets[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintCrossMargin(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintCrossMargin(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintCrossMargin(dAtA []byte, offset int, v uint64) int {
	offset -= sovCrossMargin(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *CrossMargin) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovCrossMargin(uint64(l))
	}
	if len(m.Assets) > 0 {
		for _, e := range m.Assets {
			l = e.Size()
			n += 1 + l + sovCrossMargin(uint64(l))
		}
	}
	return n
}

func sovCrossMargin(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozCrossMargin(x uint64) (n int) {
	return sovCrossMargin(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *CrossMargin) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowCrossMargin
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
			return fmt.Errorf("proto: CrossMargin: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CrossMargin: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrossMargin
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
				return ErrInvalidLengthCrossMargin
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthCrossMargin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Assets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowCrossMargin
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
				return ErrInvalidLengthCrossMargin
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthCrossMargin
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Assets = append(m.Assets, types.Coin{})
			if err := m.Assets[len(m.Assets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipCrossMargin(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthCrossMargin
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
func skipCrossMargin(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowCrossMargin
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
					return 0, ErrIntOverflowCrossMargin
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
					return 0, ErrIntOverflowCrossMargin
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
				return 0, ErrInvalidLengthCrossMargin
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupCrossMargin
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthCrossMargin
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthCrossMargin        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowCrossMargin          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupCrossMargin = fmt.Errorf("proto: unexpected end of group")
)
