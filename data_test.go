package pastemystgo

import (
	"testing"
)

func TestGetLanguage(t *testing.T) {
	tests := []string { 
		"Autodetect",
        "Plain Text",
        "APL",
        "PGP",
        "ASN.1",
        "Asterisk",
        "Brainfuck",
        "C",
        "C++",
        "Cobol",
        "C#",
        "Clojure",
        "ClojureScript",
	}

	for _, tt := range tests { 
		language, err := GetLanguageByName(DataLanguageByName, tt)
		if err != nil {
			t.Errorf("Something went wrong\nError:%v\n%s", err, tt)
		}
	
		if language == nil {
			t.Error("Unable to requested language.")
		}
		if language.Name != tt { 
			t.Errorf("Unable to get language '%v'. got=%v", tt, language.Name)
		}
	}
}

func TestGetLanguageByExtension(t *testing.T) { 
	language, err := GetLanguageByExtension("c")
	if err != nil {
		t.Error("Something went wrong")
	}

	if language == nil {
		t.Error("Unable to requested language.")
	}
}