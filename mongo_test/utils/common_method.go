package utils

import (
	"gopkg.in/mgo.v2"
	"nonolive/model/feweb"
	"gopkg.in/mgo.v2/bson"
	."./logs"
	"nonolive/nonoutils"
	"github.com/mediocregopher/radix.v2/redis"
	"strconv"
	"time"
	"regexp"
	"strings"
)

func ChkForbiddenWord(word string,country string, forbiddenWordCol *mgo.Collection) bool {
	valid := true
	var forbiddenWords []*feweb.ForbiddenWord
	err := forbiddenWordCol.Find(bson.M{
		"locations": bson.M{
			"$in": []string{"Global", country},
		},
	}).All(&forbiddenWords)
	if err != nil && err != mgo.ErrNotFound {
		MainLogger.Errorf("get Forbidden words err: %v", err.Error())
		valid = false
		return valid
	}

	for _, v := range forbiddenWords {
		if nonoutils.StringIndexOf(v.ForbiddenWords, word) > -1 {
			valid = false
		}
	}
	return valid
}

//结束抽奖上锁
func LockCloseLuckyDraw(UserId int, client *redis.Client) (r bool) {
	LuckyDrawLockSec := 20
	cacheKey := strconv.Itoa(UserId) + ":close_lucky_draw"
	now := time.Now()
	//使用redis的NX参数,当且仅当字段不存在才设置
	rsp := client.Cmd("SET", cacheKey, now.UnixNano()/1e6, "NX", "EX", LuckyDrawLockSec)
	if v, err := rsp.Str(); rsp.Err != nil || err != nil || v != "OK" {
		r = false
		MainLogger.Errorf("close lucky draw: %v another order is adding.", UserId)
	} else {
		r = true
	}
	return
}

//结束抽奖上锁
func UnlockCloseLuckyDraw(UserId int, client *redis.Client) (r bool) {
	cacheKey := strconv.Itoa(UserId) + ":close_lucky_draw"
	client.Cmd("DEL", cacheKey)
	return true
}

// 处理聊天字符串的空格
func HandleChatString(s string) string {
	s = strings.Trim(s, " ")
	pat := "\\s+" //正则
	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	s = re.ReplaceAllString(s, " ")

	return s
}