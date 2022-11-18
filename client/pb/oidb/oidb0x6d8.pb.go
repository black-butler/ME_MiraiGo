// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/oidb/oidb0x6d8.proto

package oidb

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type D6D8ReqBody struct {
	FileInfoReq       *GetFileInfoReqBody  `protobuf:"bytes,1,opt"`
	FileListInfoReq   *GetFileListReqBody  `protobuf:"bytes,2,opt"`
	GroupFileCountReq *GetFileCountReqBody `protobuf:"bytes,3,opt"`
	GroupSpaceReq     *GetSpaceReqBody     `protobuf:"bytes,4,opt"`
	_                 [0]func()
}

type D6D8RspBody struct {
	FileInfoRsp     *GetFileInfoRspBody  `protobuf:"bytes,1,opt"`
	FileListInfoRsp *GetFileListRspBody  `protobuf:"bytes,2,opt"`
	FileCountRsp    *GetFileCountRspBody `protobuf:"bytes,3,opt"`
	GroupSpaceRsp   *GetSpaceRspBody     `protobuf:"bytes,4,opt"`
	_               [0]func()
}

type GetFileInfoReqBody struct {
	GroupCode proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId     proto.Option[uint32] `protobuf:"varint,2,opt"`
	BusId     proto.Option[uint32] `protobuf:"varint,3,opt"`
	FileId    proto.Option[string] `protobuf:"bytes,4,opt"`
	FieldFlag proto.Option[uint32] `protobuf:"varint,5,opt"`
	_         [0]func()
}

type GetFileInfoRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	FileInfo      *GroupFileInfo       `protobuf:"bytes,4,opt"`
	_             [0]func()
}

type GetFileListRspBody struct {
	RetCode       proto.Option[int32]        `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string]       `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string]       `protobuf:"bytes,3,opt"`
	IsEnd         proto.Option[bool]         `protobuf:"varint,4,opt"`
	ItemList      []*GetFileListRspBody_Item `protobuf:"bytes,5,rep"`
	MaxTimestamp  *FileTimeStamp             `protobuf:"bytes,6,opt"`
	AllFileCount  proto.Option[uint32]       `protobuf:"varint,7,opt"`
	FilterCode    proto.Option[uint32]       `protobuf:"varint,8,opt"`
	SafeCheckFlag proto.Option[bool]         `protobuf:"varint,11,opt"`
	SafeCheckRes  proto.Option[uint32]       `protobuf:"varint,12,opt"`
	NextIndex     proto.Option[uint32]       `protobuf:"varint,13,opt"`
	Context       []byte                     `protobuf:"bytes,14,opt"`
	Role          proto.Option[uint32]       `protobuf:"varint,15,opt"`
	OpenFlag      proto.Option[uint32]       `protobuf:"varint,16,opt"`
}

type GroupFileInfo struct {
	FileId         proto.Option[string] `protobuf:"bytes,1,opt"`
	FileName       proto.Option[string] `protobuf:"bytes,2,opt"`
	FileSize       proto.Option[uint64] `protobuf:"varint,3,opt"`
	BusId          proto.Option[uint32] `protobuf:"varint,4,opt"`
	UploadedSize   proto.Option[uint64] `protobuf:"varint,5,opt"`
	UploadTime     proto.Option[uint32] `protobuf:"varint,6,opt"`
	DeadTime       proto.Option[uint32] `protobuf:"varint,7,opt"`
	ModifyTime     proto.Option[uint32] `protobuf:"varint,8,opt"`
	DownloadTimes  proto.Option[uint32] `protobuf:"varint,9,opt"`
	Sha            []byte               `protobuf:"bytes,10,opt"`
	Sha3           []byte               `protobuf:"bytes,11,opt"`
	Md5            []byte               `protobuf:"bytes,12,opt"`
	LocalPath      proto.Option[string] `protobuf:"bytes,13,opt"`
	UploaderName   proto.Option[string] `protobuf:"bytes,14,opt"`
	UploaderUin    proto.Option[uint64] `protobuf:"varint,15,opt"`
	ParentFolderId proto.Option[string] `protobuf:"bytes,16,opt"`
}

type GroupFolderInfo struct {
	FolderId       proto.Option[string] `protobuf:"bytes,1,opt"`
	ParentFolderId proto.Option[string] `protobuf:"bytes,2,opt"`
	FolderName     proto.Option[string] `protobuf:"bytes,3,opt"`
	CreateTime     proto.Option[uint32] `protobuf:"varint,4,opt"`
	ModifyTime     proto.Option[uint32] `protobuf:"varint,5,opt"`
	CreateUin      proto.Option[uint64] `protobuf:"varint,6,opt"`
	CreatorName    proto.Option[string] `protobuf:"bytes,7,opt"`
	TotalFileCount proto.Option[uint32] `protobuf:"varint,8,opt"`
	_              [0]func()
}

type GetFileListReqBody struct {
	GroupCode      proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId          proto.Option[uint32] `protobuf:"varint,2,opt"`
	FolderId       proto.Option[string] `protobuf:"bytes,3,opt"`
	StartTimestamp *FileTimeStamp       `protobuf:"bytes,4,opt"`
	FileCount      proto.Option[uint32] `protobuf:"varint,5,opt"`
	MaxTimestamp   *FileTimeStamp       `protobuf:"bytes,6,opt"`
	AllFileCount   proto.Option[uint32] `protobuf:"varint,7,opt"`
	ReqFrom        proto.Option[uint32] `protobuf:"varint,8,opt"`
	SortBy         proto.Option[uint32] `protobuf:"varint,9,opt"`
	FilterCode     proto.Option[uint32] `protobuf:"varint,10,opt"`
	Uin            proto.Option[uint64] `protobuf:"varint,11,opt"`
	FieldFlag      proto.Option[uint32] `protobuf:"varint,12,opt"`
	StartIndex     proto.Option[uint32] `protobuf:"varint,13,opt"`
	Context        []byte               `protobuf:"bytes,14,opt"`
	ClientVersion  proto.Option[uint32] `protobuf:"varint,15,opt"`
}

type GetFileCountReqBody struct {
	GroupCode proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId     proto.Option[uint32] `protobuf:"varint,2,opt"`
	BusId     proto.Option[uint32] `protobuf:"varint,3,opt"`
	_         [0]func()
}

type GetSpaceReqBody struct {
	GroupCode proto.Option[uint64] `protobuf:"varint,1,opt"`
	AppId     proto.Option[uint32] `protobuf:"varint,2,opt"`
	_         [0]func()
}

type GetFileCountRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	AllFileCount  proto.Option[uint32] `protobuf:"varint,4,opt"`
	FileTooMany   proto.Option[bool]   `protobuf:"varint,5,opt"`
	LimitCount    proto.Option[uint32] `protobuf:"varint,6,opt"`
	IsFull        proto.Option[bool]   `protobuf:"varint,7,opt"`
	_             [0]func()
}

type GetSpaceRspBody struct {
	RetCode       proto.Option[int32]  `protobuf:"varint,1,opt"`
	RetMsg        proto.Option[string] `protobuf:"bytes,2,opt"`
	ClientWording proto.Option[string] `protobuf:"bytes,3,opt"`
	TotalSpace    proto.Option[uint64] `protobuf:"varint,4,opt"`
	UsedSpace     proto.Option[uint64] `protobuf:"varint,5,opt"`
	_             [0]func()
}

type FileTimeStamp struct {
	UploadTime proto.Option[uint32] `protobuf:"varint,1,opt"`
	FileId     proto.Option[string] `protobuf:"bytes,2,opt"`
	_          [0]func()
}

type GetFileListRspBody_Item struct {
	Type       proto.Option[uint32] `protobuf:"varint,1,opt"`
	FolderInfo *GroupFolderInfo     `protobuf:"bytes,2,opt"`
	FileInfo   *GroupFileInfo       `protobuf:"bytes,3,opt"`
	_          [0]func()
}
