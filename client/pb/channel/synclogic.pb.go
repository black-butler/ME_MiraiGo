// Code generated by protoc-gen-golite. DO NOT EDIT.
// source: pb/channel/synclogic.proto

package channel

import (
	proto "github.com/RomiChan/protobuf/proto"
)

type ChannelMsg struct {
	GuildId     proto.Option[uint64] `protobuf:"varint,1,opt"`
	ChannelId   proto.Option[uint64] `protobuf:"varint,2,opt"`
	Result      proto.Option[uint32] `protobuf:"varint,3,opt"`
	RspBeginSeq proto.Option[uint64] `protobuf:"varint,4,opt"`
	RspEndSeq   proto.Option[uint64] `protobuf:"varint,5,opt"`
	Msgs        []*ChannelMsgContent `protobuf:"bytes,6,rep"`
}

type ChannelMsgReq struct {
	ChannelParam      *ChannelParam        `protobuf:"bytes,1,opt"`
	WithVersionFlag   proto.Option[uint32] `protobuf:"varint,2,opt"`
	DirectMessageFlag proto.Option[uint32] `protobuf:"varint,3,opt"`
	_                 [0]func()
}

type ChannelMsgRsp struct {
	Result          proto.Option[uint32] `protobuf:"varint,1,opt"`
	ErrMsg          []byte               `protobuf:"bytes,2,opt"`
	ChannelMsg      *ChannelMsg          `protobuf:"bytes,3,opt"`
	WithVersionFlag proto.Option[uint32] `protobuf:"varint,4,opt"`
	GetMsgTime      proto.Option[uint64] `protobuf:"varint,5,opt"`
}

type ChannelNode struct {
	ChannelId        proto.Option[uint64] `protobuf:"varint,1,opt"`
	Seq              proto.Option[uint64] `protobuf:"varint,2,opt"`
	CntSeq           proto.Option[uint64] `protobuf:"varint,3,opt"`
	Time             proto.Option[uint64] `protobuf:"varint,4,opt"`
	MemberReadMsgSeq proto.Option[uint64] `protobuf:"varint,5,opt"`
	MemberReadCntSeq proto.Option[uint64] `protobuf:"varint,6,opt"`
	NotifyType       proto.Option[uint32] `protobuf:"varint,7,opt"`
	ChannelName      []byte               `protobuf:"bytes,8,opt"`
	ChannelType      proto.Option[uint32] `protobuf:"varint,9,opt"`
	Meta             []byte               `protobuf:"bytes,10,opt"`
	ReadMsgMeta      []byte               `protobuf:"bytes,11,opt"`
	EventTime        proto.Option[uint32] `protobuf:"varint,12,opt"`
}

type ChannelParam struct {
	GuildId   proto.Option[uint64] `protobuf:"varint,1,opt"`
	ChannelId proto.Option[uint64] `protobuf:"varint,2,opt"`
	BeginSeq  proto.Option[uint64] `protobuf:"varint,3,opt"`
	EndSeq    proto.Option[uint64] `protobuf:"varint,4,opt"`
	Time      proto.Option[uint64] `protobuf:"varint,5,opt"`
	Version   []uint64             `protobuf:"varint,6,rep"`
	Seqs      []*MsgCond           `protobuf:"bytes,7,rep"`
}

type DirectMessageSource struct {
	TinyId     proto.Option[uint64] `protobuf:"varint,1,opt"`
	GuildId    proto.Option[uint64] `protobuf:"varint,2,opt"`
	GuildName  []byte               `protobuf:"bytes,3,opt"`
	MemberName []byte               `protobuf:"bytes,4,opt"`
	NickName   []byte               `protobuf:"bytes,5,opt"`
}

type FirstViewMsg struct {
	PushFlag                proto.Option[uint32] `protobuf:"varint,1,opt"`
	Seq                     proto.Option[uint32] `protobuf:"varint,2,opt"`
	GuildNodes              []*GuildNode         `protobuf:"bytes,3,rep"`
	ChannelMsgs             []*ChannelMsg        `protobuf:"bytes,4,rep"`
	GetMsgTime              proto.Option[uint64] `protobuf:"varint,5,opt"`
	DirectMessageGuildNodes []*GuildNode         `protobuf:"bytes,6,rep"`
}

type FirstViewReq struct {
	LastMsgTime       proto.Option[uint64] `protobuf:"varint,1,opt"`
	UdcFlag           proto.Option[uint32] `protobuf:"varint,2,opt"`
	Seq               proto.Option[uint32] `protobuf:"varint,3,opt"`
	DirectMessageFlag proto.Option[uint32] `protobuf:"varint,4,opt"`
	_                 [0]func()
}

type FirstViewRsp struct {
	Result                  proto.Option[uint32] `protobuf:"varint,1,opt"`
	ErrMsg                  []byte               `protobuf:"bytes,2,opt"`
	Seq                     proto.Option[uint32] `protobuf:"varint,3,opt"`
	UdcFlag                 proto.Option[uint32] `protobuf:"varint,4,opt"`
	GuildCount              proto.Option[uint32] `protobuf:"varint,5,opt"`
	SelfTinyid              proto.Option[uint64] `protobuf:"varint,6,opt"`
	DirectMessageSwitch     proto.Option[uint32] `protobuf:"varint,7,opt"`
	DirectMessageGuildCount proto.Option[uint32] `protobuf:"varint,8,opt"`
}

type GuildNode struct {
	GuildId      proto.Option[uint64] `protobuf:"varint,1,opt"`
	GuildCode    proto.Option[uint64] `protobuf:"varint,2,opt"`
	ChannelNodes []*ChannelNode       `protobuf:"bytes,3,rep"`
	GuildName    []byte               `protobuf:"bytes,4,opt"`
	PeerSource   *DirectMessageSource `protobuf:"bytes,5,opt"`
}

type MsgCond struct {
	Seq          proto.Option[uint64] `protobuf:"varint,1,opt"`
	EventVersion proto.Option[uint64] `protobuf:"varint,2,opt"`
	_            [0]func()
}

type MultiChannelMsg struct {
	PushFlag    proto.Option[uint32] `protobuf:"varint,1,opt"`
	Seq         proto.Option[uint32] `protobuf:"varint,2,opt"`
	ChannelMsgs []*ChannelMsg        `protobuf:"bytes,3,rep"`
	GetMsgTime  proto.Option[uint64] `protobuf:"varint,4,opt"`
}

type MultiChannelMsgReq struct {
	ChannelParams     []*ChannelParam      `protobuf:"bytes,1,rep"`
	Seq               proto.Option[uint32] `protobuf:"varint,2,opt"`
	DirectMessageFlag proto.Option[uint32] `protobuf:"varint,3,opt"`
}

type MultiChannelMsgRsp struct {
	Result proto.Option[uint32] `protobuf:"varint,1,opt"`
	ErrMsg []byte               `protobuf:"bytes,2,opt"`
	Seq    proto.Option[uint32] `protobuf:"varint,3,opt"`
}

type ReqBody struct {
	ChannelParam      *ChannelParam        `protobuf:"bytes,1,opt"`
	DirectMessageFlag proto.Option[uint32] `protobuf:"varint,2,opt"`
	_                 [0]func()
}

type RspBody struct {
	Result     proto.Option[uint32] `protobuf:"varint,1,opt"`
	ErrMsg     []byte               `protobuf:"bytes,2,opt"`
	ChannelMsg *ChannelMsg          `protobuf:"bytes,3,opt"`
}
