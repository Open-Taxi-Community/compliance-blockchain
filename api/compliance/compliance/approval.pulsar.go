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

var _ protoreflect.List = (*_Approval_7_list)(nil)

type _Approval_7_list struct {
	list *[]*ApprovalResult
}

func (x *_Approval_7_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Approval_7_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Approval_7_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ApprovalResult)
	(*x.list)[i] = concreteValue
}

func (x *_Approval_7_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*ApprovalResult)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Approval_7_list) AppendMutable() protoreflect.Value {
	v := new(ApprovalResult)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Approval_7_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Approval_7_list) NewElement() protoreflect.Value {
	v := new(ApprovalResult)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Approval_7_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Approval           protoreflect.MessageDescriptor
	fd_Approval_id        protoreflect.FieldDescriptor
	fd_Approval_serverUrl protoreflect.FieldDescriptor
	fd_Approval_driverID  protoreflect.FieldDescriptor
	fd_Approval_status    protoreflect.FieldDescriptor
	fd_Approval_creator   protoreflect.FieldDescriptor
	fd_Approval_results   protoreflect.FieldDescriptor
)

func init() {
	file_compliance_compliance_approval_proto_init()
	md_Approval = File_compliance_compliance_approval_proto.Messages().ByName("Approval")
	fd_Approval_id = md_Approval.Fields().ByName("id")
	fd_Approval_serverUrl = md_Approval.Fields().ByName("serverUrl")
	fd_Approval_driverID = md_Approval.Fields().ByName("driverID")
	fd_Approval_status = md_Approval.Fields().ByName("status")
	fd_Approval_creator = md_Approval.Fields().ByName("creator")
	fd_Approval_results = md_Approval.Fields().ByName("results")
}

var _ protoreflect.Message = (*fastReflection_Approval)(nil)

type fastReflection_Approval Approval

func (x *Approval) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Approval)(x)
}

func (x *Approval) slowProtoReflect() protoreflect.Message {
	mi := &file_compliance_compliance_approval_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Approval_messageType fastReflection_Approval_messageType
var _ protoreflect.MessageType = fastReflection_Approval_messageType{}

type fastReflection_Approval_messageType struct{}

func (x fastReflection_Approval_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Approval)(nil)
}
func (x fastReflection_Approval_messageType) New() protoreflect.Message {
	return new(fastReflection_Approval)
}
func (x fastReflection_Approval_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Approval
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Approval) Descriptor() protoreflect.MessageDescriptor {
	return md_Approval
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Approval) Type() protoreflect.MessageType {
	return _fastReflection_Approval_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Approval) New() protoreflect.Message {
	return new(fastReflection_Approval)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Approval) Interface() protoreflect.ProtoMessage {
	return (*Approval)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Approval) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Id != uint64(0) {
		value := protoreflect.ValueOfUint64(x.Id)
		if !f(fd_Approval_id, value) {
			return
		}
	}
	if x.ServerUrl != "" {
		value := protoreflect.ValueOfString(x.ServerUrl)
		if !f(fd_Approval_serverUrl, value) {
			return
		}
	}
	if x.DriverID != "" {
		value := protoreflect.ValueOfString(x.DriverID)
		if !f(fd_Approval_driverID, value) {
			return
		}
	}
	if x.Status != 0 {
		value := protoreflect.ValueOfEnum((protoreflect.EnumNumber)(x.Status))
		if !f(fd_Approval_status, value) {
			return
		}
	}
	if x.Creator != "" {
		value := protoreflect.ValueOfString(x.Creator)
		if !f(fd_Approval_creator, value) {
			return
		}
	}
	if len(x.Results) != 0 {
		value := protoreflect.ValueOfList(&_Approval_7_list{list: &x.Results})
		if !f(fd_Approval_results, value) {
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
func (x *fastReflection_Approval) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "compliance.compliance.Approval.id":
		return x.Id != uint64(0)
	case "compliance.compliance.Approval.serverUrl":
		return x.ServerUrl != ""
	case "compliance.compliance.Approval.driverID":
		return x.DriverID != ""
	case "compliance.compliance.Approval.status":
		return x.Status != 0
	case "compliance.compliance.Approval.creator":
		return x.Creator != ""
	case "compliance.compliance.Approval.results":
		return len(x.Results) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.Approval"))
		}
		panic(fmt.Errorf("message compliance.compliance.Approval does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Approval) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "compliance.compliance.Approval.id":
		x.Id = uint64(0)
	case "compliance.compliance.Approval.serverUrl":
		x.ServerUrl = ""
	case "compliance.compliance.Approval.driverID":
		x.DriverID = ""
	case "compliance.compliance.Approval.status":
		x.Status = 0
	case "compliance.compliance.Approval.creator":
		x.Creator = ""
	case "compliance.compliance.Approval.results":
		x.Results = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.Approval"))
		}
		panic(fmt.Errorf("message compliance.compliance.Approval does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Approval) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "compliance.compliance.Approval.id":
		value := x.Id
		return protoreflect.ValueOfUint64(value)
	case "compliance.compliance.Approval.serverUrl":
		value := x.ServerUrl
		return protoreflect.ValueOfString(value)
	case "compliance.compliance.Approval.driverID":
		value := x.DriverID
		return protoreflect.ValueOfString(value)
	case "compliance.compliance.Approval.status":
		value := x.Status
		return protoreflect.ValueOfEnum((protoreflect.EnumNumber)(value))
	case "compliance.compliance.Approval.creator":
		value := x.Creator
		return protoreflect.ValueOfString(value)
	case "compliance.compliance.Approval.results":
		if len(x.Results) == 0 {
			return protoreflect.ValueOfList(&_Approval_7_list{})
		}
		listValue := &_Approval_7_list{list: &x.Results}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.Approval"))
		}
		panic(fmt.Errorf("message compliance.compliance.Approval does not contain field %s", descriptor.FullName()))
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
func (x *fastReflection_Approval) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "compliance.compliance.Approval.id":
		x.Id = value.Uint()
	case "compliance.compliance.Approval.serverUrl":
		x.ServerUrl = value.Interface().(string)
	case "compliance.compliance.Approval.driverID":
		x.DriverID = value.Interface().(string)
	case "compliance.compliance.Approval.status":
		x.Status = (ApprovalStatus)(value.Enum())
	case "compliance.compliance.Approval.creator":
		x.Creator = value.Interface().(string)
	case "compliance.compliance.Approval.results":
		lv := value.List()
		clv := lv.(*_Approval_7_list)
		x.Results = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.Approval"))
		}
		panic(fmt.Errorf("message compliance.compliance.Approval does not contain field %s", fd.FullName()))
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
func (x *fastReflection_Approval) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "compliance.compliance.Approval.results":
		if x.Results == nil {
			x.Results = []*ApprovalResult{}
		}
		value := &_Approval_7_list{list: &x.Results}
		return protoreflect.ValueOfList(value)
	case "compliance.compliance.Approval.id":
		panic(fmt.Errorf("field id of message compliance.compliance.Approval is not mutable"))
	case "compliance.compliance.Approval.serverUrl":
		panic(fmt.Errorf("field serverUrl of message compliance.compliance.Approval is not mutable"))
	case "compliance.compliance.Approval.driverID":
		panic(fmt.Errorf("field driverID of message compliance.compliance.Approval is not mutable"))
	case "compliance.compliance.Approval.status":
		panic(fmt.Errorf("field status of message compliance.compliance.Approval is not mutable"))
	case "compliance.compliance.Approval.creator":
		panic(fmt.Errorf("field creator of message compliance.compliance.Approval is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.Approval"))
		}
		panic(fmt.Errorf("message compliance.compliance.Approval does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Approval) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "compliance.compliance.Approval.id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "compliance.compliance.Approval.serverUrl":
		return protoreflect.ValueOfString("")
	case "compliance.compliance.Approval.driverID":
		return protoreflect.ValueOfString("")
	case "compliance.compliance.Approval.status":
		return protoreflect.ValueOfEnum(0)
	case "compliance.compliance.Approval.creator":
		return protoreflect.ValueOfString("")
	case "compliance.compliance.Approval.results":
		list := []*ApprovalResult{}
		return protoreflect.ValueOfList(&_Approval_7_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: compliance.compliance.Approval"))
		}
		panic(fmt.Errorf("message compliance.compliance.Approval does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Approval) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in compliance.compliance.Approval", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Approval) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Approval) SetUnknown(fields protoreflect.RawFields) {
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
func (x *fastReflection_Approval) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Approval) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Approval)
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
		l = len(x.ServerUrl)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		l = len(x.DriverID)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.Status != 0 {
			n += 1 + runtime.Sov(uint64(x.Status))
		}
		l = len(x.Creator)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if len(x.Results) > 0 {
			for _, e := range x.Results {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
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
		x := input.Message.Interface().(*Approval)
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
		if len(x.Results) > 0 {
			for iNdEx := len(x.Results) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Results[iNdEx])
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
				dAtA[i] = 0x3a
			}
		}
		if len(x.Creator) > 0 {
			i -= len(x.Creator)
			copy(dAtA[i:], x.Creator)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.Creator)))
			i--
			dAtA[i] = 0x32
		}
		if x.Status != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.Status))
			i--
			dAtA[i] = 0x28
		}
		if len(x.DriverID) > 0 {
			i -= len(x.DriverID)
			copy(dAtA[i:], x.DriverID)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.DriverID)))
			i--
			dAtA[i] = 0x22
		}
		if len(x.ServerUrl) > 0 {
			i -= len(x.ServerUrl)
			copy(dAtA[i:], x.ServerUrl)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ServerUrl)))
			i--
			dAtA[i] = 0x1a
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
		x := input.Message.Interface().(*Approval)
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Approval: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Approval: illegal tag %d (wire type %d)", fieldNum, wire)
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
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ServerUrl", wireType)
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
				x.ServerUrl = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field DriverID", wireType)
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
				x.DriverID = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 5:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
				}
				x.Status = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.Status |= ApprovalStatus(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 6:
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
			case 7:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Results", wireType)
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
				x.Results = append(x.Results, &ApprovalResult{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Results[len(x.Results)-1]); err != nil {
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
// source: compliance/compliance/approval.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ApprovalStatus int32

const (
	ApprovalStatus_UNSET    ApprovalStatus = 0
	ApprovalStatus_PENDING  ApprovalStatus = 1
	ApprovalStatus_APPROVED ApprovalStatus = 2
	ApprovalStatus_REJECTED ApprovalStatus = 3
)

// Enum value maps for ApprovalStatus.
var (
	ApprovalStatus_name = map[int32]string{
		0: "UNSET",
		1: "PENDING",
		2: "APPROVED",
		3: "REJECTED",
	}
	ApprovalStatus_value = map[string]int32{
		"UNSET":    0,
		"PENDING":  1,
		"APPROVED": 2,
		"REJECTED": 3,
	}
)

func (x ApprovalStatus) Enum() *ApprovalStatus {
	p := new(ApprovalStatus)
	*p = x
	return p
}

func (x ApprovalStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ApprovalStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_compliance_compliance_approval_proto_enumTypes[0].Descriptor()
}

func (ApprovalStatus) Type() protoreflect.EnumType {
	return &file_compliance_compliance_approval_proto_enumTypes[0]
}

func (x ApprovalStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ApprovalStatus.Descriptor instead.
func (ApprovalStatus) EnumDescriptor() ([]byte, []int) {
	return file_compliance_compliance_approval_proto_rawDescGZIP(), []int{0}
}

type Approval struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64            `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	ServerUrl string            `protobuf:"bytes,3,opt,name=serverUrl,proto3" json:"serverUrl,omitempty"`
	DriverID  string            `protobuf:"bytes,4,opt,name=driverID,proto3" json:"driverID,omitempty"`
	Status    ApprovalStatus    `protobuf:"varint,5,opt,name=status,proto3,enum=compliance.compliance.ApprovalStatus" json:"status,omitempty"`
	Creator   string            `protobuf:"bytes,6,opt,name=creator,proto3" json:"creator,omitempty"`
	Results   []*ApprovalResult `protobuf:"bytes,7,rep,name=results,proto3" json:"results,omitempty"`
}

func (x *Approval) Reset() {
	*x = Approval{}
	if protoimpl.UnsafeEnabled {
		mi := &file_compliance_compliance_approval_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Approval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Approval) ProtoMessage() {}

// Deprecated: Use Approval.ProtoReflect.Descriptor instead.
func (*Approval) Descriptor() ([]byte, []int) {
	return file_compliance_compliance_approval_proto_rawDescGZIP(), []int{0}
}

func (x *Approval) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Approval) GetServerUrl() string {
	if x != nil {
		return x.ServerUrl
	}
	return ""
}

func (x *Approval) GetDriverID() string {
	if x != nil {
		return x.DriverID
	}
	return ""
}

func (x *Approval) GetStatus() ApprovalStatus {
	if x != nil {
		return x.Status
	}
	return ApprovalStatus_UNSET
}

func (x *Approval) GetCreator() string {
	if x != nil {
		return x.Creator
	}
	return ""
}

func (x *Approval) GetResults() []*ApprovalResult {
	if x != nil {
		return x.Results
	}
	return nil
}

var File_compliance_compliance_approval_proto protoreflect.FileDescriptor

var file_compliance_compliance_approval_proto_rawDesc = []byte{
	0x0a, 0x24, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d,
	0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e,
	0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x1a, 0x2b, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x2f, 0x61, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x5f, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xee, 0x01, 0x0a, 0x08, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x65,
	0x72, 0x55, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x49,
	0x44, 0x12, 0x3d, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x25, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63,
	0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x72, 0x65, 0x61, 0x74, 0x6f, 0x72, 0x12, 0x3f, 0x0a, 0x07, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x52, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x52, 0x07, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x73, 0x2a, 0x44, 0x0a, 0x0e, 0x41,
	0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x09, 0x0a,
	0x05, 0x55, 0x4e, 0x53, 0x45, 0x54, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07, 0x50, 0x45, 0x4e, 0x44,
	0x49, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x41, 0x50, 0x50, 0x52, 0x4f, 0x56, 0x45,
	0x44, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x52, 0x45, 0x4a, 0x45, 0x43, 0x54, 0x45, 0x44, 0x10,
	0x03, 0x42, 0xc7, 0x01, 0x0a, 0x19, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69,
	0x61, 0x6e, 0x63, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x42,
	0x0d, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01,
	0x5a, 0x26, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x73, 0x64, 0x6b, 0x2e, 0x69, 0x6f, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0xa2, 0x02, 0x03, 0x43, 0x43, 0x58, 0xaa, 0x02,
	0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0xca, 0x02, 0x15, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61,
	0x6e, 0x63, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0xe2, 0x02,
	0x21, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x43, 0x6f, 0x6d, 0x70,
	0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0xea, 0x02, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x3a,
	0x3a, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x69, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_compliance_compliance_approval_proto_rawDescOnce sync.Once
	file_compliance_compliance_approval_proto_rawDescData = file_compliance_compliance_approval_proto_rawDesc
)

func file_compliance_compliance_approval_proto_rawDescGZIP() []byte {
	file_compliance_compliance_approval_proto_rawDescOnce.Do(func() {
		file_compliance_compliance_approval_proto_rawDescData = protoimpl.X.CompressGZIP(file_compliance_compliance_approval_proto_rawDescData)
	})
	return file_compliance_compliance_approval_proto_rawDescData
}

var file_compliance_compliance_approval_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_compliance_compliance_approval_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_compliance_compliance_approval_proto_goTypes = []interface{}{
	(ApprovalStatus)(0),    // 0: compliance.compliance.ApprovalStatus
	(*Approval)(nil),       // 1: compliance.compliance.Approval
	(*ApprovalResult)(nil), // 2: compliance.compliance.ApprovalResult
}
var file_compliance_compliance_approval_proto_depIdxs = []int32{
	0, // 0: compliance.compliance.Approval.status:type_name -> compliance.compliance.ApprovalStatus
	2, // 1: compliance.compliance.Approval.results:type_name -> compliance.compliance.ApprovalResult
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_compliance_compliance_approval_proto_init() }
func file_compliance_compliance_approval_proto_init() {
	if File_compliance_compliance_approval_proto != nil {
		return
	}
	file_compliance_compliance_approval_result_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_compliance_compliance_approval_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Approval); i {
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
			RawDescriptor: file_compliance_compliance_approval_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_compliance_compliance_approval_proto_goTypes,
		DependencyIndexes: file_compliance_compliance_approval_proto_depIdxs,
		EnumInfos:         file_compliance_compliance_approval_proto_enumTypes,
		MessageInfos:      file_compliance_compliance_approval_proto_msgTypes,
	}.Build()
	File_compliance_compliance_approval_proto = out.File
	file_compliance_compliance_approval_proto_rawDesc = nil
	file_compliance_compliance_approval_proto_goTypes = nil
	file_compliance_compliance_approval_proto_depIdxs = nil
}