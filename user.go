package pastemystgo

import (
	"io"
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

	Stars           []string `json:"stars"`
	ServiceIds      []interface{} `json:"serviceIds"`
	Pastes          []Paste `json:"pastes"`
}

func (c *Client) GetCurrentUser() (*User, error) {
	url := UserEndpoint + "self"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, newError(err)
	}

	request.Header.Add("Content-Type", "application/json")

	if c.Token != "" {
		request.Header.Add("Authorization", c.Token)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, newError(err)
	}

	var user User
	err = c.bodyToJson(response, &user)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			newError(err)
		}
	} (response.Body)

	return &user, nil
}

// GetCurrentUserPastes gets all of the currently logged in users pastes
//
// Returns:
//  ([]string, error)
func (c *Client) GetCurrentUserPastes() ([]string, error) {
	url := SelfUserPastesEndpoint

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, newError(err)
	}

	request.Header.Add("Content-Type", "application/json")

	if c.Token != "" {
		request.Header.Add("Authorization", c.Token)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, newError(err)
	}

	var pastes []string
	err = c.bodyToJson(response, &pastes)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			newError(err)
		}
	} (response.Body)

	return pastes, nil
}

// UserExists checks if a given user exists based on a username
//  
// Returns:
//  (bool, error)
func (c *Client) UserExists(username string) (bool, error) {
	url := UserEndpoint + url.QueryEscape(username) + "/exists"

	request, err := http.Get(url)
	if err != nil { 
		return false, newError(err)
	}

	return request.StatusCode == http.StatusOK, nil
}

// GetUser gets a user by their username
//
// User will be nil if they don't have a public profile yet. 
//  
// Returns:
//  (*User, error)
func (c *Client) GetUser(username string) (*User, error) {
	var user User
	url := UserEndpoint + url.QueryEscape(username)
	// c := &Client{}
	err := c.get(url, &user)
	if err != nil { 
		return nil, newError(err)
	}

	return &user, nil
}

// TryGetUser attempts to get a user by their username
//
// User will be nil if they don't have a public profile yet. 
//  
// Returns:
//  (*User, bool, error)
func (c *Client) TryGetUser(username string) (*User, bool, error) {
	user, err := c.GetUser(username)
	if err != nil { 
		return nil, false, newError(err)
	}

	return user, true, nil
}