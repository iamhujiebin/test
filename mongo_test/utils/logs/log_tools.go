package logs

import (
	"../../utils/args"
	"bytes"
	"github.com/cihub/seelog"
	"io/ioutil"
	"strconv"
	"strings"
)

var MainLogger seelog.LoggerInterface
var StatLogger seelog.LoggerInterface

func init() {

	port := *args.ParamHttpPort
	var err error
	var datas []byte
	var contentStr string
	datas, err = ioutil.ReadFile("../etc/seelog-main.xml")
	contentStr = string(datas)
	contentStr = strings.Replace(contentStr, "##{port}", strconv.Itoa(port), -1)
	MainLogger, err = seelog.LoggerFromConfigAsString(contentStr)
	if err != nil {
		panic(err)
	}

	datas, err = ioutil.ReadFile("../etc/seelog-stat.xml")
	contentStr = string(datas)
	contentStr = strings.Replace(contentStr, "##{port}", strconv.Itoa(port), -1)
	StatLogger, err = seelog.LoggerFromConfigAsString(contentStr)
	if err != nil {
		panic(err)
	}

}

func LogFlushAndClose() {
	MainLogger.Flush()
	MainLogger.Close()
	StatLogger.Flush()
	StatLogger.Close()
}

func MapToStatLogger(statObj map[string]string) *bytes.Buffer {
	var buffer bytes.Buffer
	for k, v := range statObj {
		buffer.WriteByte('`')
		switch k {
		case "__v":
			k = "ver"
		case "__user_id":
			k = "uid"
		case "__platform":
			k = "fr"
		case "__guest_id":
			k = "gid"
		case "__bm":
			k = "bm"
		case "__z":
			k = "timezone"
		case "__location":
			k = "na"
		case "__sp":
			k = "sp"
		}
		buffer.WriteString(k)
		buffer.WriteByte('=')
		buffer.WriteString(v)
	}
	return &buffer
}
