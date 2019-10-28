package compare_test

func name() {

	now := time.Now()
	luckyDraw.EndTime = &now
	luckyDraw.IsEnd = 1
	luckyDraw.EndStatus = 2
	luckyDraw.Winners = luckyUserIds

	update := bson.M{
		"end_time":   luckyDraw.EndTime,
		"is_end":     luckyDraw.IsEnd,
		"end_status": luckyDraw.EndStatus,
		"winners":    luckyDraw.Winners,
		"join_count": luckyDraw.JoinCount,
	}
	if luckyDraw.DrawConfig.DrawType == 1 {
		// TODO: 若是金币抽奖，发奖。若中奖人数小于预设的中奖人数，则返还部分金币
		// 发奖
		resendLuckyUserData := ResendLuckyUserCoinData{
			LuckyDraw:    luckyDraw,
			LuckyUserIds: luckyUserIds,
			GuestId:      c.GuestId,
			//EventTime:    time.Now(),
		}
		DelLuckDrawCloseCoinFailKey(resendLuckyUserData)
		for _, uid := range luckyUserIds {
			LuckyDrawPayAccount(uid, luckyDraw, c.GuestId, luckyDrawJoinUsersCol, false)
		}

		// 返还金币
		returnAccount := 0
		if len(luckyUserIds) < luckyDraw.DrawConfig.WinnerCount {
			returnAccount = (luckyDraw.DrawConfig.WinnerCount - len(luckyUserIds)) * (*luckyDraw.DrawConfig.CoinsPerWinner)
			err = LuckyDrawReturnAccountRequest(luckyDraw.HostId, returnAccount, c.GuestId)
			if err != nil {
				MainLogger.Errorf("retrun account err: %v", err.Error())
			}
		}

		update["real_cost"] = len(luckyUserIds) * (*luckyDraw.DrawConfig.CoinsPerWinner)
		update["return_coins"] = returnAccount
	}
	_, err = luckyDrawCol.Find(bson.M{
		"_id": luckyDraw.DrawId,
	}).Apply(mgo.Change{
		Update: bson.M{
			"$set": update,
		},
		Upsert:    false,
		ReturnNew: true,
	}, &luckyDraw)

	// 删除redis数据
	client.Cmd("DEL", redisKey)
}
