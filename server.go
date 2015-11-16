package main

import (
	"html/template"
	"net/http"
	"github.com/mihanshilov/GO-i18n-example/i18n"
	"time"
	vubei18n "github.com/vube/i18n" // for formatting numbers, dates, currencies
)

func main() {
	SetUpI18n()

	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})

	http.ListenAndServe(":8080", nil)
}

func SetUpI18n(){
	dafaultLocaleId := "en"

	pathToStringFormattingRules := "resources/vube-i18n/rules"

	translationFiles := []string{
		"resources/en.all.json",
		"resources/fr.all.json",
	}

	i18n.SetUp(pathToStringFormattingRules, translationFiles, dafaultLocaleId)
}

func handler(w http.ResponseWriter, r *http.Request) {

	i18n := i18n.NewI18n(getLocaleId(r))

	type dataForCompositeString struct {
		Item1 string
		Item2 string
		Item3 string
	}

	formattedDate, _ := i18n.Formatter.FormatDateTime(vubei18n.DateFormatShort, time.Now().Local())

	pageContents := map[string]string{
		"Title": i18n.T("page-title"),
		"MenuItem1": i18n.T("menu-item-1"),
		"MenuItem2": i18n.T("menu-item-2"),
		"MenuItem3": i18n.T("menu-item-3"),
		"CurrenttLocale": i18n.T("current-locale:"),
		"StaticString": i18n.T("static-string"),
		"CompositeString": i18n.T("composite-string", dataForCompositeString { Item1: "item 1", Item2: "item 2", Item3: "item 3"}),
		"Today": i18n.T("today", map[string]string{ "Date": formattedDate }),
	}

	t, _:= template.ParseFiles("index.html")
	t.Execute(w, pageContents)
}

// try get selected locale from cookie
func getLocaleId(r *http.Request) string{

	defaultLocale := "en"

	cookie, err := r.Cookie("lang")

	if err != nil {
		return defaultLocale;
	}

	if cookie == nil || cookie.Value == "" {
		return "en"
	} else {
		return cookie.Value
	}
}
