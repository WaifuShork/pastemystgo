package main

import (
	"testing"
)

func TestGetLanguage(t *testing.T) {
	language, err := getLanguageByName(DataLanguageByName, "C")
	if err != nil {
		t.Error("Something went wrong")
	}

	if language == nil {
		t.Error("Unable to requested language.")
	}
}