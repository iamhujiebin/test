package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"nonolive/model/interaction"
)

type ResendLuckyUserCoinData struct {
	LuckyUserIds []int
	LuckyDraw    interaction.LuckyDraw
	GuestId      string
}

func main() {
	coins := 12
	resendData := ResendLuckyUserCoinData{
		LuckyUserIds: []int{10010, 10011, 10012, 10013, 10014},
		GuestId:      "gueestId",
		LuckyDraw: interaction.LuckyDraw{
			DrawId: bson.ObjectIdHex("5ba1d1cc806a881a8bc9a760"),
			HostId: 1111,
			DrawConfig: interaction.DrawConfig{
				CoinsPerWinner: &coins,
			},
		},
	}
	j, _ := json.Marshal(resendData)
	fmt.Println(string(j))
}
