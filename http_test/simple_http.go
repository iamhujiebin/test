package main

import (
	"fmt"
	"net/http"
)

func goodgoodstudy(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.URL.Path) //request：http请求       response.Write([]byte("day day up")) //response：http响应
	fmt.Fprintln(response, "hello world")
}

func main() {

	http.HandleFunc("/check", goodgoodstudy) //设置访问的监听路径，以及处理方法

	http.ListenAndServe(":9000", nil) //设置监听的端口
}
