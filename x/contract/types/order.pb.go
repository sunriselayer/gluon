// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/contract/order.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/msgservice"
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
	return fileDescriptor_a010d6fb1943fa40, []int{0}
}

// Order
// Encoded binary will be used as a signing message.
type Order struct {
	Id         string                `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address    string                `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DenomBase  string                `protobuf:"bytes,3,opt,name=denom_base,json=denomBase,proto3" json:"denom_base,omitempty"`
	DenomQuote string                `protobuf:"bytes,4,opt,name=denom_quote,json=denomQuote,proto3" json:"denom_quote,omitempty"`
	Direction  OrderDirection        `protobuf:"varint,5,opt,name=direction,proto3,enum=gluon.contract.OrderDirection" json:"direction,omitempty"`
	Amount     cosmossdk_io_math.Int `protobuf:"bytes,6,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	// Empty means market order
	LimitPrice *cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=limit_price,json=limitPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"limit_price,omitempty"`
	// Required due to off-chain matching & on-chain contract settlement
	Expiry time.Time `protobuf:"bytes,8,opt,name=expiry,proto3,stdtime" json:"expiry"`
}

func (m *Order) Reset()         { *m = Order{} }
func (m *Order) String() string { return proto.CompactTextString(m) }
func (*Order) ProtoMessage()    {}
func (*Order) Descriptor() ([]byte, []int) {
	return fileDescriptor_a010d6fb1943fa40, []int{0}
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

func (m *Order) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Order) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Order) GetDenomBase() string {
	if m != nil {
		return m.DenomBase
	}
	return ""
}

func (m *Order) GetDenomQuote() string {
	if m != nil {
		return m.DenomQuote
	}
	return ""
}

func (m *Order) GetDirection() OrderDirection {
	if m != nil {
		return m.Direction
	}
	return OrderDirection_ORDER_DIRECTION_UNSPECIFIED
}

func (m *Order) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

func init() {
	proto.RegisterEnum("gluon.contract.OrderDirection", OrderDirection_name, OrderDirection_value)
	proto.RegisterType((*Order)(nil), "gluon.contract.Order")
}

func init() { proto.RegisterFile("gluon/contract/order.proto", fileDescriptor_a010d6fb1943fa40) }

var fileDescriptor_a010d6fb1943fa40 = []byte{
	// 503 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0xad, 0x3b, 0xd6, 0x6d, 0x2e, 0xaa, 0x2a, 0x53, 0x98, 0xe9, 0x44, 0x52, 0x71, 0xaa, 0x86,
	0x96, 0xb0, 0x72, 0x43, 0xbb, 0xd0, 0xa6, 0x48, 0x91, 0xaa, 0x75, 0xb8, 0xdb, 0x01, 0x2e, 0x55,
	0x9a, 0x98, 0x60, 0xd1, 0xc4, 0x21, 0x76, 0xd1, 0x7a, 0xe5, 0x17, 0xec, 0x97, 0x20, 0x0e, 0xfb,
	0x11, 0x3d, 0x4e, 0x3b, 0xa1, 0x1d, 0x0a, 0x6a, 0x0f, 0xfc, 0x0d, 0x14, 0x3b, 0xd1, 0xd8, 0x38,
	0x25, 0xdf, 0x7b, 0xef, 0x7b, 0x7e, 0xfe, 0x3e, 0xc3, 0x66, 0x38, 0x9d, 0xf1, 0xd8, 0xf6, 0x79,
	0x2c, 0x53, 0xcf, 0x97, 0x36, 0x4f, 0x03, 0x9a, 0x5a, 0x49, 0xca, 0x25, 0x47, 0x35, 0xc5, 0x59,
	0x05, 0xd7, 0xdc, 0xf5, 0xb9, 0x88, 0xb8, 0xb0, 0x23, 0x11, 0xda, 0x5f, 0x0f, 0xb3, 0x8f, 0x16,
	0x36, 0x9f, 0x6a, 0x62, 0xac, 0x2a, 0x5b, 0x17, 0x39, 0xd5, 0x08, 0x79, 0xc8, 0x35, 0x9e, 0xfd,
	0xe5, 0xa8, 0x19, 0x72, 0x1e, 0x4e, 0xa9, 0xad, 0xaa, 0xc9, 0xec, 0xa3, 0x2d, 0x59, 0x44, 0x85,
	0xf4, 0xa2, 0x44, 0x0b, 0x9e, 0x7f, 0xdf, 0x80, 0x9b, 0xc3, 0x2c, 0x0a, 0xaa, 0xc1, 0x32, 0x0b,
	0x30, 0x68, 0x81, 0xf6, 0x0e, 0x29, 0xb3, 0x00, 0x75, 0xe0, 0x96, 0x17, 0x04, 0x29, 0x15, 0x02,
	0x97, 0x33, 0xb0, 0x8b, 0xaf, 0x2f, 0x0f, 0x1a, 0xf9, 0x99, 0x6f, 0x34, 0x33, 0x92, 0x29, 0x8b,
	0x43, 0x52, 0x08, 0xd1, 0x33, 0x08, 0x03, 0x1a, 0xf3, 0x68, 0x3c, 0xf1, 0x04, 0xc5, 0x1b, 0xca,
	0x6b, 0x47, 0x21, 0x5d, 0x4f, 0x50, 0x64, 0xc2, 0xaa, 0xa6, 0xbf, 0xcc, 0xb8, 0xa4, 0xf8, 0x81,
	0xe2, 0x75, 0xc7, 0xbb, 0x0c, 0x41, 0x47, 0x70, 0x27, 0x60, 0x29, 0xf5, 0x25, 0xe3, 0x31, 0xde,
	0x6c, 0x81, 0x76, 0xad, 0x63, 0x58, 0x77, 0x87, 0x63, 0xa9, 0xb4, 0x4e, 0xa1, 0x22, 0xb7, 0x0d,
	0xa8, 0x07, 0x2b, 0x5e, 0xc4, 0x67, 0xb1, 0xc4, 0x15, 0x15, 0xf8, 0xc5, 0x62, 0x69, 0x96, 0x6e,
	0x96, 0xe6, 0x63, 0x1d, 0x5a, 0x04, 0x9f, 0x2d, 0xc6, 0xed, 0xc8, 0x93, 0x9f, 0x2c, 0x37, 0x96,
	0xd7, 0x97, 0x07, 0x30, 0xbf, 0x8d, 0x1b, 0x4b, 0x92, 0xb7, 0x22, 0x02, 0xab, 0x53, 0x16, 0x31,
	0x39, 0x4e, 0x52, 0xe6, 0x53, 0xbc, 0xa5, 0x9c, 0x0e, 0x17, 0x4b, 0x13, 0xdc, 0x2c, 0xcd, 0xbd,
	0xff, 0x9d, 0x06, 0x34, 0xf4, 0xfc, 0xb9, 0x43, 0xfd, 0x7f, 0xfc, 0x1c, 0xea, 0x13, 0xa8, 0x5c,
	0x4e, 0x32, 0x13, 0x74, 0x04, 0x2b, 0xf4, 0x3c, 0x61, 0xe9, 0x1c, 0x6f, 0xb7, 0x40, 0xbb, 0xda,
	0x69, 0x5a, 0x7a, 0x2d, 0x56, 0xb1, 0x16, 0xeb, 0xb4, 0x58, 0x4b, 0x77, 0x3b, 0x0b, 0x7d, 0xf1,
	0xcb, 0x04, 0x24, 0xef, 0x79, 0xfd, 0xf0, 0xdb, 0x9f, 0x1f, 0xfb, 0xc5, 0x88, 0xf7, 0x03, 0x58,
	0xbb, 0x3b, 0x01, 0x64, 0xc2, 0xbd, 0x21, 0x71, 0xfa, 0x64, 0xec, 0xb8, 0xa4, 0xdf, 0x3b, 0x75,
	0x87, 0xc7, 0xe3, 0xb3, 0xe3, 0xd1, 0x49, 0xbf, 0xe7, 0xbe, 0x75, 0xfb, 0x4e, 0xbd, 0x84, 0x76,
	0xe1, 0xa3, 0xfb, 0x82, 0xee, 0xd9, 0xfb, 0x3a, 0x40, 0x18, 0x36, 0xee, 0x13, 0xa3, 0xfe, 0x60,
	0x50, 0x2f, 0x77, 0x5f, 0x2e, 0x56, 0x06, 0xb8, 0x5a, 0x19, 0xe0, 0xf7, 0xca, 0x00, 0x17, 0x6b,
	0xa3, 0x74, 0xb5, 0x36, 0x4a, 0x3f, 0xd7, 0x46, 0xe9, 0xc3, 0x13, 0xfd, 0x8e, 0xcf, 0x6f, 0x5f,
	0xb2, 0x9c, 0x27, 0x54, 0x4c, 0x2a, 0xea, 0x2e, 0xaf, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x28,
	0xb3, 0xad, 0x48, 0xe8, 0x02, 0x00, 0x00,
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
	n1, err1 := github_com_cosmos_gogoproto_types.StdTimeMarshalTo(m.Expiry, dAtA[i-github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry):])
	if err1 != nil {
		return 0, err1
	}
	i -= n1
	i = encodeVarintOrder(dAtA, i, uint64(n1))
	i--
	dAtA[i] = 0x42
	if m.LimitPrice != nil {
		{
			size := m.LimitPrice.Size()
			i -= size
			if _, err := m.LimitPrice.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintOrder(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	{
		size := m.Amount.Size()
		i -= size
		if _, err := m.Amount.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintOrder(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x32
	if m.Direction != 0 {
		i = encodeVarintOrder(dAtA, i, uint64(m.Direction))
		i--
		dAtA[i] = 0x28
	}
	if len(m.DenomQuote) > 0 {
		i -= len(m.DenomQuote)
		copy(dAtA[i:], m.DenomQuote)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.DenomQuote)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.DenomBase) > 0 {
		i -= len(m.DenomBase)
		copy(dAtA[i:], m.DenomBase)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.DenomBase)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintOrder(dAtA, i, uint64(len(m.Id)))
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
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.DenomBase)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	l = len(m.DenomQuote)
	if l > 0 {
		n += 1 + l + sovOrder(uint64(l))
	}
	if m.Direction != 0 {
		n += 1 + sovOrder(uint64(m.Direction))
	}
	l = m.Amount.Size()
	n += 1 + l + sovOrder(uint64(l))
	if m.LimitPrice != nil {
		l = m.LimitPrice.Size()
		n += 1 + l + sovOrder(uint64(l))
	}
	l = github_com_cosmos_gogoproto_types.SizeOfStdTime(m.Expiry)
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
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomBase", wireType)
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
			m.DenomBase = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomQuote", wireType)
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
			m.DenomQuote = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Direction", wireType)
			}
			m.Direction = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowOrder
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
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LimitPrice", wireType)
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
			var v cosmossdk_io_math.LegacyDec
			m.LimitPrice = &v
			if err := m.LimitPrice.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
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
			if err := github_com_cosmos_gogoproto_types.StdTimeUnmarshal(&m.Expiry, dAtA[iNdEx:postIndex]); err != nil {
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
