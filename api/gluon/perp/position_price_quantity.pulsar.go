// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package perp

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	_ "github.com/cosmos/cosmos-proto"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_PositionPriceQuantity                     protoreflect.MessageDescriptor
	fd_PositionPriceQuantity_owner               protoreflect.FieldDescriptor
	fd_PositionPriceQuantity_position_order_hash protoreflect.FieldDescriptor
	fd_PositionPriceQuantity_price               protoreflect.FieldDescriptor
	fd_PositionPriceQuantity_quantity            protoreflect.FieldDescriptor
)

func init() {
	file_gluon_perp_position_price_quantity_proto_init()
	md_PositionPriceQuantity = File_gluon_perp_position_price_quantity_proto.Messages().ByName("PositionPriceQuantity")
	fd_PositionPriceQuantity_owner = md_PositionPriceQuantity.Fields().ByName("owner")
	fd_PositionPriceQuantity_position_order_hash = md_PositionPriceQuantity.Fields().ByName("position_order_hash")
	fd_PositionPriceQuantity_price = md_PositionPriceQuantity.Fields().ByName("price")
	fd_PositionPriceQuantity_quantity = md_PositionPriceQuantity.Fields().ByName("quantity")
}

var _ protoreflect.Message = (*fastReflection_PositionPriceQuantity)(nil)

type fastReflection_PositionPriceQuantity PositionPriceQuantity

func (x *PositionPriceQuantity) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PositionPriceQuantity)(x)
}

func (x *PositionPriceQuantity) slowProtoReflect() protoreflect.Message {
	mi := &file_gluon_perp_position_price_quantity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PositionPriceQuantity_messageType fastReflection_PositionPriceQuantity_messageType
var _ protoreflect.MessageType = fastReflection_PositionPriceQuantity_messageType{}

type fastReflection_PositionPriceQuantity_messageType struct{}

func (x fastReflection_PositionPriceQuantity_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PositionPriceQuantity)(nil)
}
func (x fastReflection_PositionPriceQuantity_messageType) New() protoreflect.Message {
	return new(fastReflection_PositionPriceQuantity)
}
func (x fastReflection_PositionPriceQuantity_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PositionPriceQuantity
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PositionPriceQuantity) Descriptor() protoreflect.MessageDescriptor {
	return md_PositionPriceQuantity
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PositionPriceQuantity) Type() protoreflect.MessageType {
	return _fastReflection_PositionPriceQuantity_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PositionPriceQuantity) New() protoreflect.Message {
	return new(fastReflection_PositionPriceQuantity)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PositionPriceQuantity) Interface() protoreflect.ProtoMessage {
	return (*PositionPriceQuantity)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PositionPriceQuantity) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Owner != "" {
		value := protoreflect.ValueOfString(x.Owner)
		if !f(fd_PositionPriceQuantity_owner, value) {
			return
		}
	}
	if x.PositionOrderHash != "" {
		value := protoreflect.ValueOfString(x.PositionOrderHash)
		if !f(fd_PositionPriceQuantity_position_order_hash, value) {
			return
		}
	}
	if x.Price != "" {
		value := protoreflect.ValueOfString(x.Price)
		if !f(fd_PositionPriceQuantity_price, value) {
			return
		}
	}
	if x.Quantity != "" {
		value := protoreflect.ValueOfString(x.Quantity)
		if !f(fd_PositionPriceQuantity_quantity, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_PositionPriceQuantity) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "gluon.perp.PositionPriceQuantity.owner":
		return x.Owner != ""
	case "gluon.perp.PositionPriceQuantity.position_order_hash":
		return x.PositionOrderHash != ""
	case "gluon.perp.PositionPriceQuantity.price":
		return x.Price != ""
	case "gluon.perp.PositionPriceQuantity.quantity":
		return x.Quantity != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.PositionPriceQuantity"))
		}
		panic(fmt.Errorf("message gluon.perp.PositionPriceQuantity does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PositionPriceQuantity) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "gluon.perp.PositionPriceQuantity.owner":
		x.Owner = ""
	case "gluon.perp.PositionPriceQuantity.position_order_hash":
		x.PositionOrderHash = ""
	case "gluon.perp.PositionPriceQuantity.price":
		x.Price = ""
	case "gluon.perp.PositionPriceQuantity.quantity":
		x.Quantity = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.PositionPriceQuantity"))
		}
		panic(fmt.Errorf("message gluon.perp.PositionPriceQuantity does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PositionPriceQuantity) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "gluon.perp.PositionPriceQuantity.owner":
		value := x.Owner
		return protoreflect.ValueOfString(value)
	case "gluon.perp.PositionPriceQuantity.position_order_hash":
		value := x.PositionOrderHash
		return protoreflect.ValueOfString(value)
	case "gluon.perp.PositionPriceQuantity.price":
		value := x.Price
		return protoreflect.ValueOfString(value)
	case "gluon.perp.PositionPriceQuantity.quantity":
		value := x.Quantity
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.PositionPriceQuantity"))
		}
		panic(fmt.Errorf("message gluon.perp.PositionPriceQuantity does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PositionPriceQuantity) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "gluon.perp.PositionPriceQuantity.owner":
		x.Owner = value.Interface().(string)
	case "gluon.perp.PositionPriceQuantity.position_order_hash":
		x.PositionOrderHash = value.Interface().(string)
	case "gluon.perp.PositionPriceQuantity.price":
		x.Price = value.Interface().(string)
	case "gluon.perp.PositionPriceQuantity.quantity":
		x.Quantity = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.PositionPriceQuantity"))
		}
		panic(fmt.Errorf("message gluon.perp.PositionPriceQuantity does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PositionPriceQuantity) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.perp.PositionPriceQuantity.owner":
		panic(fmt.Errorf("field owner of message gluon.perp.PositionPriceQuantity is not mutable"))
	case "gluon.perp.PositionPriceQuantity.position_order_hash":
		panic(fmt.Errorf("field position_order_hash of message gluon.perp.PositionPriceQuantity is not mutable"))
	case "gluon.perp.PositionPriceQuantity.price":
		panic(fmt.Errorf("field price of message gluon.perp.PositionPriceQuantity is not mutable"))
	case "gluon.perp.PositionPriceQuantity.quantity":
		panic(fmt.Errorf("field quantity of message gluon.perp.PositionPriceQuantity is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.PositionPriceQuantity"))
		}
		panic(fmt.Errorf("message gluon.perp.PositionPriceQuantity does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PositionPriceQuantity) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.perp.PositionPriceQuantity.owner":
		return protoreflect.ValueOfString("")
	case "gluon.perp.PositionPriceQuantity.position_order_hash":
		return protoreflect.ValueOfString("")
	case "gluon.perp.PositionPriceQuantity.price":
		return protoreflect.ValueOfString("")
	case "gluon.perp.PositionPriceQuantity.quantity":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.PositionPriceQuantity"))
		}
		panic(fmt.Errorf("message gluon.perp.PositionPriceQuantity does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PositionPriceQuantity) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in gluon.perp.PositionPriceQuantity", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PositionPriceQuantity) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PositionPriceQuantity) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_PositionPriceQuantity) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PositionPriceQuantity) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PositionPriceQuantity)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		l = len(x.Owner)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PositionOrderHash)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Price)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Quantity)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*PositionPriceQuantity)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Quantity) > 0 {
			i -= len(x.Quantity)
			copy(dAtA[i:], x.Quantity)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Quantity)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.Price) > 0 {
			i -= len(x.Price)
			copy(dAtA[i:], x.Price)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Price)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.PositionOrderHash) > 0 {
			i -= len(x.PositionOrderHash)
			copy(dAtA[i:], x.PositionOrderHash)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PositionOrderHash)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.Owner) > 0 {
			i -= len(x.Owner)
			copy(dAtA[i:], x.Owner)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Owner)))
			i--
			dAtA[i] = 0xa
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*PositionPriceQuantity)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PositionPriceQuantity: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PositionPriceQuantity: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Owner = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PositionOrderHash", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.PositionOrderHash = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Price", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Price = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Quantity", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Quantity = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: gluon/perp/position_price_quantity.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// PositionPriceQuantity
type PositionPriceQuantity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Owner             string `protobuf:"bytes,1,opt,name=owner,proto3" json:"owner,omitempty"`
	PositionOrderHash string `protobuf:"bytes,2,opt,name=position_order_hash,json=positionOrderHash,proto3" json:"position_order_hash,omitempty"`
	Price             string `protobuf:"bytes,3,opt,name=price,proto3" json:"price,omitempty"`
	Quantity          string `protobuf:"bytes,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *PositionPriceQuantity) Reset() {
	*x = PositionPriceQuantity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gluon_perp_position_price_quantity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PositionPriceQuantity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PositionPriceQuantity) ProtoMessage() {}

// Deprecated: Use PositionPriceQuantity.ProtoReflect.Descriptor instead.
func (*PositionPriceQuantity) Descriptor() ([]byte, []int) {
	return file_gluon_perp_position_price_quantity_proto_rawDescGZIP(), []int{0}
}

func (x *PositionPriceQuantity) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *PositionPriceQuantity) GetPositionOrderHash() string {
	if x != nil {
		return x.PositionOrderHash
	}
	return ""
}

func (x *PositionPriceQuantity) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

func (x *PositionPriceQuantity) GetQuantity() string {
	if x != nil {
		return x.Quantity
	}
	return ""
}

var File_gluon_perp_position_price_quantity_proto protoreflect.FileDescriptor

var file_gluon_perp_position_price_quantity_proto_rawDesc = []byte{
	0x0a, 0x28, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x70, 0x2f, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x5f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x67, 0x6c, 0x75, 0x6f,
	0x6e, 0x2e, 0x70, 0x65, 0x72, 0x70, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67,
	0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x01, 0x0a, 0x15, 0x50, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x2e, 0x0a, 0x13, 0x70, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x68, 0x61, 0x73, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a,
	0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x2b, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x15, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73,
	0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68, 0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4,
	0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0x52, 0x08, 0x71, 0x75,
	0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x42, 0x8b, 0x01, 0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67,
	0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x72, 0x70, 0x42, 0x1a, 0x50, 0x6f, 0x73, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x50, 0x72, 0x69, 0x63, 0x65, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x14, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x70, 0xa2, 0x02, 0x03,
	0x47, 0x50, 0x58, 0xaa, 0x02, 0x0a, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x70,
	0xca, 0x02, 0x0a, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x50, 0x65, 0x72, 0x70, 0xe2, 0x02, 0x16,
	0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x50, 0x65, 0x72, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0b, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x3a, 0x3a,
	0x50, 0x65, 0x72, 0x70, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gluon_perp_position_price_quantity_proto_rawDescOnce sync.Once
	file_gluon_perp_position_price_quantity_proto_rawDescData = file_gluon_perp_position_price_quantity_proto_rawDesc
)

func file_gluon_perp_position_price_quantity_proto_rawDescGZIP() []byte {
	file_gluon_perp_position_price_quantity_proto_rawDescOnce.Do(func() {
		file_gluon_perp_position_price_quantity_proto_rawDescData = protoimpl.X.CompressGZIP(file_gluon_perp_position_price_quantity_proto_rawDescData)
	})
	return file_gluon_perp_position_price_quantity_proto_rawDescData
}

var file_gluon_perp_position_price_quantity_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gluon_perp_position_price_quantity_proto_goTypes = []interface{}{
	(*PositionPriceQuantity)(nil), // 0: gluon.perp.PositionPriceQuantity
}
var file_gluon_perp_position_price_quantity_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gluon_perp_position_price_quantity_proto_init() }
func file_gluon_perp_position_price_quantity_proto_init() {
	if File_gluon_perp_position_price_quantity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gluon_perp_position_price_quantity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PositionPriceQuantity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gluon_perp_position_price_quantity_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gluon_perp_position_price_quantity_proto_goTypes,
		DependencyIndexes: file_gluon_perp_position_price_quantity_proto_depIdxs,
		MessageInfos:      file_gluon_perp_position_price_quantity_proto_msgTypes,
	}.Build()
	File_gluon_perp_position_price_quantity_proto = out.File
	file_gluon_perp_position_price_quantity_proto_rawDesc = nil
	file_gluon_perp_position_price_quantity_proto_goTypes = nil
	file_gluon_perp_position_price_quantity_proto_depIdxs = nil
}
