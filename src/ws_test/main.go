package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"nonolive/nonoutils/nonohttp"
	"time"
)

type DisPatcherBody struct {
	Path      string `json:"path"`
	Server    string `json:"server"`
	WssServer string `json:"wss_server"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	roomId := flag.String("room_id", "10052", "the room you enter")
	//port := flag.Int("port", 3620, "server port")
	msg := flag.String("msg", "barrage msg", "set up your barrage")
	uid := flag.Int("uid", 10052, "uid")

	done := make(chan struct{})

	flag.Parse()

	RBody := DisPatcherBody{}
	//resp, err := http.Get("http://cloudac.livenono.com/nonolive/icewolf/interactive/ws/dispatcher?room_id=" + *roomId)
	resp, err := http.Get("http://127.0.0.1:38002/nonolive/icewolf/interactive/ws/dispatcher?room_id=" + *roomId)
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
	var respBody = nonohttp.NonoServerResponse{}
	err1 := json.Unmarshal(body, &respBody)
	if err1 != nil {
		log.Printf("parse error.body:%v,err:%v", string(body), err)
		return
	}
	if respBody.Code != 0 {
		log.Printf("resp code : %v", respBody.Code)
		return
	}
	j, returnErr := json.Marshal(respBody.Body)
	if returnErr != nil {
		return
	}
	if err := json.Unmarshal(j, &RBody); err != nil {
		return
	}
	//RBody.Server = fmt.Sprintf("127.0.0.1:%d", *port)
	log.Printf("resq body server:%v, path:%v", RBody.Server, RBody.Path)

	u := url.URL{Scheme: "ws", Host: RBody.Server, Path: RBody.Path, RawQuery: "room_id=" + *roomId}

	go build(u, *uid, msg)
	<-done
}

func build(u url.URL, uid int, msg *string) {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	isdone := make(chan struct{})

	go func() {
		defer func() {
			close(isdone)
		}()
		for {
			//_, _, err := c.ReadMessage()
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			log.Printf("recv: %s", message)
		}
	}()

	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()

	//enterRoom
	type MsgData struct {
		Nonopara string `json:"nonopara"`
		UserId   int    `json:"user_id"`
	}
	enterRoom := struct {
		Cmd     string  `json:"cmd"`
		ReqId   int     `json:"req_id"`
		MsgData MsgData `json:"msg_data"`
	}{Cmd: "enterRoom", ReqId: 1111, MsgData: MsgData{Nonopara: "fr=android`gid=hello-gid`na=Indonesia", UserId: uid}}
	enterRoomBody, _ := json.Marshal(enterRoom)
	c.WriteMessage(websocket.TextMessage, []byte(enterRoomBody))
	<-isdone

	type SendBarrageBody struct {
		Cmd     string `json:"cmd"`
		ReqId   int    `json:"req_id,omitempty"`
		MsgData struct {
			UserId  int    `json:"user_id"`
			Barrage string `json:"barrage"`
		} `json:"msg_data"`
	}
	var sendBarrageBody SendBarrageBody
	sendBarrageBody.Cmd = "sendBarrage"
	sendBarrageBody.ReqId = int(time.Now().Unix())
	sendBarrageBody.MsgData.Barrage = *msg

	sendBarrageBody.MsgData.UserId = uid
	jStr, _ := json.Marshal(sendBarrageBody)
	err = c.WriteMessage(websocket.TextMessage, []byte(jStr))
	if err != nil {
		log.Println("write:", err)
		return
	}

	<-isdone

	//for {
	//	select {
	//	case <-isdone:
	//		return
	//	case <-ticker.C:
	//		sendBarrageBody.MsgData.UserId = uid
	//		jStr, _ := json.Marshal(sendBarrageBody)
	//		err := c.WriteMessage(websocket.TextMessage, []byte(jStr))
	//		if err != nil {
	//			log.Println("write:", err)
	//			return
	//		}
	//	}
	//}

}
