// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/exciting/group.proto

package exciting

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type FileUploadExt struct {
	Unknown1   proto.Option[int32] `protobuf:"varint,1,opt"`
	Unknown2   proto.Option[int32] `protobuf:"varint,2,opt"`
	Unknown3   proto.Option[int32] `protobuf:"varint,3,opt"`
	Entry      *FileUploadEntry    `protobuf:"bytes,100,opt"`
	Unknown200 proto.Option[int32] `protobuf:"varint,200,opt"`
	_          [0]func()
}

type FileUploadEntry struct {
	BusiBuff     *ExcitingBusiInfo     `protobuf:"bytes,100,opt"`
	FileEntry    *ExcitingFileEntry    `protobuf:"bytes,200,opt"`
	ClientInfo   *ExcitingClientInfo   `protobuf:"bytes,300,opt"`
	FileNameInfo *ExcitingFileNameInfo `protobuf:"bytes,400,opt"`
	Host         *ExcitingHostConfig   `protobuf:"bytes,500,opt"`
	_            [0]func()
}

type ExcitingBusiInfo struct {
	BusId       proto.Option[int32] `protobuf:"varint,1,opt"`
	SenderUin   proto.Option[int64] `protobuf:"varint,100,opt"`
	ReceiverUin proto.Option[int64] `protobuf:"varint,200,opt"` // probable
	GroupCode   proto.Option[int64] `protobuf:"varint,400,opt"` // probable
	_           [0]func()
}

type ExcitingFileEntry struct {
	FileSize  proto.Option[int64] `protobuf:"varint,100,opt"`
	Md5       []byte              `protobuf:"bytes,200,opt"`
	Sha1      []byte              `protobuf:"bytes,300,opt"`
	FileId    []byte              `protobuf:"bytes,600,opt"`
	UploadKey []byte              `protobuf:"bytes,700,opt"`
}

type ExcitingClientInfo struct {
	ClientType   proto.Option[int32]  `protobuf:"varint,100,opt"` // probable
	AppId        proto.Option[string] `protobuf:"bytes,200,opt"`
	TerminalType proto.Option[int32]  `protobuf:"varint,300,opt"` // probable
	ClientVer    proto.Option[string] `protobuf:"bytes,400,opt"`
	Unknown      proto.Option[int32]  `protobuf:"varint,600,opt"`
	_            [0]func()
}

type ExcitingFileNameInfo struct {
	FileName proto.Option[string] `protobuf:"bytes,100,opt"`
	_        [0]func()
}

type ExcitingHostConfig struct {
	Hosts []*ExcitingHostInfo `protobuf:"bytes,200,rep"`
}

type ExcitingHostInfo struct {
	Url  *ExcitingUrlInfo    `protobuf:"bytes,1,opt"`
	Port proto.Option[int32] `protobuf:"varint,2,opt"`
	_    [0]func()
}

type ExcitingUrlInfo struct {
	Unknown proto.Option[int32]  `protobuf:"varint,1,opt"` // not https?
	Host    proto.Option[string] `protobuf:"bytes,2,opt"`
	_       [0]func()
}
