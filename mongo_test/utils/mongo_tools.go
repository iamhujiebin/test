package utils

import (

	"../config"
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	sessionMap map[string]*mgo.Session = make(map[string]*mgo.Session)
)

//获取一个session。但是不会自动关闭这个 session。因此调用之后记得使用 defer关闭它
func GetSession(url string) *mgo.Session {
	var session *mgo.Session
	var exists bool
	if session, exists = sessionMap[url]; !exists {
		var err error
		session, err = mgo.Dial(url)
		if err != nil {
			panic(err) // no, not really
		}
		poolSize := config.GetGlobalIntValue("mongodb_pool_size", 400)
		session.SetPoolLimit(poolSize)
		sessionMap[url] = session

	}
	return session.Clone()
}

//公共方法，在Session中进行某个操作
func WithinDatabase(url, db string, s func(*mgo.Database) error) error {
	session := GetSession(url)
	_db := session.DB(db)
	defer session.Close()
	return s(_db)
}

type IDS struct {
	ID       int    `bson:"id"`
	TypeName string `bson:"type_name"`
}

func GetMaxId(typeName string, db *mgo.Database) int {
	c := db.C(IDS_COLLECTION)
	change := mgo.Change{
		Update:    bson.M{"$inc": bson.M{"id": 1}},
		Upsert:    true,
		ReturnNew: true,
	}
	var id IDS
	if _, err := c.Find(bson.M{"type_name": typeName}).Apply(change, &id); err != nil {
		panic(fmt.Errorf("get counter failed: %f", err.Error()))
	}
	return id.ID
}
