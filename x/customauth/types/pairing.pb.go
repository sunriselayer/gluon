// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/customauth/pairing.proto

package types

import (
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/codec/types"
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

// Pairing
type Pairing struct {
	Id        uint64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string    `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey types.Any `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key"`
	CreatedAt time.Time `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3,stdtime" json:"created_at"`
}

func (m *Pairing) Reset()         { *m = Pairing{} }
func (m *Pairing) String() string { return proto.CompactTextString(m) }
func (*Pairing) ProtoMessage()    {}
func (*Pairing) Descriptor() ([]byte, []int) {
	return fileDescriptor_6b971e3bc1d64983, []int{0}
}
func (m *Pairing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Pairing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Pairing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Pairing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Pairing.Merge(m, src)
}
func (m *Pairing) XXX_Size() int {
	return m.Size()
}
func (m *Pairing) XXX_DiscardUnknown() {
	xxx_messageInfo_Pairing.DiscardUnknown(m)
}

var xxx_messageInfo_Pairing proto.InternalMessageInfo

func (m *Pairing) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Pairing) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Pairing) GetPublicKey() types.Any {
	if m != nil {
		return m.PublicKey
	}
	return types.Any{}
}

func (m *Pairing) GetCreatedAt() time.Time {
	if m != nil {
		return m.CreatedAt
	}
	return time.Time{}
}

func init() {
	proto.RegisterType((*Pairing)(nil), "gluon.customauth.Pairing")
}

func init() { proto.RegisterFile("gluon/customauth/pairing.proto", fileDescriptor_6b971e3bc1d64983) }

var fileDescriptor_6b971e3bc1d64983 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x3f, 0x4f, 0xb4, 0x30,
	0x1c, 0xc7, 0x29, 0x0f, 0x79, 0x4e, 0x6a, 0x62, 0x0c, 0xb9, 0xa1, 0x32, 0x14, 0xe2, 0xc4, 0x04,
	0xc9, 0x39, 0x39, 0x1e, 0x8e, 0x2e, 0x86, 0x38, 0xb9, 0x5c, 0x0a, 0xd4, 0xda, 0x08, 0x94, 0x40,
	0x49, 0xec, 0xbb, 0xb8, 0xf7, 0xe2, 0x9b, 0xb8, 0xf1, 0x46, 0x27, 0x35, 0xf0, 0x46, 0xcc, 0xb5,
	0x47, 0x34, 0xba, 0xf5, 0x97, 0xcf, 0xf7, 0x4f, 0xf3, 0x85, 0x98, 0x55, 0x83, 0x68, 0x92, 0x62,
	0xe8, 0xa5, 0xa8, 0xc9, 0x20, 0x9f, 0x92, 0x96, 0xf0, 0x8e, 0x37, 0x2c, 0x6e, 0x3b, 0x21, 0x85,
	0x77, 0xae, 0x79, 0xfc, 0xcd, 0xfd, 0x25, 0x13, 0x4c, 0x68, 0x98, 0x1c, 0x5e, 0x46, 0xe7, 0x5f,
	0x30, 0x21, 0x58, 0x45, 0x13, 0x7d, 0xe5, 0xc3, 0x63, 0x42, 0x1a, 0x75, 0x44, 0xc1, 0x6f, 0x24,
	0x79, 0x4d, 0x7b, 0x49, 0xea, 0xd6, 0x08, 0x2e, 0x5f, 0x01, 0x5c, 0xdc, 0x99, 0x56, 0xef, 0x0c,
	0xda, 0xbc, 0x44, 0x20, 0x04, 0x91, 0x93, 0xd9, 0xbc, 0xf4, 0x10, 0x5c, 0x90, 0xb2, 0xec, 0x68,
	0xdf, 0x23, 0x3b, 0x04, 0x91, 0x9b, 0xcd, 0xa7, 0x77, 0x0d, 0x61, 0x3b, 0xe4, 0x15, 0x2f, 0x36,
	0xcf, 0x54, 0xa1, 0x7f, 0x21, 0x88, 0x4e, 0x57, 0xcb, 0xd8, 0x74, 0xc5, 0x73, 0x57, 0xbc, 0x6e,
	0x54, 0xea, 0xec, 0xde, 0x03, 0x2b, 0x73, 0x8d, 0xfa, 0x96, 0x2a, 0xef, 0x06, 0xc2, 0xa2, 0xa3,
	0x44, 0xd2, 0x72, 0x43, 0x24, 0x72, 0xb4, 0xd5, 0xff, 0x63, 0xbd, 0x9f, 0xbf, 0x99, 0x9e, 0x1c,
	0x02, 0xb6, 0x1f, 0x01, 0xc8, 0xdc, 0xa3, 0x6f, 0x2d, 0xd3, 0xd5, 0x6e, 0xc4, 0x60, 0x3f, 0x62,
	0xf0, 0x39, 0x62, 0xb0, 0x9d, 0xb0, 0xb5, 0x9f, 0xb0, 0xf5, 0x36, 0x61, 0xeb, 0x01, 0x99, 0x4d,
	0x5f, 0x7e, 0xae, 0x2a, 0x55, 0x4b, 0xfb, 0xfc, 0xbf, 0x0e, 0xbf, 0xfa, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0xeb, 0x3f, 0xe0, 0xa4, 0x76, 0x01, 0x00, 0x00,
}

func (m *Pairing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Pairing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Pairing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.CreatedAt, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintPairing(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x22
	{
		size, err := m.PublicKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPairing(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintPairing(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if m.Id != 0 {
		i = encodeVarintPairing(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintPairing(dAtA []byte, offset int, v uint64) int {
	offset -= sovPairing(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Pairing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovPairing(uint64(m.Id))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovPairing(uint64(l))
	}
	l = m.PublicKey.Size()
	n += 1 + l + sovPairing(uint64(l))
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.CreatedAt)
	n += 1 + l + sovPairing(uint64(l))
	return n
}

func sovPairing(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPairing(x uint64) (n int) {
	return sovPairing(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Pairing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPairing
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
			return fmt.Errorf("proto: Pairing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Pairing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairing
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
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairing
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
				return ErrInvalidLengthPairing
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthPairing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairing
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
				return ErrInvalidLengthPairing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPairing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPairing
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
				return ErrInvalidLengthPairing
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPairing
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.CreatedAt, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPairing(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPairing
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
func skipPairing(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPairing
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
					return 0, ErrIntOverflowPairing
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
					return 0, ErrIntOverflowPairing
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
				return 0, ErrInvalidLengthPairing
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPairing
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPairing
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPairing        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPairing          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPairing = fmt.Errorf("proto: unexpected end of group")
)
