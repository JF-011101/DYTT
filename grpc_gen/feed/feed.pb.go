// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: idl/feed.proto

package feed

import (
	context "context"
	user "github.com/jf-011101/dytt/grpc_gen/user"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DouyinFeedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LatestTime *int64  `protobuf:"varint,1,opt,name=latest_time,json=latestTime,proto3,oneof" json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      *string `protobuf:"bytes,2,opt,name=token,proto3,oneof" json:"token,omitempty"`                              // 可选参数，登录用户设置
}

func (x *DouyinFeedRequest) Reset() {
	*x = DouyinFeedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_feed_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFeedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFeedRequest) ProtoMessage() {}

func (x *DouyinFeedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFeedRequest.ProtoReflect.Descriptor instead.
func (*DouyinFeedRequest) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{0}
}

func (x *DouyinFeedRequest) GetLatestTime() int64 {
	if x != nil && x.LatestTime != nil {
		return *x.LatestTime
	}
	return 0
}

func (x *DouyinFeedRequest) GetToken() string {
	if x != nil && x.Token != nil {
		return *x.Token
	}
	return ""
}

type DouyinFeedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32    `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg  *string  `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3,oneof" json:"status_msg,omitempty"` // 返回状态描述
	VideoList  []*Video `protobuf:"bytes,3,rep,name=video_list,json=videoList,proto3" json:"video_list,omitempty"`       // 视频列表
	NextTime   *int64   `protobuf:"varint,4,opt,name=next_time,json=nextTime,proto3,oneof" json:"next_time,omitempty"`   // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}

func (x *DouyinFeedResponse) Reset() {
	*x = DouyinFeedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_feed_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DouyinFeedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DouyinFeedResponse) ProtoMessage() {}

func (x *DouyinFeedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DouyinFeedResponse.ProtoReflect.Descriptor instead.
func (*DouyinFeedResponse) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{1}
}

func (x *DouyinFeedResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *DouyinFeedResponse) GetStatusMsg() string {
	if x != nil && x.StatusMsg != nil {
		return *x.StatusMsg
	}
	return ""
}

func (x *DouyinFeedResponse) GetVideoList() []*Video {
	if x != nil {
		return x.VideoList
	}
	return nil
}

func (x *DouyinFeedResponse) GetNextTime() int64 {
	if x != nil && x.NextTime != nil {
		return *x.NextTime
	}
	return 0
}

type VideoIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	VideoId  int64 `protobuf:"varint,1,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`
	SearchId int64 `protobuf:"varint,2,opt,name=search_id,json=searchId,proto3" json:"search_id,omitempty"`
}

func (x *VideoIdRequest) Reset() {
	*x = VideoIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_feed_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VideoIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VideoIdRequest) ProtoMessage() {}

func (x *VideoIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VideoIdRequest.ProtoReflect.Descriptor instead.
func (*VideoIdRequest) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{2}
}

func (x *VideoIdRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *VideoIdRequest) GetSearchId() int64 {
	if x != nil {
		return x.SearchId
	}
	return 0
}

type Video struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64      `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                            // 视频唯一标识
	Author        *user.User `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`                                     // 视频作者信息
	PlayUrl       string     `protobuf:"bytes,3,opt,name=play_url,json=playUrl,proto3" json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string     `protobuf:"bytes,4,opt,name=cover_url,json=coverUrl,proto3" json:"cover_url,omitempty"`                 // 视频封面地址
	FavoriteCount int64      `protobuf:"varint,5,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"` // 视频的点赞总数
	CommentCount  int64      `protobuf:"varint,6,opt,name=comment_count,json=commentCount,proto3" json:"comment_count,omitempty"`    // 视频的评论总数
	IsFavorite    bool       `protobuf:"varint,7,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`          // true-已点赞，false-未点赞
	Title         string     `protobuf:"bytes,8,opt,name=title,proto3" json:"title,omitempty"`                                       // 视频标题
}

func (x *Video) Reset() {
	*x = Video{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_feed_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Video) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Video) ProtoMessage() {}

func (x *Video) ProtoReflect() protoreflect.Message {
	mi := &file_idl_feed_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Video.ProtoReflect.Descriptor instead.
func (*Video) Descriptor() ([]byte, []int) {
	return file_idl_feed_proto_rawDescGZIP(), []int{3}
}

func (x *Video) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Video) GetAuthor() *user.User {
	if x != nil {
		return x.Author
	}
	return nil
}

func (x *Video) GetPlayUrl() string {
	if x != nil {
		return x.PlayUrl
	}
	return ""
}

func (x *Video) GetCoverUrl() string {
	if x != nil {
		return x.CoverUrl
	}
	return ""
}

func (x *Video) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

func (x *Video) GetCommentCount() int64 {
	if x != nil {
		return x.CommentCount
	}
	return 0
}

func (x *Video) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

func (x *Video) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

var File_idl_feed_proto protoreflect.FileDescriptor

var file_idl_feed_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x69, 0x64, 0x6c, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x66, 0x65, 0x65, 0x64, 0x1a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x6a, 0x66, 0x2d, 0x30, 0x31, 0x31, 0x31, 0x30, 0x31, 0x2f, 0x44, 0x59, 0x54,
	0x54, 0x2f, 0x69, 0x64, 0x6c, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x70, 0x0a, 0x13, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f,
	0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x24, 0x0a, 0x0b, 0x6c, 0x61, 0x74, 0x65, 0x73,
	0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0a,
	0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x19, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x01, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x88, 0x01, 0x01, 0x42, 0x0e, 0x0a, 0x0c, 0x5f, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0xc6, 0x01, 0x0a, 0x14, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65,
	0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x22, 0x0a, 0x0a,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x48, 0x00, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x88, 0x01, 0x01,
	0x12, 0x2a, 0x0a, 0x0a, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x52, 0x09, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x09,
	0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x01, 0x52, 0x08, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x69, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x0d,
	0x0a, 0x0b, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x42, 0x0c, 0x0a,
	0x0a, 0x5f, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x22, 0x4a, 0x0a, 0x10, 0x76,
	0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x65,
	0x61, 0x72, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x73,
	0x65, 0x61, 0x72, 0x63, 0x68, 0x49, 0x64, 0x22, 0xf6, 0x01, 0x0a, 0x05, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x22, 0x0a, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0a, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x06, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x5f, 0x75, 0x72,
	0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x61, 0x79, 0x55, 0x72, 0x6c,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x55, 0x72, 0x6c, 0x12, 0x25, 0x0a,
	0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f,
	0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a,
	0x69, 0x73, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65,
	0x32, 0x88, 0x01, 0x0a, 0x07, 0x46, 0x65, 0x65, 0x64, 0x53, 0x72, 0x76, 0x12, 0x46, 0x0a, 0x0b,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x46, 0x65, 0x65, 0x64, 0x12, 0x19, 0x2e, 0x66, 0x65,
	0x65, 0x64, 0x2e, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x64, 0x6f,
	0x75, 0x79, 0x69, 0x6e, 0x5f, 0x66, 0x65, 0x65, 0x64, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x56, 0x69, 0x64, 0x65, 0x6f,
	0x42, 0x79, 0x49, 0x64, 0x12, 0x16, 0x2e, 0x66, 0x65, 0x65, 0x64, 0x2e, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x5f, 0x69, 0x64, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0b, 0x2e, 0x66,
	0x65, 0x65, 0x64, 0x2e, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x22, 0x00, 0x42, 0x29, 0x5a, 0x27, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6a, 0x66, 0x2d, 0x30, 0x31, 0x31,
	0x31, 0x30, 0x31, 0x2f, 0x44, 0x59, 0x54, 0x54, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x67, 0x65,
	0x6e, 0x2f, 0x66, 0x65, 0x65, 0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_feed_proto_rawDescOnce sync.Once
	file_idl_feed_proto_rawDescData = file_idl_feed_proto_rawDesc
)

func file_idl_feed_proto_rawDescGZIP() []byte {
	file_idl_feed_proto_rawDescOnce.Do(func() {
		file_idl_feed_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_feed_proto_rawDescData)
	})
	return file_idl_feed_proto_rawDescData
}

var file_idl_feed_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_idl_feed_proto_goTypes = []interface{}{
	(*DouyinFeedRequest)(nil),  // 0: feed.douyin_feed_request
	(*DouyinFeedResponse)(nil), // 1: feed.douyin_feed_response
	(*VideoIdRequest)(nil),     // 2: feed.video_id_request
	(*Video)(nil),              // 3: feed.Video
	(*user.User)(nil),          // 4: user.User
}
var file_idl_feed_proto_depIdxs = []int32{
	3, // 0: feed.douyin_feed_response.video_list:type_name -> feed.Video
	4, // 1: feed.Video.author:type_name -> user.User
	0, // 2: feed.FeedSrv.GetUserFeed:input_type -> feed.douyin_feed_request
	2, // 3: feed.FeedSrv.GetVideoById:input_type -> feed.video_id_request
	1, // 4: feed.FeedSrv.GetUserFeed:output_type -> feed.douyin_feed_response
	3, // 5: feed.FeedSrv.GetVideoById:output_type -> feed.Video
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_idl_feed_proto_init() }
func file_idl_feed_proto_init() {
	if File_idl_feed_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_feed_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFeedRequest); i {
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
		file_idl_feed_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DouyinFeedResponse); i {
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
		file_idl_feed_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VideoIdRequest); i {
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
		file_idl_feed_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Video); i {
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
	file_idl_feed_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_idl_feed_proto_msgTypes[1].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_feed_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_feed_proto_goTypes,
		DependencyIndexes: file_idl_feed_proto_depIdxs,
		MessageInfos:      file_idl_feed_proto_msgTypes,
	}.Build()
	File_idl_feed_proto = out.File
	file_idl_feed_proto_rawDesc = nil
	file_idl_feed_proto_goTypes = nil
	file_idl_feed_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// FeedSrvClient is the client API for FeedSrv service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type FeedSrvClient interface {
	GetUserFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error)
	GetVideoById(ctx context.Context, in *VideoIdRequest, opts ...grpc.CallOption) (*Video, error)
}

type feedSrvClient struct {
	cc grpc.ClientConnInterface
}

func NewFeedSrvClient(cc grpc.ClientConnInterface) FeedSrvClient {
	return &feedSrvClient{cc}
}

func (c *feedSrvClient) GetUserFeed(ctx context.Context, in *DouyinFeedRequest, opts ...grpc.CallOption) (*DouyinFeedResponse, error) {
	out := new(DouyinFeedResponse)
	err := c.cc.Invoke(ctx, "/feed.FeedSrv/GetUserFeed", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *feedSrvClient) GetVideoById(ctx context.Context, in *VideoIdRequest, opts ...grpc.CallOption) (*Video, error) {
	out := new(Video)
	err := c.cc.Invoke(ctx, "/feed.FeedSrv/GetVideoById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FeedSrvServer is the server API for FeedSrv service.
type FeedSrvServer interface {
	GetUserFeed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error)
	GetVideoById(context.Context, *VideoIdRequest) (*Video, error)
}

// UnimplementedFeedSrvServer can be embedded to have forward compatible implementations.
type UnimplementedFeedSrvServer struct {
}

func (*UnimplementedFeedSrvServer) GetUserFeed(context.Context, *DouyinFeedRequest) (*DouyinFeedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFeed not implemented")
}
func (*UnimplementedFeedSrvServer) GetVideoById(context.Context, *VideoIdRequest) (*Video, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetVideoById not implemented")
}

func RegisterFeedSrvServer(s *grpc.Server, srv FeedSrvServer) {
	s.RegisterService(&_FeedSrv_serviceDesc, srv)
}

func _FeedSrv_GetUserFeed_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DouyinFeedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedSrvServer).GetUserFeed(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.FeedSrv/GetUserFeed",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedSrvServer).GetUserFeed(ctx, req.(*DouyinFeedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FeedSrv_GetVideoById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VideoIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FeedSrvServer).GetVideoById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/feed.FeedSrv/GetVideoById",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FeedSrvServer).GetVideoById(ctx, req.(*VideoIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _FeedSrv_serviceDesc = grpc.ServiceDesc{
	ServiceName: "feed.FeedSrv",
	HandlerType: (*FeedSrvServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUserFeed",
			Handler:    _FeedSrv_GetUserFeed_Handler,
		},
		{
			MethodName: "GetVideoById",
			Handler:    _FeedSrv_GetVideoById_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "idl/feed.proto",
}
