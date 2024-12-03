// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package customauth

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	_ "github.com/cosmos/gogoproto/gogoproto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_Pairing            protoreflect.MessageDescriptor
	fd_Pairing_id         protoreflect.FieldDescriptor
	fd_Pairing_address    protoreflect.FieldDescriptor
	fd_Pairing_public_key protoreflect.FieldDescriptor
	fd_Pairing_created_at protoreflect.FieldDescriptor
)

func init() {
	file_gluon_customauth_pairing_proto_init()
	md_Pairing = File_gluon_customauth_pairing_proto.Messages().ByName("Pairing")
	fd_Pairing_id = md_Pairing.Fields().ByName("id")
	fd_Pairing_address = md_Pairing.Fields().ByName("address")
	fd_Pairing_public_key = md_Pairing.Fields().ByName("public_key")
	fd_Pairing_created_at = md_Pairing.Fields().ByName("created_at")
}

var _ protoreflect.Message = (*fastReflection_Pairing)(nil)

type fastReflection_Pairing Pairing

func (x *Pairing) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Pairing)(x)
}

func (x *Pairing) slowProtoReflect() protoreflect.Message {
	mi := &file_gluon_customauth_pairing_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Pairing_messageType fastReflection_Pairing_messageType
var _ protoreflect.MessageType = fastReflection_Pairing_messageType{}

type fastReflection_Pairing_messageType struct{}

func (x fastReflection_Pairing_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Pairing)(nil)
}
func (x fastReflection_Pairing_messageType) New() protoreflect.Message {
	return new(fastReflection_Pairing)
}
func (x fastReflection_Pairing_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Pairing
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Pairing) Descriptor() protoreflect.MessageDescriptor {
	return md_Pairing
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Pairing) Type() protoreflect.MessageType {
	return _fastReflection_Pairing_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Pairing) New() protoreflect.Message {
	return new(fastReflection_Pairing)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Pairing) Interface() protoreflect.ProtoMessage {
	return (*Pairing)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Pairing) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Pairing_id, value) {
			return
		}
	}
	if x.Address != "" {
		value := protoreflect.ValueOfString(x.Address)
		if !f(fd_Pairing_address, value) {
			return
		}
	}
	if x.PublicKey != nil {
		value := protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
		if !f(fd_Pairing_public_key, value) {
			return
		}
	}
	if x.CreatedAt != nil {
		value := protoreflect.ValueOfMessage(x.CreatedAt.ProtoReflect())
		if !f(fd_Pairing_created_at, value) {
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
func (x *fastReflection_Pairing) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "gluon.customauth.Pairing.id":
		return x.Id != uint64(0)
	case "gluon.customauth.Pairing.address":
		return x.Address != ""
	case "gluon.customauth.Pairing.public_key":
		return x.PublicKey != nil
	case "gluon.customauth.Pairing.created_at":
		return x.CreatedAt != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.Pairing"))
		}
		panic(fmt.Errorf("message gluon.customauth.Pairing does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pairing) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "gluon.customauth.Pairing.id":
		x.Id = uint64(0)
	case "gluon.customauth.Pairing.address":
		x.Address = ""
	case "gluon.customauth.Pairing.public_key":
		x.PublicKey = nil
	case "gluon.customauth.Pairing.created_at":
		x.CreatedAt = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.Pairing"))
		}
		panic(fmt.Errorf("message gluon.customauth.Pairing does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Pairing) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "gluon.customauth.Pairing.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "gluon.customauth.Pairing.address":
		value := x.Address
		return protoreflect.ValueOfString(value)
	case "gluon.customauth.Pairing.public_key":
		value := x.PublicKey
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "gluon.customauth.Pairing.created_at":
		value := x.CreatedAt
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.Pairing"))
		}
		panic(fmt.Errorf("message gluon.customauth.Pairing does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Pairing) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "gluon.customauth.Pairing.id":
		x.Id = value.Uint()
	case "gluon.customauth.Pairing.address":
		x.Address = value.Interface().(string)
	case "gluon.customauth.Pairing.public_key":
		x.PublicKey = value.Message().Interface().(*anypb.Any)
	case "gluon.customauth.Pairing.created_at":
		x.CreatedAt = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.Pairing"))
		}
		panic(fmt.Errorf("message gluon.customauth.Pairing does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Pairing) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.customauth.Pairing.public_key":
		if x.PublicKey == nil {
			x.PublicKey = new(anypb.Any)
		}
		return protoreflect.ValueOfMessage(x.PublicKey.ProtoReflect())
	case "gluon.customauth.Pairing.created_at":
		if x.CreatedAt == nil {
			x.CreatedAt = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.CreatedAt.ProtoReflect())
	case "gluon.customauth.Pairing.id":
		panic(fmt.Errorf("field id of message gluon.customauth.Pairing is not mutable"))
	case "gluon.customauth.Pairing.address":
		panic(fmt.Errorf("field address of message gluon.customauth.Pairing is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.Pairing"))
		}
		panic(fmt.Errorf("message gluon.customauth.Pairing does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Pairing) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.customauth.Pairing.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "gluon.customauth.Pairing.address":
		return protoreflect.ValueOfString("")
	case "gluon.customauth.Pairing.public_key":
		m := new(anypb.Any)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "gluon.customauth.Pairing.created_at":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.Pairing"))
		}
		panic(fmt.Errorf("message gluon.customauth.Pairing does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Pairing) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in gluon.customauth.Pairing", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Pairing) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Pairing) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Pairing) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Pairing) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Pairing)
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
		l = len(x.Address)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.PublicKey != nil {
			l = options.Size(x.PublicKey)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.CreatedAt != nil {
			l = options.Size(x.CreatedAt)
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
		x := input.Message.Interface().(*Pairing)
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
		if x.CreatedAt != nil {
			encoded, err := options.Marshal(x.CreatedAt)
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
		if x.PublicKey != nil {
			encoded, err := options.Marshal(x.PublicKey)
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
			dAtA[i] = 0x1a
		}
		if len(x.Address) > 0 {
			i -= len(x.Address)
			copy(dAtA[i:], x.Address)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Address)))
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
		x := input.Message.Interface().(*Pairing)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pairing: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Pairing: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
				x.Address = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PublicKey", wireType)
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
				if x.PublicKey == nil {
					x.PublicKey = &anypb.Any{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.PublicKey); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field CreatedAt", wireType)
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
				if x.CreatedAt == nil {
					x.CreatedAt = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.CreatedAt); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
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
// source: gluon/customauth/pairing.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Pairing
type Pairing struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64                 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Address   string                 `protobuf:"bytes,2,opt,name=address,proto3" json:"address,omitempty"`
	PublicKey *anypb.Any             `protobuf:"bytes,3,opt,name=public_key,json=publicKey,proto3" json:"public_key,omitempty"`
	CreatedAt *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *Pairing) Reset() {
	*x = Pairing{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gluon_customauth_pairing_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pairing) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pairing) ProtoMessage() {}

// Deprecated: Use Pairing.ProtoReflect.Descriptor instead.
func (*Pairing) Descriptor() ([]byte, []int) {
	return file_gluon_customauth_pairing_proto_rawDescGZIP(), []int{0}
}

func (x *Pairing) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pairing) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *Pairing) GetPublicKey() *anypb.Any {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Pairing) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_gluon_customauth_pairing_proto protoreflect.FileDescriptor

var file_gluon_customauth_pairing_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75,
	0x74, 0x68, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb3, 0x01, 0x0a, 0x07, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x39, 0x0a, 0x0a, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x63, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x42, 0x04, 0xc8, 0xde, 0x1f, 0x00, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c,
	0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x43, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x08, 0xc8, 0xde, 0x1f, 0x00, 0x90, 0xdf, 0x1f, 0x01, 0x52,
	0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x42, 0xa1, 0x01, 0x0a, 0x14, 0x63,
	0x6f, 0x6d, 0x2e, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61,
	0x75, 0x74, 0x68, 0x42, 0x0c, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67,
	0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0xa2,
	0x02, 0x03, 0x47, 0x43, 0x58, 0xaa, 0x02, 0x10, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x43, 0x75,
	0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0xca, 0x02, 0x10, 0x47, 0x6c, 0x75, 0x6f, 0x6e,
	0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0xe2, 0x02, 0x1c, 0x47, 0x6c,
	0x75, 0x6f, 0x6e, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x47, 0x6c, 0x75,
	0x6f, 0x6e, 0x3a, 0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gluon_customauth_pairing_proto_rawDescOnce sync.Once
	file_gluon_customauth_pairing_proto_rawDescData = file_gluon_customauth_pairing_proto_rawDesc
)

func file_gluon_customauth_pairing_proto_rawDescGZIP() []byte {
	file_gluon_customauth_pairing_proto_rawDescOnce.Do(func() {
		file_gluon_customauth_pairing_proto_rawDescData = protoimpl.X.CompressGZIP(file_gluon_customauth_pairing_proto_rawDescData)
	})
	return file_gluon_customauth_pairing_proto_rawDescData
}

var file_gluon_customauth_pairing_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gluon_customauth_pairing_proto_goTypes = []interface{}{
	(*Pairing)(nil),               // 0: gluon.customauth.Pairing
	(*anypb.Any)(nil),             // 1: google.protobuf.Any
	(*timestamppb.Timestamp)(nil), // 2: google.protobuf.Timestamp
}
var file_gluon_customauth_pairing_proto_depIdxs = []int32{
	1, // 0: gluon.customauth.Pairing.public_key:type_name -> google.protobuf.Any
	2, // 1: gluon.customauth.Pairing.created_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_gluon_customauth_pairing_proto_init() }
func file_gluon_customauth_pairing_proto_init() {
	if File_gluon_customauth_pairing_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gluon_customauth_pairing_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pairing); i {
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
			RawDescriptor: file_gluon_customauth_pairing_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gluon_customauth_pairing_proto_goTypes,
		DependencyIndexes: file_gluon_customauth_pairing_proto_depIdxs,
		MessageInfos:      file_gluon_customauth_pairing_proto_msgTypes,
	}.Build()
	File_gluon_customauth_pairing_proto = out.File
	file_gluon_customauth_pairing_proto_rawDesc = nil
	file_gluon_customauth_pairing_proto_goTypes = nil
	file_gluon_customauth_pairing_proto_depIdxs = nil
}
