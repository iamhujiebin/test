package main

import (
	_ "./config"
	"./utils"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"nonolive/model/interaction"
)

func main() {
	utils.WithinInteractionDB(func(database *mgo.Database) error {
		//drawId := "5ba1d1cc806a881a8bc9a760"
		var luckyDrawJoinUser *interaction.LuckyDrawJoinUser
		//luckyDrawCol := database.C(interaction.LUCKY_DRAW_COLLECTION)
		luckyDrawJoinUsersCol := database.C(interaction.LUCKY_DRAW_JOIN_USER_COLLECTION)
		err := luckyDrawJoinUsersCol.Find(bson.M{
			"user_id": 1001099,
		}).One(&luckyDrawJoinUser)
		if err != nil && err != mgo.ErrNotFound {
			fmt.Printf("get join user err: %v", err.Error())
			return nil
		}
		fmt.Printf("user:%v", luckyDrawJoinUser == nil)
		return nil
	})
}
