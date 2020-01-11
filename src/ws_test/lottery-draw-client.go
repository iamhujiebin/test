package main

import (
	"./utils"
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/url"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	roomId := flag.String("room_id", "22102", "the room you enter")
	msg := flag.String("msg", "1111", "set up your barrage")
	reconnect := flag.Bool("reconnect", true, "socket automatic reconnect")

	done := make(chan struct{})

	flag.Parse()

	resBody, err := utils.Dispatcher(*roomId)
	if err != nil {
		log.Printf("dispatcher failed. %v", err)
		return
	}
	resBody.Server = "127.0.0.1:38000"
	log.Printf("resq body server:%v, path:%v", resBody.Server, resBody.Path)

	u := url.URL{Scheme: "ws", Host: resBody.Server, Path: resBody.Path, RawQuery: "room_id=" + *roomId}

	for i := 0; i < 1000; i++ {
		time.Sleep(time.Millisecond * 5)
		uid := rand.Intn(9999999)
		go runDrawClient(u, uid, msg, reconnect)
	}
	<-done
}

func runDrawClient(u url.URL, uid int, msg *string, reconnect *bool) {
	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer func() {
			close(done)
			if *reconnect {
				go runDrawClient(u, uid, msg, reconnect)
			}
		}()
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			//log.Printf("recv: %s", message)
		}
	}()
	n := rand.Intn(10) + 1

	ticker := time.NewTicker(time.Second * 5)
	closeTicker := time.NewTicker(time.Second * time.Duration(n*5))
	defer ticker.Stop()

	//enterRoom
	type MsgData struct {
		Nonopara  string `json:"nonopara"`
		UserId    int    `json:"user_id"`
		Loginname string `json:"loginname"`
	}
	enterRoom := struct {
		Cmd     string  `json:"cmd"`
		ReqId   int     `json:"req_id"`
		MsgData MsgData `json:"msg_data"`
	}{Cmd: "enterRoom", ReqId: 1111, MsgData: MsgData{Nonopara: "fr=android`gid=hello-gid`na=Indonesia", UserId: uid, Loginname: "robot"}}
	enterRoomBody, _ := json.Marshal(enterRoom)
	c.WriteMessage(websocket.TextMessage, []byte(enterRoomBody))

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

	//<-done

	for {
		select {
		case <-done:
			return
		case <-ticker.C:
			jStr, _ := json.Marshal(sendBarrageBody)
			err := c.WriteMessage(websocket.TextMessage, []byte(jStr))
			if err != nil {
				log.Println("write:", err)
				return
			}
		case <-closeTicker.C:
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return
		}
	}

}
