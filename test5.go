package main

import (
	"github.com/cihub/seelog"
	"io/ioutil"
	"strconv"
	"strings"
)

var logger seelog.LoggerInterface

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
	test()
}

func test() {
	logger.Info("test in info")
}
