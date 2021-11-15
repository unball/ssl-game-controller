// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0-devel
// 	protoc        v3.14.0
// source: ssl_gc_rcon_team.proto

package rcon

import (
	state "github.com/RoboCup-SSL/ssl-game-controller/internal/app/state"
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

type AdvantageResponse int32

const (
	// stop the game
	AdvantageResponse_STOP AdvantageResponse = 0
	// continue the game a bit more
	AdvantageResponse_CONTINUE AdvantageResponse = 1
	// no choice -> will default to STOP
	AdvantageResponse_UNDECIDED AdvantageResponse = 0
)

// Enum value maps for AdvantageResponse.
var (
	AdvantageResponse_name = map[int32]string{
		0: "STOP",
		1: "CONTINUE",
		// Duplicate value: 0: "UNDECIDED",
	}
	AdvantageResponse_value = map[string]int32{
		"STOP":      0,
		"CONTINUE":  1,
		"UNDECIDED": 0,
	}
)

func (x AdvantageResponse) Enum() *AdvantageResponse {
	p := new(AdvantageResponse)
	*p = x
	return p
}

func (x AdvantageResponse) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdvantageResponse) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_rcon_team_proto_enumTypes[0].Descriptor()
}

func (AdvantageResponse) Type() protoreflect.EnumType {
	return &file_ssl_gc_rcon_team_proto_enumTypes[0]
}

func (x AdvantageResponse) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AdvantageResponse) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AdvantageResponse(num)
	return nil
}

// Deprecated: Use AdvantageResponse.Descriptor instead.
func (AdvantageResponse) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{0}
}

type AdvantageChoice_Foul int32

const (
	// default value when not set
	AdvantageChoice_UNKNOWN AdvantageChoice_Foul = 0
)

// Enum value maps for AdvantageChoice_Foul.
var (
	AdvantageChoice_Foul_name = map[int32]string{
		0: "UNKNOWN",
	}
	AdvantageChoice_Foul_value = map[string]int32{
		"UNKNOWN": 0,
	}
)

func (x AdvantageChoice_Foul) Enum() *AdvantageChoice_Foul {
	p := new(AdvantageChoice_Foul)
	*p = x
	return p
}

func (x AdvantageChoice_Foul) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (AdvantageChoice_Foul) Descriptor() protoreflect.EnumDescriptor {
	return file_ssl_gc_rcon_team_proto_enumTypes[1].Descriptor()
}

func (AdvantageChoice_Foul) Type() protoreflect.EnumType {
	return &file_ssl_gc_rcon_team_proto_enumTypes[1]
}

func (x AdvantageChoice_Foul) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Do not use.
func (x *AdvantageChoice_Foul) UnmarshalJSON(b []byte) error {
	num, err := protoimpl.X.UnmarshalJSONEnum(x.Descriptor(), b)
	if err != nil {
		return err
	}
	*x = AdvantageChoice_Foul(num)
	return nil
}

// Deprecated: Use AdvantageChoice_Foul.Descriptor instead.
func (AdvantageChoice_Foul) EnumDescriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{3, 0}
}

// a registration that must be send by teams to the controller as the very first message
type TeamRegistration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the exact team name as published by the game-controller
	TeamName *string `protobuf:"bytes,1,req,name=team_name,json=teamName" json:"team_name,omitempty"`
	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	// the team (relevant only if a team plays against itself)
	Team *state.Team `protobuf:"varint,3,opt,name=team,enum=Team" json:"team,omitempty"`
}

func (x *TeamRegistration) Reset() {
	*x = TeamRegistration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamRegistration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamRegistration) ProtoMessage() {}

func (x *TeamRegistration) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamRegistration.ProtoReflect.Descriptor instead.
func (*TeamRegistration) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{0}
}

func (x *TeamRegistration) GetTeamName() string {
	if x != nil && x.TeamName != nil {
		return *x.TeamName
	}
	return ""
}

func (x *TeamRegistration) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (x *TeamRegistration) GetTeam() state.Team {
	if x != nil && x.Team != nil {
		return *x.Team
	}
	return state.Team_UNKNOWN
}

// wrapper for all messages from a team's computer to the controller
type TeamToController struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// signature can optionally be specified to enable secure communication
	Signature *Signature `protobuf:"bytes,1,opt,name=signature" json:"signature,omitempty"`
	// Types that are assignable to Msg:
	//	*TeamToController_DesiredKeeper
	//	*TeamToController_AdvantageResponse
	//	*TeamToController_SubstituteBot
	//	*TeamToController_Ping
	Msg isTeamToController_Msg `protobuf_oneof:"msg"`
}

func (x *TeamToController) Reset() {
	*x = TeamToController{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TeamToController) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TeamToController) ProtoMessage() {}

func (x *TeamToController) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TeamToController.ProtoReflect.Descriptor instead.
func (*TeamToController) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{1}
}

func (x *TeamToController) GetSignature() *Signature {
	if x != nil {
		return x.Signature
	}
	return nil
}

func (m *TeamToController) GetMsg() isTeamToController_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *TeamToController) GetDesiredKeeper() int32 {
	if x, ok := x.GetMsg().(*TeamToController_DesiredKeeper); ok {
		return x.DesiredKeeper
	}
	return 0
}

func (x *TeamToController) GetAdvantageResponse() AdvantageResponse {
	if x, ok := x.GetMsg().(*TeamToController_AdvantageResponse); ok {
		return x.AdvantageResponse
	}
	return AdvantageResponse_STOP
}

func (x *TeamToController) GetSubstituteBot() bool {
	if x, ok := x.GetMsg().(*TeamToController_SubstituteBot); ok {
		return x.SubstituteBot
	}
	return false
}

func (x *TeamToController) GetPing() bool {
	if x, ok := x.GetMsg().(*TeamToController_Ping); ok {
		return x.Ping
	}
	return false
}

type isTeamToController_Msg interface {
	isTeamToController_Msg()
}

type TeamToController_DesiredKeeper struct {
	// request a new desired keeper id
	DesiredKeeper int32 `protobuf:"varint,2,opt,name=desired_keeper,json=desiredKeeper,oneof"`
}

type TeamToController_AdvantageResponse struct {
	// response to an advantage choice request
	AdvantageResponse AdvantageResponse `protobuf:"varint,3,opt,name=advantage_response,json=advantageResponse,enum=AdvantageResponse,oneof"`
}

type TeamToController_SubstituteBot struct {
	// request to substitute a robot at the next possibility
	SubstituteBot bool `protobuf:"varint,4,opt,name=substitute_bot,json=substituteBot,oneof"`
}

type TeamToController_Ping struct {
	// send a ping to the GC to test if the connection is still open.
	// the value is ignored and a reply is sent back
	Ping bool `protobuf:"varint,5,opt,name=ping,oneof"`
}

func (*TeamToController_DesiredKeeper) isTeamToController_Msg() {}

func (*TeamToController_AdvantageResponse) isTeamToController_Msg() {}

func (*TeamToController_SubstituteBot) isTeamToController_Msg() {}

func (*TeamToController_Ping) isTeamToController_Msg() {}

// wrapper for all messages from controller to a team's computer
type ControllerToTeam struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Msg:
	//	*ControllerToTeam_ControllerReply
	//	*ControllerToTeam_AdvantageChoice
	Msg isControllerToTeam_Msg `protobuf_oneof:"msg"`
}

func (x *ControllerToTeam) Reset() {
	*x = ControllerToTeam{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControllerToTeam) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControllerToTeam) ProtoMessage() {}

func (x *ControllerToTeam) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControllerToTeam.ProtoReflect.Descriptor instead.
func (*ControllerToTeam) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{2}
}

func (m *ControllerToTeam) GetMsg() isControllerToTeam_Msg {
	if m != nil {
		return m.Msg
	}
	return nil
}

func (x *ControllerToTeam) GetControllerReply() *ControllerReply {
	if x, ok := x.GetMsg().(*ControllerToTeam_ControllerReply); ok {
		return x.ControllerReply
	}
	return nil
}

func (x *ControllerToTeam) GetAdvantageChoice() *AdvantageChoice {
	if x, ok := x.GetMsg().(*ControllerToTeam_AdvantageChoice); ok {
		return x.AdvantageChoice
	}
	return nil
}

type isControllerToTeam_Msg interface {
	isControllerToTeam_Msg()
}

type ControllerToTeam_ControllerReply struct {
	// a reply from the controller
	ControllerReply *ControllerReply `protobuf:"bytes,1,opt,name=controller_reply,json=controllerReply,oneof"`
}

type ControllerToTeam_AdvantageChoice struct {
	// the team is offered an advantage choice
	AdvantageChoice *AdvantageChoice `protobuf:"bytes,2,opt,name=advantage_choice,json=advantageChoice,oneof"`
}

func (*ControllerToTeam_ControllerReply) isControllerToTeam_Msg() {}

func (*ControllerToTeam_AdvantageChoice) isControllerToTeam_Msg() {}

// information about the advantage choice that is offered to a team
// there are currently no information, just a "ping"
type AdvantageChoice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// the type of foul that occurred
	//
	// Deprecated: Do not use.
	Foul *AdvantageChoice_Foul `protobuf:"varint,1,req,name=foul,enum=AdvantageChoice_Foul,def=0" json:"foul,omitempty"`
}

// Default values for AdvantageChoice fields.
const (
	Default_AdvantageChoice_Foul = AdvantageChoice_UNKNOWN
)

func (x *AdvantageChoice) Reset() {
	*x = AdvantageChoice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_ssl_gc_rcon_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AdvantageChoice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvantageChoice) ProtoMessage() {}

func (x *AdvantageChoice) ProtoReflect() protoreflect.Message {
	mi := &file_ssl_gc_rcon_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvantageChoice.ProtoReflect.Descriptor instead.
func (*AdvantageChoice) Descriptor() ([]byte, []int) {
	return file_ssl_gc_rcon_team_proto_rawDescGZIP(), []int{3}
}

// Deprecated: Do not use.
func (x *AdvantageChoice) GetFoul() AdvantageChoice_Foul {
	if x != nil && x.Foul != nil {
		return *x.Foul
	}
	return Default_AdvantageChoice_Foul
}

var File_ssl_gc_rcon_team_proto protoreflect.FileDescriptor

var file_ssl_gc_rcon_team_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63, 0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x5f, 0x74, 0x65,
	0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x11, 0x73, 0x73, 0x6c, 0x5f, 0x67, 0x63,
	0x5f, 0x72, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x13, 0x73, 0x73, 0x6c,
	0x5f, 0x67, 0x63, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x74, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1b, 0x0a, 0x09, 0x74, 0x65, 0x61, 0x6d, 0x5f, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x01, 0x20, 0x02, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x61, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x28, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x74,
	0x65, 0x61, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x05, 0x2e, 0x54, 0x65, 0x61, 0x6d,
	0x52, 0x04, 0x74, 0x65, 0x61, 0x6d, 0x22, 0xf0, 0x01, 0x0a, 0x10, 0x54, 0x65, 0x61, 0x6d, 0x54,
	0x6f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x12, 0x28, 0x0a, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0a,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64,
	0x5f, 0x6b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52,
	0x0d, 0x64, 0x65, 0x73, 0x69, 0x72, 0x65, 0x64, 0x4b, 0x65, 0x65, 0x70, 0x65, 0x72, 0x12, 0x43,
	0x0a, 0x12, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12, 0x2e, 0x41, 0x64, 0x76,
	0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x48, 0x00,
	0x52, 0x11, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x27, 0x0a, 0x0e, 0x73, 0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74,
	0x65, 0x5f, 0x62, 0x6f, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x0d, 0x73,
	0x75, 0x62, 0x73, 0x74, 0x69, 0x74, 0x75, 0x74, 0x65, 0x42, 0x6f, 0x74, 0x12, 0x14, 0x0a, 0x04,
	0x70, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x48, 0x00, 0x52, 0x04, 0x70, 0x69,
	0x6e, 0x67, 0x42, 0x05, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x97, 0x01, 0x0a, 0x10, 0x43, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x54, 0x6f, 0x54, 0x65, 0x61, 0x6d, 0x12, 0x3d,
	0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x72, 0x65, 0x70,
	0x6c, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72,
	0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x48, 0x00, 0x52, 0x0f, 0x63, 0x6f,
	0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3d, 0x0a,
	0x10, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x68, 0x6f, 0x69, 0x63,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74,
	0x61, 0x67, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x48, 0x00, 0x52, 0x0f, 0x61, 0x64, 0x76,
	0x61, 0x6e, 0x74, 0x61, 0x67, 0x65, 0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x42, 0x05, 0x0a, 0x03,
	0x6d, 0x73, 0x67, 0x22, 0x70, 0x0a, 0x0f, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x12, 0x36, 0x0a, 0x04, 0x66, 0x6f, 0x75, 0x6c, 0x18, 0x01,
	0x20, 0x02, 0x28, 0x0e, 0x32, 0x15, 0x2e, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61, 0x67, 0x65,
	0x43, 0x68, 0x6f, 0x69, 0x63, 0x65, 0x2e, 0x46, 0x6f, 0x75, 0x6c, 0x3a, 0x07, 0x55, 0x4e, 0x4b,
	0x4e, 0x4f, 0x57, 0x4e, 0x42, 0x02, 0x18, 0x01, 0x52, 0x04, 0x66, 0x6f, 0x75, 0x6c, 0x22, 0x13,
	0x0a, 0x04, 0x46, 0x6f, 0x75, 0x6c, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57,
	0x4e, 0x10, 0x00, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x03, 0x10, 0x04, 0x4a,
	0x04, 0x08, 0x04, 0x10, 0x05, 0x2a, 0x3e, 0x0a, 0x11, 0x41, 0x64, 0x76, 0x61, 0x6e, 0x74, 0x61,
	0x67, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x08, 0x0a, 0x04, 0x53, 0x54,
	0x4f, 0x50, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x4f, 0x4e, 0x54, 0x49, 0x4e, 0x55, 0x45,
	0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e, 0x44, 0x45, 0x43, 0x49, 0x44, 0x45, 0x44, 0x10,
	0x00, 0x1a, 0x02, 0x10, 0x01, 0x42, 0x3e, 0x5a, 0x3c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x52, 0x6f, 0x62, 0x6f, 0x43, 0x75, 0x70, 0x2d, 0x53, 0x53, 0x4c, 0x2f,
	0x73, 0x73, 0x6c, 0x2d, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c,
	0x6c, 0x65, 0x72, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x70, 0x70,
	0x2f, 0x72, 0x63, 0x6f, 0x6e,
}

var (
	file_ssl_gc_rcon_team_proto_rawDescOnce sync.Once
	file_ssl_gc_rcon_team_proto_rawDescData = file_ssl_gc_rcon_team_proto_rawDesc
)

func file_ssl_gc_rcon_team_proto_rawDescGZIP() []byte {
	file_ssl_gc_rcon_team_proto_rawDescOnce.Do(func() {
		file_ssl_gc_rcon_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_ssl_gc_rcon_team_proto_rawDescData)
	})
	return file_ssl_gc_rcon_team_proto_rawDescData
}

var file_ssl_gc_rcon_team_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_ssl_gc_rcon_team_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_ssl_gc_rcon_team_proto_goTypes = []interface{}{
	(AdvantageResponse)(0),    // 0: AdvantageResponse
	(AdvantageChoice_Foul)(0), // 1: AdvantageChoice.Foul
	(*TeamRegistration)(nil),  // 2: TeamRegistration
	(*TeamToController)(nil),  // 3: TeamToController
	(*ControllerToTeam)(nil),  // 4: ControllerToTeam
	(*AdvantageChoice)(nil),   // 5: AdvantageChoice
	(*Signature)(nil),         // 6: Signature
	(state.Team)(0),           // 7: Team
	(*ControllerReply)(nil),   // 8: ControllerReply
}
var file_ssl_gc_rcon_team_proto_depIdxs = []int32{
	6, // 0: TeamRegistration.signature:type_name -> Signature
	7, // 1: TeamRegistration.team:type_name -> Team
	6, // 2: TeamToController.signature:type_name -> Signature
	0, // 3: TeamToController.advantage_response:type_name -> AdvantageResponse
	8, // 4: ControllerToTeam.controller_reply:type_name -> ControllerReply
	5, // 5: ControllerToTeam.advantage_choice:type_name -> AdvantageChoice
	1, // 6: AdvantageChoice.foul:type_name -> AdvantageChoice.Foul
	7, // [7:7] is the sub-list for method output_type
	7, // [7:7] is the sub-list for method input_type
	7, // [7:7] is the sub-list for extension type_name
	7, // [7:7] is the sub-list for extension extendee
	0, // [0:7] is the sub-list for field type_name
}

func init() { file_ssl_gc_rcon_team_proto_init() }
func file_ssl_gc_rcon_team_proto_init() {
	if File_ssl_gc_rcon_team_proto != nil {
		return
	}
	file_ssl_gc_rcon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_ssl_gc_rcon_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamRegistration); i {
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
		file_ssl_gc_rcon_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TeamToController); i {
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
		file_ssl_gc_rcon_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControllerToTeam); i {
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
		file_ssl_gc_rcon_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AdvantageChoice); i {
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
	file_ssl_gc_rcon_team_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*TeamToController_DesiredKeeper)(nil),
		(*TeamToController_AdvantageResponse)(nil),
		(*TeamToController_SubstituteBot)(nil),
		(*TeamToController_Ping)(nil),
	}
	file_ssl_gc_rcon_team_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*ControllerToTeam_ControllerReply)(nil),
		(*ControllerToTeam_AdvantageChoice)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_ssl_gc_rcon_team_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_ssl_gc_rcon_team_proto_goTypes,
		DependencyIndexes: file_ssl_gc_rcon_team_proto_depIdxs,
		EnumInfos:         file_ssl_gc_rcon_team_proto_enumTypes,
		MessageInfos:      file_ssl_gc_rcon_team_proto_msgTypes,
	}.Build()
	File_ssl_gc_rcon_team_proto = out.File
	file_ssl_gc_rcon_team_proto_rawDesc = nil
	file_ssl_gc_rcon_team_proto_goTypes = nil
	file_ssl_gc_rcon_team_proto_depIdxs = nil
}
