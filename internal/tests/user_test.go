package tests

import (
	"github.com/waifushork/pastemystgo"
	"os"
	"testing"
	"time"
)

func TestUserExists(t *testing.T) {
	var client = pastemystgo.NewClient(os.Getenv("TOKEN"))

	user, err := client.UserExists("waifushork")
	if err != nil {
		t.Fatal("unable to get user waifushork")
	}
	if !user {
		t.Errorf("problem locating user.")
	}
	time.Sleep(time.Second)
}

func TestTryGetUser(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	user, ok := client.TryGetUser("waifushork")
	if !ok {
		t.Fatalf("unable to locate user. want=%s, got=%s", "waifushork", user.Username)
	}

	if user.Username != "WaifuShork" {
		t.Errorf("unable to locate user. want=%s, got=%s", "waifushork", user.Username)
	}
	time.Sleep(time.Second)
}

func TestGetUser(t *testing.T) {
	var client = pastemystgo.NewClient(os.Getenv("TOKEN"))

	tests := []string{
		"codemyst",
		"waifushork",
	}

	for _, tt := range tests { 
		user, err := client.GetUser(tt)
		if err != nil {
			t.Fatalf("unable to get user %s", tt)
		}
		if !user.PublicProfile {
			t.Errorf("could not properly get user %s, please ensure their profile is public", tt)
		}
	}
	time.Sleep(time.Second)
}

func TestGetSelfUser(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	user, err := client.GetSelfUser()
	if err != nil {
		t.Fatal("unable to get self user")
	}

	if user.Username != "WaifuShork" {
		t.Errorf("username was not correct. want=%v, got=%v", "WaifuShork", user.Username)
	}
	time.Sleep(time.Second)
}

func TestTryGetSelfUser(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	user, ok := client.TryGetSelfUser()
	if !ok {
		t.Fatalf("unable to get self user")
	}

	if user.Username != "WaifuShork" {
		t.Errorf("username was not correct. want=%v, got=%v", "WaifuShork", user.Username)
	}
	time.Sleep(time.Second)
}

func TestGetSelfPastesByAmount(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, err := client.GetSelfPastesByAmount(5)
	if err != nil {
		t.Fatal("unable get pastes")
	}

	if len(pastes) != 5 {
		t.Errorf("wrong paste count. want=%d. got=%d", 5, len(pastes))
	}
	time.Sleep(time.Second)
}

func TestTryGetSelfPastesByAmount(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, ok := client.TryGetSelfPastesByAmount(5)
	if !ok {
		t.Fatalf("unable to get expected pastes")
	}

	if len(pastes) != 5 {
		t.Errorf("wrong paste count. want=%d. got=%d", 5, len(pastes))
	}
	time.Sleep(time.Second)
}

func TestGetSelfPasteIdsByAmount(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, _ := client.GetSelfPasteIdsByAmount(5)

	if len(pastes) != 5 {
		t.Errorf("wrong paste count. want=%d, got=%d", 5, len(pastes))
	}
	time.Sleep(time.Second)
}

func TestTryGetSelfPasteIdsByAmount(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, ok := client.TryGetSelfPasteIdsByAmount(5)
	if !ok {
		t.Fatal("unable to get pastes")
	}

	if len(pastes) != 5 {
		t.Errorf("wrong paste count. want=%d, got=%d", 5, len(pastes))
	}
	time.Sleep(time.Second)
}

func TestTryGetSelfPasteIds(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, ok := client.TryGetSelfPasteIds()
	if !ok {
		t.Fatal("unable to get paste ids")
	}

	if pastes == nil  {
		t.Errorf("pastes was nil\npastes: %+v", pastes)
	}

	if len(pastes) == 0 {
		t.Errorf("pastes was empty\npastes: %+v", pastes)
	}
	for _, paste := range pastes {
		if paste == "" {
			t.Error("paste id was empty.")
		}
	}
	time.Sleep(time.Second)
}


// Ensure you have at least 1 paste on your account, not sure how to test this otherwise.
func TestGetSelfPasteIds(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, _ := client.GetSelfPasteIds()

	if pastes == nil  {
		t.Errorf("pastes was nil\npastes: %+v", pastes)
	}

	if len(pastes) == 0 {
		t.Errorf("pastes was empty\npastes: %+v", pastes)
	}
	for _, paste := range pastes {
		if paste == "" {
			t.Error("paste id was empty.")
		}
	}
	time.Sleep(time.Second)
}

func TestTryGetSelfPasteId(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, ok := client.TryGetSelfPasteIds()
	if !ok {
		t.Error("unable to get expected paste ids")
	}

	if pastes == nil  {
		t.Errorf("pastes was nil\npastes: %+v", pastes)
	}

	if len(pastes) == 0 {
		t.Errorf("pastes was empty\npastes: %+v", pastes)
	}
	for _, paste := range pastes {
		if paste == "" {
			t.Error("paste id was empty.")
		}
	}
	time.Sleep(time.Second)
}

func TestGetSelfPastes(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, _ := client.GetSelfPastes()

	if pastes == nil  {
		t.Errorf("pastes was nil\npastes: %+v", pastes)
	}

	if len(pastes) == 0 {
		t.Errorf("pastes was empty\npastes: %+v", pastes)
	}

	for _, paste := range pastes {
		if paste.Id == "" {
			t.Error("paste id was empty.")
		}
	}
	time.Sleep(time.Second)
}

func TestTryGetSelfPastes(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	pastes, ok := client.TryGetSelfPastes()
	if !ok {
		t.Error("unable to get expected pastes")
	}

	if pastes == nil  {
		t.Errorf("pastes was nil\npastes: %+v", pastes)
	}

	if len(pastes) == 0 {
		t.Errorf("pastes was empty\npastes: %+v", pastes)
	}

	for _, paste := range pastes {
		if paste.Id == "" {
			t.Error("paste id was empty.")
		}
	}
	time.Sleep(time.Second)
}