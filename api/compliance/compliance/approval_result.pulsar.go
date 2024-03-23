// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package compliance

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
	md_ApprovalResult           protoreflect.MessageDescriptor
	fd_ApprovalResult_id        protoreflect.FieldDescriptor
	fd_ApprovalResult_testerUrl protoreflect.FieldDescriptor
	fd_ApprovalResult_testerID  protoreflect.FieldDescriptor
	fd_ApprovalResult_approved  protoreflect.FieldDescriptor
	fd_ApprovalResult_creator   protoreflect.FieldDescriptor
)

func init() {
	file_compliance_compliance_approval_result_proto_init()
	md_ApprovalResult = File_compliance_compliance_approval_result_proto.Messages().ByName("ApprovalResult")
	fd_ApprovalResult_id = md_ApprovalResult.Fields().ByName("id")
	fd_ApprovalResult_testerUrl = md_ApprovalResult.Fields().ByName("testerUrl")
	fd_ApprovalResult_testerID = md_ApprovalResult.Fields().ByName("testerID")
	fd_ApprovalResult_approved = md_ApprovalResult.Fields().ByName("approved")
	fd_ApprovalResult_creator = md_ApprovalResult.Fields().ByName("creator")
}

var _ protoreflect.Message = (*fastReflection_ApprovalResult)(nil)

type fastReflection_ApprovalResult ApprovalResult

func (x *ApprovalResult) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ApprovalResult)(x)
}

func (x *ApprovalResult) slowProtoReflect() protoreflect.Message {
	mi := &file_compliance_compliance_approval_result_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ApprovalResult_messageType fastReflection_ApprovalResult_messageType
var _ protoreflect.MessageType = fastReflection_ApprovalResult_messageType{}

type fastReflection_ApprovalResult_messageType struct{}

func (x fastReflection_ApprovalResult_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ApprovalResult)(nil)
}
func (x fastReflection_ApprovalResult_messageType) New() protoreflect.Message {
	return new(fastReflection_ApprovalResult)
}
func (x fastReflection_ApprovalResult_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ApprovalResult
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ApprovalResult) Descriptor() protoreflect.MessageDescriptor {
	return md_ApprovalResult
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ApprovalResult) Type() protoreflect.MessageType {
	return _fastReflection_ApprovalResult_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ApprovalResult) New() protoreflect.Message {
	return new(fastReflection_ApprovalResult)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ApprovalResult) Interface() protoreflect.ProtoMessage {
	return (*ApprovalResult)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ApprovalResult) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_ApprovalResult_id, value) {
			return
		}
	}
	if x.TesterUrl != "" {
		value := protoreflect.ValueOfString(x.TesterUrl)
		if !f(fd_ApprovalResult_testerUrl, value) {
			return
		}
	}
	if x.TesterID != "" {
		value := protoreflect.ValueOfString(x.TesterID)
		if !f(fd_ApprovalResult_testerID, value) {
			return
		}
	}
	if x.Approved != false {
		value := protoreflect.ValueOfBool(x.Approved)
		if !f(fd_ApprovalResult_approved, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_ApprovalResult_creator, value) {
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
func (x *fastReflection_ApprovalResult) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "compliance.compliance.ApprovalResult.id":
		return x.Id != uint64(0)
	case "compliance.compliance.ApprovalResult.testerUrl":
		return x.TesterUrl != ""
	case "compliance.compliance.ApprovalResult.testerID":
		return x.TesterID != ""
	case "compliance.compliance.ApprovalResult.approved":
		return x.Approved != false
	case "compliance.compliance.ApprovalResult.creator":
		return x.Creator != ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.ApprovalResult"))
		}
		panic(fmt.Errorf("message compliance.compliance.ApprovalResult does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ApprovalResult) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "compliance.compliance.ApprovalResult.id":
		x.Id = uint64(0)
	case "compliance.compliance.ApprovalResult.testerUrl":
		x.TesterUrl = ""
	case "compliance.compliance.ApprovalResult.testerID":
		x.TesterID = ""
	case "compliance.compliance.ApprovalResult.approved":
		x.Approved = false
	case "compliance.compliance.ApprovalResult.creator":
		x.Creator = ""
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.ApprovalResult"))
		}
		panic(fmt.Errorf("message compliance.compliance.ApprovalResult does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ApprovalResult) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "compliance.compliance.ApprovalResult.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "compliance.compliance.ApprovalResult.testerUrl":
		value := x.TesterUrl
		return protoreflect.ValueOfString(value)
	case "compliance.compliance.ApprovalResult.testerID":
		value := x.TesterID
		return protoreflect.ValueOfString(value)
	case "compliance.compliance.ApprovalResult.approved":
		value := x.Approved
		return protoreflect.ValueOfBool(value)
	case "compliance.compliance.ApprovalResult.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.ApprovalResult"))
		}
		panic(fmt.Errorf("message compliance.compliance.ApprovalResult does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_ApprovalResult) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "compliance.compliance.ApprovalResult.id":
		x.Id = value.Uint()
	case "compliance.compliance.ApprovalResult.testerUrl":
		x.TesterUrl = value.Interface().(string)
	case "compliance.compliance.ApprovalResult.testerID":
		x.TesterID = value.Interface().(string)
	case "compliance.compliance.ApprovalResult.approved":
		x.Approved = value.Bool()
	case "compliance.compliance.ApprovalResult.creator":
		x.Creator = value.Interface().(string)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.ApprovalResult"))
		}
		panic(fmt.Errorf("message compliance.compliance.ApprovalResult does not contain field %s", fd.FullName()))
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
func (x *fastReflection_ApprovalResult) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "compliance.compliance.ApprovalResult.id":
		panic(fmt.Errorf("field id of message compliance.compliance.ApprovalResult is not mutable"))
	case "compliance.compliance.ApprovalResult.testerUrl":
		panic(fmt.Errorf("field testerUrl of message compliance.compliance.ApprovalResult is not mutable"))
	case "compliance.compliance.ApprovalResult.testerID":
		panic(fmt.Errorf("field testerID of message compliance.compliance.ApprovalResult is not mutable"))
	case "compliance.compliance.ApprovalResult.approved":
		panic(fmt.Errorf("field approved of message compliance.compliance.ApprovalResult is not mutable"))
	case "compliance.compliance.ApprovalResult.creator":
		panic(fmt.Errorf("field creator of message compliance.compliance.ApprovalResult is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.ApprovalResult"))
		}
		panic(fmt.Errorf("message compliance.compliance.ApprovalResult does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ApprovalResult) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "compliance.compliance.ApprovalResult.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "compliance.compliance.ApprovalResult.testerUrl":
		return protoreflect.ValueOfString("")
	case "compliance.compliance.ApprovalResult.testerID":
		return protoreflect.ValueOfString("")
	case "compliance.compliance.ApprovalResult.approved":
		return protoreflect.ValueOfBool(false)
	case "compliance.compliance.ApprovalResult.creator":
		return protoreflect.ValueOfString("")
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.ApprovalResult"))
		}
		panic(fmt.Errorf("message compliance.compliance.ApprovalResult does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ApprovalResult) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in compliance.compliance.ApprovalResult", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ApprovalResult) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ApprovalResult) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_ApprovalResult) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ApprovalResult) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ApprovalResult)
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
		l = len(x.TesterUrl)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.TesterID)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Approved {
			n += 2
		}
		l = len(x.Creator)
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
		x := input.Message.Interface().(*ApprovalResult)
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
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x2a
		}
		if x.Approved {
			i--
			if x.Approved {
				dAtA[i] = 1
			} else {
				dAtA[i] = 0
			}
			i--
			dAtA[i] = 0x20
		}
		if len(x.TesterID) > 0 {
			i -= len(x.TesterID)
			copy(dAtA[i:], x.TesterID)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TesterID)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.TesterUrl) > 0 {
			i -= len(x.TesterUrl)
			copy(dAtA[i:], x.TesterUrl)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.TesterUrl)))
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
		x := input.Message.Interface().(*ApprovalResult)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ApprovalResult: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ApprovalResult: illegal tag %d (wire type %d)", fieldNum, wire)
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
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TesterUrl", wireType)
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
				x.TesterUrl = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field TesterID", wireType)
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
				x.TesterID = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Approved", wireType)
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
				x.Approved = bool(v != 0)
			case 5:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
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
				x.Creator = string(dAtA[iNdEx:postIndex])
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
// source: compliance/compliance/approval_result.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApprovalResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TesterUrl string `protobuf:"bytes,2,opt,name=testerUrl,proto3" json:"testerUrl,omitempty"`
	TesterID  string `protobuf:"bytes,3,opt,name=testerID,proto3" json:"testerID,omitempty"`
	Approved  bool   `protobuf:"varint,4,opt,name=approved,proto3" json:"approved,omitempty"`
	Creator   string `protobuf:"bytes,5,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (x *ApprovalResult) Reset() {
	*x = ApprovalResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compliance_compliance_approval_result_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ApprovalResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApprovalResult) ProtoMessage() {}

// Deprecated: Use ApprovalResult.ProtoReflect.Descriptor instead.
func (*ApprovalResult) Descriptor() ([]byte, []int) {
	return file_compliance_compliance_approval_result_proto_rawDescGZIP(), []int{0}
}

func (x *ApprovalResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ApprovalResult) GetTesterUrl() string {
	if x != nil {
		return x.TesterUrl
	}
	return ""
}

func (x *ApprovalResult) GetTesterID() string {
	if x != nil {
		return x.TesterID
	}
	return ""
}

func (x *ApprovalResult) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

func (x *ApprovalResult) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

var File_compliance_compliance_approval_result_proto protoreflect.FileDescriptor

var file_compliance_compliance_approval_result_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x22, 0x90, 0x01, 0x0a, 0x0e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61,
	0x6c, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x65, 0x73, 0x74, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x74, 0x65, 0x73, 0x74,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x1a, 0x0a, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x65, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x42, 0xcd, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e,
	0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c,
	0x69, 0x61, 0x6e, 0x63, 0x65, 0x42, 0x13, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x26, 0x63, 0x6f,
	0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x15, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0xca, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5c,
	0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0xe2, 0x02, 0x21, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02,
	0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x3a, 0x3a, 0x43, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_compliance_compliance_approval_result_proto_rawDescOnce sync.Once
	file_compliance_compliance_approval_result_proto_rawDescData = file_compliance_compliance_approval_result_proto_rawDesc
)

func file_compliance_compliance_approval_result_proto_rawDescGZIP() []byte {
	file_compliance_compliance_approval_result_proto_rawDescOnce.Do(func() {
		file_compliance_compliance_approval_result_proto_rawDescData = protoimpl.X.CompressGZIP(file_compliance_compliance_approval_result_proto_rawDescData)
	})
	return file_compliance_compliance_approval_result_proto_rawDescData
}

var file_compliance_compliance_approval_result_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_compliance_compliance_approval_result_proto_goTypes = []interface{}{
	(*ApprovalResult)(nil), // 0: compliance.compliance.ApprovalResult
}
var file_compliance_compliance_approval_result_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_compliance_compliance_approval_result_proto_init() }
func file_compliance_compliance_approval_result_proto_init() {
	if File_compliance_compliance_approval_result_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_compliance_compliance_approval_result_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ApprovalResult); i {
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
			RawDescriptor: file_compliance_compliance_approval_result_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compliance_compliance_approval_result_proto_goTypes,
		DependencyIndexes: file_compliance_compliance_approval_result_proto_depIdxs,
		MessageInfos:      file_compliance_compliance_approval_result_proto_msgTypes,
	}.Build()
	File_compliance_compliance_approval_result_proto = out.File
	file_compliance_compliance_approval_result_proto_rawDesc = nil
	file_compliance_compliance_approval_result_proto_goTypes = nil
	file_compliance_compliance_approval_result_proto_depIdxs = nil
}