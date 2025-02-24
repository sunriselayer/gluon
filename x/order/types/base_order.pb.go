// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/order/base_order.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// OrderDirection
type OrderDirection int32

const (
	// Unspecified
	OrderDirection_ORDER_DIRECTION_UNSPECIFIED OrderDirection = 0
	// Buy
	OrderDirection_ORDER_DIRECTION_BUY OrderDirection = 1
	// Sell
	OrderDirection_ORDER_DIRECTION_SELL OrderDirection = 2
)

var OrderDirection_name = map[int32]string{
	0: "ORDER_DIRECTION_UNSPECIFIED",
	1: "ORDER_DIRECTION_BUY",
	2: "ORDER_DIRECTION_SELL",
}

var OrderDirection_value = map[string]int32{
	"ORDER_DIRECTION_UNSPECIFIED": 0,
	"ORDER_DIRECTION_BUY":         1,
	"ORDER_DIRECTION_SELL":        2,
}

func (x OrderDirection) String() string {
	return proto.EnumName(OrderDirection_name, int32(x))
}

func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_4ad0c3a6a0688c54, []int{0}
}

// BaseOrder
type BaseOrder struct {
	AddressString string                `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	DenomBase     string                `protobuf:"bytes,2,opt,name=denom_base,json=denomBase,proto3" json:"denom_base,omitempty"`
	DenomQuote    string                `protobuf:"bytes,3,opt,name=denom_quote,json=denomQuote,proto3" json:"denom_quote,omitempty"`
	Direction     OrderDirection        `protobuf:"varint,4,opt,name=direction,proto3,enum=gluon.order.OrderDirection" json:"direction,omitempty"`
	Amount        cosmossdk_io_math.Int `protobuf:"bytes,5,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	// Empty means market order
	LimitPriceString string `protobuf:"bytes,6,opt,name=limit_price,json=limitPrice,proto3" json:"limit_price,omitempty"`
	// Required to prevent operator from utilizing old order
	Expiry time.Time `protobuf:"bytes,7,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *BaseOrder) Reset()         { *m = BaseOrder{} }
func (m *BaseOrder) String() string { return proto.CompactTextString(m) }
func (*BaseOrder) ProtoMessage()    {}
func (*BaseOrder) Descriptor() ([]byte, []int) {
	return fileDescriptor_4ad0c3a6a0688c54, []int{0}
}
func (m *BaseOrder) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseOrder) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseOrder.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseOrder) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseOrder.Merge(m, src)
}
func (m *BaseOrder) XXX_Size() int {
	return m.Size()
}
func (m *BaseOrder) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseOrder.DiscardUnknown(m)
}

var xxx_messageInfo_BaseOrder proto.InternalMessageInfo

func (m *BaseOrder) GetAddressString() string {
	if m != nil {
		return m.AddressString
	}
	return ""
}

func (m *BaseOrder) GetDenomBase() string {
	if m != nil {
		return m.DenomBase
	}
	return ""
}

func (m *BaseOrder) GetDenomQuote() string {
	if m != nil {
		return m.DenomQuote
	}
	return ""
}

func (m *BaseOrder) GetDirection() OrderDirection {
	if m != nil {
		return m.Direction
	}
	return OrderDirection_ORDER_DIRECTION_UNSPECIFIED
}

func (m *BaseOrder) GetLimitPriceString() string {
	if m != nil {
		return m.LimitPriceString
	}
	return ""
}

func (m *BaseOrder) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("gluon.order.OrderDirection", OrderDirection_name, OrderDirection_value)
	proto.RegisterType((*BaseOrder)(nil), "gluon.order.BaseOrder")
}

func init() { proto.RegisterFile("gluon/order/base_order.proto", fileDescriptor_4ad0c3a6a0688c54) }

var fileDescriptor_4ad0c3a6a0688c54 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x92, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xe3, 0x0d, 0x3a, 0xea, 0x8a, 0xa9, 0xf2, 0x8a, 0x08, 0x1d, 0x24, 0x15, 0xa7, 0x02,
	0xaa, 0x23, 0x0d, 0x71, 0x40, 0xe2, 0x42, 0xda, 0x20, 0x45, 0xaa, 0xd6, 0x91, 0x6e, 0x07, 0xb8,
	0x44, 0x69, 0x63, 0x82, 0x45, 0x13, 0x07, 0xdb, 0x95, 0xb6, 0x07, 0xe0, 0xbe, 0x87, 0xd9, 0x43,
	0xec, 0x38, 0xed, 0x84, 0x38, 0x14, 0x94, 0xbe, 0x08, 0xb2, 0x9d, 0xc2, 0xb6, 0x9b, 0xfd, 0xfb,
	0xff, 0xbf, 0xcf, 0xff, 0xef, 0x93, 0xe1, 0xd3, 0x6c, 0xb1, 0x64, 0x85, 0xc7, 0x78, 0x4a, 0xb8,
	0x37, 0x4b, 0x04, 0x89, 0xf5, 0x11, 0x97, 0x9c, 0x49, 0x86, 0x5a, 0x5a, 0xc5, 0x1a, 0x75, 0x9f,
	0xcc, 0x99, 0xc8, 0x99, 0x88, 0xb5, 0xe4, 0x99, 0x8b, 0xf1, 0x75, 0x3b, 0x19, 0xcb, 0x98, 0xe1,
	0xea, 0x54, 0x53, 0x37, 0x63, 0x2c, 0x5b, 0x10, 0x4f, 0xdf, 0x66, 0xcb, 0x2f, 0x9e, 0xa4, 0x39,
	0x11, 0x32, 0xc9, 0x4b, 0x63, 0x78, 0xfe, 0x63, 0x1b, 0x36, 0xfd, 0x44, 0x90, 0x89, 0xea, 0x8f,
	0x86, 0x70, 0x27, 0x49, 0x53, 0x4e, 0x84, 0xb0, 0x41, 0x0f, 0xf4, 0x9b, 0xfe, 0x8b, 0x6a, 0xe5,
	0x3e, 0x7c, 0x6f, 0xd0, 0x54, 0x72, 0x5a, 0x64, 0xd7, 0x17, 0x83, 0x4e, 0xfd, 0xf0, 0x2d, 0x1e,
	0x6d, 0x2a, 0xd1, 0x33, 0x08, 0x53, 0x52, 0xb0, 0x3c, 0x56, 0xb3, 0xd8, 0x5b, 0xaa, 0x4f, 0xd4,
	0xd4, 0x44, 0x3d, 0x84, 0x5c, 0xd8, 0x32, 0xf2, 0xf7, 0x25, 0x93, 0xc4, 0xde, 0xd6, 0xba, 0xa9,
	0xf8, 0xa8, 0x08, 0x7a, 0x0b, 0x9b, 0x29, 0xe5, 0x64, 0x2e, 0x29, 0x2b, 0xec, 0x7b, 0x3d, 0xd0,
	0xdf, 0x3d, 0xd8, 0xc7, 0x37, 0xb6, 0x80, 0x75, 0xd6, 0xd1, 0xc6, 0x12, 0xfd, 0x77, 0xa3, 0x21,
	0x6c, 0x24, 0x39, 0x5b, 0x16, 0xd2, 0xbe, 0xaf, 0xe3, 0xbf, 0xba, 0x5c, 0xb9, 0xd6, 0xaf, 0x95,
	0xfb, 0xc8, 0x24, 0x16, 0xe9, 0x37, 0x4c, 0x99, 0x97, 0x27, 0xf2, 0x2b, 0x0e, 0x0b, 0x79, 0x7d,
	0x31, 0x80, 0xf5, 0x28, 0x61, 0x21, 0xa3, 0xba, 0x14, 0xbd, 0x81, 0xad, 0x05, 0xcd, 0xa9, 0x8c,
	0x4b, 0x4e, 0xe7, 0xc4, 0x6e, 0xe8, 0x4e, 0x9d, 0x6a, 0xe5, 0xb6, 0xc7, 0x0a, 0x1f, 0x29, 0x5a,
	0xcf, 0x0c, 0x17, 0xff, 0x08, 0x7a, 0x07, 0x1b, 0xe4, 0xb4, 0xa4, 0xfc, 0xcc, 0xde, 0xe9, 0x81,
	0x7e, 0xeb, 0xa0, 0x8b, 0xcd, 0xee, 0xf1, 0x66, 0xf7, 0xf8, 0x78, 0xb3, 0x7b, 0xff, 0x81, 0xca,
	0x75, 0xfe, 0xdb, 0x05, 0x51, 0x5d, 0xf3, 0x32, 0x85, 0xbb, 0xb7, 0xc7, 0x42, 0x2e, 0xdc, 0x9f,
	0x44, 0xa3, 0x20, 0x8a, 0x47, 0x61, 0x14, 0x0c, 0x8f, 0xc3, 0xc9, 0x61, 0x7c, 0x72, 0x38, 0x3d,
	0x0a, 0x86, 0xe1, 0x87, 0x30, 0x18, 0xb5, 0x2d, 0xf4, 0x18, 0xee, 0xdd, 0x35, 0xf8, 0x27, 0x9f,
	0xda, 0x00, 0xd9, 0xb0, 0x73, 0x57, 0x98, 0x06, 0xe3, 0x71, 0x7b, 0xcb, 0x1f, 0x5c, 0x56, 0x0e,
	0xb8, 0xaa, 0x1c, 0xf0, 0xa7, 0x72, 0xc0, 0xf9, 0xda, 0xb1, 0xae, 0xd6, 0x8e, 0xf5, 0x73, 0xed,
	0x58, 0x9f, 0xf7, 0xcc, 0x27, 0x3c, 0xad, 0xbf, 0xa1, 0x3c, 0x2b, 0x89, 0x98, 0x35, 0x74, 0xf4,
	0xd7, 0x7f, 0x03, 0x00, 0x00, 0xff, 0xff, 0x83, 0xc7, 0x6e, 0x64, 0xa2, 0x02, 0x00, 0x00,
}

func (m *BaseOrder) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseOrder) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseOrder) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintBaseOrder(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x3a
	if len(m.LimitPriceString) > 0 {
		i -= len(m.LimitPriceString)
		copy(dAtA[i:], m.LimitPriceString)
		i = encodeVarintBaseOrder(dAtA, i, uint64(len(m.LimitPriceString)))
		i--
		dAtA[i] = 0x32
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintBaseOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	if m.Direction != 0 {
		i = encodeVarintBaseOrder(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x20
	}
	if len(m.DenomQuote) > 0 {
		i -= len(m.DenomQuote)
		copy(dAtA[i:], m.DenomQuote)
		i = encodeVarintBaseOrder(dAtA, i, uint64(len(m.DenomQuote)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.DenomBase) > 0 {
		i -= len(m.DenomBase)
		copy(dAtA[i:], m.DenomBase)
		i = encodeVarintBaseOrder(dAtA, i, uint64(len(m.DenomBase)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.AddressString) > 0 {
		i -= len(m.AddressString)
		copy(dAtA[i:], m.AddressString)
		i = encodeVarintBaseOrder(dAtA, i, uint64(len(m.AddressString)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintBaseOrder(dAtA []byte, offset int, v uint64) int {
	offset -= sovBaseOrder(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseOrder) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.AddressString)
	if l > 0 {
		n += 1 + l + sovBaseOrder(uint64(l))
	}
	l = len(m.DenomBase)
	if l > 0 {
		n += 1 + l + sovBaseOrder(uint64(l))
	}
	l = len(m.DenomQuote)
	if l > 0 {
		n += 1 + l + sovBaseOrder(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovBaseOrder(uint64(m.Direction))
	}
	l = m.Amount.Size()
	n += 1 + l + sovBaseOrder(uint64(l))
	l = len(m.LimitPriceString)
	if l > 0 {
		n += 1 + l + sovBaseOrder(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry)
	n += 1 + l + sovBaseOrder(uint64(l))
	return n
}

func sovBaseOrder(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBaseOrder(x uint64) (n int) {
	return sovBaseOrder(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseOrder) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBaseOrder
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
			return fmt.Errorf("proto: BaseOrder: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseOrder: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AddressString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
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
				return ErrInvalidLengthBaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AddressString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomBase", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
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
				return ErrInvalidLengthBaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomQuote", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
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
				return ErrInvalidLengthBaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomQuote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Direction |= OrderDirection(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
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
				return ErrInvalidLengthBaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitPriceString", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
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
				return ErrInvalidLengthBaseOrder
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBaseOrder
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.LimitPriceString = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBaseOrder
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
				return ErrInvalidLengthBaseOrder
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBaseOrder
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
			skippy, err := skipBaseOrder(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBaseOrder
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
func skipBaseOrder(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBaseOrder
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
					return 0, ErrIntOverflowBaseOrder
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
					return 0, ErrIntOverflowBaseOrder
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
				return 0, ErrInvalidLengthBaseOrder
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBaseOrder
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBaseOrder
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBaseOrder        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBaseOrder          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBaseOrder = fmt.Errorf("proto: unexpected end of group")
)
