// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package pairing

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	io "io"
	reflect "reflect"
	sync "sync"
)

var (
	md_PubKey               protoreflect.MessageDescriptor
	fd_PubKey_user          protoreflect.FieldDescriptor
	fd_PubKey_pairing_index protoreflect.FieldDescriptor
)

func init() {
	file_gluon_customauth_pairing_pubkey_proto_init()
	md_PubKey = File_gluon_customauth_pairing_pubkey_proto.Messages().ByName("PubKey")
	fd_PubKey_user = md_PubKey.Fields().ByName("user")
	fd_PubKey_pairing_index = md_PubKey.Fields().ByName("pairing_index")
}

var _ protoreflect.Message = (*fastReflection_PubKey)(nil)

type fastReflection_PubKey PubKey

func (x *PubKey) ProtoReflect() protoreflect.Message {
	return (*fastReflection_PubKey)(x)
}

func (x *PubKey) slowProtoReflect() protoreflect.Message {
	mi := &file_gluon_customauth_pairing_pubkey_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_PubKey_messageType fastReflection_PubKey_messageType
var _ protoreflect.MessageType = fastReflection_PubKey_messageType{}

type fastReflection_PubKey_messageType struct{}

func (x fastReflection_PubKey_messageType) Zero() protoreflect.Message {
	return (*fastReflection_PubKey)(nil)
}
func (x fastReflection_PubKey_messageType) New() protoreflect.Message {
	return new(fastReflection_PubKey)
}
func (x fastReflection_PubKey_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_PubKey
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_PubKey) Descriptor() protoreflect.MessageDescriptor {
	return md_PubKey
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_PubKey) Type() protoreflect.MessageType {
	return _fastReflection_PubKey_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_PubKey) New() protoreflect.Message {
	return new(fastReflection_PubKey)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_PubKey) Interface() protoreflect.ProtoMessage {
	return (*PubKey)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_PubKey) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.User) != 0 {
		value := protoreflect.ValueOfBytes(x.User)
		if !f(fd_PubKey_user, value) {
			return
		}
	}
	if x.PairingIndex != "" {
		value := protoreflect.ValueOfString(x.PairingIndex)
		if !f(fd_PubKey_pairing_index, value) {
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
func (x *fastReflection_PubKey) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "gluon.customauth.pairing.PubKey.user":
		return len(x.User) != 0
	case "gluon.customauth.pairing.PubKey.pairing_index":
		return x.PairingIndex != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.pairing.PubKey"))
		}
		panic(fmt.Errorf("message gluon.customauth.pairing.PubKey does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PubKey) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "gluon.customauth.pairing.PubKey.user":
		x.User = nil
	case "gluon.customauth.pairing.PubKey.pairing_index":
		x.PairingIndex = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.pairing.PubKey"))
		}
		panic(fmt.Errorf("message gluon.customauth.pairing.PubKey does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_PubKey) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "gluon.customauth.pairing.PubKey.user":
		value := x.User
		return protoreflect.ValueOfBytes(value)
	case "gluon.customauth.pairing.PubKey.pairing_index":
		value := x.PairingIndex
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.pairing.PubKey"))
		}
		panic(fmt.Errorf("message gluon.customauth.pairing.PubKey does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_PubKey) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "gluon.customauth.pairing.PubKey.user":
		x.User = value.Bytes()
	case "gluon.customauth.pairing.PubKey.pairing_index":
		x.PairingIndex = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.pairing.PubKey"))
		}
		panic(fmt.Errorf("message gluon.customauth.pairing.PubKey does not contain field %s", fd.FullName()))
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
func (x *fastReflection_PubKey) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.customauth.pairing.PubKey.user":
		panic(fmt.Errorf("field user of message gluon.customauth.pairing.PubKey is not mutable"))
	case "gluon.customauth.pairing.PubKey.pairing_index":
		panic(fmt.Errorf("field pairing_index of message gluon.customauth.pairing.PubKey is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.pairing.PubKey"))
		}
		panic(fmt.Errorf("message gluon.customauth.pairing.PubKey does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_PubKey) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "gluon.customauth.pairing.PubKey.user":
		return protoreflect.ValueOfBytes(nil)
	case "gluon.customauth.pairing.PubKey.pairing_index":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: gluon.customauth.pairing.PubKey"))
		}
		panic(fmt.Errorf("message gluon.customauth.pairing.PubKey does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_PubKey) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in gluon.customauth.pairing.PubKey", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_PubKey) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_PubKey) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_PubKey) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_PubKey) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*PubKey)
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
		l = len(x.User)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.PairingIndex)
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
		x := input.Message.Interface().(*PubKey)
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
		if len(x.PairingIndex) > 0 {
			i -= len(x.PairingIndex)
			copy(dAtA[i:], x.PairingIndex)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.PairingIndex)))
			i--
			dAtA[i] = 0x12
		}
		if len(x.User) > 0 {
			i -= len(x.User)
			copy(dAtA[i:], x.User)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.User)))
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
		x := input.Message.Interface().(*PubKey)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PubKey: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: PubKey: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field User", wireType)
				}
				var byteLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					byteLen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if byteLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + byteLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.User = append(x.User[:0], dAtA[iNdEx:postIndex]...)
				if x.User == nil {
					x.User = []byte{}
				}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field PairingIndex", wireType)
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
				x.PairingIndex = string(dAtA[iNdEx:postIndex])
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
// source: gluon/customauth/pairing/pubkey.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Pairing PubKey
// This is only for adapting PubKey interface of Cosmos SDK
type PubKey struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Address binary
	User         []byte `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	PairingIndex string `protobuf:"bytes,2,opt,name=pairing_index,json=pairingIndex,proto3" json:"pairing_index,omitempty"`
}

func (x *PubKey) Reset() {
	*x = PubKey{}
	if protoimpl.UnsafeEnabled {
		mi := &file_gluon_customauth_pairing_pubkey_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PubKey) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PubKey) ProtoMessage() {}

// Deprecated: Use PubKey.ProtoReflect.Descriptor instead.
func (*PubKey) Descriptor() ([]byte, []int) {
	return file_gluon_customauth_pairing_pubkey_proto_rawDescGZIP(), []int{0}
}

func (x *PubKey) GetUser() []byte {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PubKey) GetPairingIndex() string {
	if x != nil {
		return x.PairingIndex
	}
	return ""
}

var File_gluon_customauth_pairing_pubkey_proto protoreflect.FileDescriptor

var file_gluon_customauth_pairing_pubkey_proto_rawDesc = []byte{
	0x0a, 0x25, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75,
	0x74, 0x68, 0x2f, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x2f, 0x70, 0x75, 0x62, 0x6b, 0x65,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x18, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x63,
	0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e,
	0x67, 0x22, 0x41, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x49,
	0x6e, 0x64, 0x65, 0x78, 0x42, 0xd1, 0x01, 0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6c, 0x75,
	0x6f, 0x6e, 0x2e, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x61,
	0x69, 0x72, 0x69, 0x6e, 0x67, 0x42, 0x0b, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x22, 0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x67, 0x6c, 0x75, 0x6f, 0x6e, 0x2f, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68,
	0x2f, 0x70, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0xa2, 0x02, 0x03, 0x47, 0x43, 0x50, 0xaa, 0x02,
	0x18, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x2e, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74,
	0x68, 0x2e, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0xca, 0x02, 0x18, 0x47, 0x6c, 0x75, 0x6f,
	0x6e, 0x5c, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x50, 0x61, 0x69,
	0x72, 0x69, 0x6e, 0x67, 0xe2, 0x02, 0x24, 0x47, 0x6c, 0x75, 0x6f, 0x6e, 0x5c, 0x43, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x5c, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x1a, 0x47, 0x6c,
	0x75, 0x6f, 0x6e, 0x3a, 0x3a, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x61, 0x75, 0x74, 0x68, 0x3a,
	0x3a, 0x50, 0x61, 0x69, 0x72, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gluon_customauth_pairing_pubkey_proto_rawDescOnce sync.Once
	file_gluon_customauth_pairing_pubkey_proto_rawDescData = file_gluon_customauth_pairing_pubkey_proto_rawDesc
)

func file_gluon_customauth_pairing_pubkey_proto_rawDescGZIP() []byte {
	file_gluon_customauth_pairing_pubkey_proto_rawDescOnce.Do(func() {
		file_gluon_customauth_pairing_pubkey_proto_rawDescData = protoimpl.X.CompressGZIP(file_gluon_customauth_pairing_pubkey_proto_rawDescData)
	})
	return file_gluon_customauth_pairing_pubkey_proto_rawDescData
}

var file_gluon_customauth_pairing_pubkey_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_gluon_customauth_pairing_pubkey_proto_goTypes = []interface{}{
	(*PubKey)(nil), // 0: gluon.customauth.pairing.PubKey
}
var file_gluon_customauth_pairing_pubkey_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_gluon_customauth_pairing_pubkey_proto_init() }
func file_gluon_customauth_pairing_pubkey_proto_init() {
	if File_gluon_customauth_pairing_pubkey_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_gluon_customauth_pairing_pubkey_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PubKey); i {
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
			RawDescriptor: file_gluon_customauth_pairing_pubkey_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_gluon_customauth_pairing_pubkey_proto_goTypes,
		DependencyIndexes: file_gluon_customauth_pairing_pubkey_proto_depIdxs,
		MessageInfos:      file_gluon_customauth_pairing_pubkey_proto_msgTypes,
	}.Build()
	File_gluon_customauth_pairing_pubkey_proto = out.File
	file_gluon_customauth_pairing_pubkey_proto_rawDesc = nil
	file_gluon_customauth_pairing_pubkey_proto_goTypes = nil
	file_gluon_customauth_pairing_pubkey_proto_depIdxs = nil
}
