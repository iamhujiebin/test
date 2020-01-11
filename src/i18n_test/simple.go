package main

import (
	"fmt"
	"github.com/nicksnyder/go-i18n/i18n"
)

func main() {
	i18n.MustLoadTranslationFile("./i18n/en.all.json")
	i18n.MustLoadTranslationFile("./i18n/id.all.json")
	i18n.MustLoadTranslationFile("./i18n/tr.all.json")
	i18n.MustLoadTranslationFile("./i18n/vi.all.json")
	i18n.MustLoadTranslationFile("./i18n/es.all.json")
	i18n.MustLoadTranslationFile("./i18n/th.all.json")
	i18n.MustLoadTranslationFile("./i18n/zh.all.json")
	T, _ := i18n.Tfunc("vi", "en")
	fmt.Println(T("hello"))
	fmt.Println(T("your_unread_email_count", 0))
}
