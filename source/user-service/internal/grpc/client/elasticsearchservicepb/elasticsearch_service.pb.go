// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.0
// source: elasticsearch_service.proto

package elasticsearchservicepb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type GetUsersRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Offset        int32                  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit         int32                  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	SortBy        string                 `protobuf:"bytes,3,opt,name=sort_by,json=sortBy,proto3" json:"sort_by,omitempty"`
	FullName      string                 `protobuf:"bytes,4,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,5,opt,name=email,proto3" json:"email,omitempty"`
	Username      string                 `protobuf:"bytes,6,opt,name=username,proto3" json:"username,omitempty"`
	Address       string                 `protobuf:"bytes,7,opt,name=address,proto3" json:"address,omitempty"`
	RoleName      string                 `protobuf:"bytes,8,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	CreatedAtGte  string                 `protobuf:"bytes,9,opt,name=created_at_gte,json=createdAtGte,proto3" json:"created_at_gte,omitempty"`
	CreatedAtLte  string                 `protobuf:"bytes,10,opt,name=created_at_lte,json=createdAtLte,proto3" json:"created_at_lte,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersRequest) Reset() {
	*x = GetUsersRequest{}
	mi := &file_elasticsearch_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersRequest) ProtoMessage() {}

func (x *GetUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersRequest.ProtoReflect.Descriptor instead.
func (*GetUsersRequest) Descriptor() ([]byte, []int) {
	return file_elasticsearch_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetUsersRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetUsersRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetUsersRequest) GetSortBy() string {
	if x != nil {
		return x.SortBy
	}
	return ""
}

func (x *GetUsersRequest) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *GetUsersRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetUsersRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *GetUsersRequest) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *GetUsersRequest) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *GetUsersRequest) GetCreatedAtGte() string {
	if x != nil {
		return x.CreatedAtGte
	}
	return ""
}

func (x *GetUsersRequest) GetCreatedAtLte() string {
	if x != nil {
		return x.CreatedAtLte
	}
	return ""
}

type GetUsersResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Code          string                 `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Status        int32                  `protobuf:"varint,2,opt,name=status,proto3" json:"status,omitempty"`
	Message       string                 `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	Users         []*User                `protobuf:"bytes,4,rep,name=users,proto3" json:"users,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetUsersResponse) Reset() {
	*x = GetUsersResponse{}
	mi := &file_elasticsearch_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUsersResponse) ProtoMessage() {}

func (x *GetUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUsersResponse.ProtoReflect.Descriptor instead.
func (*GetUsersResponse) Descriptor() ([]byte, []int) {
	return file_elasticsearch_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetUsersResponse) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *GetUsersResponse) GetStatus() int32 {
	if x != nil {
		return x.Status
	}
	return 0
}

func (x *GetUsersResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	FullName      string                 `protobuf:"bytes,2,opt,name=full_name,json=fullName,proto3" json:"full_name,omitempty"`
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Username      string                 `protobuf:"bytes,4,opt,name=username,proto3" json:"username,omitempty"`
	Address       string                 `protobuf:"bytes,5,opt,name=address,proto3" json:"address,omitempty"`
	RoleName      string                 `protobuf:"bytes,6,opt,name=role_name,json=roleName,proto3" json:"role_name,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *User) Reset() {
	*x = User{}
	mi := &file_elasticsearch_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_elasticsearch_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_elasticsearch_service_proto_rawDescGZIP(), []int{2}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetFullName() string {
	if x != nil {
		return x.FullName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *User) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *User) GetRoleName() string {
	if x != nil {
		return x.RoleName
	}
	return ""
}

func (x *User) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *User) GetUpdatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.UpdatedAt
	}
	return nil
}

var File_elasticsearch_service_proto protoreflect.FileDescriptor

const file_elasticsearch_service_proto_rawDesc = "" +
	"\n" +
	"\x1belasticsearch_service.proto\x12\x16elasticsearchservicepb\x1a\x1fgoogle/protobuf/timestamp.proto\"\xaa\x02\n" +
	"\x0fGetUsersRequest\x12\x16\n" +
	"\x06offset\x18\x01 \x01(\x05R\x06offset\x12\x14\n" +
	"\x05limit\x18\x02 \x01(\x05R\x05limit\x12\x17\n" +
	"\asort_by\x18\x03 \x01(\tR\x06sortBy\x12\x1b\n" +
	"\tfull_name\x18\x04 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x05 \x01(\tR\x05email\x12\x1a\n" +
	"\busername\x18\x06 \x01(\tR\busername\x12\x18\n" +
	"\aaddress\x18\a \x01(\tR\aaddress\x12\x1b\n" +
	"\trole_name\x18\b \x01(\tR\broleName\x12$\n" +
	"\x0ecreated_at_gte\x18\t \x01(\tR\fcreatedAtGte\x12$\n" +
	"\x0ecreated_at_lte\x18\n" +
	" \x01(\tR\fcreatedAtLte\"\x8c\x01\n" +
	"\x10GetUsersResponse\x12\x12\n" +
	"\x04code\x18\x01 \x01(\tR\x04code\x12\x16\n" +
	"\x06status\x18\x02 \x01(\x05R\x06status\x12\x18\n" +
	"\amessage\x18\x03 \x01(\tR\amessage\x122\n" +
	"\x05users\x18\x04 \x03(\v2\x1c.elasticsearchservicepb.UserR\x05users\"\x92\x02\n" +
	"\x04User\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1b\n" +
	"\tfull_name\x18\x02 \x01(\tR\bfullName\x12\x14\n" +
	"\x05email\x18\x03 \x01(\tR\x05email\x12\x1a\n" +
	"\busername\x18\x04 \x01(\tR\busername\x12\x18\n" +
	"\aaddress\x18\x05 \x01(\tR\aaddress\x12\x1b\n" +
	"\trole_name\x18\x06 \x01(\tR\broleName\x129\n" +
	"\n" +
	"created_at\x18\a \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x129\n" +
	"\n" +
	"updated_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tupdatedAt2y\n" +
	"\x18ElasticsearchServiceGRPC\x12]\n" +
	"\bGetUsers\x12'.elasticsearchservicepb.GetUsersRequest\x1a(.elasticsearchservicepb.GetUsersResponseB\x19Z\x17elasticsearchservicepb/b\x06proto3"

var (
	file_elasticsearch_service_proto_rawDescOnce sync.Once
	file_elasticsearch_service_proto_rawDescData []byte
)

func file_elasticsearch_service_proto_rawDescGZIP() []byte {
	file_elasticsearch_service_proto_rawDescOnce.Do(func() {
		file_elasticsearch_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_elasticsearch_service_proto_rawDesc), len(file_elasticsearch_service_proto_rawDesc)))
	})
	return file_elasticsearch_service_proto_rawDescData
}

var file_elasticsearch_service_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_elasticsearch_service_proto_goTypes = []any{
	(*GetUsersRequest)(nil),       // 0: elasticsearchservicepb.GetUsersRequest
	(*GetUsersResponse)(nil),      // 1: elasticsearchservicepb.GetUsersResponse
	(*User)(nil),                  // 2: elasticsearchservicepb.User
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_elasticsearch_service_proto_depIdxs = []int32{
	2, // 0: elasticsearchservicepb.GetUsersResponse.users:type_name -> elasticsearchservicepb.User
	3, // 1: elasticsearchservicepb.User.created_at:type_name -> google.protobuf.Timestamp
	3, // 2: elasticsearchservicepb.User.updated_at:type_name -> google.protobuf.Timestamp
	0, // 3: elasticsearchservicepb.ElasticsearchServiceGRPC.GetUsers:input_type -> elasticsearchservicepb.GetUsersRequest
	1, // 4: elasticsearchservicepb.ElasticsearchServiceGRPC.GetUsers:output_type -> elasticsearchservicepb.GetUsersResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_elasticsearch_service_proto_init() }
func file_elasticsearch_service_proto_init() {
	if File_elasticsearch_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_elasticsearch_service_proto_rawDesc), len(file_elasticsearch_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_elasticsearch_service_proto_goTypes,
		DependencyIndexes: file_elasticsearch_service_proto_depIdxs,
		MessageInfos:      file_elasticsearch_service_proto_msgTypes,
	}.Build()
	File_elasticsearch_service_proto = out.File
	file_elasticsearch_service_proto_goTypes = nil
	file_elasticsearch_service_proto_depIdxs = nil
}
