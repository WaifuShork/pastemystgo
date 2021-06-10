package pastemystgo

import (
	"net/http"
	"net/url"
)

// User represents a single pastemyst user
type User struct {
	// Id of the user
	Id string `json:"_id"`
	// Username of the user
	Username string `json:"username"`
	// Url of the avatar
	AvatarUrl string `json:"avatarUrl"`
	// The users default language
	DefaultLang string `json:"defaultLang"`
	// Specifies if their profile is public
	PublicProfile bool `json:"publicProfile"`
	// Specifies how long the user has been a support for, 0 if not a supporter
	SupporterLength uint64 `json:"supporterLength"`
	// Specifies if the user is a contributor
	IsContributor bool `json:"contributor"`

	// These are additional fields for self user features

	// Stars represents a list of paste ids the user has starred
	Stars []string `json:"stars,omitempty"`
	// ServiceIds represents user ids of the service the user used to create an account
	ServiceIds map[string]string `json:"serviceIds,omitempty"`
}

// GetSelfUser gets the currently logged in user, this function is not available if no token is available
//
// Returns:
//  (*User, error)
func (c *Client) GetSelfUser() (user *User, err error) {
	if !c.IsAuthorized() {
		return nil, err
	}

	err = c.get(EndpointSelfUser, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// TryGetSelfUser attempts to get the currently logged in user, this function is not available if no token is available
//
// Returns:
//  (*User, bool)
func (c *Client) TryGetSelfUser() (user *User, ok bool) {
	user, err := c.GetSelfUser()
	if err != nil {
		return nil, false
	}
	return user, true
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
		return nil, err
	}

	err = c.get(EndpointSelfUserPastes, &pastes)
	if err != nil {
		return nil, err
	}
	return pastes, nil
}

// TryGetSelfPasteIds attempts to get all of the currently logged in users paste Ids,
// this function is not available if no token is available
//
// Returns:
//  ([]string, bool)
//
// Remarks: this will return ALL paste ids, use with caution.
func (c *Client) TryGetSelfPasteIds() (pastes []string, ok bool) {
	pastes, err := c.GetSelfPasteIds()
	if err != nil {
		return nil, false
	}

	return pastes, true
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
		return nil, err
	}

	pasteIds, err := c.GetSelfPasteIds()
	if err != nil {
		return nil, err
	}

	for i := range pasteIds {
		paste, err := c.GetPaste(pasteIds[i])
		if err != nil {
			return nil, err
		}

		pastes = append(pastes, paste)
	}

	return pastes, nil
}

// TryGetSelfPastes attempts to get all of the currently logged in users pastes,
// this function is not available if no token is available
//
// Returns:
//  ([]*Paste, bool)
//
// Remarks: this is a HEAVY function depending on the amount of pastes you have on your account,
// it will fetch all the paste ids, and convert them into actual paste objects.
// use with caution, you will be rate-limited.
func (c *Client) TryGetSelfPastes() (pastes []*Paste, ok bool) {
	pastes, err := c.GetSelfPastes()
	if err != nil {
		return nil, false
	}
	return pastes, true
}

// GetSelfPastesByAmount gets a specific amount of the currently logged in users pastes,
// this function is not available if no token is available
//
// Params:
//  (amount uint)
//
// Returns:
//  ([]*Paste, error)
//
// Remarks: this is a HEAVY function depending on the amount of pastes you specify,
// use with caution, you will be rate-limited.
//
// Addendum: uint parameter because there will never be negative pastes on your account
func (c *Client) GetSelfPastesByAmount(amount uint) (pastes []*Paste, err error) {
	if !c.IsAuthorized() {
		return nil, err
	}

	pasteIds, err := c.GetSelfPasteIdsByAmount(amount)
	if err != nil {
		return nil, err
	}

	for i := range pasteIds {
		paste, err := c.GetPaste(pasteIds[i])
		if err != nil {
			return nil, err
		}

		pastes = append(pastes, paste)
	}
	return pastes, nil
}

// TryGetSelfPastesByAmount attempts to get a specific amount of the currently logged in users pastes,
// this function is not available if no token is available
//
// Params:
//  (amount uint)
//
// Returns:
//  ([]*Paste, error)
//
// Remarks: this is a HEAVY function depending on the amount of pastes you specify,
// use with caution, you will be rate-limited.
//
// Addendum: uint parameter because there will never be negative pastes on your account
func (c *Client) TryGetSelfPastesByAmount(amount uint) (pastes []*Paste, ok bool) {
	pastes, err := c.GetSelfPastesByAmount(amount)
	if err != nil {
		return nil, false
	}
	return pastes, true
}

// GetSelfPasteIdsByAmount gets a specific amount of the currently logged in user paste ids,
// this function is not available if no token is available
//
// Params:
//  (amount uint)
//
// Returns:
//  ([]string, error)
//
// Remarks: this is a HEAVY function depending on the amount of pastes you specify,
// use with caution, you will be rate-limited.
//
// Addendum: uint parameter because there will never be negative pastes on your account
func (c *Client) GetSelfPasteIdsByAmount(amount uint) (pastes []string, err error) {
	if !c.IsAuthorized() {
		return nil, err
	}

	pastes, err = c.GetSelfPasteIds()
	if err != nil {
		return nil, err
	}

	return pastes[:amount], nil
}

// TryGetSelfPasteIdsByAmount attempts to get a specific amount of the currently logged in user paste ids,
// this function is not available if no token is available
//
// Params:
//  (amount uint)
//
// Returns:
//  ([]string, bool)
//
// Remarks: this is a HEAVY function depending on the amount of pastes you specify,
// use with caution, you will be rate-limited.
//
// Addendum: uint parameter because there will never be negative pastes on your account
func (c *Client) TryGetSelfPasteIdsByAmount(amount uint) (pastes []string, ok bool) {
	pastes, err := c.GetSelfPasteIdsByAmount(amount)
	if err != nil {
		return nil, false
	}
	return pastes, true
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
	endpointUrl := EndpointUser + url.QueryEscape(username) + "/exists"

	request, err := http.Get(endpointUrl)
	if err != nil {
		return false, err
	}

	return request.StatusCode == http.StatusOK, nil
}

// GetUser gets a user by their username
//
// User will be nil if they don't have a public profile.
//
// Params:
// 	(username string)
//
// Returns:
//  (*User, error)
func (c *Client) GetUser(username string) (user *User, err error) {
	endpointUrl := EndpointUser + url.QueryEscape(username)

	err = c.get(endpointUrl, &user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// TryGetUser attempts to get a user by their username
//
// User will be nil if they don't have a public profile.
//
// Params:
//  (username string)
//
// Returns:
//  (*User, bool)
func (c *Client) TryGetUser(username string) (*User, bool) {
	user, err := c.GetUser(username)
	if err != nil {
		return nil, false
	}

	return user, true
}
