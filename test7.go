package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/i18n"
)

const (
	Indonesia = "Indonesia"
	Turkey    = "Turkey"
	Russian   = "Russian"
	Thailand  = "Thailand"
	Vietnam   = "Vietnam"
	Malaysia  = "Malaysia"
	China     = "China"
	Singapore = "Singapore"
	Spain     = "Spain"
)

type Person struct {
	Name string
	Age  int
}

type Man struct {
	Person
	ManName string
}

func main() {
	persion := Person{Name: "hjb", Age: 11}
	man := new(Man)
	man.Person = persion
	fmt.Printf("My Name is %s\n", man.Name)
	la := GetLanguageByCountry(Turkey)
	T, _ := i18n.Tfunc(la, "en")
	a := T("你好")
	fmt.Println(a)
}

func GetLanguageByCountry(country string) string {
	if country == Turkey {
		return "tr"
	} else if country == Russian {
		return "en"
	} else if country == Vietnam {
		return "vi"
	} else if country == Thailand {
		return "th"
	} else if country == Indonesia {
		return "id"
	} else if country == Malaysia {
		return "id"
	} else if country == Spain {
		return "es"
	} else {
		return "en"
	}
	return "en"
}
