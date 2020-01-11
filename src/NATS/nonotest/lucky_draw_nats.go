package main

import (
	"encoding/json"
	//"github.com/nats-io/nats"
	"nonolive/model/interaction"
	"nonolive/nonoutils/nononats"
	//"runtime"
	"fmt"
)

const (
	NatCmdNotifyHostLuckyDrawJoinUser = "notifyHostLuckyDrawJoinUser"
	NatCmdNotifyLuckyDrawLuckyUsers   = "notifyLuckyDrawLuckyUser"

	NatCmdSendMessageToUserIds  = "sendMessageToUserIds"
	NatCmdBroadcastMessage      = "broadcastMessage"
	NatCmdWriteRoomCtxLuckyDraw = "writeRoomCtxLuckyDraw"
	NatCmdReleaseRoomContext    = "releaseRoomContext"
)

type WsCommonMessage struct {
	WsCommonCommand
	MsgData interface{} `json:"msg_data,omitempty"`
	MsgType int         `json:"type"`
	Rst     int         `json:"rst"`
	ErrMsg  string      `json:"err_msg,omitempty"`
}

type WsCommonCommand struct {
	MessageType  int           `json:"-"` //表示传过来的是binary message 还是 text message
	ProtocolInfo *ProtocolInfo `json:"-"` //协议版本
	Cmd          string        `json:"cmd"`
	ReqId        int           `json:"req_id,omitempty"`
}

type ProtocolInfo struct {
	ProtocolVersion uint16 `json:"-"` //协议版本
	CompressSupport uint8  `json:"-"` //暂时没用到
	CompressUse     uint8  `json:"-"` //暂时没用到
	BodyLength      uint32 `json:"-"` //body的长度
}

type WsRoomCanRelease interface {
	//返回是否能释放房间,如果不能释放,原因是什么
	CanRelease() (bool, string)
	//释放房间资源
	Release()
}

type InteractiveLive struct {
	MemberNumber int `json:"-"`
	//Room         *wsm.WsRoom        `json:"-"`
	ch        chan int `json:"-"`
	isRunning bool     `json:"-"`
	//LinkClients  []*LiveClient      `json:"link_guests"`
	//WaitingQueue []*LiveWaitingUser `json:"waiting_queue"`
}

type WsRoomContext struct {
	canRelease      bool                   //能否释放room了
	interactiveLive *InteractiveLive       //多嘉宾直播
	luckyDraw       *interaction.LuckyDraw //房间的抽奖信息
}

func (p *WsRoomContext) CanRelease() (bool, string) {
	return p.canRelease, ""
}

//释放房间资源，需要释放的资源在此释放
func (p *WsRoomContext) Release() {
	//todo 当前是释放所有房间资源，因为目前资源不冲突，如果以后有变动，需要按需释放
	//释放多嘉宾广播
	if p.interactiveLive != nil {
		//p.interactiveLive.Close()
		p.interactiveLive = nil
	}
	//释放抽奖
	if p.luckyDraw != nil {
		p.luckyDraw = nil
	}
}

type _RoomSendMessage struct {
	RoomId  string           `json:"room_id"`
	UserIds []int            `json:"user_ids"`
	Msg     *WsCommonMessage `json:"msg"`
}

type Result struct {
	interaction.LuckyDraw `json:",inline"`
	IsWinner              int `json:"is_winner"`
}

type NatCommandMsg struct {
	Cmd             string           `json:"cmd"`
	Msg             []byte           `json:"msg"`
	UserIds         []int            `json:"user_ids"`
	WsCommonMessage *WsCommonMessage `json:"ws_common_message"`
	//RoomCtx         WsRoomCanRelease `json:"ws_room_context"`
	LuckyDraw       *interaction.LuckyDraw `json:"lucky_draw"`
	InteractiveLive *InteractiveLive
}

func main() {
	luckyDraw := &interaction.LuckyDraw{
		HostId: 1111,
	}
	lj, _ := json.Marshal(luckyDraw)
	fmt.Println(string(lj))
	wcm := WsCommonMessage{
		WsCommonCommand: WsCommonCommand{
			MessageType: 11,
			ProtocolInfo: &ProtocolInfo{
				ProtocolVersion: 0,
			},
			Cmd:   "cmdooo----",
			ReqId: 12321,
		},
		//MsgData: Result{IsWinner: 1},
		MsgData: "hello cmd",
		MsgType: 13,
		Rst:     0,
		ErrMsg:  "errmsg here",
	}
	wj, _ := json.Marshal(wcm)
	fmt.Println(string(wj))
	rsm := _RoomSendMessage{
		RoomId:  "1111",
		UserIds: []int{1112, 1116},
		Msg:     &wcm,
	}
	rj, _ := json.Marshal(rsm)
	fmt.Println(string(rj))
	helper := new(nononats.NatsConnHelper)
	urls := []string{"nats://192.168.16.20:4222"}
	helper.SetIsMonitor(false)
	helper.Init(urls, nil, nil)
	nMsg := NatCommandMsg{
		Cmd:             NatCmdWriteRoomCtxLuckyDraw,
		Msg:             wj,
		WsCommonMessage: &wcm,
		UserIds:         []int{1111, 1112, 1113},
		LuckyDraw:       luckyDraw,
	}
	j, _ := json.Marshal(nMsg)
	helper.WrapPublish("nats_ws_room_subj:1111", "", j)
	fmt.Println("\n")
	fmt.Println(string(j))
	//runtime.Goexit()
	helper.Close()
}
