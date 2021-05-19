package tests

import (
	"os"
	"testing"

	"github.com/waifushork/pastemystgo"
)
var client = pastemystgo.NewClient(os.Getenv("TOKEN"))

// Keep the list small due to rate-limiting
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
		language, err := client.GetLanguageByName(tt)
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

func TestGetLanguageByExtension(t *testing.T) { 
	tests := []string { 
		"c",
		"go",
		"cs",
	}
	
	for _, tt := range tests {
		language, err := client.GetLanguageByExtension(tt)
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