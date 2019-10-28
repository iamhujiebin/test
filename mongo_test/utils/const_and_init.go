package utils

import (
	"gopkg.in/mgo.v2"
	"nonolive/model/interaction"
)

//import "gopkg.in/mgo.v2"

const (
	IDS_COLLECTION                   = "ids"
	USERS_COLLECTION                 = "users"
	GIFTS_COLLECTION                 = "gifts"
	VOTE_INTERACTION_COLLECTION      = "vote_interactions"
	CHALLENGE_INTERACTION_COLLECTION = "challenge_interactions"
	EFFECTED_INTERACTION_COLLECTION  = "effected_interactions"
	INTERACTION_CATE_COLLECTION      = "interaction_cates"
	WEEX_PAGE_INFO_COLLECTION        = "weex_page_infos"
)

//MAX ID TABLE TYPE
const (
	VIDEO_TYPE = "video"
)

func ensureIndex() {
	//WithinVideoContentDB(func(db *mgo.Database) error {
	//
	//	c := db.C(VIDEO_COLLECTION)
	//	c.EnsureIndex(mgo.Index{
	//		Key:        []string{"v_id"},
	//		Unique:     true,
	//		Background: true,
	//	})
	//	c.EnsureIndex(mgo.Index{
	//		Key:        []string{"-publish_time", "author_id"},
	//		Unique:     false,
	//		Background: true,
	//	})
	//	c.EnsureIndex(mgo.Index{
	//		Key:        []string{"author_id", "-publish_time"},
	//		Unique:     false,
	//		Background: true,
	//	})
	//	c.EnsureIndex(mgo.Index{
	//		Key:        []string{"cv_id"},
	//		Unique:     false,
	//		Background: true,
	//	})
	//	return nil
	//})

	WithinInteractionDB(func(db *mgo.Database) error {

		// 投票互动索引
		voteInteractionCol := db.C(VOTE_INTERACTION_COLLECTION)
		voteInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"cate_type", "cate_key"},
			Unique:     false,
			Background: true,
		})
		voteInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"host_id", "end_status", "-end_limit"},
			Unique:     false,
			Background: true,
		})
		voteInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"host_id", "run_status", "-end_limit"},
			Unique:     false,
			Background: true,
		})
		voteInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"-create_time"},
			Unique:     false,
			Background: true,
		})
		voteInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"organizer"},
			Unique:     false,
			Background: true,
		})
		voteInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"-end_time"},
			Unique:     false,
			Background: true,
		})

		// 挑战互动索引
		challengeInteractionCol := db.C(CHALLENGE_INTERACTION_COLLECTION)
		challengeInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"cate_type", "cate_key"},
			Unique:     false,
			Background: true,
		})
		challengeInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"host_id", "end_status", "-end_limit"},
			Unique:     false,
			Background: true,
		})
		challengeInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"host_id", "run_status", "-end_limit"},
			Unique:     false,
			Background: true,
		})
		challengeInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"-create_time"},
			Unique:     false,
			Background: true,
		})
		challengeInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"organizer"},
			Unique:     false,
			Background: true,
		})
		challengeInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"-end_time"},
			Unique:     false,
			Background: true,
		})

		//好友PK索引
		hostPkCol := db.C(interaction.HOST_PK_RECREATION_COLLECTION)
		hostPkCol.EnsureIndex(mgo.Index{
			Key:        []string{"-create_time"},
			Unique:     false,
			Background: true,
		})
		hostPkCol.EnsureIndex(mgo.Index{
			Key:        []string{"participants","run_status"},
			Unique:     false,
			Background: true,
		})
		hostPkCol.EnsureIndex(mgo.Index{
			Key:        []string{"participants","result"},
			Unique:     false,
			Background: true,
		})
		hostPkCol.EnsureIndex(mgo.Index{
			Key:        []string{"participants","host_pk_recreation_status"},
			Unique:     false,
			Background: true,
		})

		// 通用互动索引
		effectedInteractionCol := db.C(EFFECTED_INTERACTION_COLLECTION)
		effectedInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"cate_type", "cate_key"},
			Unique:     false,
			Background: true,
		})
		effectedInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"host_id", "end_status", "-end_limit"},
			Unique:     false,
			Background: true,
		})
		effectedInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"host_id", "run_status", "-end_limit"},
			Unique:     false,
			Background: true,
		})
		effectedInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"-create_time"},
			Unique:     false,
			Background: true,
		})
		effectedInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"organizer"},
			Unique:     false,
			Background: true,
		})
		effectedInteractionCol.EnsureIndex(mgo.Index{
			Key:        []string{"-end_time"},
			Unique:     false,
			Background: true,
		})
		return nil
	})
}
func init() {
	ensureIndex()
}
