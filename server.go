package main

import (
	"html/template"
	"net/http"
	"github.com/mihanshilov/GO-i18n-example/i18n"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {

	T, _ := i18n.GetTranslator(getLocaleId(r))

	type dataForCompositeString struct {
		Item1 string
		Item2 string
		Item3 string
	}

	pageContents := map[string]string{
		"Title": T("page-title"),
		"MenuItem1": T("menu-item-1"),
		"MenuItem2": T("menu-item-2"),
		"MenuItem3": T("menu-item-3"),
		"CurrenttLocale": T("current-locale:"),
		"StaticString": T("static-string"),
		"CompositeString": T("composite-string", dataForCompositeString { Item1: "item 1", Item2: "item 2", Item3: "item 3"}),
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
