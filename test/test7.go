package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"net/http"
	"sort"
	"strings"
	"time"
)

func main() {
	a := []int{3, 5, 4, -1, 9, 11, -14}
	sort.Ints(a)
	fmt.Println(a)

	ss := []string{"surface", "ipad", "mac pro", "mac air", "think pad", "idea pad"}
	sort.Strings(ss)
	fmt.Println(ss)
	sort.Sort(sort.StringSlice(ss))
	fmt.Printf("After reverse: %v\n", ss)
	min := math.MaxInt64
	fmt.Println(min)

	params := make(map[string]string)
	params["host_id"] = "22315"
	params["user_id"] = "3990582"
	params["from"] = "weex"
	fmt.Println(createSign(params))
	outStr := "hello my friend"
	outStrArr := strings.Split(outStr, "my")
	fmt.Println(outStrArr)

	res, err := http.Head("http://video.nonolive.com/download/media/sg/nonolive/video/ff43b7398aa6f0f2ac8b0cf12893a342.mp4")
	if err != nil {
		panic(err)
	}
	contentlength := res.ContentLength
	fmt.Printf("ContentLength:%v", contentlength)
	fmt.Printf("res:%v", res.Header)
	fmt.Println("======================")
	fmt.Println(randomCursor(-2))

}

func randomCursor(maxNum int) int {
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(maxNum)
}

func createSign(params map[string]string) string {
	sign_secret_key := "PhRyzxzTYKJ5kPA8k6tRa25NPIQk5HY5la0uYMPjhtQubRnnQJHMxtPc5ZkXWTu"
	var paramsKeyList []string
	for k, _ := range params {
		paramsKeyList = append(paramsKeyList, k)
	}
	sort.Strings(paramsKeyList)
	var buffer bytes.Buffer
	buffer.WriteString(sign_secret_key)
	for _, k := range paramsKeyList {
		buffer.WriteString(params[k])
	}
	h := md5.New()
	h.Write(buffer.Bytes())
	expectSign := hex.EncodeToString(h.Sum(nil))
	return expectSign
}
