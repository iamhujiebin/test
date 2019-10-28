package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type ResendLuckyMsgData struct {
	Param     MessageSendPara
	EventTime time.Time
}

type MessageSendPara struct {
	RoomId  string          `json:"room_id"`
	UserIds []int           `json:"user_ids"`
	Msg     WsCommonMessage `json:"msg"`
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

type Person struct {
	Name string
	Age  int
}

func main() {
	per := Person{
		Name: "jiebin",
		Age:  25,
	}
	msp := MessageSendPara{
		RoomId:  "10519",
		UserIds: []int{1, 2, 3, 4, 6},
		Msg: WsCommonMessage{
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
			MsgData: per,
		},
	}
	rm := ResendLuckyMsgData{
		Param:     msp,
		EventTime: time.Now(),
	}
	j, _ := json.Marshal(rm)
	fmt.Println(string(j))
}
