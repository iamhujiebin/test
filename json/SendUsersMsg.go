package main

import (
	"encoding/json"
	"fmt"
)

type _RoomSendMessage struct {
	RoomId  string      `json:"room_id"`
	UserIds []int       `json:"user_ids"`
	Msg     interface{} `json:"msg"`
}

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

func main() {
	w := WsCommonMessage{
		WsCommonCommand: WsCommonCommand{
			MessageType: 1,
			ProtocolInfo: &ProtocolInfo{
				ProtocolVersion: 1,
				CompressSupport: 2,
				CompressUse:     3,
				BodyLength:      16,
			},
			Cmd:   "testcmd",
			ReqId: 2},
		Rst:     11,
		ErrMsg:  "error message",
		MsgType: 109,
		MsgData: "test msg",
	}
	msg := _RoomSendMessage{
		RoomId:  "1111",
		UserIds: []int{1112},
		Msg:     w,
	}
	j, _ := json.Marshal(msg)
	fmt.Println(string(j))
}
