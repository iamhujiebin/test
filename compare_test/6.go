package compare_test

func main() {
	// 补全中奖用户信息
	hostInfo := new(interaction.LuckyDrawUserModel)
	luckyDraw.WinnersWithUserModels = make([]*interaction.LuckyDrawUserModel, 0)
	utils.WithinFEWebDB(func(fewebDb *mgo.Database) error {
		users := make([]feweb.BasicUser, 0)
		userIds := append(luckyUserIds, luckyDraw.HostId)
		userCol := fewebDb.C(feweb.USERS_COLLECTION)
		err := userCol.Find(bson.M{
			"user_id": bson.M{
				"$in": userIds,
			},
		}).Select(bson.M{"user_id": 1, "avatar": 1, "level": 1, "loginname": 1}).All(&users)
		if err != nil && err != mgo.ErrNotFound {
			MainLogger.Errorf("get winner info err: %v", err.Error())
			utils.UnlockCloseLuckyDraw(record.HostId, client)
			return nil
		}

		for _, v := range users {
			u := new(interaction.LuckyDrawUserModel)
			u.UserId = v.UserId
			u.Level = v.Level
			u.LoginName = v.LoginName
			u.Avatar = v.Avatar
			if luckyDraw.HostId == v.UserId {
				hostInfo = u
			} else {
				luckyDraw.WinnersWithUserModels = append(luckyDraw.WinnersWithUserModels, u)
			}
		}

		return nil
	})

	luckyDraw.CreateTimeMs = luckyDraw.CreateTime.UnixNano() / 1e6
	luckyDraw.EndTimeMs = luckyDraw.EndTime.UnixNano() / 1e6
}
