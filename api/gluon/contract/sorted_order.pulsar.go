// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package contract

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
	md_SortedOrder                   protoreflect.MessageDescriptor
	fd_SortedOrder_expiry            protoreflect.FieldDescriptor
	fd_SortedOrder_id                protoreflect.FieldDescriptor
	fd_SortedOrder_cancelled         protoreflect.FieldDescriptor
	fd_SortedOrder_contracted_amount protoreflect.FieldDescriptor
)

func init() {
	file_gluon_contract_sorted_order_proto_init()
	md_SortedOrder = File_gluon_contract_sorted_order_proto.Messages().ByName("SortedOrder")
	fd_SortedOrder_expiry = md_SortedOrder.Fields().ByName("expiry")
	fd_SortedOrder_id = md_SortedOrder.Fields().ByName("id")
	fd_SortedOrder_cancelled = md_SortedOrder.Fields().ByName("cancelled")
	fd_SortedOrder_contracted_amount = md_SortedOrder.Fields().ByName("contracted_amount")
}

var _ protoreflect.Message = (*fastReflection_SortedOrder)(nil)

type fastReflection_SortedOrder SortedOrder

func (x *SortedOrder) ProtoReflect() protoreflect.Message {
	return (*fastReflection_SortedOrder)(x)
}

func (x *SortedOrder) slowProtoReflect() protoreflect.Message {
	mi := &file_gluon_contract_sorted_order_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_SortedOrder_messageType fastReflection_SortedOrder_messageType
var _ protoreflect.MessageType = fastReflection_SortedOrder_messageType{}

type fastReflection_SortedOrder_messageType struct{}

func (x fastReflection_SortedOrder_messageType) Zero() protoreflect.Message {
	return (*fastReflection_SortedOrder)(nil)
}
func (x fastReflection_SortedOrder_messageType) New() protoreflect.Message {
	return new(fastReflection_SortedOrder)
}
func (x fastReflection_SortedOrder_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_SortedOrder
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_SortedOrder) Descriptor() protoreflect.MessageDescriptor {
	return md_SortedOrder
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_SortedOrder) Type() protoreflect.MessageType {
	return _fastReflection_SortedOrder_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_SortedOrder) New() protoreflect.Message {
	return new(fastReflection_SortedOrder)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_SortedOrder) Interface() protoreflect.ProtoMessage {
	return (*SortedOrder)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_SortedOrder) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Expiry != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Expiry)
		if !f(fd_SortedOrder_expiry, value) {
			return
		}
	}
	if x.Id != "" {
		value := protoreflect.ValueOfString(x.Id)
		if !f(fd_SortedOrder_id, value) {
			return
		}
	}
	if x.Cancelled != false {
		value := protoreflect.ValueOfBool(x.Cancelled)
		if !f(fd_SortedOrder_cancelled, value) {
			return
		}
	}
	if x.ContractedAmount != "" {
		value := protoreflect.ValueOfString(x.ContractedAmount)
		if !f(fd_SortedOrder_contracted_amount, value) {
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
func (x *fastReflection_SortedOrder) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "gluon.contract.SortedOrder.expiry":
		return x.Expiry != uint64(0)
	case "gluon.contract.SortedOrder.id":
		return x.Id != ""
	case "gluon.contract.SortedOrder.cancelled":
		return x.Cancelled != false
	case "gluon.contract.SortedOrder.contracted_amount":
		return x.ContractedAmount != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.SortedOrder"))
		}
		panic(fmt.Errorf("message gluon.contract.SortedOrder does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SortedOrder) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "gluon.contract.SortedOrder.expiry":
		x.Expiry = uint64(0)
	case "gluon.contract.SortedOrder.id":
		x.Id = ""
	case "gluon.contract.SortedOrder.cancelled":
		x.Cancelled = false
	case "gluon.contract.SortedOrder.contracted_amount":
		x.ContractedAmount = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.SortedOrder"))
		}
		panic(fmt.Errorf("message gluon.contract.SortedOrder does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_SortedOrder) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "gluon.contract.SortedOrder.expiry":
		value := x.Expiry
		return protoreflect.ValueOfUint64(value)
	case "gluon.contract.SortedOrder.id":
		value := x.Id
		return protoreflect.ValueOfString(value)
	case "gluon.contract.SortedOrder.cancelled":
		value := x.Cancelled
		return protoreflect.ValueOfBool(value)
	case "gluon.contract.SortedOrder.contracted_amount":
		value := x.ContractedAmount
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.SortedOrder"))
		}
		panic(fmt.Errorf("message gluon.contract.SortedOrder does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_SortedOrder) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "gluon.contract.SortedOrder.expiry":
		x.Expiry = value.Uint()
	case "gluon.contract.SortedOrder.id":
		x.Id = value.Interface().(string)
	case "gluon.contract.SortedOrder.cancelled":
		x.Cancelled = value.Bool()
	case "gluon.contract.SortedOrder.contracted_amount":
		x.ContractedAmount = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.SortedOrder"))
		}
		panic(fmt.Errorf("message gluon.contract.SortedOrder does not contain field %s", fd.FullName()))
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
func (x *fastReflection_SortedOrder) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.contract.SortedOrder.expiry":
		panic(fmt.Errorf("field expiry of message gluon.contract.SortedOrder is not mutable"))
	case "gluon.contract.SortedOrder.id":
		panic(fmt.Errorf("field id of message gluon.contract.SortedOrder is not mutable"))
	case "gluon.contract.SortedOrder.cancelled":
		panic(fmt.Errorf("field cancelled of message gluon.contract.SortedOrder is not mutable"))
	case "gluon.contract.SortedOrder.contracted_amount":
		panic(fmt.Errorf("field contracted_amount of message gluon.contract.SortedOrder is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.SortedOrder"))
		}
		panic(fmt.Errorf("message gluon.contract.SortedOrder does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_SortedOrder) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.contract.SortedOrder.expiry":
		return protoreflect.ValueOfUint64(uint64(0))
	case "gluon.contract.SortedOrder.id":
		return protoreflect.ValueOfString("")
	case "gluon.contract.SortedOrder.cancelled":
		return protoreflect.ValueOfBool(false)
	case "gluon.contract.SortedOrder.contracted_amount":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.contract.SortedOrder"))
		}
		panic(fmt.Errorf("message gluon.contract.SortedOrder does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_SortedOrder) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in gluon.contract.SortedOrder", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_SortedOrder) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_SortedOrder) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_SortedOrder) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_SortedOrder) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*SortedOrder)
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
		if x.Expiry != 0 {
			n += 1 + runtime.Sov(uint64(x.Expiry))
		}
		l = len(x.Id)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Cancelled {
			n += 2
		}
		l = len(x.ContractedAmount)
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
		x := input.Message.Interface().(*SortedOrder)
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
		if len(x.ContractedAmount) > 0 {
			i -= len(x.ContractedAmount)
			copy(dAtA[i:], x.ContractedAmount)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ContractedAmount)))
			i--
			dAtA[i] = 0x22
		}
		if x.Cancelled {
			i--
			if x.Cancelled {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x18
		}
		if len(x.Id) > 0 {
			i -= len(x.Id)
			copy(dAtA[i:], x.Id)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Id)))
			i--
			dAtA[i] = 0x12
		}
		if x.Expiry != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Expiry))
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
		x := input.Message.Interface().(*SortedOrder)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SortedOrder: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: SortedOrder: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
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
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
				x.Id = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Cancelled", wireType)
				}
				var v int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				x.Cancelled = bool(v != 0)
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ContractedAmount", wireType)
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
				x.ContractedAmount = string(dAtA[iNdEx:postIndex])
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
// source: gluon/contract/sorted_order.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// SortedOrder
type SortedOrder struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Expiry           uint64 `protobuf:"varint,1,opt,name=expiry,proto3" json:"expiry,omitempty"`
	Id               string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Cancelled        bool   `protobuf:"varint,3,opt,name=cancelled,proto3" json:"cancelled,omitempty"`
	ContractedAmount string `protobuf:"bytes,4,opt,name=contracted_amount,json=contractedAmount,proto3" json:"contracted_amount,omitempty"`
}

func (x *SortedOrder) Reset() {
	*x = SortedOrder{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gluon_contract_sorted_order_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SortedOrder) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SortedOrder) ProtoMessage() {}

// Deprecated: Use SortedOrder.ProtoReflect.Descriptor instead.
func (*SortedOrder) Descriptor() ([]byte, []int) {
	return file_gluon_contract_sorted_order_proto_rawDescGZIP(), []int{0}
}

func (x *SortedOrder) GetExpiry() uint64 {
	if x != nil {
		return x.Expiry
	}
	return 0
}

func (x *SortedOrder) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SortedOrder) GetCancelled() bool {
	if x != nil {
		return x.Cancelled
	}
	return false
}

func (x *SortedOrder) GetContractedAmount() string {
	if x != nil {
		return x.ContractedAmount
	}
	return ""
}

var File_gluon_contract_sorted_order_proto protoreflect.FileDescriptor

var file_gluon_contract_sorted_order_proto_rawDesc = []byte{
	0x0a, 0x21, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x2f, 0x73, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x5f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72,
	0x61, 0x63, 0x74, 0x1a, 0x19, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x5f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x14,
	0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xad, 0x01, 0x0a, 0x0b, 0x53, 0x6f, 0x72, 0x74, 0x65, 0x64, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x65, 0x78, 0x70, 0x69, 0x72, 0x79, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x09, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x12, 0x58, 0x0a, 0x11, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x42, 0x2b, 0xc8, 0xde, 0x1f, 0x00, 0xda, 0xde, 0x1f, 0x15, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x6d, 0x61, 0x74, 0x68,
	0x2e, 0x49, 0x6e, 0x74, 0xd2, 0xb4, 0x2d, 0x0a, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2e, 0x49,
	0x6e, 0x74, 0x52, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x65, 0x64, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x42, 0x99, 0x01, 0x0a, 0x12, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6c, 0x75,
	0x6f, 0x6e, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0x42, 0x10, 0x53, 0x6f, 0x72,
	0x74, 0x65, 0x64, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x18, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6c, 0x75, 0x6f, 0x6e,
	0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74, 0xa2, 0x02, 0x03, 0x47, 0x43, 0x58, 0xaa,
	0x02, 0x0e, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0xca, 0x02, 0x0e, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63,
	0x74, 0xe2, 0x02, 0x1a, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61,
	0x63, 0x74, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x0f, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x61, 0x63, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gluon_contract_sorted_order_proto_rawDescOnce sync.Once
	file_gluon_contract_sorted_order_proto_rawDescData = file_gluon_contract_sorted_order_proto_rawDesc
)

func file_gluon_contract_sorted_order_proto_rawDescGZIP() []byte {
	file_gluon_contract_sorted_order_proto_rawDescOnce.Do(func() {
		file_gluon_contract_sorted_order_proto_rawDescData = protoimpl.X.CompressGZIP(file_gluon_contract_sorted_order_proto_rawDescData)
	})
	return file_gluon_contract_sorted_order_proto_rawDescData
}

var file_gluon_contract_sorted_order_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gluon_contract_sorted_order_proto_goTypes = []interface{}{
	(*SortedOrder)(nil), // 0: gluon.contract.SortedOrder
}
var file_gluon_contract_sorted_order_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gluon_contract_sorted_order_proto_init() }
func file_gluon_contract_sorted_order_proto_init() {
	if File_gluon_contract_sorted_order_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gluon_contract_sorted_order_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SortedOrder); i {
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
			RawDescriptor: file_gluon_contract_sorted_order_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gluon_contract_sorted_order_proto_goTypes,
		DependencyIndexes: file_gluon_contract_sorted_order_proto_depIdxs,
		MessageInfos:      file_gluon_contract_sorted_order_proto_msgTypes,
	}.Build()
	File_gluon_contract_sorted_order_proto = out.File
	file_gluon_contract_sorted_order_proto_rawDesc = nil
	file_gluon_contract_sorted_order_proto_goTypes = nil
	file_gluon_contract_sorted_order_proto_depIdxs = nil
}
