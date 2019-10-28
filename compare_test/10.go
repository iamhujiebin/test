func otherLinkUserLeave(hostId int, linkUserId int) (err error) {
	__start := time.Now()
	defer func() {
		cost := time.Now().Sub(__start).Nanoseconds() / 1e6
		if cost >= 5000 {
			logs.MainLogger.Warnf("handling time takes longer than 5s (%v ms). There may be a problem.", cost)
		}
		logs.StatLogger.Infof("updateSizeWindowInteraction cost :%v", cost)
	}()
	utils.WithinInteractionDB(func(database *mgo.Database) error {
		ei := new(interaction.EffectedInteraction)
		effectCol := database.C(utils.EFFECTED_INTERACTION_COLLECTION)
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
			"interaction_detail.linking_users": linkUserId,
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
			logs.MainLogger.Errorf("effectCol.Update fail,err:%v", err)
		}
		return nil
	})
	return
}
