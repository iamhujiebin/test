func getRandomIds(userIds []int, randomCount int) []int {
	if len(userIds) <= randomCount {
		return userIds
	}

	resultIds := make([]int, 0)
	randomIndex := generateRandomNumber(0, len(userIds), randomCount)
	for _, i := range randomIndex {
		resultIds = append(resultIds, userIds[i])
	}

	return resultIds
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	return nums
}

func getWhiteList(drawId bson.ObjectId, luckyDrawCol *mgo.Collection) []int {
	whiteList := make([]int, 0)
	var luckyDraw *interaction.LuckyDraw
	err := luckyDrawCol.Find(bson.M{
		"_id": drawId,
	}).One(&luckyDraw)
	if err != nil {
		MainLogger.Errorf("get lucky draw white list err: %v", err.Error())
		return whiteList
	}

	whiteList = luckyDraw.WhiteList
	return whiteList
}

func NotifyLuckyDrawMessage(UserId int, luckyDraw *interaction.LuckyDraw, hostInfo *interaction.LuckyDrawUserModel) (returnErr error) {
	type NotifyLuckyDrawMessageBody struct {
		DrawId     string                            `json:"draw_id"`
		Type       int                               `json:"type"`
		DrawConfig interaction.DrawConfig            `json:"draw_config"`
		HostInfo   interaction.LuckyDrawUserModel    `json:"host_info"`
		UserId     int                               `json:"user_id"`
		IsWinner   int                               `json:"is_winner"`
		Winners    []*interaction.LuckyDrawUserModel `json:"winners"`
		Link       string                            `json:"link"`
	}

	messageBody := new(NotifyLuckyDrawMessageBody)
	messageBody.DrawId = luckyDraw.DrawId.Hex()
	messageBody.UserId = UserId
	messageBody.DrawConfig = luckyDraw.DrawConfig
	messageBody.HostInfo = *hostInfo
	messageBody.Winners = luckyDraw.WinnersWithUserModels
	messageBody.Link = config.GetGlobalStringValue("hybrid", "http://m.nonolive.com") + "/views/winner_list.html?draw_id=" + messageBody.DrawId
	if UserId == luckyDraw.HostId {
		messageBody.Type = 1
		messageBody.IsWinner = 0
	} else {
		messageBody.Type = 2
		messageBody.IsWinner = 1
	}

	tryNum := 3
	url := config.GetGlobalStringValue("cloudac_server", "") + "/nonolive/chatdisp/chat/ws/notifyLuckyDrawMessage"
	b, _ := json.Marshal(messageBody)
	for tryNum > 0 {
		tryNum--
		returnErr = nil
		nonohttp.DoPostJsonSync(url, nil, b, func(resp *http.Response, body []byte, err error) {
			if err != nil {
				returnErr = err
				return
			}
			if resp.StatusCode != http.StatusOK {
				returnErr = errors.New("resp code error:" + strconv.Itoa(resp.StatusCode))
				return
			}
			var respBody = nonohttp.NonoServerResponse{}
			err = json.Unmarshal(body, &respBody)
			if err != nil {
				MainLogger.Errorf("parse chat broadcast resp error.server:%v, %v", url, err)
				returnErr = err
				return
			}
			returnErr = nil
			if respBody.Code != 0 {
				returnErr = errors.New(respBody.Message)
				MainLogger.Errorf("callChatBroadcast response error.server:%v, %v", url, err)
			}
		})
		if returnErr == nil {
			return
		}
	}
	return
}

func luckyDrawPayAccount(userId int, luckyDraw *interaction.LuckyDraw, guestId string, luckyDrawJoinUsersCol *mgo.Collection) {
	err := luckyDrawPayAccountRequest(userId, luckyDraw.HostId, *luckyDraw.DrawConfig.CoinsPerWinner, guestId)
	if err != nil {
		MainLogger.Errorf("pay lucky draw account err: %v, draw_id: %v, user_id: %v", err.Error(), luckyDraw.DrawId.Hex(), userId)
		return
	}

	var luckyDrawJoinUser *interaction.LuckyDrawJoinUser
	_, err = luckyDrawJoinUsersCol.Find(bson.M{
		"draw_id": luckyDraw.DrawId,
		"user_id": userId,
	}).Apply(mgo.Change{
		Upsert: false,
		Update: bson.M{
			"$set": bson.M{
				"pay_success": 1,
			},
		},
		ReturnNew: true,
	}, &luckyDrawJoinUser)
	if err != nil && err != mgo.ErrNotFound {
		MainLogger.Errorf("set lucky user err: %v", err.Error())
	}
}

func luckyDrawPayAccountRequest(userId int, hostId int, account int, guestId string) (returnErr error) {
	params := map[string]string{
		"user_id":  strconv.Itoa(userId),
		"host_id":  strconv.Itoa(hostId),
		"account":  strconv.Itoa(account),
		"guest_id": guestId,
	}

	tryNum := 3
	url := config.GetGlobalStringValue("local_cloudac_server", "") + "/nonolive/gareaserv/luckyDraw/payAccount"
	for tryNum > 0 {
		tryNum--
		returnErr = nil
		nonohttp.DoGetSync(url, params, func(resp *http.Response, body []byte, err error) {
			if err != nil {
				returnErr = err
				return
			}
			if resp.StatusCode != http.StatusOK {
				returnErr = errors.New("resp code error:" + strconv.Itoa(resp.StatusCode))
				return
			}
			var respBody = nonohttp.NonoServerResponse{}
			err = json.Unmarshal(body, &respBody)
			if err != nil {
				MainLogger.Errorf("parse luckyDraw/payAccount resp error.server:%v, %v", url, err)
				returnErr = err
				return
			}
			returnErr = nil
			if respBody.Code != 0 {
				returnErr = errors.New(respBody.Message)
				MainLogger.Errorf("call luckyDraw/payAccount response error.server:%v, %v", url, err)
			}
		})
		if returnErr == nil {
			return
		}
	}
	return
}

func luckyDrawReturnAccountRequest(userId int, account int, guestId string) (returnErr error) {
	params := map[string]string{
		"user_id":  strconv.Itoa(userId),
		"account":  strconv.Itoa(account),
		"guest_id": guestId,
	}

	tryNum := 3
	url := config.GetGlobalStringValue("local_cloudac_server", "") + "/nonolive/gareaserv/luckyDraw/returnAccount"
	for tryNum > 0 {
		tryNum--
		returnErr = nil
		nonohttp.DoGetSync(url, params, func(resp *http.Response, body []byte, err error) {
			if err != nil {
				returnErr = err
				return
			}
			if resp.StatusCode != http.StatusOK {
				returnErr = errors.New("resp code error:" + strconv.Itoa(resp.StatusCode))
				return
			}
			var respBody = nonohttp.NonoServerResponse{}
			err = json.Unmarshal(body, &respBody)
			if err != nil {
				MainLogger.Errorf("parse luckyDraw/returnAccount resp error.server:%v, %v", url, err)
				returnErr = err
				return
			}
			returnErr = nil
			if respBody.Code != 0 {
				returnErr = errors.New(respBody.Message)
				MainLogger.Errorf("call luckyDraw/returnAccount response error.server:%v, %v", url, err)
			}
		})
		if returnErr == nil {
			return
		}
	}
	return
}