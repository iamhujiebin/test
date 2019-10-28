package compare_test

func name() {
	now := time.Now()
	luckyDraw.EndTime = &now
	luckyDraw.IsEnd = 1
	luckyDraw.EndStatus = 1
	luckyDraw.Winners = luckyUserIds
	update := bson.M{
		"end_time":   luckyDraw.EndTime,
		"is_end":     luckyDraw.IsEnd,
		"end_status": luckyDraw.EndStatus,
		"winners":    luckyDraw.Winners,
		"join_count": luckyDraw.JoinCount,
	}
	luckyDrawJoinUsersCol := database.C(interaction.LUCKY_DRAW_JOIN_USER_COLLECTION)

	if luckyDraw.DrawConfig.DrawType == 1 {
		// TODO: 若是金币抽奖，发奖。若中奖人数小于预设的中奖人数，则返还部分金币
		// 发奖
		resendLuckyUserData := ResendLuckyUserCoinData{
			LuckyDraw:    luckyDraw,
			LuckyUserIds: luckyUserIds,
			GuestId:      "",
			//EventTime:    time.Now(),
		}
		DelLuckDrawCloseCoinFailKey(resendLuckyUserData)
		for _, uid := range luckyUserIds {
			LuckyDrawPayAccount(uid, luckyDraw, "", luckyDrawJoinUsersCol, true)
		}

		// 返还金币
		returnAccount := 0
		if len(luckyUserIds) < luckyDraw.DrawConfig.WinnerCount {
			returnAccount := (luckyDraw.DrawConfig.WinnerCount - len(luckyUserIds)) * (*luckyDraw.DrawConfig.CoinsPerWinner)
			err := LuckyDrawReturnAccountRequest(luckyDraw.HostId, returnAccount, "")
			if err != nil {
				MainLogger.Errorf("retrun account err: %v", err.Error())
			}
		}

		update["real_cost"] = len(luckyUserIds) * (*luckyDraw.DrawConfig.CoinsPerWinner)
		update["return_coins"] = returnAccount
	}

	// 结果写入数据库
	luckyDrawCol := database.C(interaction.LUCKY_DRAW_COLLECTION)
	luckyDrawCol.Find(bson.M{
		//"_id": record.DrawId,//
		"_id": luckyDraw.DrawId, //TODO: 此处的record.DrawId 能否改成luckyDraw.DrawId(在自动结束抽奖的位置用的是record.DrawId)
	}).Apply(mgo.Change{
		Update: bson.M{
			"$set": update,
		},
		Upsert:    false,
		ReturnNew: true,
	}, &luckyDraw)

	// 删除redis数据
	client.Cmd("DEL", luckyDrawRedisKey)
}
