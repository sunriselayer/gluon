// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package perp

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	durationpb "google.golang.org/protobuf/types/known/durationpb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_FundingRate             protoreflect.MessageDescriptor
	fd_FundingRate_denom_base  protoreflect.FieldDescriptor
	fd_FundingRate_denom_quote protoreflect.FieldDescriptor
	fd_FundingRate_seconds     protoreflect.FieldDescriptor
	fd_FundingRate_nanos       protoreflect.FieldDescriptor
	fd_FundingRate_interval    protoreflect.FieldDescriptor
	fd_FundingRate_rate        protoreflect.FieldDescriptor
	fd_FundingRate_price       protoreflect.FieldDescriptor
)

func init() {
	file_gluon_perp_funding_rate_proto_init()
	md_FundingRate = File_gluon_perp_funding_rate_proto.Messages().ByName("FundingRate")
	fd_FundingRate_denom_base = md_FundingRate.Fields().ByName("denom_base")
	fd_FundingRate_denom_quote = md_FundingRate.Fields().ByName("denom_quote")
	fd_FundingRate_seconds = md_FundingRate.Fields().ByName("seconds")
	fd_FundingRate_nanos = md_FundingRate.Fields().ByName("nanos")
	fd_FundingRate_interval = md_FundingRate.Fields().ByName("interval")
	fd_FundingRate_rate = md_FundingRate.Fields().ByName("rate")
	fd_FundingRate_price = md_FundingRate.Fields().ByName("price")
}

var _ protoreflect.Message = (*fastReflection_FundingRate)(nil)

type fastReflection_FundingRate FundingRate

func (x *FundingRate) ProtoReflect() protoreflect.Message {
	return (*fastReflection_FundingRate)(x)
}

func (x *FundingRate) slowProtoReflect() protoreflect.Message {
	mi := &file_gluon_perp_funding_rate_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_FundingRate_messageType fastReflection_FundingRate_messageType
var _ protoreflect.MessageType = fastReflection_FundingRate_messageType{}

type fastReflection_FundingRate_messageType struct{}

func (x fastReflection_FundingRate_messageType) Zero() protoreflect.Message {
	return (*fastReflection_FundingRate)(nil)
}
func (x fastReflection_FundingRate_messageType) New() protoreflect.Message {
	return new(fastReflection_FundingRate)
}
func (x fastReflection_FundingRate_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_FundingRate
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_FundingRate) Descriptor() protoreflect.MessageDescriptor {
	return md_FundingRate
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_FundingRate) Type() protoreflect.MessageType {
	return _fastReflection_FundingRate_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_FundingRate) New() protoreflect.Message {
	return new(fastReflection_FundingRate)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_FundingRate) Interface() protoreflect.ProtoMessage {
	return (*FundingRate)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_FundingRate) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.DenomBase != "" {
		value := protoreflect.ValueOfString(x.DenomBase)
		if !f(fd_FundingRate_denom_base, value) {
			return
		}
	}
	if x.DenomQuote != "" {
		value := protoreflect.ValueOfString(x.DenomQuote)
		if !f(fd_FundingRate_denom_quote, value) {
			return
		}
	}
	if x.Seconds != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Seconds)
		if !f(fd_FundingRate_seconds, value) {
			return
		}
	}
	if x.Nanos != uint32(0) {
		value := protoreflect.ValueOfUint32(x.Nanos)
		if !f(fd_FundingRate_nanos, value) {
			return
		}
	}
	if x.Interval != nil {
		value := protoreflect.ValueOfMessage(x.Interval.ProtoReflect())
		if !f(fd_FundingRate_interval, value) {
			return
		}
	}
	if x.Rate != "" {
		value := protoreflect.ValueOfString(x.Rate)
		if !f(fd_FundingRate_rate, value) {
			return
		}
	}
	if x.Price != "" {
		value := protoreflect.ValueOfString(x.Price)
		if !f(fd_FundingRate_price, value) {
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
func (x *fastReflection_FundingRate) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "gluon.perp.FundingRate.denom_base":
		return x.DenomBase != ""
	case "gluon.perp.FundingRate.denom_quote":
		return x.DenomQuote != ""
	case "gluon.perp.FundingRate.seconds":
		return x.Seconds != uint64(0)
	case "gluon.perp.FundingRate.nanos":
		return x.Nanos != uint32(0)
	case "gluon.perp.FundingRate.interval":
		return x.Interval != nil
	case "gluon.perp.FundingRate.rate":
		return x.Rate != ""
	case "gluon.perp.FundingRate.price":
		return x.Price != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.FundingRate"))
		}
		panic(fmt.Errorf("message gluon.perp.FundingRate does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FundingRate) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "gluon.perp.FundingRate.denom_base":
		x.DenomBase = ""
	case "gluon.perp.FundingRate.denom_quote":
		x.DenomQuote = ""
	case "gluon.perp.FundingRate.seconds":
		x.Seconds = uint64(0)
	case "gluon.perp.FundingRate.nanos":
		x.Nanos = uint32(0)
	case "gluon.perp.FundingRate.interval":
		x.Interval = nil
	case "gluon.perp.FundingRate.rate":
		x.Rate = ""
	case "gluon.perp.FundingRate.price":
		x.Price = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.FundingRate"))
		}
		panic(fmt.Errorf("message gluon.perp.FundingRate does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_FundingRate) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "gluon.perp.FundingRate.denom_base":
		value := x.DenomBase
		return protoreflect.ValueOfString(value)
	case "gluon.perp.FundingRate.denom_quote":
		value := x.DenomQuote
		return protoreflect.ValueOfString(value)
	case "gluon.perp.FundingRate.seconds":
		value := x.Seconds
		return protoreflect.ValueOfUint64(value)
	case "gluon.perp.FundingRate.nanos":
		value := x.Nanos
		return protoreflect.ValueOfUint32(value)
	case "gluon.perp.FundingRate.interval":
		value := x.Interval
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "gluon.perp.FundingRate.rate":
		value := x.Rate
		return protoreflect.ValueOfString(value)
	case "gluon.perp.FundingRate.price":
		value := x.Price
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.FundingRate"))
		}
		panic(fmt.Errorf("message gluon.perp.FundingRate does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_FundingRate) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "gluon.perp.FundingRate.denom_base":
		x.DenomBase = value.Interface().(string)
	case "gluon.perp.FundingRate.denom_quote":
		x.DenomQuote = value.Interface().(string)
	case "gluon.perp.FundingRate.seconds":
		x.Seconds = value.Uint()
	case "gluon.perp.FundingRate.nanos":
		x.Nanos = uint32(value.Uint())
	case "gluon.perp.FundingRate.interval":
		x.Interval = value.Message().Interface().(*durationpb.Duration)
	case "gluon.perp.FundingRate.rate":
		x.Rate = value.Interface().(string)
	case "gluon.perp.FundingRate.price":
		x.Price = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.FundingRate"))
		}
		panic(fmt.Errorf("message gluon.perp.FundingRate does not contain field %s", fd.FullName()))
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
func (x *fastReflection_FundingRate) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.perp.FundingRate.interval":
		if x.Interval == nil {
			x.Interval = new(durationpb.Duration)
		}
		return protoreflect.ValueOfMessage(x.Interval.ProtoReflect())
	case "gluon.perp.FundingRate.denom_base":
		panic(fmt.Errorf("field denom_base of message gluon.perp.FundingRate is not mutable"))
	case "gluon.perp.FundingRate.denom_quote":
		panic(fmt.Errorf("field denom_quote of message gluon.perp.FundingRate is not mutable"))
	case "gluon.perp.FundingRate.seconds":
		panic(fmt.Errorf("field seconds of message gluon.perp.FundingRate is not mutable"))
	case "gluon.perp.FundingRate.nanos":
		panic(fmt.Errorf("field nanos of message gluon.perp.FundingRate is not mutable"))
	case "gluon.perp.FundingRate.rate":
		panic(fmt.Errorf("field rate of message gluon.perp.FundingRate is not mutable"))
	case "gluon.perp.FundingRate.price":
		panic(fmt.Errorf("field price of message gluon.perp.FundingRate is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.FundingRate"))
		}
		panic(fmt.Errorf("message gluon.perp.FundingRate does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_FundingRate) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.perp.FundingRate.denom_base":
		return protoreflect.ValueOfString("")
	case "gluon.perp.FundingRate.denom_quote":
		return protoreflect.ValueOfString("")
	case "gluon.perp.FundingRate.seconds":
		return protoreflect.ValueOfUint64(uint64(0))
	case "gluon.perp.FundingRate.nanos":
		return protoreflect.ValueOfUint32(uint32(0))
	case "gluon.perp.FundingRate.interval":
		m := new(durationpb.Duration)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "gluon.perp.FundingRate.rate":
		return protoreflect.ValueOfString("")
	case "gluon.perp.FundingRate.price":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.perp.FundingRate"))
		}
		panic(fmt.Errorf("message gluon.perp.FundingRate does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_FundingRate) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in gluon.perp.FundingRate", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_FundingRate) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_FundingRate) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_FundingRate) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_FundingRate) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*FundingRate)
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
		l = len(x.DenomBase)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DenomQuote)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Seconds != 0 {
			n += 1 + runtime.Sov(uint64(x.Seconds))
		}
		if x.Nanos != 0 {
			n += 1 + runtime.Sov(uint64(x.Nanos))
		}
		if x.Interval != nil {
			l = options.Size(x.Interval)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Rate)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Price)
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
		x := input.Message.Interface().(*FundingRate)
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
		if len(x.Price) > 0 {
			i -= len(x.Price)
			copy(dAtA[i:], x.Price)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Price)))
			i--
			dAtA[i] = 0x3a
		}
		if len(x.Rate) > 0 {
			i -= len(x.Rate)
			copy(dAtA[i:], x.Rate)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Rate)))
			i--
			dAtA[i] = 0x32
		}
		if x.Interval != nil {
			encoded, err := options.Marshal(x.Interval)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Nanos != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Nanos))
			i--
			dAtA[i] = 0x20
		}
		if x.Seconds != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Seconds))
			i--
			dAtA[i] = 0x18
		}
		if len(x.DenomQuote) > 0 {
			i -= len(x.DenomQuote)
			copy(dAtA[i:], x.DenomQuote)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DenomQuote)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.DenomBase) > 0 {
			i -= len(x.DenomBase)
			copy(dAtA[i:], x.DenomBase)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DenomBase)))
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
		x := input.Message.Interface().(*FundingRate)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FundingRate: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: FundingRate: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DenomBase", wireType)
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
				x.DenomBase = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DenomQuote", wireType)
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
				x.DenomQuote = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Seconds", wireType)
				}
				x.Seconds = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Seconds |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Nanos", wireType)
				}
				x.Nanos = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Nanos |= uint32(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.Interval == nil {
					x.Interval = &durationpb.Duration{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Interval); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Rate", wireType)
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
				x.Rate = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 7:
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
// source: gluon/perp/funding_rate.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// FundingRate
type FundingRate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DenomBase  string               `protobuf:"bytes,1,opt,name=denom_base,json=denomBase,proto3" json:"denom_base,omitempty"`
	DenomQuote string               `protobuf:"bytes,2,opt,name=denom_quote,json=denomQuote,proto3" json:"denom_quote,omitempty"`
	Seconds    uint64               `protobuf:"varint,3,opt,name=seconds,proto3" json:"seconds,omitempty"`
	Nanos      uint32               `protobuf:"varint,4,opt,name=nanos,proto3" json:"nanos,omitempty"`
	Interval   *durationpb.Duration `protobuf:"bytes,5,opt,name=interval,proto3" json:"interval,omitempty"`
	Rate       string               `protobuf:"bytes,6,opt,name=rate,proto3" json:"rate,omitempty"`
	Price      string               `protobuf:"bytes,7,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *FundingRate) Reset() {
	*x = FundingRate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gluon_perp_funding_rate_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FundingRate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FundingRate) ProtoMessage() {}

// Deprecated: Use FundingRate.ProtoReflect.Descriptor instead.
func (*FundingRate) Descriptor() ([]byte, []int) {
	return file_gluon_perp_funding_rate_proto_rawDescGZIP(), []int{0}
}

func (x *FundingRate) GetDenomBase() string {
	if x != nil {
		return x.DenomBase
	}
	return ""
}

func (x *FundingRate) GetDenomQuote() string {
	if x != nil {
		return x.DenomQuote
	}
	return ""
}

func (x *FundingRate) GetSeconds() uint64 {
	if x != nil {
		return x.Seconds
	}
	return 0
}

func (x *FundingRate) GetNanos() uint32 {
	if x != nil {
		return x.Nanos
	}
	return 0
}

func (x *FundingRate) GetInterval() *durationpb.Duration {
	if x != nil {
		return x.Interval
	}
	return nil
}

func (x *FundingRate) GetRate() string {
	if x != nil {
		return x.Rate
	}
	return ""
}

func (x *FundingRate) GetPrice() string {
	if x != nil {
		return x.Price
	}
	return ""
}

var File_gluon_perp_funding_rate_proto protoreflect.FileDescriptor

var file_gluon_perp_funding_rate_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x70, 0x2f, 0x66, 0x75, 0x6e,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0a, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x72, 0x70, 0x1a, 0x14, 0x67, 0x6f, 0x67,
	0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x1a, 0x1e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xe8, 0x01, 0x0a, 0x0b, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74,
	0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f, 0x62, 0x61, 0x73, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x42, 0x61, 0x73, 0x65,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x5f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x51, 0x75, 0x6f, 0x74,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x73, 0x65, 0x63, 0x6f, 0x6e, 0x64, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6e,
	0x61, 0x6e, 0x6f, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x6e, 0x61, 0x6e, 0x6f,
	0x73, 0x12, 0x3f, 0x0a, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x61, 0x6c, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x44, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x08,
	0xc8, 0xde, 0x1f, 0x00, 0x98, 0xdf, 0x1f, 0x01, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x61, 0x6c, 0x12, 0x12, 0x0a, 0x04, 0x72, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x72, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x42, 0x81, 0x01, 0x0a,
	0x0e, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x70, 0x65, 0x72, 0x70, 0x42,
	0x10, 0x46, 0x75, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x14, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x70, 0x65, 0x72, 0x70, 0xa2, 0x02, 0x03, 0x47, 0x50, 0x58, 0xaa,
	0x02, 0x0a, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x50, 0x65, 0x72, 0x70, 0xca, 0x02, 0x0a, 0x47,
	0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x50, 0x65, 0x72, 0x70, 0xe2, 0x02, 0x16, 0x47, 0x6c, 0x75, 0x6f,
	0x6e, 0x5c, 0x50, 0x65, 0x72, 0x70, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x0b, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x3a, 0x3a, 0x50, 0x65, 0x72, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gluon_perp_funding_rate_proto_rawDescOnce sync.Once
	file_gluon_perp_funding_rate_proto_rawDescData = file_gluon_perp_funding_rate_proto_rawDesc
)

func file_gluon_perp_funding_rate_proto_rawDescGZIP() []byte {
	file_gluon_perp_funding_rate_proto_rawDescOnce.Do(func() {
		file_gluon_perp_funding_rate_proto_rawDescData = protoimpl.X.CompressGZIP(file_gluon_perp_funding_rate_proto_rawDescData)
	})
	return file_gluon_perp_funding_rate_proto_rawDescData
}

var file_gluon_perp_funding_rate_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gluon_perp_funding_rate_proto_goTypes = []interface{}{
	(*FundingRate)(nil),         // 0: gluon.perp.FundingRate
	(*durationpb.Duration)(nil), // 1: google.protobuf.Duration
}
var file_gluon_perp_funding_rate_proto_depIdxs = []int32{
	1, // 0: gluon.perp.FundingRate.interval:type_name -> google.protobuf.Duration
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_gluon_perp_funding_rate_proto_init() }
func file_gluon_perp_funding_rate_proto_init() {
	if File_gluon_perp_funding_rate_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gluon_perp_funding_rate_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FundingRate); i {
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
			RawDescriptor: file_gluon_perp_funding_rate_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gluon_perp_funding_rate_proto_goTypes,
		DependencyIndexes: file_gluon_perp_funding_rate_proto_depIdxs,
		MessageInfos:      file_gluon_perp_funding_rate_proto_msgTypes,
	}.Build()
	File_gluon_perp_funding_rate_proto = out.File
	file_gluon_perp_funding_rate_proto_rawDesc = nil
	file_gluon_perp_funding_rate_proto_goTypes = nil
	file_gluon_perp_funding_rate_proto_depIdxs = nil
}
