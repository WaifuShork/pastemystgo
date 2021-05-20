package tests

import (
	"os"
	"testing"

	"github.com/waifushork/pastemystgo"
)

// Keep the list small due to rate-limiting
func TestGetLanguage(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

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
			t.Fatalf("something went wrong\nerror:%v\n%s", err, tt)
		}
	
		if language == nil {
			t.Fatal("unable to get requested language.")
		}

		if language.Name != tt { 
			t.Errorf("unable to get language. want=%s, got=%s", tt, language.Name)
		}
	}
}

func TestTryGetLanguage(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

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
		language, ok := client.TryGetLanguageByName(tt)
		if !ok {
			t.Fatalf("something went wrong\n%s", tt)
		}

		if language == nil {
			t.Fatal("unable to get requested language.")
		}

		if language.Name != tt {
			t.Errorf("unable to get language. want=%s, got=%s", tt, language.Name)
		}
	}
}

func TestGetLanguageByExtension(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	tests := []string { 
		"c",
		"go",
		"cs",
	}
	
	for _, tt := range tests {
		language, err := client.GetLanguageByExtension(tt)
		if err != nil {
			t.Fatal("unable to get language by name")
		}
	
		if language == nil {
			t.Fatal("unable to get requested language.")
		}

		if language.Extensions[0] != tt { 
			t.Errorf("unable to get language. want=%s, got=%s", tt, language.Name)
		}
	}
}

func TestTryGetLanguageByExtension(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	tests := []string {
		"c",
		"go",
		"cs",
	}

	for _, tt := range tests {
		language, ok := client.TryGetLanguageByExtension(tt)
		if !ok {
			t.Fatal("unable to get language by extension")
		}

		if language == nil {
			t.Fatal("unable to get requested language.")
		}

		if language.Extensions[0] != tt {
			t.Errorf("unable to get language. want=%s, got=%s", tt, language.Name)
		}
	}
}