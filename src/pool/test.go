package main

import (
	"fmt"
	"github.com/wangyaofenghist/go-Call/call"
	"github.com/wangyaofenghist/go-Call/test"
	"github.com/wangyaofenghist/go-worker-base/worker"
	"time"
)

//声明一号池子
var poolOne worker.WorkPool

//声明回调变量
var funcs call.CallMap

//以结构体方式调用
type runWorker struct{}

//初始化协程池 和回调参数
func init() {
	poolOne.InitPool()
	funcs = call.CreateCall()

}

//通用回调
func (f *runWorker) Run(param []interface{}) {
	name := param[0].(string)
	//调用回调并拿回结果
	funcs.Call(name, param[1:]...)
}

//主函数
func main() {
	var runFunc runWorker = runWorker{}
	funcs.AddCall("test4", test.Test4)
	for i := 0; i < 10000; i++ {
		poolOne.Run(runFunc.Run, "test4", " aa ", " BB")
		poolOne.Run(runFunc.Run, "test4", " cc ", " dd")
		//RunAuto 可以在压力大的时候自动扩充
		poolOne.RunAuto(runFunc.Run, "test4", " ee ", " ff")
		fmt.Println("come", i)
	}
	time.Sleep(time.Millisecond * 1000)
	poolOne.Stop()
}
