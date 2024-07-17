// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: comments.proto

package genproto

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

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	MemoryId  string `protobuf:"bytes,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,6,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{0}
}

func (x *Comment) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *Comment) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *Comment) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Comment) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type ByMemoryId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
}

func (x *ByMemoryId) Reset() {
	*x = ByMemoryId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ByMemoryId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ByMemoryId) ProtoMessage() {}

func (x *ByMemoryId) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ByMemoryId.ProtoReflect.Descriptor instead.
func (*ByMemoryId) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{1}
}

func (x *ByMemoryId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *ByMemoryId) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *ByMemoryId) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *ByMemoryId) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

type AddCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	MemoryId  string `protobuf:"bytes,2,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content   string `protobuf:"bytes,4,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *AddCommentRequest) Reset() {
	*x = AddCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentRequest) ProtoMessage() {}

func (x *AddCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentRequest.ProtoReflect.Descriptor instead.
func (*AddCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{2}
}

func (x *AddCommentRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *AddCommentRequest) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *AddCommentRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *AddCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type AddCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AddCommentResponse) Reset() {
	*x = AddCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddCommentResponse) ProtoMessage() {}

func (x *AddCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddCommentResponse.ProtoReflect.Descriptor instead.
func (*AddCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{3}
}

type GetByIdMemoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId string `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
}

func (x *GetByIdMemoryRequest) Reset() {
	*x = GetByIdMemoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdMemoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdMemoryRequest) ProtoMessage() {}

func (x *GetByIdMemoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdMemoryRequest.ProtoReflect.Descriptor instead.
func (*GetByIdMemoryRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{4}
}

func (x *GetByIdMemoryRequest) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

type GetByIdMemoryResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comments []*ByMemoryId `protobuf:"bytes,1,rep,name=comments,proto3" json:"comments,omitempty"`
}

func (x *GetByIdMemoryResponse) Reset() {
	*x = GetByIdMemoryResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdMemoryResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdMemoryResponse) ProtoMessage() {}

func (x *GetByIdMemoryResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdMemoryResponse.ProtoReflect.Descriptor instead.
func (*GetByIdMemoryResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{5}
}

func (x *GetByIdMemoryResponse) GetComments() []*ByMemoryId {
	if x != nil {
		return x.Comments
	}
	return nil
}

type UpdateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MemoryId  string `protobuf:"bytes,1,opt,name=memory_id,json=memoryId,proto3" json:"memory_id,omitempty"`
	CommentId string `protobuf:"bytes,2,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
	Content   string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *UpdateCommentRequest) Reset() {
	*x = UpdateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentRequest) ProtoMessage() {}

func (x *UpdateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{6}
}

func (x *UpdateCommentRequest) GetMemoryId() string {
	if x != nil {
		return x.MemoryId
	}
	return ""
}

func (x *UpdateCommentRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

func (x *UpdateCommentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type UpdateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCommentResponse) Reset() {
	*x = UpdateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCommentResponse) ProtoMessage() {}

func (x *UpdateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCommentResponse.ProtoReflect.Descriptor instead.
func (*UpdateCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{7}
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{8}
}

func (x *DeleteCommentRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type DeleteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{9}
}

type GetByCommentIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CommentId string `protobuf:"bytes,1,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"`
}

func (x *GetByCommentIdRequest) Reset() {
	*x = GetByCommentIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByCommentIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByCommentIdRequest) ProtoMessage() {}

func (x *GetByCommentIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByCommentIdRequest.ProtoReflect.Descriptor instead.
func (*GetByCommentIdRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{10}
}

func (x *GetByCommentIdRequest) GetCommentId() string {
	if x != nil {
		return x.CommentId
	}
	return ""
}

type GetByCommentIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment *Comment `protobuf:"bytes,1,opt,name=comment,proto3" json:"comment,omitempty"`
}

func (x *GetByCommentIdResponse) Reset() {
	*x = GetByCommentIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByCommentIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByCommentIdResponse) ProtoMessage() {}

func (x *GetByCommentIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByCommentIdResponse.ProtoReflect.Descriptor instead.
func (*GetByCommentIdResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{11}
}

func (x *GetByCommentIdResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type GetCommentsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Limit  string `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Ofset  string `protobuf:"bytes,3,opt,name=ofset,proto3" json:"ofset,omitempty"`
}

func (x *GetCommentsRequest) Reset() {
	*x = GetCommentsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsRequest) ProtoMessage() {}

func (x *GetCommentsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsRequest.ProtoReflect.Descriptor instead.
func (*GetCommentsRequest) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{12}
}

func (x *GetCommentsRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetCommentsRequest) GetLimit() string {
	if x != nil {
		return x.Limit
	}
	return ""
}

func (x *GetCommentsRequest) GetOfset() string {
	if x != nil {
		return x.Ofset
	}
	return ""
}

type GetCommentsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comment []*Comment `protobuf:"bytes,1,rep,name=comment,proto3" json:"comment,omitempty"`
}

func (x *GetCommentsResponse) Reset() {
	*x = GetCommentsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_comments_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCommentsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCommentsResponse) ProtoMessage() {}

func (x *GetCommentsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_comments_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCommentsResponse.ProtoReflect.Descriptor instead.
func (*GetCommentsResponse) Descriptor() ([]byte, []int) {
	return file_comments_proto_rawDescGZIP(), []int{13}
}

func (x *GetCommentsResponse) GetComment() []*Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

var File_comments_proto protoreflect.FileDescriptor

var file_comments_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x08, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x22, 0xb6, 0x01, 0x0a, 0x07, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63,
	0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x6e, 0x0a, 0x0a, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x41, 0x74, 0x22, 0x82, 0x01, 0x0a, 0x11, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d,
	0x6f, 0x72, 0x79, 0x49, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x18,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x14, 0x0a, 0x12, 0x41, 0x64, 0x64, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x33,
	0x0a, 0x14, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x22, 0x49, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49, 0x64, 0x4d, 0x65,
	0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x08,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x6f,
	0x72, 0x79, 0x49, 0x64, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x6c,
	0x0a, 0x14, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x6d, 0x65, 0x6d, 0x6f, 0x72, 0x79,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x6d, 0x6f, 0x72,
	0x79, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x22, 0x17, 0x0a, 0x15,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x26, 0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x17, 0x0a,
	0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x36, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x45,
	0x0a, 0x16, 0x47, 0x65, 0x74, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65,
	0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x59, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x66,
	0x73, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x66, 0x73, 0x65, 0x74,
	0x22, 0x42, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c,
	0x69, 0x6e, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x32, 0xeb, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x47, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1b, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x41, 0x64, 0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x41, 0x64,
	0x64, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x50, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x42, 0x79, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x49,
	0x64, 0x12, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x49, 0x64, 0x4d, 0x65, 0x6d, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x79, 0x49,
	0x64, 0x12, 0x1f, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x74, 0x73, 0x12, 0x1c, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x6c, 0x69, 0x6e, 0x65, 0x2e, 0x47,
	0x65, 0x74, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_comments_proto_rawDescOnce sync.Once
	file_comments_proto_rawDescData = file_comments_proto_rawDesc
)

func file_comments_proto_rawDescGZIP() []byte {
	file_comments_proto_rawDescOnce.Do(func() {
		file_comments_proto_rawDescData = protoimpl.X.CompressGZIP(file_comments_proto_rawDescData)
	})
	return file_comments_proto_rawDescData
}

var file_comments_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_comments_proto_goTypes = []interface{}{
	(*Comment)(nil),                // 0: timeline.Comment
	(*ByMemoryId)(nil),             // 1: timeline.ByMemoryId
	(*AddCommentRequest)(nil),      // 2: timeline.AddCommentRequest
	(*AddCommentResponse)(nil),     // 3: timeline.AddCommentResponse
	(*GetByIdMemoryRequest)(nil),   // 4: timeline.GetByIdMemoryRequest
	(*GetByIdMemoryResponse)(nil),  // 5: timeline.GetByIdMemoryResponse
	(*UpdateCommentRequest)(nil),   // 6: timeline.UpdateCommentRequest
	(*UpdateCommentResponse)(nil),  // 7: timeline.UpdateCommentResponse
	(*DeleteCommentRequest)(nil),   // 8: timeline.DeleteCommentRequest
	(*DeleteCommentResponse)(nil),  // 9: timeline.DeleteCommentResponse
	(*GetByCommentIdRequest)(nil),  // 10: timeline.GetByCommentIdRequest
	(*GetByCommentIdResponse)(nil), // 11: timeline.GetByCommentIdResponse
	(*GetCommentsRequest)(nil),     // 12: timeline.GetCommentsRequest
	(*GetCommentsResponse)(nil),    // 13: timeline.GetCommentsResponse
}
var file_comments_proto_depIdxs = []int32{
	1,  // 0: timeline.GetByIdMemoryResponse.comments:type_name -> timeline.ByMemoryId
	0,  // 1: timeline.GetByCommentIdResponse.comment:type_name -> timeline.Comment
	0,  // 2: timeline.GetCommentsResponse.comment:type_name -> timeline.Comment
	2,  // 3: timeline.CommentService.AddComment:input_type -> timeline.AddCommentRequest
	4,  // 4: timeline.CommentService.GetByMemoryId:input_type -> timeline.GetByIdMemoryRequest
	6,  // 5: timeline.CommentService.UpdateComment:input_type -> timeline.UpdateCommentRequest
	8,  // 6: timeline.CommentService.DeleteComment:input_type -> timeline.DeleteCommentRequest
	10, // 7: timeline.CommentService.GetById:input_type -> timeline.GetByCommentIdRequest
	12, // 8: timeline.CommentService.GetAllCommets:input_type -> timeline.GetCommentsRequest
	3,  // 9: timeline.CommentService.AddComment:output_type -> timeline.AddCommentResponse
	5,  // 10: timeline.CommentService.GetByMemoryId:output_type -> timeline.GetByIdMemoryResponse
	7,  // 11: timeline.CommentService.UpdateComment:output_type -> timeline.UpdateCommentResponse
	9,  // 12: timeline.CommentService.DeleteComment:output_type -> timeline.DeleteCommentResponse
	11, // 13: timeline.CommentService.GetById:output_type -> timeline.GetByCommentIdResponse
	13, // 14: timeline.CommentService.GetAllCommets:output_type -> timeline.GetCommentsResponse
	9,  // [9:15] is the sub-list for method output_type
	3,  // [3:9] is the sub-list for method input_type
	3,  // [3:3] is the sub-list for extension type_name
	3,  // [3:3] is the sub-list for extension extendee
	0,  // [0:3] is the sub-list for field type_name
}

func init() { file_comments_proto_init() }
func file_comments_proto_init() {
	if File_comments_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_comments_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
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
		file_comments_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ByMemoryId); i {
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
		file_comments_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentRequest); i {
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
		file_comments_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddCommentResponse); i {
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
		file_comments_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdMemoryRequest); i {
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
		file_comments_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdMemoryResponse); i {
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
		file_comments_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentRequest); i {
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
		file_comments_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCommentResponse); i {
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
		file_comments_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
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
		file_comments_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentResponse); i {
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
		file_comments_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByCommentIdRequest); i {
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
		file_comments_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByCommentIdResponse); i {
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
		file_comments_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsRequest); i {
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
		file_comments_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCommentsResponse); i {
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
			RawDescriptor: file_comments_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_comments_proto_goTypes,
		DependencyIndexes: file_comments_proto_depIdxs,
		MessageInfos:      file_comments_proto_msgTypes,
	}.Build()
	File_comments_proto = out.File
	file_comments_proto_rawDesc = nil
	file_comments_proto_goTypes = nil
	file_comments_proto_depIdxs = nil
}
