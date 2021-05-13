package tests

import (
	"testing"

	"github.com/waifushork/pastemystgo"
)

// For new pastes ensure that the max paste lifetime is 1m for courtesy of paste space.
func TestGetPaste(t *testing.T) {
	tests := []struct{
		id    string
		token string
	}{
		{
			id: "9dj6x301",
			token: "",
		},
		{
			id: "cwy615yg",
			token: "",
		},
		{
			id: "98xqsist",
			token: "",
		},
	}

	for _, tt := range tests { 
		paste, err := client.GetPaste(tt.id)
		if err != nil { 
			t.Error(err)
		}

		if paste == nil { 
			t.Errorf("Paste was nil.\n%+v", paste)
		}
		if paste.Id != tt.id {
			t.Errorf("Incorrect Paste Id. \nExpected: (%s)\nGot: (%s)", tt.id, paste.Id)
		}	
	}
}

func TestCreatePaste(t *testing.T) { 
	pastyCreateInfo := []pastemystgo.PastyCreateInfo{
		{
			Title: "pasty1",
			Language: "plain text",
			Code: "asd asd asd",
		},
	}
	
	createInfo := pastemystgo.PasteCreateInfo{
		Title:     "pastemystgotest",
		ExpiresIn: "1d",
		IsPrivate: false,
		IsPublic:  false,
		Tags:      "",
		Pasties: pastyCreateInfo,
	}

	paste, _ := client.CreatePaste(createInfo)

	if paste.Title != createInfo.Title { 
		t.Errorf("Could not create paste\n%+v\nTitle %s", paste, paste.Title)
	}
}

func TestCreatePrivatePaste(t *testing.T) { 
	pastyCreateInfo := []pastemystgo.PastyCreateInfo{
		{
			Title: "pasty1",
			Language: "plain text",
			Code: "asd asd asd",
		},
	}
	
	createInfo := pastemystgo.PasteCreateInfo{
		Title:     "api test paste",
		ExpiresIn: "1d",
		IsPrivate: true,
		IsPublic:  false,
		Tags:      "",
		Pasties: pastyCreateInfo,
	}
	
	paste, _ := client.CreatePaste(createInfo)

	if paste.Title != createInfo.Title {
		t.Errorf("Could not create paste\n%+v\nTitle %s", paste, paste.Title)
	}

	if !paste.IsPrivate {
		t.Errorf("Paste was not private.")
	}
}

func TestDeletePaste(t *testing.T) { 
	pastyCreateInfo := []pastemystgo.PastyCreateInfo{
		{
			Title: "pasty1",
			Language: "plain text",
			Code: "asd asd asd",
		},
	}
	
	createInfo := pastemystgo.PasteCreateInfo{
		Title:     "api test paste",
		ExpiresIn: "1d",
		IsPrivate: false,
		IsPublic:  false,
		Tags:      "",
		Pasties: pastyCreateInfo,
	}
	
	paste, _ := client.CreatePaste(createInfo)
	err := client.DeletePaste(paste.Id)
	if err != nil { 
		t.Errorf("Paste was not deleted.\nPaste Id=%v\nError=\n%v", paste.Id, err)
	}
}

func TestEditPaste(t *testing.T) { 
	pastyCreateInfo := []pastemystgo.PastyCreateInfo{
		{
			Title: "pasty1",
			Language: "plain text",
			Code: "asd asd asd",
		},
	}
	
	createInfo := pastemystgo.PasteCreateInfo{
		Title:     "api test paste",
		ExpiresIn: "1d",
		IsPrivate: false,
		IsPublic:  false,
		Tags:      "",
		Pasties: pastyCreateInfo,
	}

	paste, _ := client.CreatePaste(createInfo)
	paste.Title = "edited title"

	newPaste, _ := client.EditPaste(*paste)
	if newPaste.Title != "edited title" {
		t.Errorf("Paste was not edited")
	}
}