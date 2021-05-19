package pastemystgo

import (
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
	IsContributor   bool   `json:"contributor,omitempty"`

	// These are additional fields for self user features

	// Stars represents a list of paste ids the user has starred
	Stars           []string 	      `json:"stars,omitempty"`
	// ServiceIds represents user ids of the service the user used to create an account
	ServiceIds      map[string]string `json:"serviceIds,omitempty"`
}

// GetSelfUser gets the currently logged in user, this function is not available if no token is available
//
// Returns:
//  ([]string, error)
func (c *Client) GetSelfUser() (user *User, err error) {
	if !c.IsAuthorized() {
		return nil, err//newErrorf("error: no API token has been provided")
	}

	err = c.get(EndpointSelfUser, &user)
	if err != nil {
		return nil, err//newError(err)
	}

	return user, nil
}

// GetSelfPasteIds gets all of the currently logged in users paste Ids,
// this function is not available if no token is available
//
// Returns:
//  ([]string, error)
//
// Remarks: this will return ALL paste ids, use with caution.
func (c *Client) GetSelfPasteIds() (pastes []string, err error) {
	if !c.IsAuthorized() {
		return nil, err//newErrorf("error: no API token has been provided")
	}

	err = c.get(EndpointSelfUserPastes, &pastes)
	if err != nil {
		return nil, err //newError(err)
	}
	return pastes, nil
}

// GetSelfPastes gets all of the currently logged in users pastes,
// this function is not available if no token is available
//
// Returns:
//  ([]*Paste, error)
//
// Remarks: this is a HEAVY function depending on the amount of pastes you have on your account,
// it will fetch all the paste ids, and convert them into actual paste objects.
// use with caution, you will be rate-limited.
func (c *Client) GetSelfPastes() (pastes []*Paste, err error) {
	if !c.IsAuthorized() {
		return nil, err//newErrorf("error: no API token has been provided")
	}

	pasteIds, err := c.GetSelfPasteIds()
	if err != nil {
		return nil, err//newError(err)
	}

	for i := range pasteIds {
		paste, err := c.GetPaste(pasteIds[i])
		if err != nil {
			return nil, err//newError(err)
		}

		pastes = append(pastes, paste)
	}

	return pastes, nil
}

func (c *Client) GetSelfPastesByAmount(amount int) (pastes []*Paste, err error) {
	if !c.IsAuthorized() {
		return nil, err//newErrorf("error: no API token has been provided")
	}

	pastes, err = c.GetSelfPastes()
	if err != nil {
		return nil, err
	}

	return pastes[:amount], nil
}

// UserExists checks if a given user exists based on a username
//
// Params:
// 	(username string)
//
// Returns:
//  (bool, error)
//
// Remarks: the user account MUST be public, or you must be accessing your own
// account while signed in with your API token.
func (c *Client) UserExists(username string) (bool, error) {
	url := EndpointUser + url.QueryEscape(username) + "/exists"

	request, err := http.Get(url)
	if err != nil { 
		return false, err//newError(err)
	}

	return request.StatusCode == http.StatusOK, nil
}

// GetUser gets a user by their username
//
// User will be nil if they don't have a public profile yet. 
//
// Params:
// 	(username string)
//
// Returns:
//  (*User, error)
func (c *Client) GetUser(username string) (user *User, err error) {
	url := EndpointUser + url.QueryEscape(username)

	err = c.get(url, &user)
	if err != nil { 
		return nil, err//newError(err)
	}

	return user, nil
}

// TryGetUser attempts to get a user by their username
//
// User will be nil if they don't have a public profile yet. 
//
// Params:
//  (username string)
//
// Returns:
//  (*User, bool, error)
func (c *Client) TryGetUser(username string) (*User, bool) {
	user, err := c.GetUser(username)
	if err != nil { 
		return nil, false
	}

	return user, true
}