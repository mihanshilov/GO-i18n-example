package i18n

import (
	nicksnyderi18n "github.com/nicksnyder/go-i18n/i18n"
)

var defaultLang string = "en"

func init(){
	nicksnyderi18n.MustLoadTranslationFile("resources/en.all.json")
	nicksnyderi18n.MustLoadTranslationFile("resources/fr.all.json")
}

func GetTranslator(lng string) (nicksnyderi18n.TranslateFunc, error) {

	return nicksnyderi18n.Tfunc(lng, "", defaultLang)
}