// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: actor.proto

package actor

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TerminatedReason int32

const (
	TerminatedReason_Stopped           TerminatedReason = 0
	TerminatedReason_AddressTerminated TerminatedReason = 1
	TerminatedReason_NotFound          TerminatedReason = 2
)

// Enum value maps for TerminatedReason.
var (
	TerminatedReason_name = map[int32]string{
		0: "Stopped",
		1: "AddressTerminated",
		2: "NotFound",
	}
	TerminatedReason_value = map[string]int32{
		"Stopped":           0,
		"AddressTerminated": 1,
		"NotFound":          2,
	}
)

func (x TerminatedReason) Enum() *TerminatedReason {
	p := new(TerminatedReason)
	*p = x
	return p
}

func (x TerminatedReason) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TerminatedReason) Descriptor() protoreflect.EnumDescriptor {
	return file_actor_proto_enumTypes[0].Descriptor()
}

func (TerminatedReason) Type() protoreflect.EnumType {
	return &file_actor_proto_enumTypes[0]
}

func (x TerminatedReason) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TerminatedReason.Descriptor instead.
func (TerminatedReason) EnumDescriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

type PID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Address   string `protobuf:"bytes,1,opt,name=Address,proto3" json:"Address,omitempty"`
	Id        string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	RequestId uint32 `protobuf:"varint,3,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`

	//manually added
	p *Process
}

func (x *PID) Reset() {
	*x = PID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PID) ProtoMessage() {}

func (x *PID) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PID.ProtoReflect.Descriptor instead.
func (*PID) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{0}
}

func (x *PID) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PID) GetRequestId() uint32 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

// user messages
type PoisonPill struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *PoisonPill) Reset() {
	*x = PoisonPill{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PoisonPill) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PoisonPill) ProtoMessage() {}

func (x *PoisonPill) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PoisonPill.ProtoReflect.Descriptor instead.
func (*PoisonPill) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{1}
}

type DeadLetterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Target *PID `protobuf:"bytes,1,opt,name=Target,proto3" json:"Target,omitempty"`
}

func (x *DeadLetterResponse) Reset() {
	*x = DeadLetterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeadLetterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeadLetterResponse) ProtoMessage() {}

func (x *DeadLetterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeadLetterResponse.ProtoReflect.Descriptor instead.
func (*DeadLetterResponse) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{2}
}

func (x *DeadLetterResponse) GetTarget() *PID {
	if x != nil {
		return x.Target
	}
	return nil
}

// system messages
type Watch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Watcher *PID `protobuf:"bytes,1,opt,name=Watcher,proto3" json:"Watcher,omitempty"`
}

func (x *Watch) Reset() {
	*x = Watch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Watch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Watch) ProtoMessage() {}

func (x *Watch) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Watch.ProtoReflect.Descriptor instead.
func (*Watch) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{3}
}

func (x *Watch) GetWatcher() *PID {
	if x != nil {
		return x.Watcher
	}
	return nil
}

type Unwatch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Watcher *PID `protobuf:"bytes,1,opt,name=Watcher,proto3" json:"Watcher,omitempty"`
}

func (x *Unwatch) Reset() {
	*x = Unwatch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Unwatch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Unwatch) ProtoMessage() {}

func (x *Unwatch) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Unwatch.ProtoReflect.Descriptor instead.
func (*Unwatch) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{4}
}

func (x *Unwatch) GetWatcher() *PID {
	if x != nil {
		return x.Watcher
	}
	return nil
}

type Terminated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Who *PID             `protobuf:"bytes,1,opt,name=who,proto3" json:"who,omitempty"`
	Why TerminatedReason `protobuf:"varint,2,opt,name=Why,proto3,enum=actor.TerminatedReason" json:"Why,omitempty"`
}

func (x *Terminated) Reset() {
	*x = Terminated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Terminated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Terminated) ProtoMessage() {}

func (x *Terminated) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Terminated.ProtoReflect.Descriptor instead.
func (*Terminated) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{5}
}

func (x *Terminated) GetWho() *PID {
	if x != nil {
		return x.Who
	}
	return nil
}

func (x *Terminated) GetWhy() TerminatedReason {
	if x != nil {
		return x.Why
	}
	return TerminatedReason_Stopped
}

type Stop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Stop) Reset() {
	*x = Stop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Stop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Stop) ProtoMessage() {}

func (x *Stop) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Stop.ProtoReflect.Descriptor instead.
func (*Stop) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{6}
}

type Touch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Touch) Reset() {
	*x = Touch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Touch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Touch) ProtoMessage() {}

func (x *Touch) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Touch.ProtoReflect.Descriptor instead.
func (*Touch) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{7}
}

type Touched struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Who *PID `protobuf:"bytes,1,opt,name=who,proto3" json:"who,omitempty"`
}

func (x *Touched) Reset() {
	*x = Touched{}
	if protoimpl.UnsafeEnabled {
		mi := &file_actor_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Touched) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Touched) ProtoMessage() {}

func (x *Touched) ProtoReflect() protoreflect.Message {
	mi := &file_actor_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Touched.ProtoReflect.Descriptor instead.
func (*Touched) Descriptor() ([]byte, []int) {
	return file_actor_proto_rawDescGZIP(), []int{8}
}

func (x *Touched) GetWho() *PID {
	if x != nil {
		return x.Who
	}
	return nil
}

var File_actor_proto protoreflect.FileDescriptor

var file_actor_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x22, 0x4e, 0x0a, 0x03, 0x50, 0x49, 0x44, 0x12, 0x18, 0x0a, 0x07, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x41, 0x64,
	0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x09, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x49, 0x64, 0x22, 0x0c, 0x0a, 0x0a, 0x50, 0x6f, 0x69, 0x73, 0x6f, 0x6e, 0x50, 0x69,
	0x6c, 0x6c, 0x22, 0x38, 0x0a, 0x12, 0x44, 0x65, 0x61, 0x64, 0x4c, 0x65, 0x74, 0x74, 0x65, 0x72,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x06, 0x54, 0x61, 0x72, 0x67,
	0x65, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x2e, 0x50, 0x49, 0x44, 0x52, 0x06, 0x54, 0x61, 0x72, 0x67, 0x65, 0x74, 0x22, 0x2d, 0x0a, 0x05,
	0x57, 0x61, 0x74, 0x63, 0x68, 0x12, 0x24, 0x0a, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50,
	0x49, 0x44, 0x52, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x2f, 0x0a, 0x07, 0x55,
	0x6e, 0x77, 0x61, 0x74, 0x63, 0x68, 0x12, 0x24, 0x0a, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x50, 0x49, 0x44, 0x52, 0x07, 0x57, 0x61, 0x74, 0x63, 0x68, 0x65, 0x72, 0x22, 0x55, 0x0a, 0x0a,
	0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x03, 0x77, 0x68,
	0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e,
	0x50, 0x49, 0x44, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x12, 0x29, 0x0a, 0x03, 0x57, 0x68, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2e, 0x54, 0x65,
	0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f, 0x6e, 0x52, 0x03,
	0x57, 0x68, 0x79, 0x22, 0x06, 0x0a, 0x04, 0x53, 0x74, 0x6f, 0x70, 0x22, 0x07, 0x0a, 0x05, 0x54,
	0x6f, 0x75, 0x63, 0x68, 0x22, 0x27, 0x0a, 0x07, 0x54, 0x6f, 0x75, 0x63, 0x68, 0x65, 0x64, 0x12,
	0x1c, 0x0a, 0x03, 0x77, 0x68, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x61,
	0x63, 0x74, 0x6f, 0x72, 0x2e, 0x50, 0x49, 0x44, 0x52, 0x03, 0x77, 0x68, 0x6f, 0x2a, 0x44, 0x0a,
	0x10, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x61, 0x73, 0x6f,
	0x6e, 0x12, 0x0b, 0x0a, 0x07, 0x53, 0x74, 0x6f, 0x70, 0x70, 0x65, 0x64, 0x10, 0x00, 0x12, 0x15,
	0x0a, 0x11, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x54, 0x65, 0x72, 0x6d, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x0c, 0x0a, 0x08, 0x4e, 0x6f, 0x74, 0x46, 0x6f, 0x75, 0x6e,
	0x64, 0x10, 0x02, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x41, 0x73, 0x79, 0x6e, 0x6b, 0x72, 0x6f, 0x6e, 0x49, 0x54, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x2d, 0x67, 0x6f, 0x2f, 0x61, 0x63, 0x74, 0x6f, 0x72,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_actor_proto_rawDescOnce sync.Once
	file_actor_proto_rawDescData = file_actor_proto_rawDesc
)

func file_actor_proto_rawDescGZIP() []byte {
	file_actor_proto_rawDescOnce.Do(func() {
		file_actor_proto_rawDescData = protoimpl.X.CompressGZIP(file_actor_proto_rawDescData)
	})
	return file_actor_proto_rawDescData
}

var file_actor_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_actor_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_actor_proto_goTypes = []interface{}{
	(TerminatedReason)(0),      // 0: actor.TerminatedReason
	(*PID)(nil),                // 1: actor.PID
	(*PoisonPill)(nil),         // 2: actor.PoisonPill
	(*DeadLetterResponse)(nil), // 3: actor.DeadLetterResponse
	(*Watch)(nil),              // 4: actor.Watch
	(*Unwatch)(nil),            // 5: actor.Unwatch
	(*Terminated)(nil),         // 6: actor.Terminated
	(*Stop)(nil),               // 7: actor.Stop
	(*Touch)(nil),              // 8: actor.Touch
	(*Touched)(nil),            // 9: actor.Touched
}
var file_actor_proto_depIdxs = []int32{
	1, // 0: actor.DeadLetterResponse.Target:type_name -> actor.PID
	1, // 1: actor.Watch.Watcher:type_name -> actor.PID
	1, // 2: actor.Unwatch.Watcher:type_name -> actor.PID
	1, // 3: actor.Terminated.who:type_name -> actor.PID
	0, // 4: actor.Terminated.Why:type_name -> actor.TerminatedReason
	1, // 5: actor.Touched.who:type_name -> actor.PID
	6, // [6:6] is the sub-list for method output_type
	6, // [6:6] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_actor_proto_init() }
func file_actor_proto_init() {
	if File_actor_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_actor_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PID); i {
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
		file_actor_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PoisonPill); i {
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
		file_actor_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeadLetterResponse); i {
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
		file_actor_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Watch); i {
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
		file_actor_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Unwatch); i {
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
		file_actor_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Terminated); i {
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
		file_actor_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Stop); i {
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
		file_actor_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Touch); i {
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
		file_actor_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Touched); i {
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
			RawDescriptor: file_actor_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_actor_proto_goTypes,
		DependencyIndexes: file_actor_proto_depIdxs,
		EnumInfos:         file_actor_proto_enumTypes,
		MessageInfos:      file_actor_proto_msgTypes,
	}.Build()
	File_actor_proto = out.File
	file_actor_proto_rawDesc = nil
	file_actor_proto_goTypes = nil
	file_actor_proto_depIdxs = nil
}
