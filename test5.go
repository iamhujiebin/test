package main

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

var logger seelog.LoggerInterface

type _Msg struct {
	startTime int `json:"starttime"`
	endTime   int `json:"endtime"`
	index     int `json:"index"`
}

func main() {
	port := 1111
	datas, err := ioutil.ReadFile("./seelog-main.xml")
	contentStr := string(datas)
	contentStr = strings.Replace(contentStr, "##{port}", strconv.Itoa(port), -1)
	logger, err = seelog.LoggerFromConfigAsString(contentStr)
	if err != nil {
		panic(err)
	}
	defer logger.Flush()
	logger.Trace("trace")
	logger.Debugf("var = %s", "debug")
	logger.Info("info")
	logger.Error("error")
	logger.Errorf("hehe%+v","fuck")
	test()
	startTime := time.Now()
	time.Sleep(time.Second)
	endtime := time.Now()
	du := endtime.Sub(startTime)
	logger.Info(du.Nanoseconds())
	logger.Info(du.Seconds())
	logger.Error(DJBHash("10519"))
	_msg := _Msg{
		startTime: time.Now().Nanosecond(),
		endTime:   time.Now().Nanosecond() + 11111,
		index:     100,
	}
	logger.Debug(_msg)
	j, _ := json.Marshal(_msg)
	var _msg2 _Msg
	json.Unmarshal(j, &_msg2)
	logger.Error(_msg2)
}

func test() {
	logger.Info("test in info")
}

func DJBHash(str string) uint {
	if len(str) == 0 {
		return 0
	}
	hash := uint(5381)
	for i := 0; i < len(str); i++ {
		hash = ((hash << 5) + hash) + uint(str[i])
	}
	return hash
}
