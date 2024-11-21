// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package contract

import (
	fmt "fmt"
	io "io"
	reflect "reflect"
	sync "sync"

	v1beta1 "cosmossdk.io/api/cosmos/base/v1beta1"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

var (
	md_LazyContract               protoreflect.MessageDescriptor
	fd_LazyContract_id            protoreflect.FieldDescriptor
	fd_LazyContract_creditor      protoreflect.FieldDescriptor
	fd_LazyContract_debtor        protoreflect.FieldDescriptor
	fd_LazyContract_amountEscrow  protoreflect.FieldDescriptor
	fd_LazyContract_amountPending protoreflect.FieldDescriptor
	fd_LazyContract_expiry        protoreflect.FieldDescriptor
)

func init() {
	file_gluon_contract_lazy_settlement_proto_init()
	md_LazyContract = File_gluon_contract_lazy_settlement_proto.Messages().ByName("LazyContract")
	fd_LazyContract_id = md_LazyContract.Fields().ByName("id")
	fd_LazyContract_creditor = md_LazyContract.Fields().ByName("creditor")
	fd_LazyContract_debtor = md_LazyContract.Fields().ByName("debtor")
	fd_LazyContract_amountEscrow = md_LazyContract.Fields().ByName("amountEscrow")
	fd_LazyContract_amountPending = md_LazyContract.Fields().ByName("amountPending")
	fd_LazyContract_expiry = md_LazyContract.Fields().ByName("expiry")
}

var _ protoreflect.Message = (*fastReflection_LazyContract)(nil)

type fastReflection_LazyContract LazyContract

func (x *LazyContract) ProtoReflect() protoreflect.Message {
	return (*fastReflection_LazyContract)(x)
}

func (x *LazyContract) slowProtoReflect() protoreflect.Message {
	mi := &file_gluon_contract_lazy_settlement_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_LazyContract_messageType fastReflection_LazyContract_messageType
var _ protoreflect.MessageType = fastReflection_LazyContract_messageType{}

type fastReflection_LazyContract_messageType struct{}

func (x fastReflection_LazyContract_messageType) Zero() protoreflect.Message {
	return (*fastReflection_LazyContract)(nil)
}
func (x fastReflection_LazyContract_messageType) New() protoreflect.Message {
	return new(fastReflection_LazyContract)
}
func (x fastReflection_LazyContract_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_LazyContract
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_LazyContract) Descriptor() protoreflect.MessageDescriptor {
	return md_LazyContract
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_LazyContract) Type() protoreflect.MessageType {
	return _fastReflection_LazyContract_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_LazyContract) New() protoreflect.Message {
	return new(fastReflection_LazyContract)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_LazyContract) Interface() protoreflect.ProtoMessage {
	return (*LazyContract)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_LazyContract) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_LazyContract_id, value) {
			return
		}
	}
	if x.Creditor != "" {
		value := protoreflect.ValueOfString(x.Creditor)
		if !f(fd_LazyContract_creditor, value) {
			return
		}
	}
	if x.Debtor != "" {
		value := protoreflect.ValueOfString(x.Debtor)
		if !f(fd_LazyContract_debtor, value) {
			return
		}
	}
	if x.AmountEscrow != nil {
		value := protoreflect.ValueOfMessage(x.AmountEscrow.ProtoReflect())
		if !f(fd_LazyContract_amountEscrow, value) {
			return
		}
	}
	if x.AmountPending != nil {
		value := protoreflect.ValueOfMessage(x.AmountPending.ProtoReflect())
		if !f(fd_LazyContract_amountPending, value) {
			return
		}
	}
	if x.Expiry != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Expiry)
		if !f(fd_LazyContract_expiry, value) {
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
func (x *fastReflection_LazyContract) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "gluon.contract.LazyContract.id":
		return x.Id != uint64(0)
	case "gluon.contract.LazyContract.creditor":
		return x.Creditor != ""
	case "gluon.contract.LazyContract.debtor":
		return x.Debtor != ""
	case "gluon.contract.LazyContract.amountEscrow":
		return x.AmountEscrow != nil
	case "gluon.contract.LazyContract.amountPending":
		return x.AmountPending != nil
	case "gluon.contract.LazyContract.expiry":
		return x.Expiry != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.LazyContract"))
		}
		panic(fmt.Errorf("message gluon.contract.LazyContract does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LazyContract) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "gluon.contract.LazyContract.id":
		x.Id = uint64(0)
	case "gluon.contract.LazyContract.creditor":
		x.Creditor = ""
	case "gluon.contract.LazyContract.debtor":
		x.Debtor = ""
	case "gluon.contract.LazyContract.amountEscrow":
		x.AmountEscrow = nil
	case "gluon.contract.LazyContract.amountPending":
		x.AmountPending = nil
	case "gluon.contract.LazyContract.expiry":
		x.Expiry = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.LazyContract"))
		}
		panic(fmt.Errorf("message gluon.contract.LazyContract does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_LazyContract) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "gluon.contract.LazyContract.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "gluon.contract.LazyContract.creditor":
		value := x.Creditor
		return protoreflect.ValueOfString(value)
	case "gluon.contract.LazyContract.debtor":
		value := x.Debtor
		return protoreflect.ValueOfString(value)
	case "gluon.contract.LazyContract.amountEscrow":
		value := x.AmountEscrow
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "gluon.contract.LazyContract.amountPending":
		value := x.AmountPending
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "gluon.contract.LazyContract.expiry":
		value := x.Expiry
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.LazyContract"))
		}
		panic(fmt.Errorf("message gluon.contract.LazyContract does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_LazyContract) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "gluon.contract.LazyContract.id":
		x.Id = value.Uint()
	case "gluon.contract.LazyContract.creditor":
		x.Creditor = value.Interface().(string)
	case "gluon.contract.LazyContract.debtor":
		x.Debtor = value.Interface().(string)
	case "gluon.contract.LazyContract.amountEscrow":
		x.AmountEscrow = value.Message().Interface().(*v1beta1.Coin)
	case "gluon.contract.LazyContract.amountPending":
		x.AmountPending = value.Message().Interface().(*v1beta1.Coin)
	case "gluon.contract.LazyContract.expiry":
		x.Expiry = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.LazyContract"))
		}
		panic(fmt.Errorf("message gluon.contract.LazyContract does not contain field %s", fd.FullName()))
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
func (x *fastReflection_LazyContract) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.contract.LazyContract.amountEscrow":
		if x.AmountEscrow == nil {
			x.AmountEscrow = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AmountEscrow.ProtoReflect())
	case "gluon.contract.LazyContract.amountPending":
		if x.AmountPending == nil {
			x.AmountPending = new(v1beta1.Coin)
		}
		return protoreflect.ValueOfMessage(x.AmountPending.ProtoReflect())
	case "gluon.contract.LazyContract.id":
		panic(fmt.Errorf("field id of message gluon.contract.LazyContract is not mutable"))
	case "gluon.contract.LazyContract.creditor":
		panic(fmt.Errorf("field creditor of message gluon.contract.LazyContract is not mutable"))
	case "gluon.contract.LazyContract.debtor":
		panic(fmt.Errorf("field debtor of message gluon.contract.LazyContract is not mutable"))
	case "gluon.contract.LazyContract.expiry":
		panic(fmt.Errorf("field expiry of message gluon.contract.LazyContract is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.LazyContract"))
		}
		panic(fmt.Errorf("message gluon.contract.LazyContract does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_LazyContract) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.contract.LazyContract.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "gluon.contract.LazyContract.creditor":
		return protoreflect.ValueOfString("")
	case "gluon.contract.LazyContract.debtor":
		return protoreflect.ValueOfString("")
	case "gluon.contract.LazyContract.amountEscrow":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "gluon.contract.LazyContract.amountPending":
		m := new(v1beta1.Coin)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "gluon.contract.LazyContract.expiry":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.LazyContract"))
		}
		panic(fmt.Errorf("message gluon.contract.LazyContract does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_LazyContract) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in gluon.contract.LazyContract", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_LazyContract) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_LazyContract) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_LazyContract) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_LazyContract) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*LazyContract)
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
		if x.Id != 0 {
			n += 1 + runtime.Sov(uint64(x.Id))
		}
		l = len(x.Creditor)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.Debtor)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AmountEscrow != nil {
			l = options.Size(x.AmountEscrow)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.AmountPending != nil {
			l = options.Size(x.AmountPending)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Expiry != 0 {
			n += 1 + runtime.Sov(uint64(x.Expiry))
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
		x := input.Message.Interface().(*LazyContract)
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
		if x.Expiry != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Expiry))
			i--
			dAtA[i] = 0x30
		}
		if x.AmountPending != nil {
			encoded, err := options.Marshal(x.AmountPending)
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
		if x.AmountEscrow != nil {
			encoded, err := options.Marshal(x.AmountEscrow)
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
			dAtA[i] = 0x22
		}
		if len(x.Debtor) > 0 {
			i -= len(x.Debtor)
			copy(dAtA[i:], x.Debtor)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Debtor)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.Creditor) > 0 {
			i -= len(x.Creditor)
			copy(dAtA[i:], x.Creditor)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creditor)))
			i--
			dAtA[i] = 0x12
		}
		if x.Id != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Id))
			i--
			dAtA[i] = 0x8
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
		x := input.Message.Interface().(*LazyContract)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LazyContract: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: LazyContract: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
				}
				x.Id = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Id |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creditor", wireType)
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
				x.Creditor = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Debtor", wireType)
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
				x.Debtor = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AmountEscrow", wireType)
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
				if x.AmountEscrow == nil {
					x.AmountEscrow = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AmountEscrow); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field AmountPending", wireType)
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
				if x.AmountPending == nil {
					x.AmountPending = &v1beta1.Coin{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.AmountPending); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 6:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Expiry", wireType)
				}
				x.Expiry = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Expiry |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
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
// source: gluon/contract/lazy_settlement.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type LazyContract struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            uint64        `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Creditor      string        `protobuf:"bytes,2,opt,name=creditor,proto3" json:"creditor,omitempty"`
	Debtor        string        `protobuf:"bytes,3,opt,name=debtor,proto3" json:"debtor,omitempty"`
	AmountEscrow  *v1beta1.Coin `protobuf:"bytes,4,opt,name=amountEscrow,proto3" json:"amountEscrow,omitempty"`
	AmountPending *v1beta1.Coin `protobuf:"bytes,5,opt,name=amountPending,proto3" json:"amountPending,omitempty"`
	Expiry        uint64        `protobuf:"varint,6,opt,name=expiry,proto3" json:"expiry,omitempty"`
}

func (x *LazyContract) Reset() {
	*x = LazyContract{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gluon_contract_lazy_settlement_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LazyContract) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LazyContract) ProtoMessage() {}

// Deprecated: Use LazyContract.ProtoReflect.Descriptor instead.
func (*LazyContract) Descriptor() ([]byte, []int) {
	return file_gluon_contract_lazy_settlement_proto_rawDescGZIP(), []int{0}
}

func (x *LazyContract) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *LazyContract) GetCreditor() string {
	if x != nil {
		return x.Creditor
	}
	return ""
}

func (x *LazyContract) GetDebtor() string {
	if x != nil {
		return x.Debtor
	}
	return ""
}

func (x *LazyContract) GetAmountEscrow() *v1beta1.Coin {
	if x != nil {
		return x.AmountEscrow
	}
	return nil
}

func (x *LazyContract) GetAmountPending() *v1beta1.Coin {
	if x != nil {
		return x.AmountPending
	}
	return nil
}

func (x *LazyContract) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

var File_gluon_contract_lazy_settlement_proto protoreflect.FileDescriptor

var file_gluon_contract_lazy_settlement_proto_rawDesc = []byte{
	0x0a, 0x24, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2f, 0x6c, 0x61, 0x7a, 0x79, 0x5f, 0x73, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2f, 0x63, 0x6f, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xf8, 0x01, 0x0a,
	0x0e, 0x4c, 0x61, 0x7a, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x64,
	0x65, 0x62, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x62,
	0x74, 0x6f, 0x72, 0x12, 0x43, 0x0a, 0x0c, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x73, 0x63,
	0x72, 0x6f, 0x77, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d,
	0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x0c, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x73, 0x63, 0x72, 0x6f, 0x77, 0x12, 0x45, 0x0a, 0x0d, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x69, 0x6e, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00,
	0x52, 0x0d, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x12,
	0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x42, 0x9c, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e,
	0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x13,
	0x4c, 0x61, 0x7a, 0x79, 0x53, 0x65, 0x74, 0x74, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x18, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69,
	0x2f, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0xa2,
	0x02, 0x03, 0x47, 0x43, 0x58, 0xaa, 0x02, 0x0e, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0xca, 0x02, 0x0e, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0xe2, 0x02, 0x1a, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0f, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x3a, 0x3a, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gluon_contract_lazy_settlement_proto_rawDescOnce sync.Once
	file_gluon_contract_lazy_settlement_proto_rawDescData = file_gluon_contract_lazy_settlement_proto_rawDesc
)

func file_gluon_contract_lazy_settlement_proto_rawDescGZIP() []byte {
	file_gluon_contract_lazy_settlement_proto_rawDescOnce.Do(func() {
		file_gluon_contract_lazy_settlement_proto_rawDescData = protoimpl.X.CompressGZIP(file_gluon_contract_lazy_settlement_proto_rawDescData)
	})
	return file_gluon_contract_lazy_settlement_proto_rawDescData
}

var file_gluon_contract_lazy_settlement_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gluon_contract_lazy_settlement_proto_goTypes = []interface{}{
	(*LazyContract)(nil), // 0: gluon.contract.LazyContract
	(*v1beta1.Coin)(nil),   // 1: cosmos.base.v1beta1.Coin
}
var file_gluon_contract_lazy_settlement_proto_depIdxs = []int32{
	1, // 0: gluon.contract.LazyContract.amountEscrow:type_name -> cosmos.base.v1beta1.Coin
	1, // 1: gluon.contract.LazyContract.amountPending:type_name -> cosmos.base.v1beta1.Coin
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gluon_contract_lazy_settlement_proto_init() }
func file_gluon_contract_lazy_settlement_proto_init() {
	if File_gluon_contract_lazy_settlement_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gluon_contract_lazy_settlement_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LazyContract); i {
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
			RawDescriptor: file_gluon_contract_lazy_settlement_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gluon_contract_lazy_settlement_proto_goTypes,
		DependencyIndexes: file_gluon_contract_lazy_settlement_proto_depIdxs,
		MessageInfos:      file_gluon_contract_lazy_settlement_proto_msgTypes,
	}.Build()
	File_gluon_contract_lazy_settlement_proto = out.File
	file_gluon_contract_lazy_settlement_proto_rawDesc = nil
	file_gluon_contract_lazy_settlement_proto_goTypes = nil
	file_gluon_contract_lazy_settlement_proto_depIdxs = nil
}
