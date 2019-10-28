package utils

import (
	"../config"
	. "./logs"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"nonolive/nonoutils"
	"nonolive/nonoutils/nonohttp"
)

func RoomBroadcast(countries *[]string, roomId *[]string, t int, msg string) (error, int) {
	roomServer := config.GetGlobalStringValue("room_notification_server", "")
	params := map[string]interface{}{
		"type": t,
		"msg":  base64.StdEncoding.EncodeToString([]byte(msg)),
	}
	if countries != nil {
		// 国家需要urlencode
		encodedLocs := make([]string, len(*countries))
		for i, v := range *countries {
			loc, _ := nonoutils.UrlEncoded(v)
			encodedLocs[i] = loc
		}
		params["na"] = encodedLocs
	}
	if roomId != nil {
		params["roomId"] = roomId
	}
	postBody, _ := json.Marshal(params)
	urlparams := map[string]string{}

	var broadcastErr error
	var statusCode int
	nonohttp.DoPostJsonSync(roomServer+"/nonolive/room/broadcast", urlparams, postBody, func(resp *http.Response, body []byte, err error) {
		if err != nil {
			broadcastErr = err
			return
		}
		statusCode = resp.StatusCode
		if resp.StatusCode != http.StatusOK {
			return
		}

		MainLogger.Infof("RoomBroadcast success, params: %v, respBody: %s", params, string(body))
	})

	return broadcastErr, statusCode
}
