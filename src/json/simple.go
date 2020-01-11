package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
}

func main() {
	//name := Person{Name: "jeibin"}
	//j, _ := json.Marshal(name)
	j := `{"name":"jeibin"}122`
	var name *Person
	json.Unmarshal([]byte(j), &name)
	fmt.Println(name)
}
