package i18n

import (
	"testing"
)

type Person struct {
	Name string
}

func TestStaticString_En(t *testing.T) {

	i18n := createI18n("en")

	translated := i18n.T("hello_world");

	expected := "Hello world"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func TestComposedString_En(t *testing.T) {

	i18n := createI18n("en")

	translated := i18n.T("greeting", &Person{ Name: "John"  });

	expected := "Hello John"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func TestStaticString_Fr(t *testing.T) {

	i18n := createI18n("fr")

	translated := i18n.T("hello_world");

	expected := "Bonjour le monde"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func TestComposedString_Fr(t *testing.T) {

	i18n := createI18n("fr")

	translated := i18n.T("greeting", &Person{ Name: "John"  });

	expected := "Bonjour John"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func createI18n(localeId string) I18n{

	dafaultLocaleId := "en"

	pathToStringFormattingRules := "resources/vube-i18n/rules"

	translationFiles := []string{
		"resources/en.all.json",
		"resources/fr.all.json",
	}

	SetUp(pathToStringFormattingRules, translationFiles, dafaultLocaleId)
	i18n := NewI18n(localeId)
	return i18n
}
