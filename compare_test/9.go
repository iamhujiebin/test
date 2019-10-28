func lastLinkUserLeave(hostId int, lastLinkUserId int) (err error) {
	__start := time.Now()
	defer func() {
		cost := time.Now().Sub(__start).Nanoseconds() / 1e6
		if cost >= 5000 {
			logs.MainLogger.Warnf("handling time takes longer than 5s (%v ms). There may be a problem.", cost)
		}
		logs.StatLogger.Infof("createSizeWindowInteraction cost :%v", cost)
	}()
	utils.WithinInteractionDB(func(database *mgo.Database) error {
		effectCol := database.C(utils.EFFECTED_INTERACTION_COLLECTION)
		ei := new(interaction.EffectedInteraction)
		where := bson.M{
			"host_id":  hostId,
			"cate_key": interaction.InteractionCateKeySizeWindowLink,
		}
		effectCol.Find(where).One(ei)
		if ei == nil {
			err = errors.New("no effected size window interaction")
			return nil
		}
		pull := bson.M{
			"interaction_detail.linking_users": lastLinkUserId,
		}
		where = bson.M{
			"_id":     ei.ObjectId,
			"cluster": ei.Cluster,
		}
		_, err = effectCol.Find(where).Apply(mgo.Change{
			Upsert:    false,
			ReturnNew: true,
			Update: bson.M{
				"$pull": pull,
			},
		}, &ei)
		if err != nil {
			logs.MainLogger.Errorf("get EffectedInteraction fail:%v", err)
			return nil
		}
		j, _ := json.Marshal(ei.InteractionDetail)
		var detail interaction.EffectedInteractionSizeWindowDetail
		json.Unmarshal(j, &detail)
		sw := new(interaction.SizeWindowInteraction)
		sizeWindowCol := database.C(utils.SIZE_WINDOW_INTERACTION_COLLECTION)
		update := bson.M{
			"end_time":     time.Now(),
			"run_status":   interaction.InteractionRunStatusEnd,
			"participants": detail.Participants,
		}
		where = bson.M{
			"_id":     bson.ObjectIdHex(ei.OriginInteractionId),
			"cluster": ei.Cluster,
		}
		_, err = sizeWindowCol.Find(where).Apply(mgo.Change{
			Upsert:    false,
			ReturnNew: true,
			Update: bson.M{
				"$set": update,
			},
		}, &sw)
		if err != nil {
			logs.MainLogger.Errorf("sizeWindowCol.Upsert fail,err:%v", err)
			return nil
		}
		_, err = effectCol.RemoveAll(bson.M{
			"_id": ei.ObjectId,
		})
		if err != nil {
			logs.MainLogger.Errorf("effectCol.Remove fail,err:%v", err)
			return nil
		}
		//删除数据中心的effectInteraction互动数据 //todo mongos后删除
		err, _ = datacenter.RemoveUserAllEffectIntraction(hostId)
		if err != nil {
			logs.MainLogger.Errorf("host_id :%v, remove data center effect interaction fail.err:%v", hostId, err)
		}
		return nil
	})
	return
}
