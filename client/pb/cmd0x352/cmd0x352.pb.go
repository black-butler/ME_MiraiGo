// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/cmd0x352/cmd0x352.proto

package cmd0x352

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type ReqBody struct {
	Subcmd      proto.Option[uint32] `protobuf:"varint,1,opt"`
	TryupImgReq []*D352TryUpImgReq   `protobuf:"bytes,2,rep"`
	// repeated GetImgUrlReq getimgUrlReq = 3;
	// repeated DelImgReq delImgReq = 4;
	NetType proto.Option[uint32] `protobuf:"varint,10,opt"`
}

type RspBody struct {
	Subcmd      proto.Option[uint32] `protobuf:"varint,1,opt"`
	TryupImgRsp []*TryUpImgRsp       `protobuf:"bytes,2,rep"`
	// repeated GetImgUrlRsp getimgUrlRsp = 3;
	NewBigchan proto.Option[bool] `protobuf:"varint,4,opt"`
	// repeated DelImgRsp delImgRsp = 5;
	FailMsg []byte `protobuf:"bytes,10,opt"`
}

type D352TryUpImgReq struct {
	SrcUin        proto.Option[uint64] `protobuf:"varint,1,opt"`
	DstUin        proto.Option[uint64] `protobuf:"varint,2,opt"`
	FileId        proto.Option[uint64] `protobuf:"varint,3,opt"`
	FileMd5       []byte               `protobuf:"bytes,4,opt"`
	FileSize      proto.Option[uint64] `protobuf:"varint,5,opt"`
	FileName      []byte               `protobuf:"bytes,6,opt"`
	SrcTerm       proto.Option[uint32] `protobuf:"varint,7,opt"`
	PlatformType  proto.Option[uint32] `protobuf:"varint,8,opt"`
	InnerIp       proto.Option[uint32] `protobuf:"varint,9,opt"`
	AddressBook   proto.Option[bool]   `protobuf:"varint,10,opt"`
	Retry         proto.Option[uint32] `protobuf:"varint,11,opt"`
	BuType        proto.Option[uint32] `protobuf:"varint,12,opt"`
	PicOriginal   proto.Option[bool]   `protobuf:"varint,13,opt"`
	PicWidth      proto.Option[uint32] `protobuf:"varint,14,opt"`
	PicHeight     proto.Option[uint32] `protobuf:"varint,15,opt"`
	PicType       proto.Option[uint32] `protobuf:"varint,16,opt"`
	BuildVer      []byte               `protobuf:"bytes,17,opt"`
	FileIndex     []byte               `protobuf:"bytes,18,opt"`
	StoreDays     proto.Option[uint32] `protobuf:"varint,19,opt"`
	TryupStepflag proto.Option[uint32] `protobuf:"varint,20,opt"`
	RejectTryfast proto.Option[bool]   `protobuf:"varint,21,opt"`
	SrvUpload     proto.Option[uint32] `protobuf:"varint,22,opt"`
	TransferUrl   []byte               `protobuf:"bytes,23,opt"`
}

type TryUpImgRsp struct {
	FileId   proto.Option[uint64] `protobuf:"varint,1,opt"`
	ClientIp proto.Option[uint32] `protobuf:"varint,2,opt"`
	Result   proto.Option[uint32] `protobuf:"varint,3,opt"`
	FailMsg  []byte               `protobuf:"bytes,4,opt"`
	FileExit proto.Option[bool]   `protobuf:"varint,5,opt"`
	// optional ImgInfo imgInfo = 6;
	UpIp         []uint32             `protobuf:"varint,7,rep"`
	UpPort       []uint32             `protobuf:"varint,8,rep"`
	UpUkey       []byte               `protobuf:"bytes,9,opt"`
	UpResid      []byte               `protobuf:"bytes,10,opt"`
	UpUuid       []byte               `protobuf:"bytes,11,opt"`
	UpOffset     proto.Option[uint64] `protobuf:"varint,12,opt"`
	BlockSize    proto.Option[uint64] `protobuf:"varint,13,opt"`
	EncryptDstip []byte               `protobuf:"bytes,14,opt"`
	Roamdays     proto.Option[uint32] `protobuf:"varint,15,opt"`
	// repeated IPv6Info upIp6 = 26;
	ClientIp6        []byte               `protobuf:"bytes,27,opt"`
	ThumbDownPara    []byte               `protobuf:"bytes,60,opt"`
	OriginalDownPara []byte               `protobuf:"bytes,61,opt"`
	DownDomain       []byte               `protobuf:"bytes,62,opt"`
	BigDownPara      []byte               `protobuf:"bytes,64,opt"`
	BigThumbDownPara []byte               `protobuf:"bytes,65,opt"`
	HttpsUrlFlag     proto.Option[uint32] `protobuf:"varint,66,opt"` // optional TryUpInfo4Busi info4Busi = 1001;
}
