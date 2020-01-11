package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type DisPatcherRespBody struct {
	Path      string `json:"path"`
	Server    string `json:"server"`
	WssServer string `json:"wss_server"`
}

var _RespBody struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Body    interface{} `json:"body,omitempty"`
}

func Dispatcher(roomId string) (ResBody *DisPatcherRespBody, err error) {
	//ResBody := new(DisPatcherRespBody)
	resp, err := http.Get("http://cloudac.livenono.com/nonolive/icewolf/interactive/ws/dispatcher?room_id=" + roomId)
	if err != nil {
		log.Printf("http get failed. %v", err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("parse body error. %v", err)
		return
	}
	err = json.Unmarshal(body, &_RespBody)
	if err != nil {
		log.Printf("parse error.body:%v,err:%v", string(body), err)
		return
	}
	if _RespBody.Code != 0 {
		log.Printf("resp code : %v", _RespBody.Code)
		err = errors.New("resp code error, code : " + strconv.Itoa(_RespBody.Code))
		return
	}
	j, err := json.Marshal(_RespBody.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(j, &ResBody)
	if err != nil {
		return
	}
	return ResBody, err
}
