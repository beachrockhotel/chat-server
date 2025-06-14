// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v3.21.12
// source: chat_v1/chat.proto

package chat_v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Chat struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Info          *ChatInfo              `protobuf:"bytes,2,opt,name=info,proto3" json:"info,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chat) Reset() {
	*x = Chat{}
	mi := &file_chat_v1_chat_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chat) ProtoMessage() {}

func (x *Chat) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chat.ProtoReflect.Descriptor instead.
func (*Chat) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{0}
}

func (x *Chat) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Chat) GetInfo() *ChatInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

func (x *Chat) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Chat) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

type ChatInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Title         string                 `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	Usernames     string                 `protobuf:"bytes,2,opt,name=usernames,proto3" json:"usernames,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChatInfo) Reset() {
	*x = ChatInfo{}
	mi := &file_chat_v1_chat_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChatInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChatInfo) ProtoMessage() {}

func (x *ChatInfo) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChatInfo.ProtoReflect.Descriptor instead.
func (*ChatInfo) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{1}
}

func (x *ChatInfo) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *ChatInfo) GetUsernames() string {
	if x != nil {
		return x.Usernames
	}
	return ""
}

type CreateRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Info          *ChatInfo              `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	mi := &file_chat_v1_chat_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{2}
}

func (x *CreateRequest) GetInfo() *ChatInfo {
	if x != nil {
		return x.Info
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	mi := &file_chat_v1_chat_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{3}
}

func (x *CreateResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	mi := &file_chat_v1_chat_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Chat          *Chat                  `protobuf:"bytes,1,opt,name=chat,proto3" json:"chat,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	mi := &file_chat_v1_chat_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{5}
}

func (x *GetResponse) GetChat() *Chat {
	if x != nil {
		return x.Chat
	}
	return nil
}

type DeleteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *DeleteRequest) Reset() {
	*x = DeleteRequest{}
	mi := &file_chat_v1_chat_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *DeleteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteRequest) ProtoMessage() {}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteRequest.ProtoReflect.Descriptor instead.
func (*DeleteRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type SendMessageRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	From          string                 `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	Text          string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SendMessageRequest) Reset() {
	*x = SendMessageRequest{}
	mi := &file_chat_v1_chat_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SendMessageRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendMessageRequest) ProtoMessage() {}

func (x *SendMessageRequest) ProtoReflect() protoreflect.Message {
	mi := &file_chat_v1_chat_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendMessageRequest.ProtoReflect.Descriptor instead.
func (*SendMessageRequest) Descriptor() ([]byte, []int) {
	return file_chat_v1_chat_proto_rawDescGZIP(), []int{7}
}

func (x *SendMessageRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *SendMessageRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *SendMessageRequest) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

var File_chat_v1_chat_proto protoreflect.FileDescriptor

const file_chat_v1_chat_proto_rawDesc = "" +
	"\n" +
	"\x12chat_v1/chat.proto\x12\achat_v1\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1bgoogle/protobuf/empty.proto\"\xb3\x01\n" +
	"\x04Chat\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12%\n" +
	"\x04info\x18\x02 \x01(\v2\x11.chat_v1.ChatInfoR\x04info\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\x04 \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt\">\n" +
	"\bChatInfo\x12\x14\n" +
	"\x05title\x18\x01 \x01(\tR\x05title\x12\x1c\n" +
	"\tusernames\x18\x02 \x01(\tR\tusernames\"6\n" +
	"\rCreateRequest\x12%\n" +
	"\x04info\x18\x01 \x01(\v2\x11.chat_v1.ChatInfoR\x04info\" \n" +
	"\x0eCreateResponse\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"\x1c\n" +
	"\n" +
	"GetRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"0\n" +
	"\vGetResponse\x12!\n" +
	"\x04chat\x18\x01 \x01(\v2\r.chat_v1.ChatR\x04chat\"\x1f\n" +
	"\rDeleteRequest\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\"w\n" +
	"\x12SendMessageRequest\x12\x12\n" +
	"\x04from\x18\x01 \x01(\tR\x04from\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt2\xf3\x01\n" +
	"\x06ChatV1\x129\n" +
	"\x06Create\x12\x16.chat_v1.CreateRequest\x1a\x17.chat_v1.CreateResponse\x120\n" +
	"\x03Get\x12\x13.chat_v1.GetRequest\x1a\x14.chat_v1.GetResponse\x128\n" +
	"\x06Delete\x12\x16.chat_v1.DeleteRequest\x1a\x16.google.protobuf.Empty\x12B\n" +
	"\vSendMessage\x12\x1b.chat_v1.SendMessageRequest\x1a\x16.google.protobuf.EmptyB;Z9github.com/beachrockhotel/chat-server/pkg/chat_v1;chat_v1b\x06proto3"

var (
	file_chat_v1_chat_proto_rawDescOnce sync.Once
	file_chat_v1_chat_proto_rawDescData []byte
)

func file_chat_v1_chat_proto_rawDescGZIP() []byte {
	file_chat_v1_chat_proto_rawDescOnce.Do(func() {
		file_chat_v1_chat_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_chat_v1_chat_proto_rawDesc), len(file_chat_v1_chat_proto_rawDesc)))
	})
	return file_chat_v1_chat_proto_rawDescData
}

var file_chat_v1_chat_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_chat_v1_chat_proto_goTypes = []any{
	(*Chat)(nil),                  // 0: chat_v1.Chat
	(*ChatInfo)(nil),              // 1: chat_v1.ChatInfo
	(*CreateRequest)(nil),         // 2: chat_v1.CreateRequest
	(*CreateResponse)(nil),        // 3: chat_v1.CreateResponse
	(*GetRequest)(nil),            // 4: chat_v1.GetRequest
	(*GetResponse)(nil),           // 5: chat_v1.GetResponse
	(*DeleteRequest)(nil),         // 6: chat_v1.DeleteRequest
	(*SendMessageRequest)(nil),    // 7: chat_v1.SendMessageRequest
	(*timestamppb.Timestamp)(nil), // 8: google.protobuf.Timestamp
	(*emptypb.Empty)(nil),         // 9: google.protobuf.Empty
}
var file_chat_v1_chat_proto_depIdxs = []int32{
	1,  // 0: chat_v1.Chat.info:type_name -> chat_v1.ChatInfo
	8,  // 1: chat_v1.Chat.created_at:type_name -> google.protobuf.Timestamp
	8,  // 2: chat_v1.Chat.updated_at:type_name -> google.protobuf.Timestamp
	1,  // 3: chat_v1.CreateRequest.info:type_name -> chat_v1.ChatInfo
	0,  // 4: chat_v1.GetResponse.chat:type_name -> chat_v1.Chat
	8,  // 5: chat_v1.SendMessageRequest.created_at:type_name -> google.protobuf.Timestamp
	2,  // 6: chat_v1.ChatV1.Create:input_type -> chat_v1.CreateRequest
	4,  // 7: chat_v1.ChatV1.Get:input_type -> chat_v1.GetRequest
	6,  // 8: chat_v1.ChatV1.Delete:input_type -> chat_v1.DeleteRequest
	7,  // 9: chat_v1.ChatV1.SendMessage:input_type -> chat_v1.SendMessageRequest
	3,  // 10: chat_v1.ChatV1.Create:output_type -> chat_v1.CreateResponse
	5,  // 11: chat_v1.ChatV1.Get:output_type -> chat_v1.GetResponse
	9,  // 12: chat_v1.ChatV1.Delete:output_type -> google.protobuf.Empty
	9,  // 13: chat_v1.ChatV1.SendMessage:output_type -> google.protobuf.Empty
	10, // [10:14] is the sub-list for method output_type
	6,  // [6:10] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_chat_v1_chat_proto_init() }
func file_chat_v1_chat_proto_init() {
	if File_chat_v1_chat_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_chat_v1_chat_proto_rawDesc), len(file_chat_v1_chat_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_chat_v1_chat_proto_goTypes,
		DependencyIndexes: file_chat_v1_chat_proto_depIdxs,
		MessageInfos:      file_chat_v1_chat_proto_msgTypes,
	}.Build()
	File_chat_v1_chat_proto = out.File
	file_chat_v1_chat_proto_goTypes = nil
	file_chat_v1_chat_proto_depIdxs = nil
}
