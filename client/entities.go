package client

import (
	"github.com/pkg/errors"

	"github.com/black-butler/ME_MiraiGo/binary/jce"
	"github.com/black-butler/ME_MiraiGo/client/internal/auth"
	"github.com/black-butler/ME_MiraiGo/message"
)

var (
	ErrAlreadyOnline  = errors.New("already online")
	ErrMemberNotFound = errors.New("member not found")
	ErrNotExists      = errors.New("not exists")
)

type (
	LoginError int

	QRCodeLoginState int

	MemberPermission int

	UserOnlineStatus int

	ClientProtocol = auth.Protocol

	LoginResponse struct {
		Success bool
		Error   LoginError

		// Captcha info
		CaptchaImage []byte
		CaptchaSign  []byte

		// Unsafe device
		VerifyUrl string

		// SMS needed
		SMSPhone string

		// other error
		ErrorMessage string
	}

	QRCodeLoginResponse struct {
		State QRCodeLoginState

		ImageData []byte
		Sig       []byte

		LoginInfo *QRCodeLoginInfo
	}

	QRCodeLoginInfo struct {
		TmpPwd      []byte
		TmpNoPicSig []byte
		TgtQR       []byte
	}

	FriendInfo struct {
		Uin      int64
		Nickname string
		Remark   string
		FaceId   int16
		// msgSeqList *utils.Cache
	}

	SummaryCardInfo struct {
		Uin       int64
		Sex       byte
		Age       uint8
		Nickname  string
		Level     int32
		City      string
		Sign      string
		Mobile    string
		LoginDays int64
		Qid       string
		VipLevel  string
	}

	OtherClientInfo struct {
		AppId      int64
		DeviceName string
		DeviceKind string
	}

	FriendListResponse struct {
		TotalCount int32
		List       []*FriendInfo
	}

	OtherClientStatusChangedEvent struct {
		Client *OtherClientInfo
		Online bool
	}

	GroupMuteEvent struct {
		GroupCode   int64
		OperatorUin int64
		TargetUin   int64
		Time        int32
	}

	GroupMessageRecalledEvent struct {
		GroupCode   int64
		OperatorUin int64
		AuthorUin   int64
		MessageId   int32
		Time        int32
	}

	FriendMessageRecalledEvent struct {
		FriendUin int64
		MessageId int32
		Time      int64
	}

	TempMessageEvent struct {
		Message *message.TempMessage
		Session *TempSessionInfo
	}

	GroupLeaveEvent struct {
		Group    *GroupInfo
		Operator *GroupMemberInfo
	}

	MemberJoinGroupEvent struct {
		Group  *GroupInfo
		Member *GroupMemberInfo
	}

	GroupNameUpdatedEvent struct {
		Group       *GroupInfo
		OldName     string
		NewName     string
		OperatorUin int64
	}

	MemberCardUpdatedEvent struct {
		Group   *GroupInfo
		OldCard string
		Member  *GroupMemberInfo
	}

	INotifyEvent interface {
		From() int64
		Content() string
	}

	MemberLeaveGroupEvent struct {
		Group    *GroupInfo
		Member   *GroupMemberInfo
		Operator *GroupMemberInfo
	}

	MemberPermissionChangedEvent struct {
		Group         *GroupInfo
		Member        *GroupMemberInfo
		OldPermission MemberPermission
		NewPermission MemberPermission
	}

	ClientDisconnectedEvent struct {
		Message string
	}

	NewFriendRequest struct {
		RequestId     int64
		Message       string
		RequesterUin  int64
		RequesterNick string

		client *QQClient
	}

	ServerUpdatedEvent struct {
		Servers []jce.SsoServerInfo
	}

	NewFriendEvent struct {
		Friend *FriendInfo
	}

	OfflineFileEvent struct {
		FileName    string
		FileSize    int64
		Sender      int64
		DownloadUrl string
	}

	// GroupDigest ???????????????
	GroupDigest struct {
		GroupCode         int64  `json:"group_code,string"`
		MessageID         uint32 `json:"msg_seq"`
		InternalMessageID uint32 `json:"msg_random"`
		SenderUin         int64  `json:"sender_uin,string"`
		SenderNick        string `json:"sender_nick"`
		SenderTime        int64  `json:"sender_time"`
		AddDigestUin      int64  `json:"add_digest_uin,string"`
		AddDigestNick     string `json:"add_digest_nick"`
		AddDigestTime     int64  `json:"add_digest_time"`
	}

	// GroupDigestEvent ??????????????? ?????????tx????????????????????????
	GroupDigestEvent struct {
		GroupCode         int64
		MessageID         int32
		InternalMessageID int32
		OperationType     int32 // 1 -> ??????????????????, 2 -> ??????????????????
		OperateTime       uint32
		SenderUin         int64
		OperatorUin       int64
		SenderNick        string
		OperatorNick      string
	}

	GuildMessageReactionsUpdatedEvent struct {
		OperatorId uint64 // OperatorId ?????????TinyId, ???????????????????????????????????????
		EmojiId    int32  // EmojiId ???????????????, ??????????????????????????????????????????
		GuildId    uint64
		ChannelId  uint64
		MessageId  uint64
		// MessageSenderUin int64 // MessageSenderUin ??????????????????????????????QQ???
		CurrentReactions []*message.GuildMessageEmojiReaction
	}

	GuildChannelUpdatedEvent struct {
		OperatorId     uint64
		GuildId        uint64
		ChannelId      uint64
		OldChannelInfo *ChannelInfo
		NewChannelInfo *ChannelInfo
	}

	GuildMessageRecalledEvent struct {
		OperatorId uint64
		GuildId    uint64
		ChannelId  uint64
		MessageId  uint64
		RecallTime int64
	}

	GuildChannelOperationEvent struct {
		OperatorId  uint64
		GuildId     uint64
		ChannelInfo *ChannelInfo
	}

	MemberJoinGuildEvent struct {
		Guild  *GuildInfo
		Member *GuildMemberInfo
	}

	OcrResponse struct {
		Texts    []*TextDetection `json:"texts"`
		Language string           `json:"language"`
	}

	TextDetection struct {
		Text        string        `json:"text"`
		Confidence  int32         `json:"confidence"`
		Coordinates []*Coordinate `json:"coordinates"`
	}

	Coordinate struct {
		X int32 `json:"x"`
		Y int32 `json:"y"`
	}

	AtAllRemainInfo struct {
		CanAtAll                 bool   `json:"can_at_all"`
		RemainAtAllCountForGroup uint32 `json:"remain_at_all_count_for_group"`
		RemainAtAllCountForUin   uint32 `json:"remain_at_all_count_for_uin"`
	}

	groupMemberListResponse struct {
		NextUin int64
		list    []*GroupMemberInfo
	}

	groupMessageReceiptEvent struct {
		Rand int32
		Seq  int32
		Msg  *message.GroupMessage
	}

	bigDataSessionInfo struct {
		SigSession []byte
		SessionKey []byte
	}

	// unit is an alias for struct{}, like `()` in rust
	unit = struct{}
)

//go:generate stringer -type=LoginError
const (
	NeedCaptcha            LoginError = 1
	OtherLoginError        LoginError = 3
	UnsafeDeviceError      LoginError = 4
	SMSNeededError         LoginError = 5
	TooManySMSRequestError LoginError = 6
	SMSOrVerifyNeededError LoginError = 7
	SliderNeededError      LoginError = 8
	UnknownLoginError      LoginError = -1

	QRCodeImageFetch        QRCodeLoginState = 1
	QRCodeWaitingForScan    QRCodeLoginState = 2
	QRCodeWaitingForConfirm QRCodeLoginState = 3
	QRCodeTimeout           QRCodeLoginState = 4
	QRCodeConfirmed         QRCodeLoginState = 5
	QRCodeCanceled          QRCodeLoginState = 6

	StatusOnline        UserOnlineStatus = 11   // ??????
	StatusOffline       UserOnlineStatus = 21   // ??????
	StatusAway          UserOnlineStatus = 31   // ??????
	StatusInvisible     UserOnlineStatus = 41   // ??????
	StatusBusy          UserOnlineStatus = 50   // ???
	StatusBattery       UserOnlineStatus = 1000 // ????????????
	StatusListening     UserOnlineStatus = 1028 // ?????????
	StatusConstellation UserOnlineStatus = 1040 // ????????????
	StatusWeather       UserOnlineStatus = 1030 // ????????????
	StatusMeetSpring    UserOnlineStatus = 1069 // ????????????
	StatusTimi          UserOnlineStatus = 1027 // Timi???
	StatusEatChicken    UserOnlineStatus = 1064 // ?????????
	StatusLoving        UserOnlineStatus = 1051 // ?????????
	StatusWangWang      UserOnlineStatus = 1053 // ?????????
	StatusCookedRice    UserOnlineStatus = 1019 // ?????????
	StatusStudy         UserOnlineStatus = 1018 // ?????????
	StatusStayUp        UserOnlineStatus = 1032 // ?????????
	StatusPlayBall      UserOnlineStatus = 1050 // ?????????
	StatusSignal        UserOnlineStatus = 1011 // ?????????
	StatusStudyOnline   UserOnlineStatus = 1024 // ????????????
	StatusGaming        UserOnlineStatus = 1017 // ?????????
	StatusVacationing   UserOnlineStatus = 1022 // ?????????
	StatusWatchingTV    UserOnlineStatus = 1021 // ?????????
	StatusFitness       UserOnlineStatus = 1020 // ?????????

	Owner         MemberPermission = 1
	Administrator MemberPermission = 2
	Member        MemberPermission = 3

	Unset        = auth.Unset
	AndroidPhone = auth.AndroidPhone
	AndroidWatch = auth.AndroidWatch
	MacOS        = auth.MacOS
	QiDian       = auth.QiDian
	IPad         = auth.IPad
)

func (r *UserJoinGroupRequest) Accept() {
	r.client.SolveGroupJoinRequest(r, true, false, "")
}

func (r *UserJoinGroupRequest) Reject(block bool, reason string) {
	r.client.SolveGroupJoinRequest(r, false, block, reason)
}

func (r *GroupInvitedRequest) Accept() {
	r.client.SolveGroupJoinRequest(r, true, false, "")
}

func (r *GroupInvitedRequest) Reject(block bool, reason string) {
	r.client.SolveGroupJoinRequest(r, false, block, reason)
}

func (r *NewFriendRequest) Accept() {
	r.client.SolveFriendRequest(r, true)
}

func (r *NewFriendRequest) Reject() {
	r.client.SolveFriendRequest(r, false)
}
