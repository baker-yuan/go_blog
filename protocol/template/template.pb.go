// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v3.21.12
// source: template/template.proto

package template

import (
	reflect "reflect"
	sync "sync"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 资源类型
type ResourceType int32

const (
	ResourceType_RT_UNKNOWN ResourceType = 0 // 未知
)

// Enum value maps for ResourceType.
var (
	ResourceType_name = map[int32]string{
		0: "RT_UNKNOWN",
	}
	ResourceType_value = map[string]int32{
		"RT_UNKNOWN": 0,
	}
)

func (x ResourceType) Enum() *ResourceType {
	p := new(ResourceType)
	*p = x
	return p
}

func (x ResourceType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ResourceType) Descriptor() protoreflect.EnumDescriptor {
	return file_template_template_proto_enumTypes[0].Descriptor()
}

func (ResourceType) Type() protoreflect.EnumType {
	return &file_template_template_proto_enumTypes[0]
}

func (x ResourceType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ResourceType.Descriptor instead.
func (ResourceType) EnumDescriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{0}
}

// 空响应
type EmptyRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptyRsp) Reset() {
	*x = EmptyRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptyRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptyRsp) ProtoMessage() {}

func (x *EmptyRsp) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptyRsp.ProtoReflect.Descriptor instead.
func (*EmptyRsp) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{0}
}

// 添加修改资源返回值
type AddOrUpdateRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"` // 资源id
}

func (x *AddOrUpdateRsp) Reset() {
	*x = AddOrUpdateRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateRsp) ProtoMessage() {}

func (x *AddOrUpdateRsp) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateRsp.ProtoReflect.Descriptor instead.
func (*AddOrUpdateRsp) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{1}
}

func (x *AddOrUpdateRsp) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

// 友链
type FriendLink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          *uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"`                                     // 唯一标识
	LinkName    *string `protobuf:"bytes,2,opt,name=link_name,json=linkName,proto3,oneof" json:"link_name,omitempty"`          // 链接名
	LinkAvatar  *string `protobuf:"bytes,3,opt,name=link_avatar,json=linkAvatar,proto3,oneof" json:"link_avatar,omitempty"`    // 链接头像
	LinkAddress *string `protobuf:"bytes,4,opt,name=link_address,json=linkAddress,proto3,oneof" json:"link_address,omitempty"` // 链接地址
	LinkIntro   *string `protobuf:"bytes,5,opt,name=link_intro,json=linkIntro,proto3,oneof" json:"link_intro,omitempty"`       // 链接介绍
	Sort        *int32  `protobuf:"varint,6,opt,name=sort,proto3,oneof" json:"sort,omitempty"`                                 // 友链排序
	Status      *int32  `protobuf:"varint,7,opt,name=status,proto3,oneof" json:"status,omitempty"`                             // 友链状态（0-未发布，1-已发布，2-已下线）
	IsDeleted   *bool   `protobuf:"varint,8,opt,name=is_deleted,json=isDeleted,proto3,oneof" json:"is_deleted,omitempty"`      // 是否删除
	CreateTime  *uint32 `protobuf:"varint,9,opt,name=create_time,json=createTime,proto3,oneof" json:"create_time,omitempty"`   // 创建时间
	UpdateTime  *uint32 `protobuf:"varint,10,opt,name=update_time,json=updateTime,proto3,oneof" json:"update_time,omitempty"`  // 修改时间
}

func (x *FriendLink) Reset() {
	*x = FriendLink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendLink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendLink) ProtoMessage() {}

func (x *FriendLink) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendLink.ProtoReflect.Descriptor instead.
func (*FriendLink) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{2}
}

func (x *FriendLink) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

func (x *FriendLink) GetLinkName() string {
	if x != nil && x.LinkName != nil {
		return *x.LinkName
	}
	return ""
}

func (x *FriendLink) GetLinkAvatar() string {
	if x != nil && x.LinkAvatar != nil {
		return *x.LinkAvatar
	}
	return ""
}

func (x *FriendLink) GetLinkAddress() string {
	if x != nil && x.LinkAddress != nil {
		return *x.LinkAddress
	}
	return ""
}

func (x *FriendLink) GetLinkIntro() string {
	if x != nil && x.LinkIntro != nil {
		return *x.LinkIntro
	}
	return ""
}

func (x *FriendLink) GetSort() int32 {
	if x != nil && x.Sort != nil {
		return *x.Sort
	}
	return 0
}

func (x *FriendLink) GetStatus() int32 {
	if x != nil && x.Status != nil {
		return *x.Status
	}
	return 0
}

func (x *FriendLink) GetIsDeleted() bool {
	if x != nil && x.IsDeleted != nil {
		return *x.IsDeleted
	}
	return false
}

func (x *FriendLink) GetCreateTime() uint32 {
	if x != nil && x.CreateTime != nil {
		return *x.CreateTime
	}
	return 0
}

func (x *FriendLink) GetUpdateTime() uint32 {
	if x != nil && x.UpdateTime != nil {
		return *x.UpdateTime
	}
	return 0
}

// 友链搜索
type SearchFriendLinkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageNum  *uint32 `protobuf:"varint,100,opt,name=page_num,json=pageNum,proto3,oneof" json:"page_num,omitempty"`    // 分页查询
	PageSize *uint32 `protobuf:"varint,101,opt,name=page_size,json=pageSize,proto3,oneof" json:"page_size,omitempty"` // 分页查询
}

func (x *SearchFriendLinkReq) Reset() {
	*x = SearchFriendLinkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFriendLinkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFriendLinkReq) ProtoMessage() {}

func (x *SearchFriendLinkReq) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFriendLinkReq.ProtoReflect.Descriptor instead.
func (*SearchFriendLinkReq) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{3}
}

func (x *SearchFriendLinkReq) GetPageNum() uint32 {
	if x != nil && x.PageNum != nil {
		return *x.PageNum
	}
	return 0
}

func (x *SearchFriendLinkReq) GetPageSize() uint32 {
	if x != nil && x.PageSize != nil {
		return *x.PageSize
	}
	return 0
}

type SearchFriendLinkRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total *uint32       `protobuf:"varint,1,opt,name=total,proto3,oneof" json:"total,omitempty"` // 分页查询
	Data  []*FriendLink `protobuf:"bytes,2,rep,name=data,proto3" json:"data,omitempty"`          // 友链集合
}

func (x *SearchFriendLinkRsp) Reset() {
	*x = SearchFriendLinkRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SearchFriendLinkRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SearchFriendLinkRsp) ProtoMessage() {}

func (x *SearchFriendLinkRsp) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SearchFriendLinkRsp.ProtoReflect.Descriptor instead.
func (*SearchFriendLinkRsp) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{4}
}

func (x *SearchFriendLinkRsp) GetTotal() uint32 {
	if x != nil && x.Total != nil {
		return *x.Total
	}
	return 0
}

func (x *SearchFriendLinkRsp) GetData() []*FriendLink {
	if x != nil {
		return x.Data
	}
	return nil
}

// 友链详情
type FriendLinkDetailReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"` // 编号
}

func (x *FriendLinkDetailReq) Reset() {
	*x = FriendLinkDetailReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FriendLinkDetailReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FriendLinkDetailReq) ProtoMessage() {}

func (x *FriendLinkDetailReq) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FriendLinkDetailReq.ProtoReflect.Descriptor instead.
func (*FriendLinkDetailReq) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{5}
}

func (x *FriendLinkDetailReq) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

// 添加修改友链
type AddOrUpdateFriendLinkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddOrUpdateFriendLinkReq) Reset() {
	*x = AddOrUpdateFriendLinkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddOrUpdateFriendLinkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddOrUpdateFriendLinkReq) ProtoMessage() {}

func (x *AddOrUpdateFriendLinkReq) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddOrUpdateFriendLinkReq.ProtoReflect.Descriptor instead.
func (*AddOrUpdateFriendLinkReq) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{6}
}

// 删除友链
type DeleteFriendLinkReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id *uint32 `protobuf:"varint,1,opt,name=id,proto3,oneof" json:"id,omitempty"` // 编号
}

func (x *DeleteFriendLinkReq) Reset() {
	*x = DeleteFriendLinkReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_template_template_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteFriendLinkReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteFriendLinkReq) ProtoMessage() {}

func (x *DeleteFriendLinkReq) ProtoReflect() protoreflect.Message {
	mi := &file_template_template_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteFriendLinkReq.ProtoReflect.Descriptor instead.
func (*DeleteFriendLinkReq) Descriptor() ([]byte, []int) {
	return file_template_template_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteFriendLinkReq) GetId() uint32 {
	if x != nil && x.Id != nil {
		return *x.Id
	}
	return 0
}

var File_template_template_proto protoreflect.FileDescriptor

var file_template_template_proto_rawDesc = []byte{
	0x0a, 0x17, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61,
	0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0a, 0x0a, 0x08, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x73, 0x70, 0x22, 0x2c, 0x0a,
	0x0e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x12,
	0x13, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x69,
	0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22, 0xe3, 0x03, 0x0a, 0x0a,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x13, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x12,
	0x20, 0x0a, 0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x48, 0x01, 0x52, 0x08, 0x6c, 0x69, 0x6e, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01,
	0x01, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x02, 0x52, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x76,
	0x61, 0x74, 0x61, 0x72, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x0c, 0x6c, 0x69, 0x6e, 0x6b, 0x5f,
	0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x0b, 0x6c, 0x69, 0x6e, 0x6b, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x88, 0x01, 0x01, 0x12,
	0x22, 0x0a, 0x0a, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x09, 0x6c, 0x69, 0x6e, 0x6b, 0x49, 0x6e, 0x74, 0x72, 0x6f,
	0x88, 0x01, 0x01, 0x12, 0x17, 0x0a, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x05, 0x52, 0x04, 0x73, 0x6f, 0x72, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1b, 0x0a, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x48, 0x06, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x88, 0x01, 0x01, 0x12, 0x22, 0x0a, 0x0a, 0x69, 0x73, 0x5f,
	0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x08, 0x48, 0x07, 0x52,
	0x09, 0x69, 0x73, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x88, 0x01, 0x01, 0x12, 0x24, 0x0a,
	0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x0d, 0x48, 0x08, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65,
	0x88, 0x01, 0x01, 0x12, 0x24, 0x0a, 0x0b, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69,
	0x6d, 0x65, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x09, 0x52, 0x0a, 0x75, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64,
	0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x0e,
	0x0a, 0x0c, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x42, 0x0f,
	0x0a, 0x0d, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x42,
	0x0d, 0x0a, 0x0b, 0x5f, 0x6c, 0x69, 0x6e, 0x6b, 0x5f, 0x69, 0x6e, 0x74, 0x72, 0x6f, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x73, 0x6f, 0x72, 0x74, 0x42, 0x09, 0x0a, 0x07, 0x5f, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0d, 0x0a, 0x0b, 0x5f, 0x69, 0x73, 0x5f, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x64, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x74, 0x69, 0x6d,
	0x65, 0x22, 0x84, 0x01, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x69, 0x65,
	0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x27, 0x0a, 0x08, 0x70, 0x61, 0x67,
	0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x64, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x2a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x88,
	0x01, 0x01, 0x12, 0x29, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x48, 0x01,
	0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0b, 0x0a,
	0x09, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x42, 0x0c, 0x0a, 0x0a, 0x5f, 0x70,
	0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x22, 0x69, 0x0a, 0x13, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x73, 0x70, 0x12,
	0x19, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x48, 0x00,
	0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x88, 0x01, 0x01, 0x12, 0x2d, 0x0a, 0x04, 0x64, 0x61,
	0x74, 0x61, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e,
	0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x22, 0x3a, 0x0a, 0x13, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e,
	0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x52, 0x65, 0x71, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x48,
	0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x22,
	0x1a, 0x0a, 0x18, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72,
	0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x22, 0x3a, 0x0a, 0x13, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x1c, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x42, 0x07,
	0xfa, 0x42, 0x04, 0x2a, 0x02, 0x20, 0x00, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x88, 0x01, 0x01,
	0x42, 0x05, 0x0a, 0x03, 0x5f, 0x69, 0x64, 0x2a, 0x1e, 0x0a, 0x0c, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x0a, 0x52, 0x54, 0x5f, 0x55, 0x4e,
	0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x10, 0x00, 0x32, 0xf6, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x41, 0x70, 0x69, 0x12, 0x5c, 0x0a, 0x10, 0x53, 0x65, 0x61, 0x72, 0x63,
	0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x22, 0x2e, 0x62, 0x6c,
	0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b,
	0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x61, 0x0a, 0x15, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x27,
	0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41,
	0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64,
	0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x74,
	0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x41, 0x64, 0x64, 0x4f, 0x72, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x22, 0x2e, 0x62,
	0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x1a, 0x17, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x73, 0x70, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x10, 0x46,
	0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x12,
	0x22, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c, 0x61, 0x74, 0x65, 0x2e,
	0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c,
	0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x62, 0x6c, 0x6f, 0x67, 0x2e, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x2e, 0x46, 0x72, 0x69, 0x65, 0x6e, 0x64, 0x4c, 0x69, 0x6e, 0x6b, 0x22, 0x00,
	0x42, 0x31, 0x5a, 0x2f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x62,
	0x61, 0x6b, 0x65, 0x72, 0x2d, 0x79, 0x75, 0x61, 0x6e, 0x2f, 0x67, 0x6f, 0x2d, 0x62, 0x6c, 0x6f,
	0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2f, 0x74, 0x65, 0x6d, 0x70, 0x6c,
	0x61, 0x74, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_template_template_proto_rawDescOnce sync.Once
	file_template_template_proto_rawDescData = file_template_template_proto_rawDesc
)

func file_template_template_proto_rawDescGZIP() []byte {
	file_template_template_proto_rawDescOnce.Do(func() {
		file_template_template_proto_rawDescData = protoimpl.X.CompressGZIP(file_template_template_proto_rawDescData)
	})
	return file_template_template_proto_rawDescData
}

var file_template_template_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_template_template_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_template_template_proto_goTypes = []interface{}{
	(ResourceType)(0),                // 0: blog.template.ResourceType
	(*EmptyRsp)(nil),                 // 1: blog.template.EmptyRsp
	(*AddOrUpdateRsp)(nil),           // 2: blog.template.AddOrUpdateRsp
	(*FriendLink)(nil),               // 3: blog.template.FriendLink
	(*SearchFriendLinkReq)(nil),      // 4: blog.template.SearchFriendLinkReq
	(*SearchFriendLinkRsp)(nil),      // 5: blog.template.SearchFriendLinkRsp
	(*FriendLinkDetailReq)(nil),      // 6: blog.template.FriendLinkDetailReq
	(*AddOrUpdateFriendLinkReq)(nil), // 7: blog.template.AddOrUpdateFriendLinkReq
	(*DeleteFriendLinkReq)(nil),      // 8: blog.template.DeleteFriendLinkReq
}
var file_template_template_proto_depIdxs = []int32{
	3, // 0: blog.template.SearchFriendLinkRsp.data:type_name -> blog.template.FriendLink
	4, // 1: blog.template.TemplateApi.SearchFriendLink:input_type -> blog.template.SearchFriendLinkReq
	7, // 2: blog.template.TemplateApi.AddOrUpdateFriendLink:input_type -> blog.template.AddOrUpdateFriendLinkReq
	8, // 3: blog.template.TemplateApi.DeleteFriendLink:input_type -> blog.template.DeleteFriendLinkReq
	6, // 4: blog.template.TemplateApi.FriendLinkDetail:input_type -> blog.template.FriendLinkDetailReq
	5, // 5: blog.template.TemplateApi.SearchFriendLink:output_type -> blog.template.SearchFriendLinkRsp
	2, // 6: blog.template.TemplateApi.AddOrUpdateFriendLink:output_type -> blog.template.AddOrUpdateRsp
	1, // 7: blog.template.TemplateApi.DeleteFriendLink:output_type -> blog.template.EmptyRsp
	3, // 8: blog.template.TemplateApi.FriendLinkDetail:output_type -> blog.template.FriendLink
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_template_template_proto_init() }
func file_template_template_proto_init() {
	if File_template_template_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_template_template_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptyRsp); i {
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
		file_template_template_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateRsp); i {
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
		file_template_template_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendLink); i {
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
		file_template_template_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFriendLinkReq); i {
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
		file_template_template_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SearchFriendLinkRsp); i {
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
		file_template_template_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FriendLinkDetailReq); i {
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
		file_template_template_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddOrUpdateFriendLinkReq); i {
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
		file_template_template_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteFriendLinkReq); i {
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
	file_template_template_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_template_template_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_template_template_proto_msgTypes[3].OneofWrappers = []interface{}{}
	file_template_template_proto_msgTypes[4].OneofWrappers = []interface{}{}
	file_template_template_proto_msgTypes[5].OneofWrappers = []interface{}{}
	file_template_template_proto_msgTypes[7].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_template_template_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_template_template_proto_goTypes,
		DependencyIndexes: file_template_template_proto_depIdxs,
		EnumInfos:         file_template_template_proto_enumTypes,
		MessageInfos:      file_template_template_proto_msgTypes,
	}.Build()
	File_template_template_proto = out.File
	file_template_template_proto_rawDesc = nil
	file_template_template_proto_goTypes = nil
	file_template_template_proto_depIdxs = nil
}
