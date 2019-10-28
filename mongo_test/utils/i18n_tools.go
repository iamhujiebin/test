package utils

import (
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

func GetTranslateText(country string) i18n.TranslateFunc {
	la := GetLanguageByCountry(country)
	T, _ := i18n.Tfunc(la, "en")
	return T
}

func GetLanguageByCountry(country string) string{
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