package main

import (
	vubei18n "github.com/mihanshilov/GO-i18n-example/Godeps/_workspace/src/github.com/vube/i18n" // for formatting numbers, dates, currencies
	"github.com/mihanshilov/GO-i18n-example/i18n"
	"html/template"
	"net/http"
	"time"
)

func main() {
	SetUpI18n()

	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.ListenAndServe(":8080", nil)
}

func SetUpI18n() {
	dafaultLocaleId := "en"

	pathToStringFormattingRules := "resources/vube-i18n/rules"

	translationFiles := []string{
		"resources/en.all.json",
		"resources/fr.all.json",
	}

	i18n.SetUp(pathToStringFormattingRules, translationFiles, dafaultLocaleId)
}

func getI18n(r *http.Request) (i18n.I18n) {

	var defaultLang string = "en"
	var langCookieName string = "lang"

	return i18n.GetLanguageFromRequestCookie(r, langCookieName, defaultLang)
}

func handler(w http.ResponseWriter, r *http.Request) {

	i18n := getI18n(r);

	type dataForCompositeString struct {
		Item1 string
		Item2 string
		Item3 string
	}

	formattedDate, _ := i18n.Formatter.FormatDateTime(vubei18n.DateFormatShort, time.Now().Local())

	type menuItems struct{
		Item1 string
		Item2 string
		Item3 string
	}

	type pageData struct {
		Title           string
		Menu			menuItems
		CurrenttLocale  string
		StaticString    string
		CompositeString string
		Today           string
		SelectedLanguage string
	}

	pageContents := pageData{
		Title:            i18n.T("page-title"),
		Menu:             menuItems{Item1: i18n.T("menu-item-1"), Item2: i18n.T("menu-item-2"), Item3: i18n.T("menu-item-3")},
		CurrenttLocale:   i18n.T("current-locale:"),
		StaticString:     i18n.T("static-string"),
		CompositeString:  i18n.T("composite-string", dataForCompositeString{Item1: "item 1", Item2: "item 2", Item3: "item 3"}),
		Today:            i18n.T("today", map[string]string{"Date": formattedDate}),
		SelectedLanguage: i18n.CurrentLanguageId,
	}

	t, _ := template.ParseFiles("index.html")
	t.Execute(w, pageContents)
}
