package test

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

type Greeting func(name string) string

func (g Greeting) say(str string) {
	fmt.Println(g(str))
}

func english(name string) string {
	return "Hello," + name
}

func french(name string) string {
	return "Bonjour," + name
}

func main() {
	g := Greeting(english)
	g.say("World")
	g = Greeting(french)
	g.say("World")
	j := simplejson.New()
	j.Set("key", "val")
	body := `{"test":{"testinit":123}}`
	j2, _ := simplejson.NewJson([]byte(body))
	fmt.Println(j.Get("key").MustString())
	fmt.Println(j2.Get("test").Get("testinit").MustString("123defalut"))
	mj2, _ := j2.MarshalJSON()
	fmt.Println(string(mj2))
}
