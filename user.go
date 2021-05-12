package pastemystgo

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

// User represents a single pastemyst user
type User struct {
	// Id of the user
	Id              string `json:"_id"`
	// Username of the user
	Username        string `json:"username"`
	// Url of the avatar
	AvatarUrl       string `json:"avatarUrl"`
	// The users default language
	DefaultLang     string `json:"defaultLang"`
	// Specifies if their profile is public
	PublicProfile   bool   `json:"publicProfile"`
	// Specifies how long the user has been a support for, 0 if not a supporter
	SupporterLength uint64 `json:"supporterLength"`
	// Specifies if the user is a contributor
	IsContributor   bool   `json:"contributor"`
}

// Checks if a given user exists based on a username
//  
// Returns:
//  (bool, error)
func UserExists(username string) (bool, error) {
	url := UserEndpoint + url.QueryEscape(username) + "/exists"

	request, err := http.Get(url)
	if err != nil { 
		return false, sadness("%v", err)
	}

	return request.StatusCode == http.StatusOK, nil
}

// Gets a user by their username
//
// User will be nil if they don't have a public profile yet. 
//  
// Returns:
//  (*User, error)
func GetUser(username string) (*User, error) {
	var user User
	url := UserEndpoint + url.QueryEscape(username)

	request, err := http.Get(url)
	if err != nil { 
		return nil, sadness("%v", err)
	}

	if request.StatusCode != http.StatusNotFound {
		// Read the responses body to get the raw text
		bytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, sadness("Error reading Request Body.\n%v", err)
		}

		err = DeserializeJson(bytes, &user)
		if err != nil {
			return nil, sadness("Error deserializing the Request Body.\n%v", err)
		}
		return &user, nil
	}

	return &user, nil
}