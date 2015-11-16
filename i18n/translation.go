package i18n

import (
	vubei18n "github.com/vube/i18n" // for formatting numbers, dates, currencies
	nicksnyderi18n "github.com/nicksnyder/go-i18n/i18n" // for translations
)

var defaultLang string = "en"
var translationFactory *vubei18n.TranslatorFactory

func SetUp(formattingRulesPath string , translatedStringsPath string, defaultLang string){

	// setup formatting
	translationFactory, _ = vubei18n.NewTranslatorFactory(
		[]string{formattingRulesPath},
		[]string{""},
		defaultLang,
	)

	// setup translations
	nicksnyderi18n.MustLoadTranslationFile("resources/en.all.json")
	nicksnyderi18n.MustLoadTranslationFile("resources/fr.all.json")
}

type I18n struct  {
	T nicksnyderi18n.TranslateFunc
	Formatter *vubei18n.Translator
}



func NewI18n (lng string) I18n{
	T, _ := nicksnyderi18n.Tfunc(lng, "", defaultLang)
	F, _ := translationFactory.GetTranslator(lng)

	return I18n{
		T: T,
		Formatter: F,
	}
}