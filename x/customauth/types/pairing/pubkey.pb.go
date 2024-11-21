// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/customauth/pairing/pubkey.proto

package pairing

import (
	fmt "fmt"
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

type PubKeyNotVerifiable struct {
	User      []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PairingId uint64 `protobuf:"varint,2,opt,name=pairing_id,json=pairingId,proto3" json:"pairing_id,omitempty"`
}

func (m *PubKeyNotVerifiable) Reset()         { *m = PubKeyNotVerifiable{} }
func (m *PubKeyNotVerifiable) String() string { return proto.CompactTextString(m) }
func (*PubKeyNotVerifiable) ProtoMessage()    {}
func (*PubKeyNotVerifiable) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa9902169785e6, []int{0}
}
func (m *PubKeyNotVerifiable) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKeyNotVerifiable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKeyNotVerifiable.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKeyNotVerifiable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKeyNotVerifiable.Merge(m, src)
}
func (m *PubKeyNotVerifiable) XXX_Size() int {
	return m.Size()
}
func (m *PubKeyNotVerifiable) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKeyNotVerifiable.DiscardUnknown(m)
}

var xxx_messageInfo_PubKeyNotVerifiable proto.InternalMessageInfo

func (m *PubKeyNotVerifiable) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PubKeyNotVerifiable) GetPairingId() uint64 {
	if m != nil {
		return m.PairingId
	}
	return 0
}

type PubKey struct {
	User              []byte    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PairingPublicKey  types.Any `protobuf:"bytes,2,opt,name=pairing_public_key,json=pairingPublicKey,proto3" json:"pairing_public_key"`
	OperatorPublicKey types.Any `protobuf:"bytes,3,opt,name=operator_public_key,json=operatorPublicKey,proto3" json:"operator_public_key"`
}

func (m *PubKey) Reset()         { *m = PubKey{} }
func (m *PubKey) String() string { return proto.CompactTextString(m) }
func (*PubKey) ProtoMessage()    {}
func (*PubKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_98aa9902169785e6, []int{1}
}
func (m *PubKey) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PubKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PubKey.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PubKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubKey.Merge(m, src)
}
func (m *PubKey) XXX_Size() int {
	return m.Size()
}
func (m *PubKey) XXX_DiscardUnknown() {
	xxx_messageInfo_PubKey.DiscardUnknown(m)
}

var xxx_messageInfo_PubKey proto.InternalMessageInfo

func (m *PubKey) GetUser() []byte {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *PubKey) GetPairingPublicKey() types.Any {
	if m != nil {
		return m.PairingPublicKey
	}
	return types.Any{}
}

func (m *PubKey) GetOperatorPublicKey() types.Any {
	if m != nil {
		return m.OperatorPublicKey
	}
	return types.Any{}
}

func init() {
	proto.RegisterType((*PubKeyNotVerifiable)(nil), "gluon.customauth.pairing.PubKeyNotVerifiable")
	proto.RegisterType((*PubKey)(nil), "gluon.customauth.pairing.PubKey")
}

func init() {
	proto.RegisterFile("gluon/customauth/pairing/pubkey.proto", fileDescriptor_98aa9902169785e6)
}

var fileDescriptor_98aa9902169785e6 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xcf, 0x29, 0xcd,
	0xcf, 0xd3, 0x4f, 0x2e, 0x2d, 0x2e, 0xc9, 0xcf, 0x4d, 0x2c, 0x2d, 0xc9, 0xd0, 0x2f, 0x48, 0xcc,
	0x2c, 0xca, 0xcc, 0x4b, 0xd7, 0x2f, 0x28, 0x4d, 0xca, 0x4e, 0xad, 0xd4, 0x2b, 0x28, 0xca, 0x2f,
	0xc9, 0x17, 0x92, 0x00, 0x2b, 0xd3, 0x43, 0x28, 0xd3, 0x83, 0x2a, 0x93, 0x12, 0x49, 0xcf, 0x4f,
	0xcf, 0x07, 0x2b, 0xd2, 0x07, 0xb1, 0x20, 0xea, 0xa5, 0x24, 0xd3, 0xf3, 0xf3, 0xd3, 0x73, 0x52,
	0xf5, 0xc1, 0xbc, 0xa4, 0xd2, 0x34, 0xfd, 0xc4, 0x3c, 0xa8, 0x51, 0x4a, 0x1e, 0x5c, 0xc2, 0x01,
	0xa5, 0x49, 0xde, 0xa9, 0x95, 0x7e, 0xf9, 0x25, 0x61, 0xa9, 0x45, 0x99, 0x69, 0x99, 0x89, 0x49,
	0x39, 0xa9, 0x42, 0x42, 0x5c, 0x2c, 0xa5, 0xc5, 0xa9, 0x45, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x3c,
	0x41, 0x60, 0xb6, 0x90, 0x2c, 0x17, 0x17, 0xd4, 0x9a, 0xf8, 0xcc, 0x14, 0x09, 0x26, 0x05, 0x46,
	0x0d, 0x96, 0x20, 0x4e, 0xa8, 0x88, 0x67, 0x8a, 0xd2, 0x26, 0x46, 0x2e, 0x36, 0x88, 0x51, 0x58,
	0x75, 0x7b, 0x70, 0x09, 0xc1, 0x74, 0x17, 0x94, 0x26, 0xe5, 0x64, 0x26, 0xc7, 0x67, 0xa7, 0x56,
	0x82, 0x4d, 0xe1, 0x36, 0x12, 0xd1, 0x83, 0x38, 0x50, 0x0f, 0xe6, 0x40, 0x3d, 0xc7, 0xbc, 0x4a,
	0x27, 0x96, 0x13, 0xf7, 0xe4, 0x19, 0x82, 0x04, 0xa0, 0xba, 0x02, 0xc0, 0x9a, 0x40, 0xa6, 0x7b,
	0x71, 0x09, 0xe7, 0x17, 0xa4, 0x16, 0x25, 0x96, 0xe4, 0x17, 0x21, 0x1b, 0xc5, 0x4c, 0xd0, 0x28,
	0x41, 0x98, 0x36, 0xb8, 0x59, 0x4e, 0x56, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0xa5, 0x00, 0x89, 0x8a, 0x0a, 0xe4, 0xc8, 0x28, 0xa9, 0x2c, 0x48, 0x2d, 0x86, 0x45, 0x49,
	0x12, 0x1b, 0xd8, 0x0a, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x84, 0xad, 0xae, 0x42, 0xb5,
	0x01, 0x00, 0x00,
}

func (m *PubKeyNotVerifiable) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKeyNotVerifiable) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKeyNotVerifiable) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PairingId != 0 {
		i = encodeVarintPubkey(dAtA, i, uint64(m.PairingId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintPubkey(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *PubKey) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PubKey) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.OperatorPublicKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPubkey(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	{
		size, err := m.PairingPublicKey.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintPubkey(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.User) > 0 {
		i -= len(m.User)
		copy(dAtA[i:], m.User)
		i = encodeVarintPubkey(dAtA, i, uint64(len(m.User)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintPubkey(dAtA []byte, offset int, v uint64) int {
	offset -= sovPubkey(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PubKeyNotVerifiable) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovPubkey(uint64(l))
	}
	if m.PairingId != 0 {
		n += 1 + sovPubkey(uint64(m.PairingId))
	}
	return n
}

func (m *PubKey) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.User)
	if l > 0 {
		n += 1 + l + sovPubkey(uint64(l))
	}
	l = m.PairingPublicKey.Size()
	n += 1 + l + sovPubkey(uint64(l))
	l = m.OperatorPublicKey.Size()
	n += 1 + l + sovPubkey(uint64(l))
	return n
}

func sovPubkey(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozPubkey(x uint64) (n int) {
	return sovPubkey(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PubKeyNotVerifiable) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPubkey
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
			return fmt.Errorf("proto: PubKeyNotVerifiable: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKeyNotVerifiable: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPubkey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = append(m.User[:0], dAtA[iNdEx:postIndex]...)
			if m.User == nil {
				m.User = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairingId", wireType)
			}
			m.PairingId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.PairingId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipPubkey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPubkey
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
func (m *PubKey) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowPubkey
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
			return fmt.Errorf("proto: PubKey: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PubKey: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkey
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthPubkey
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthPubkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.User = append(m.User[:0], dAtA[iNdEx:postIndex]...)
			if m.User == nil {
				m.User = []byte{}
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PairingPublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkey
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
				return ErrInvalidLengthPubkey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPubkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.PairingPublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OperatorPublicKey", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowPubkey
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
				return ErrInvalidLengthPubkey
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthPubkey
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.OperatorPublicKey.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipPubkey(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthPubkey
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
func skipPubkey(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowPubkey
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
					return 0, ErrIntOverflowPubkey
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
					return 0, ErrIntOverflowPubkey
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
				return 0, ErrInvalidLengthPubkey
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupPubkey
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthPubkey
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthPubkey        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowPubkey          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupPubkey = fmt.Errorf("proto: unexpected end of group")
)
