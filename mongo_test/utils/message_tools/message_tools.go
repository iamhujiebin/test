package message_tools

import (
	"../../config"
	"../../wsm"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"
	"nonolive/nonoutils/nonohttp"
)

type BroadcastPara struct {
	RoomIds []string             `json:"room_ids"`
	Msg     *wsm.WsCommonMessage `json:"msg"`
}

type MessageSendPara struct {
	RoomId  string               `json:"room_id"`
	UserIds []int                `json:"user_ids"`
	Msg     *wsm.WsCommonMessage `json:"msg"`
}

func InteractionRoomBroadcast(msg *BroadcastPara) (returnJson *simplejson.Json, returnErr error) {
	tryNum := 3
	postBody, err := json.Marshal(msg)
	if err != nil {
		returnErr = err
		return
	}
	for tryNum > 0 {
		tryNum--
		returnErr = nil
		nonohttp.DoPostJsonSync(config.LOCAL_CLOUDAC+"/nonolive/icewolf/interactive/ws/broadcast", nil, postBody, func(resp *http.Response, body []byte, err error) {
			if err != nil {
				returnErr = err
				return
			}
			var respBody = nonohttp.NonoServerResponse{}
			err = json.Unmarshal(body, &respBody)
			if err != nil {
				returnErr = err
				return
			}
			if respBody.Code != 0 && respBody.Code != 404 {
				returnErr = errors.New(fmt.Sprintf("broadcast fail. %v %v", msg, respBody.Message))
			}
			returnJson, returnErr = simplejson.NewJson(body)
		})
		if returnErr == nil {
			return
		}
	}
	return
}

func InteractionRoomMessage(msg *MessageSendPara) (returnJson *simplejson.Json, returnErr error) {
	tryNum := 3
	postBody, err := json.Marshal(msg)
	if err != nil {
		returnErr = err
		return
	}
	for tryNum > 0 {
		tryNum--
		returnErr = nil
		nonohttp.DoPostJsonSync(config.LOCAL_CLOUDAC+"/nonolive/icewolf/interactive/ws/sendMessage", nil, postBody, func(resp *http.Response, body []byte, err error) {
			if err != nil {
				returnErr = err
				return
			}
			var respBody = nonohttp.NonoServerResponse{}
			err = json.Unmarshal(body, &respBody)
			if err != nil {
				returnErr = err
				return
			}
			if respBody.Code != 0 && respBody.Code != 404 {
				returnErr = errors.New(fmt.Sprintf("send msg fail. %v %v", msg, respBody.Message))
			}
			returnJson, returnErr = simplejson.NewJson(body)
		})
		if returnErr == nil {
			return
		}
	}
	return
}

func LuckyDrawRoomMessage(msg *MessageSendPara) (returnJson *simplejson.Json, returnErr error) {
	tryNum := 3
	postBody, err := json.Marshal(msg)
	if err != nil {
		returnErr = err
		return
	}
	for tryNum > 0 {
		tryNum--
		returnErr = nil
		nonohttp.DoPostJsonSync(config.LOCAL_CLOUDAC+"/nonolive/icewolf/luckyDraw/ws/sendMessage", nil, postBody, func(resp *http.Response, body []byte, err error) {
			if err != nil {
				returnErr = err
				return
			}
			var respBody = nonohttp.NonoServerResponse{}
			err = json.Unmarshal(body, &respBody)
			if err != nil {
				returnErr = err
				return
			}
			if respBody.Code != 0 && respBody.Code != 404 {
				returnErr = errors.New(fmt.Sprintf("send msg fail. %v %v", msg, respBody.Message))
			}
			returnJson, returnErr = simplejson.NewJson(body)
		})
		if returnErr == nil {
			return
		}
	}
	return
}