package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type WsRoomRelease interface {
}

type WsRoomContext struct {
	CanRelease      bool //能否释放room了
	InteractiveLive string
	LuckyDraw       string
	User            *User
}

func (w *WsRoomContext) SetContext() {
	defer func() {
		r := recover()
		fmt.Println(r)
	}()
	getType := reflect.TypeOf(*w)
	fmt.Println("get Type is :", getType.Name())
	getValuePointer := reflect.ValueOf(w)
	getValue := reflect.ValueOf(*w)
	fmt.Println("get all Fields is:", getValue)

	fmt.Println("numfield", getType.NumField())
	// 获取方法字段
	// 1. 先获取interface的reflect.Type，然后通过NumField进行遍历
	// 2. 再通过reflect.Type的Field获取其Field
	// 3. 最后通过Field的Interface()得到对应的value
	for i := 0; i < getType.NumField(); i++ {
		field := getType.Field(i)
		value := getValue.Field(i).Interface()
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
		getValuePointer.Elem().FieldByName("User").Set("sssssfudk")
	}
	fmt.Println(w.LuckyDraw)
}

func main() {
	w := WsRoomContext{}
	w.SetContext()
}
