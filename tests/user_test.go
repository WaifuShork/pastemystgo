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