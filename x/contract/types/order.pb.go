// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: gluon/contract/order.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/cosmos-sdk/types/tx/amino"
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

type OrderDirection int32

const (
	OrderDirection_UNKNOWN OrderDirection = 0
	OrderDirection_BUY     OrderDirection = 1
	OrderDirection_SELL    OrderDirection = 2
)

var OrderDirection_name = map[int32]string{
	0: "UNKNOWN",
	1: "BUY",
	2: "SELL",
}

var OrderDirection_value = map[string]int32{
	"UNKNOWN": 0,
	"BUY":     1,
	"SELL":    2,
}

func (x OrderDirection) String() string {
	return proto.EnumName(OrderDirection_name, int32(x))
}

func (OrderDirection) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_a010d6fb1943fa40, []int{0}
}

type Order struct {
	Id                    string                       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Address               string                       `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	DenomBase             string                       `protobuf:"bytes,3,opt,name=denom_base,json=denomBase,proto3" json:"denom_base,omitempty"`
	DenomQuote            string                       `protobuf:"bytes,4,opt,name=denom_quote,json=denomQuote,proto3" json:"denom_quote,omitempty"`
	Direction             OrderDirection               `protobuf:"varint,5,opt,name=direction,proto3,enum=gluon.contract.OrderDirection" json:"direction,omitempty"`
	Amount                cosmossdk_io_math.Int        `protobuf:"bytes,6,opt,name=amount,proto3,customtype=cosmossdk.io/math.Int" json:"amount"`
	LimitPrice            *cosmossdk_io_math.LegacyDec `protobuf:"bytes,7,opt,name=limit_price,json=limitPrice,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"limit_price,omitempty"`
	Expiry                time.Time                    `protobuf:"bytes,8,opt,name=expiry,proto3,stdtime" json:"expiry"`
	LazyContract          bool                         `protobuf:"varint,9,opt,name=lazy_contract,json=lazyContract,proto3" json:"lazy_contract,omitempty"`
	AllowLazyCounterparty bool                         `protobuf:"varint,10,opt,name=allow_lazy_counterparty,json=allowLazyCounterparty,proto3" json:"allow_lazy_counterparty,omitempty"`
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
	return OrderDirection_UNKNOWN
}

func (m *Order) GetExpiry() time.Time {
	if m != nil {
		return m.Expiry
	}
	return time.Time{}
}

func (m *Order) GetLazyContract() bool {
	if m != nil {
		return m.LazyContract
	}
	return false
}

func (m *Order) GetAllowLazyCounterparty() bool {
	if m != nil {
		return m.AllowLazyCounterparty
	}
	return false
}

func init() {
	proto.RegisterEnum("gluon.contract.OrderDirection", OrderDirection_name, OrderDirection_value)
	proto.RegisterType((*Order)(nil), "gluon.contract.Order")
}

func init() { proto.RegisterFile("gluon/contract/order.proto", fileDescriptor_a010d6fb1943fa40) }

var fileDescriptor_a010d6fb1943fa40 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x52, 0xcd, 0x6e, 0xd3, 0x4c,
	0x14, 0xcd, 0xa4, 0x69, 0x7e, 0x26, 0xdf, 0x17, 0x85, 0x51, 0x0b, 0x43, 0x10, 0x76, 0x04, 0x9b,
	0xa8, 0x52, 0xed, 0x28, 0x48, 0x5d, 0x75, 0x43, 0x08, 0x12, 0x15, 0x51, 0x0a, 0x2e, 0x55, 0x05,
	0x9b, 0x68, 0x62, 0x0f, 0x66, 0x84, 0xed, 0x31, 0xe3, 0x89, 0xa8, 0x79, 0x01, 0xb6, 0x7d, 0x0c,
	0x96, 0x2c, 0xfa, 0x10, 0x59, 0x56, 0x5d, 0xa1, 0x2e, 0x02, 0x4a, 0x16, 0xbc, 0x06, 0xf2, 0x8c,
	0xad, 0x50, 0xb1, 0xb1, 0x7c, 0xcf, 0x39, 0xf7, 0xcc, 0x99, 0x7b, 0x07, 0x76, 0xfc, 0x60, 0xce,
	0x23, 0xdb, 0xe5, 0x91, 0x14, 0xc4, 0x95, 0x36, 0x17, 0x1e, 0x15, 0x56, 0x2c, 0xb8, 0xe4, 0xa8,
	0xa5, 0x38, 0xab, 0xe0, 0x3a, 0x77, 0x48, 0xc8, 0x22, 0x6e, 0xab, 0xaf, 0x96, 0x74, 0xee, 0xbb,
	0x3c, 0x09, 0x79, 0x32, 0x55, 0x95, 0xad, 0x8b, 0x9c, 0xda, 0xf1, 0xb9, 0xcf, 0x35, 0x9e, 0xfd,
	0xe5, 0xa8, 0xe9, 0x73, 0xee, 0x07, 0xd4, 0x56, 0xd5, 0x6c, 0xfe, 0xde, 0x96, 0x2c, 0xa4, 0x89,
	0x24, 0x61, 0xac, 0x05, 0x8f, 0xbe, 0x56, 0xe0, 0xf6, 0x71, 0x16, 0x02, 0xb5, 0x60, 0x99, 0x79,
	0x18, 0x74, 0x41, 0xaf, 0xe1, 0x94, 0x99, 0x87, 0x06, 0xb0, 0x46, 0x3c, 0x4f, 0xd0, 0x24, 0xc1,
	0xe5, 0x0c, 0x1c, 0xe2, 0xeb, 0xcb, 0xfd, 0x9d, 0xfc, 0xcc, 0xa7, 0x9a, 0x39, 0x91, 0x82, 0x45,
	0xbe, 0x53, 0x08, 0xd1, 0x43, 0x08, 0x3d, 0x1a, 0xf1, 0x70, 0x3a, 0x23, 0x09, 0xc5, 0x5b, 0xca,
	0xab, 0xa1, 0x90, 0x21, 0x49, 0x28, 0x32, 0x61, 0x53, 0xd3, 0x9f, 0xe6, 0x5c, 0x52, 0x5c, 0x51,
	0xbc, 0xee, 0x78, 0x9d, 0x21, 0xe8, 0x10, 0x36, 0x3c, 0x26, 0xa8, 0x2b, 0x19, 0x8f, 0xf0, 0x76,
	0x17, 0xf4, 0x5a, 0x03, 0xc3, 0xba, 0x3d, 0x16, 0x4b, 0xa5, 0x1d, 0x15, 0x2a, 0x67, 0xd3, 0x80,
	0x5e, 0xc0, 0x2a, 0x09, 0xf9, 0x3c, 0x92, 0xb8, 0xaa, 0x02, 0xf7, 0x17, 0x4b, 0xb3, 0x74, 0xb3,
	0x34, 0x77, 0x75, 0xe8, 0xc4, 0xfb, 0x68, 0x31, 0x6e, 0x87, 0x44, 0x7e, 0xb0, 0x8e, 0x22, 0x79,
	0x7d, 0xb9, 0x0f, 0xf3, 0xdb, 0x1c, 0x45, 0xf2, 0xdb, 0xef, 0xef, 0x7b, 0xc0, 0xc9, 0xfb, 0xd1,
	0x19, 0x6c, 0x06, 0x2c, 0x64, 0x72, 0x1a, 0x0b, 0xe6, 0x52, 0x5c, 0x53, 0x76, 0x07, 0x8b, 0xa5,
	0x09, 0x6e, 0x96, 0xe6, 0x83, 0x7f, 0xed, 0xc6, 0xd4, 0x27, 0x6e, 0x3a, 0xa2, 0xee, 0x5f, 0xa6,
	0x23, 0xea, 0x6a, 0x53, 0xa8, 0xac, 0x5e, 0x65, 0x4e, 0xe8, 0x10, 0x56, 0xe9, 0x79, 0xcc, 0x44,
	0x8a, 0xeb, 0x5d, 0xd0, 0x6b, 0x0e, 0x3a, 0x96, 0x5e, 0x90, 0x55, 0x2c, 0xc8, 0x7a, 0x53, 0x2c,
	0x68, 0x58, 0xcf, 0xe2, 0x5f, 0xfc, 0x34, 0x81, 0x93, 0xf7, 0xa0, 0xc7, 0xf0, 0xff, 0x80, 0x7c,
	0x49, 0xa7, 0xc5, 0x2c, 0x70, 0xa3, 0x0b, 0x7a, 0x75, 0xe7, 0xbf, 0x0c, 0x7c, 0x96, 0x63, 0xe8,
	0x00, 0xde, 0x23, 0x41, 0xc0, 0x3f, 0x4f, 0x73, 0xe9, 0x3c, 0x92, 0x54, 0xc4, 0x44, 0xc8, 0x14,
	0x43, 0x25, 0xdf, 0x55, 0xf4, 0x58, 0xf5, 0x6c, 0xc8, 0xbd, 0x3e, 0x6c, 0xdd, 0x1e, 0x2d, 0x6a,
	0xc2, 0xda, 0xe9, 0xe4, 0xe5, 0xe4, 0xf8, 0x6c, 0xd2, 0x2e, 0xa1, 0x1a, 0xdc, 0x1a, 0x9e, 0xbe,
	0x6d, 0x03, 0x54, 0x87, 0x95, 0x93, 0xe7, 0xe3, 0x71, 0xbb, 0x3c, 0xec, 0x2f, 0x56, 0x06, 0xb8,
	0x5a, 0x19, 0xe0, 0xd7, 0xca, 0x00, 0x17, 0x6b, 0xa3, 0x74, 0xb5, 0x36, 0x4a, 0x3f, 0xd6, 0x46,
	0xe9, 0xdd, 0x5d, 0xfd, 0xcc, 0xcf, 0x37, 0x0f, 0x5d, 0xa6, 0x31, 0x4d, 0x66, 0x55, 0x75, 0xcd,
	0x27, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x01, 0x7a, 0xff, 0xf2, 0x07, 0x03, 0x00, 0x00,
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
	if m.AllowLazyCounterparty {
		i--
		if m.AllowLazyCounterparty {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x50
	}
	if m.LazyContract {
		i--
		if m.LazyContract {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
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
	if m.LazyContract {
		n += 2
	}
	if m.AllowLazyCounterparty {
		n += 2
	}
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
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LazyContract", wireType)
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
			m.LazyContract = bool(v != 0)
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AllowLazyCounterparty", wireType)
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
			m.AllowLazyCounterparty = bool(v != 0)
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
