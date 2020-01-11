package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	url := "http://127.0.0.1:8500/v1/health/state/critical"
	for {
		req, _ := http.NewRequest("GET", url, nil)
		res, _ := http.DefaultClient.Do(req)
		body, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()
		fmt.Println("monitor:" + string(body))
		time.Sleep(3 * time.Second)
	}
}
