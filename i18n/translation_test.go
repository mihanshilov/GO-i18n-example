package i18n

import (
	"testing"
)

type Person struct {
	Name string
}

func TestStaticString_En(t *testing.T) {

	T, _ := GetTranslator("en");

	translated := T("hello_world");

	expected := "Hello world"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func TestComposedString_En(t *testing.T) {

	T, _ := GetTranslator("en");

	translated := T("greeting", &Person{ Name: "John"  });

	expected := "Hello John"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func TestStaticString_Fr(t *testing.T) {

	T, _ := GetTranslator("fr");

	translated := T("hello_world");

	expected := "Bonjour le monde"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}

func TestComposedString_Fr(t *testing.T) {

	T, _ := GetTranslator("fr");

	translated := T("greeting", &Person{ Name: "John"  });

	expected := "Bonjour John"
	if translated != expected {
		t.Error("Incortrect, expected " + expected + " but " + translated);
	}
}