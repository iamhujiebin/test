package utils

import (
	"encoding/json"
	"fmt"
	"github.com/mediocregopher/radix.v2/pool"
	"github.com/mediocregopher/radix.v2/redis"
	//"gopkg.in/mgo.v2/bson"
	"nonolive/nonoutils/config"
	"nonolive/nonoutils/cron"
	"runtime"
	"time"
)

var _pool *pool.Pool

type RestartFunc func()

var restartFuncMap map[string]RestartFunc

func init() {
	var err error
	url := config.GetGlobalStringValue("redis_url", "192.168.16.20:6379")
	_pool, err = pool.New("tcp", url, 20)
	if err != nil {
		fmt.Printf("connect to redis(%v) fail.\n", url)
		panic(err)
	}
	fmt.Printf("connect to redis(%v) ok.\n", url)
}

func WithinRedis(f func(*redis.Client) error) error {
	conn, err := _pool.Get()
	if err != nil {
		return err
	}
	defer func() {
		if conn != nil {
			conn.PipeClear()
		}
		_pool.Put(conn)
	}()
	return f(conn)
}

type MessageSendPara struct {
	RoomId  string      `json:"room_id"`
	UserIds []int       `json:"user_ids"`
	Msg     interface{} `json:"msg"`
}

func Run() {
	initJob()
	restartFuncMap = make(map[string]RestartFunc)
	restartFuncMap["lucky_draw_fail_resend"] = _checkLuckyDrawResend
	redisKey := config.GetStringValue("rediskey", "lucky_draw_fail_resend", "LUCKY_DRAW_FAIL_RESEND")
	drawId := "1"
	WithinRedis(func(client *redis.Client) error {
		msg := "sgData"
		param := MessageSendPara{
			RoomId:  "1300921",
			UserIds: []int{1, 2, 3, 4, 5},
			Msg:     msg,
		}
		j, _ := json.Marshal(param)
		client.Cmd("HSET", redisKey, drawId, j)
		return nil
	})
	go _delLuckDrawCloseFailKey(drawId)
	runtime.Goexit()
}

func _delLuckDrawCloseFailKey(drawId string) {
	fmt.Println("_delLuckDrawCloseFailKey start")
	redisKey := config.GetStringValue("rediskey", "lucky_draw_fail_resend", "LUCKY_DRAW_FAIL_RESEND")
	luckyDarwFailInterval := 5
	time.Sleep(time.Duration(luckyDarwFailInterval) * time.Second)
	WithinRedis(func(client *redis.Client) error {
		client.Cmd("HDEL", redisKey, drawId)
		return nil
	})
}

func _checkLuckyDrawResend() {
	fmt.Println("_checkLuckyDrawResend start")
	redisKey := config.GetStringValue("rediskey", "lucky_draw_fail_resend", "LUCKY_DRAW_FAIL_RESEND")
	WithinRedis(func(client *redis.Client) error {
		dataArr, err := client.Cmd("HGETALL", redisKey).Array()
		if err != nil {
			fmt.Println(err)
			return err
		}
		for index, key := range dataArr {
			fmt.Printf("index:%d,key:%s", index, key.String())
			var msg MessageSendPara
			b, _ := key.Bytes()
			json.Unmarshal(b, &msg)
			fmt.Println(msg)
		}
		return nil
	})
}

func _weexSurviceCheck() {
	fmt.Println("_weexSurviceCheck start")
	redisKey := config.GetStringValue("rediskey", "weex_survive_key", "WEEX_SURVIVE_KEY")
	WithinRedis(func(client *redis.Client) error {
		rsp := client.Cmd("GET", redisKey)
		isSurvive, _ := rsp.Int()
		fmt.Printf("isSurvive:%d", isSurvive)
		if isSurvive > 0 {
			client.Cmd("EXPIRE", redisKey, 10)
			return nil
		} else {
			if len(restartFuncMap) > 0 {
				for name, f := range restartFuncMap {
					fmt.Printf("%s func start\n", name)
					f()
				}
			}
			client.Cmd("SET", redisKey, 1)
			client.Cmd("EXPIRE", redisKey, 3)
		}
		return nil
	})
}

func initJob() {
	c := cron.New()
	c.AddFunc("*/3 * * * * ?", func() {
		_weexSurviceCheck()
	})
	c.Start()
}
