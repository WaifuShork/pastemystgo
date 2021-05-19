package tests

import (
	"testing"
)

func TestUserExists(t *testing.T) {
	user, _ := client.UserExists("codemyst")
	if !user { 
		t.Errorf("Problem locating user.")
	}
}

func TestGetUser(t *testing.T) { 
	tests := []string{
		"codemyst",
		"waifushork",
	}

	for _, tt := range tests { 
		user, _ := client.GetUser(tt)
		if !user.PublicProfile {
			t.Errorf("Could not properly get user %s, please ensure their profile is public.", tt)
		}
	}
}

func TestGetSelfUser(t *testing.T) {
	user, _ := client.GetSelfUser()

	if user.Username != "WaifuShork" {
		t.Errorf("Username was not correct, want=%v. got=%v", "WaifuShork", user.Username)
	}
}

func TestGetSelfPastesByAmount(t *testing.T) {
	pastes, err := client.GetSelfPastesByAmount(5)
	if err != nil {
		panic(err)
	}

	if len(pastes) != 5 {
		t.Errorf("wrong paste count. want=%d. got=%d", 5, len(pastes))
	}
}

// Ensure you have at least 1 paste on your account, not sure how to test this otherwise.
func TestGetSelfPasteIds(t *testing.T) {
	pastes, _ := client.GetSelfPasteIds()

	if pastes == nil || len(pastes) == 0 {
		t.Errorf("paste(s) count was 0 or nil. %+v", pastes)
	}
}

func TestGetSelfPastes(t *testing.T) {
	pastes, _ := client.GetSelfPastes()

	if pastes == nil || len(pastes) == 0 {
		t.Errorf("Pastes was nil\npastes: %+v", pastes)
	}

	for _, paste := range pastes {
		if paste.Id == "" {
			t.Error("Paste Id was empty.")
		}
	}
}

func TestIsAuthorized(t *testing.T) {
	isAuthorized := client.IsAuthorized()

	if !isAuthorized {
		t.Errorf("unable to access API token")
	}
}