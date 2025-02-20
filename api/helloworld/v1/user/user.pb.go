// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.27.0--rc3
// source: api/helloworld/v1/user/user.proto

package user

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Account struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`      // 自增ID
	Pass  string `protobuf:"bytes,2,opt,name=pass,proto3" json:"pass,omitempty"`   // 密码
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"` // email
	Phone string `protobuf:"bytes,4,opt,name=phone,proto3" json:"phone,omitempty"` // 手机号
}

func (x *Account) Reset() {
	*x = Account{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_helloworld_v1_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Account) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Account) ProtoMessage() {}

func (x *Account) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Account.ProtoReflect.Descriptor instead.
func (*Account) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *Account) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Account) GetPass() string {
	if x != nil {
		return x.Pass
	}
	return ""
}

func (x *Account) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Account) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type GetAccountReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Phone string `protobuf:"bytes,1,opt,name=phone,proto3" json:"phone,omitempty"`
}

func (x *GetAccountReq) Reset() {
	*x = GetAccountReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_helloworld_v1_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccountReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountReq) ProtoMessage() {}

func (x *GetAccountReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountReq.ProtoReflect.Descriptor instead.
func (*GetAccountReq) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccountReq) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

type Player struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	AccountId int64  `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"` // 账户ID
	PlayerId  int64  `protobuf:"varint,3,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`    // 玩家ID
	Nickname  string `protobuf:"bytes,4,opt,name=nickname,proto3" json:"nickname,omitempty"`
	Gender    int32  `protobuf:"varint,5,opt,name=gender,proto3" json:"gender,omitempty"` // 性别
	Avatar    string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`  // 头像地址
}

func (x *Player) Reset() {
	*x = Player{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_helloworld_v1_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Player) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Player) ProtoMessage() {}

func (x *Player) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Player.ProtoReflect.Descriptor instead.
func (*Player) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *Player) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Player) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *Player) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *Player) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *Player) GetGender() int32 {
	if x != nil {
		return x.Gender
	}
	return 0
}

func (x *Player) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

type GetPlayerByIdReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetPlayerByIdReq) Reset() {
	*x = GetPlayerByIdReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_helloworld_v1_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetPlayerByIdReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPlayerByIdReq) ProtoMessage() {}

func (x *GetPlayerByIdReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_helloworld_v1_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPlayerByIdReq.ProtoReflect.Descriptor instead.
func (*GetPlayerByIdReq) Descriptor() ([]byte, []int) {
	return file_api_helloworld_v1_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *GetPlayerByIdReq) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

var File_api_helloworld_v1_user_user_proto protoreflect.FileDescriptor

var file_api_helloworld_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x68, 0x65, 0x6c, 0x6c, 0x6f, 0x77, 0x6f, 0x72, 0x6c, 0x64,
	0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x59, 0x0a, 0x07, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x73, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x70, 0x61, 0x73, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x70,
	0x68, 0x6f, 0x6e, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e,
	0x65, 0x22, 0x25, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x22, 0xa0, 0x01, 0x0a, 0x06, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x6e, 0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67,
	0x65, 0x6e, 0x64, 0x65, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x67, 0x65, 0x6e,
	0x64, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x22, 0x22, 0x0a, 0x10, 0x47,
	0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x32,
	0x76, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x30,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x13, 0x2e, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x1a, 0x0d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x35, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0c, 0x2e, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x42, 0x15, 0x5a, 0x13, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_helloworld_v1_user_user_proto_rawDescOnce sync.Once
	file_api_helloworld_v1_user_user_proto_rawDescData = file_api_helloworld_v1_user_user_proto_rawDesc
)

func file_api_helloworld_v1_user_user_proto_rawDescGZIP() []byte {
	file_api_helloworld_v1_user_user_proto_rawDescOnce.Do(func() {
		file_api_helloworld_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_helloworld_v1_user_user_proto_rawDescData)
	})
	return file_api_helloworld_v1_user_user_proto_rawDescData
}

var file_api_helloworld_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_api_helloworld_v1_user_user_proto_goTypes = []any{
	(*Account)(nil),          // 0: user.Account
	(*GetAccountReq)(nil),    // 1: user.GetAccountReq
	(*Player)(nil),           // 2: user.Player
	(*GetPlayerByIdReq)(nil), // 3: user.GetPlayerByIdReq
	(*emptypb.Empty)(nil),    // 4: google.protobuf.Empty
}
var file_api_helloworld_v1_user_user_proto_depIdxs = []int32{
	1, // 0: user.UserService.GetAccount:input_type -> user.GetAccountReq
	4, // 1: user.UserService.GetPlayerById:input_type -> google.protobuf.Empty
	0, // 2: user.UserService.GetAccount:output_type -> user.Account
	2, // 3: user.UserService.GetPlayerById:output_type -> user.Player
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_helloworld_v1_user_user_proto_init() }
func file_api_helloworld_v1_user_user_proto_init() {
	if File_api_helloworld_v1_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_helloworld_v1_user_user_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Account); i {
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
		file_api_helloworld_v1_user_user_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetAccountReq); i {
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
		file_api_helloworld_v1_user_user_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*Player); i {
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
		file_api_helloworld_v1_user_user_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*GetPlayerByIdReq); i {
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
			RawDescriptor: file_api_helloworld_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_helloworld_v1_user_user_proto_goTypes,
		DependencyIndexes: file_api_helloworld_v1_user_user_proto_depIdxs,
		MessageInfos:      file_api_helloworld_v1_user_user_proto_msgTypes,
	}.Build()
	File_api_helloworld_v1_user_user_proto = out.File
	file_api_helloworld_v1_user_user_proto_rawDesc = nil
	file_api_helloworld_v1_user_user_proto_goTypes = nil
	file_api_helloworld_v1_user_user_proto_depIdxs = nil
}
