package tests

import (
	"github.com/waifushork/pastemystgo"
	"os"
	"testing"
	"time"
)

// For new pastes ensure that the max paste lifetime is 1h for courtesy of paste space.

var createInfo = pastemystgo.PasteCreateInfo{
	Title:     "pastemystgotest",
	ExpiresIn: "1h",
	IsPrivate: false,
	IsPublic:  false,
	Tags:      "",
	Pasties: []pastemystgo.PastyCreateInfo{
		{
			Title: "pasty1",
			Language: "plain text",
			Code: "asd asd asd",
		},
	},
}

func TestGetPaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	tests := []struct{
		id    string
		token string
	}{
		{
			id: "i3dcx8ab",
			token: "",
		},
		{
			id: "cwy615yg",
			token: "",
		},
	}

	for _, tt := range tests { 
		paste, err := client.GetPaste(tt.id)
		if err != nil { 
			panic(err)
		}

		if paste == nil { 
			t.Errorf("paste was nil.\n%+v", paste)
		}
		if paste.Id != tt.id {
			t.Errorf("incorrect paste id. want=%s, got=%s", tt.id, paste.Id)
		}
	}
	time.Sleep(time.Second)
}

func TestTryGetPaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	tests := []struct{
		id    string
		token string
	}{
		{
			id: "i3dcx8ab",
			token: "",
		},
		{
			id: "cwy615yg",
			token: "",
		},
	}

	for _, tt := range tests {
		paste, ok := client.TryGetPaste(tt.id)
		if !ok {
			t.Fatal("unable to get paste")
		}

		if paste == nil {
			t.Fatalf("paste was nil.\n%+v", paste)
		}
		if paste.Id != tt.id {
			t.Errorf("incorrect paste id. want=%s, got=%s", tt.id, paste.Id)
		}
	}
	time.Sleep(time.Second)
}

func TestCreatePaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	paste, err := client.CreatePaste(createInfo)
	if err != nil {
		t.Fatal("unable to create paste")
	}

	if paste.Title != createInfo.Title {
		t.Errorf("could not create paste\n%+v\ntitle %s", paste, paste.Title)
	}
	time.Sleep(time.Second)
}

func TestTryCreatePaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	paste, ok := client.TryCreatePaste(createInfo)
	if !ok {
		t.Fatal("unable to create paste")
	}

	if paste.Title != createInfo.Title {
		t.Errorf("could not create paste\n%+v\ntitle %s", paste, paste.Title)
	}
	time.Sleep(time.Second)
}

func TestCreatePrivatePaste(t *testing.T) {
	var client = pastemystgo.NewClient(os.Getenv("TOKEN"))

	pasteInfo := pastemystgo.PasteCreateInfo{
		Title:     "pastemystgotest",
		ExpiresIn: "1h",
		IsPrivate: true,
		IsPublic:  false,
		Tags:      "",
		Pasties: []pastemystgo.PastyCreateInfo{
			{
				Title: "pasty1",
				Language: "plain text",
				Code: "asd asd asd",
			},
		},
	}

	paste, err := client.CreatePaste(pasteInfo)
	if err != nil {
		t.Fatal("unable to create paste for TestCreatePrivatePaste")
	}

	if paste.Title != createInfo.Title {
		t.Errorf("could not create paste\n%+v\ntitle %s", paste, paste.Title)
	}

	if !paste.IsPrivate {
		t.Errorf("paste was meant to be private, but wasn't")
	}
	time.Sleep(time.Second)
}

func TestDeletePaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	paste, err := client.CreatePaste(createInfo)
	if err != nil {
		t.Fatal("unable to create paste for TestDeletePaste")
	}

	err = client.DeletePaste(paste.Id)
	if err != nil {
		t.Errorf("paste was not deleted.\npaste id=%s\nerror=\n%s", paste.Id, err)
	}
	time.Sleep(time.Second)
}

func TestTryDeletePaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	paste, err := client.CreatePaste(createInfo)
	if err != nil {
		t.Fatal("unable to create paste for TestTryDeletePaste")
	}

	ok := client.TryDeletePaste(paste.Id)
	if !ok {
		t.Errorf("paste was not deleted.\npaste id=%s", paste.Id)
	}
	time.Sleep(time.Second)
}

func TestEditPaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	paste, err := client.CreatePaste(createInfo)
	if err != nil {
		t.Fatal("unable to create paste for TestEditPaste")
	}

	paste.Title = "edited title"

	newPaste, _ := client.EditPaste(paste)
	if newPaste.Title != "edited title" {
		t.Errorf("paste titles did not match. want=%s, got=%s", "edited title", newPaste.Title)
	}
	time.Sleep(time.Second)
}

func TestTryEditPaste(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))

	paste, err := client.CreatePaste(createInfo)
	if err != nil {
		t.Fatal("unable to create paste for TestTryEditPaste")
	}

	paste.Title = "edited title"

	newPaste, ok := client.TryEditPaste(paste)
	if !ok {
		t.Fatal("unable to edit paste")
	}

	if newPaste.Title != "edited title" {
		t.Errorf("paste titles did not match. want=%s, got=%s", "edited title", newPaste.Title)
	}
	time.Sleep(time.Second)
}

func TestBulkDeletePastes(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))
	var pastes []string
	for i := 0; i <= 5; i++ {
		paste, err := client.CreatePaste(createInfo)
		if err != nil {
			t.Fatal("unable to create paste for TestBulkDeletePastes")
		}
		pastes = append(pastes, paste.Id)
	}

	err := client.BulkDeletePastes(pastes)
	if err != nil {
		t.Fatal("unable to delete pastes, an error occurred")
	}

	if pastes != nil {
		paste, _ := client.TryGetPaste(pastes[0])

		if paste.Id != "" {
			t.Errorf("paste was not successfully deleted. paste id=%s", paste.Id)
		}
	}
	time.Sleep(time.Second)
}

func TestTryBulkDeletePastes(t *testing.T) {
	client := pastemystgo.NewClient(os.Getenv("TOKEN"))
	var pastes []string
	for i := 0; i <= 5; i++ {
		paste, err := client.CreatePaste(createInfo)
		if err != nil {
			t.Fatal("unable to create paste for TestTryBulkDeletePastes")
		}
		pastes = append(pastes, paste.Id)
	}

	ok := client.TryBulkDeletePastes(pastes)
	if !ok {
		t.Fatal("unable to delete pastes, an error occurred")
	}

	if pastes != nil {
		paste, _ := client.TryGetPaste(pastes[0])
		if paste.Id != "" {
			t.Errorf("paste was not successfully deleted. paste id=%s", paste.Id)
		}
	}
	time.Sleep(time.Second)
}