package compare_test

func main() {
	anchorId, _ := strconv.Atoi(room.RoomId)
	anchorCol := database.C(feweb.LIVE_ANCHOR_COLLECTION) //获取主播的信息
	err = anchorCol.Update(bson.M{"user_id": anchorId}, bson.M{"$addToSet": bson.M{"room_guests": data.UserId}})
	if err != nil {
		if err != mgo.ErrNotFound {
			errCode = _DatabaseServerError
			errMessage = c.CreateLogPrefix() + "update live anchor room_guests data failed. ErrMessage: " + err.Error()
			logs.MainLogger.Error(errMessage)
			c.ResponseWrongMessage(wcc, _DatabaseServerError, errMessage)
			return nil
		}
	}
	err = userCol.Update(bson.M{"user_id": data.UserId}, bson.M{"$set": bson.M{"is_being_guest": 1}})
	if err != nil {
		if err != mgo.ErrNotFound {
			errCode = _DatabaseServerError
			errMessage = c.CreateLogPrefix() + "update user's is_being_guest data failed. ErrMessage: " + err.Error()
			logs.MainLogger.Error(errMessage)
			c.ResponseWrongMessage(wcc, _DatabaseServerError, errMessage)
			return nil
		}
	}
}
