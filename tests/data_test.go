package tests

import (
	"testing"

	"github.com/WaifuShork/pastemystgo"
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
		language, err := pastemystgo.GetLanguageByName(pastemystgo.DataLanguageByName, tt)
		if err != nil {
			t.Errorf("Something went wrong\nError:%v\n%s", err, tt)
		}
	
		if language == nil {
			t.Error("Unable to requested language.")
		}
		if language.Name != tt { 
			t.Errorf("Unable to get language '%s'.\nGot=%s", tt, language.Name)
		}
	}
}

// TODO: Solidify security of test.
func TestGetLanguageByExtension(t *testing.T) { 
	tests := []string { 
		"c",
		"go",
		"cs",
	}
	
	for _, tt := range tests {
		language, err := pastemystgo.GetLanguageByExtension(tt)
		if err != nil {
			t.Errorf("An error occurred.\n%v", err)
		}
	
		if language == nil {
			t.Error("Language was nil.")
		}

		if language.Extensions[0] != tt { 
			t.Errorf("Unable to get language:\nGot=%s\nExpected=%s", language.Name, tt)
		}
	}
}