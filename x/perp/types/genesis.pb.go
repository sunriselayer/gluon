// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/perp/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the perp module's genesis state.
type GenesisState struct {
	// params defines all the parameters of the module.
	Params                  Params                  `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Positions               []Position              `protobuf:"bytes,2,rep,name=positions,proto3" json:"positions"`
	PositionPriceQuantities []PositionPriceQuantity `protobuf:"bytes,3,rep,name=position_price_quantities,json=positionPriceQuantities,proto3" json:"position_price_quantities"`
	CrossMargins            []CrossMargin           `protobuf:"bytes,4,rep,name=cross_margins,json=crossMargins,proto3" json:"cross_margins"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdc64cc6f71f903e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetPositions() []Position {
	if m != nil {
		return m.Positions
	}
	return nil
}

func (m *GenesisState) GetPositionPriceQuantities() []PositionPriceQuantity {
	if m != nil {
		return m.PositionPriceQuantities
	}
	return nil
}

func (m *GenesisState) GetCrossMargins() []CrossMargin {
	if m != nil {
		return m.CrossMargins
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "gluon.perp.GenesisState")
}

func init() { proto.RegisterFile("gluon/perp/genesis.proto", fileDescriptor_cdc64cc6f71f903e) }

var fileDescriptor_cdc64cc6f71f903e = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0x93, 0xb6, 0x14, 0xdc, 0xd6, 0xcb, 0x52, 0x48, 0x5a, 0x70, 0xad, 0x9e, 0x72, 0x90,
	0x44, 0xea, 0xc5, 0x73, 0x3c, 0x78, 0x12, 0xaa, 0xde, 0xbc, 0x84, 0x18, 0x96, 0xb0, 0x60, 0xb3,
	0xeb, 0xce, 0x16, 0xec, 0x5b, 0xf8, 0x0e, 0xbe, 0x4c, 0x8f, 0x3d, 0x7a, 0x12, 0x49, 0x5e, 0x44,
	0xb2, 0x19, 0x49, 0x0a, 0xb9, 0x0d, 0xf3, 0xfd, 0xff, 0xcf, 0xfc, 0x43, 0xfc, 0xfc, 0x6d, 0x2b,
	0x8b, 0x48, 0x71, 0xad, 0xa2, 0x9c, 0x17, 0x1c, 0x04, 0x84, 0x4a, 0x4b, 0x23, 0x29, 0xb1, 0x24,
	0xac, 0xc9, 0xc2, 0xeb, 0xa8, 0x54, 0xaa, 0xd3, 0x0d, 0x8a, 0x16, 0xf3, 0x2e, 0x90, 0x20, 0x8c,
	0xa8, 0x2d, 0x16, 0x05, 0x3d, 0x28, 0x51, 0x5a, 0x64, 0x3c, 0x79, 0xdf, 0xa6, 0x85, 0x11, 0x66,
	0x87, 0xca, 0x59, 0x2e, 0x73, 0x69, 0xc7, 0xa8, 0x9e, 0x70, 0x7b, 0xd6, 0xf1, 0x67, 0x5a, 0x02,
	0x24, 0x9b, 0x54, 0xe7, 0x02, 0xe3, 0x2f, 0xbf, 0x06, 0x64, 0x7a, 0xdf, 0x1c, 0xfc, 0x6c, 0x52,
	0xc3, 0xe9, 0x35, 0x19, 0x37, 0xa7, 0xf9, 0xee, 0xd2, 0x0d, 0x26, 0x2b, 0x1a, 0xb6, 0x05, 0xc2,
	0xb5, 0x25, 0xf1, 0x68, 0xff, 0x73, 0xee, 0x3c, 0xa1, 0x8e, 0xde, 0x92, 0x93, 0xff, 0xc3, 0xc0,
	0x1f, 0x2c, 0x87, 0xc1, 0x64, 0x35, 0x3b, 0x32, 0x21, 0x44, 0x5b, 0x2b, 0xa6, 0x19, 0x99, 0xf7,
	0x57, 0x12, 0x1c, 0xfc, 0xa1, 0x4d, 0xba, 0xe8, 0x4b, 0x5a, 0xd7, 0xda, 0x47, 0x6c, 0x8f, 0xb1,
	0x9e, 0xea, 0x81, 0x82, 0x03, 0x8d, 0xc9, 0x69, 0xb7, 0x37, 0xf8, 0x23, 0x1b, 0xec, 0x75, 0x83,
	0xef, 0x6a, 0xc1, 0x83, 0xe5, 0x18, 0x37, 0xcd, 0xda, 0x15, 0xc4, 0x57, 0xfb, 0x92, 0xb9, 0x87,
	0x92, 0xb9, 0xbf, 0x25, 0x73, 0x3f, 0x2b, 0xe6, 0x1c, 0x2a, 0xe6, 0x7c, 0x57, 0xcc, 0x79, 0xa1,
	0xcd, 0x7b, 0x3f, 0x9a, 0x07, 0x9b, 0x9d, 0xe2, 0xf0, 0x3a, 0xb6, 0xaf, 0xbd, 0xf9, 0x0b, 0x00,
	0x00, 0xff, 0xff, 0xb7, 0x8e, 0xca, 0x48, 0x15, 0x02, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.CrossMargins) > 0 {
		for iNdEx := len(m.CrossMargins) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.CrossMargins[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.PositionPriceQuantities) > 0 {
		for iNdEx := len(m.PositionPriceQuantities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PositionPriceQuantities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Positions) > 0 {
		for iNdEx := len(m.Positions) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Positions[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Positions) > 0 {
		for _, e := range m.Positions {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PositionPriceQuantities) > 0 {
		for _, e := range m.PositionPriceQuantities {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.CrossMargins) > 0 {
		for _, e := range m.CrossMargins {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Positions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Positions = append(m.Positions, Position{})
			if err := m.Positions[len(m.Positions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PositionPriceQuantities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PositionPriceQuantities = append(m.PositionPriceQuantities, PositionPriceQuantity{})
			if err := m.PositionPriceQuantities[len(m.PositionPriceQuantities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CrossMargins", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CrossMargins = append(m.CrossMargins, CrossMargin{})
			if err := m.CrossMargins[len(m.CrossMargins)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
