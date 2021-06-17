package pastemystgo

import (
	"errors"
)

// ExpiresIn represents all the possible time formats
// for when a paste will expire.
type ExpiresIn int

const (
	Never    ExpiresIn = iota // string form -> "never"
	OneHour                   // string form -> "1h"
	TwoHours                  // string form -> "2h"
	TenHours                  // string form -> "10h"
	OneDay                    // string form -> "1d"
	TwoDays                   // string form -> "2d"
	OneWeek                   // string form -> "1w"
	OneMonth                  // string form -> "1m"
	OneYear                   // string form -> "1y"
)

// Paste represents a single paste, containing all edits and pasties attached.
//
// If you're accessing a private paste you need to provide the Authorization header.
type Paste struct {
	// Paste Id
	Id string `json:"_id"`
	// Owner of the paste, if none then will be " "
	OwnerId string `json:"ownerId"`
	// Title of the paste
	Title string `json:"title"`
	// Date in unix time when the paste was created
	CreatedAt uint64 `json:"createdAt"`
	// When the paste expires
	ExpiresIn string `json:"expiresIn"`
	// Date in unix time when the paste will be deleted
	DeletesAt uint64 `json:"deletesAt"`
	// Amount of stars the paste has
	Stars uint64 `json:"stars"`
	// Is the paste private?
	IsPrivate bool `json:"isPrivate"`
	// Is the paste public?
	IsPublic bool `json:"isPublic"`
	// Is the paste encrypted?
	IsEncrypted bool `json:"encrypted"`
	// Slices of all tags for this paste
	Tags []string `json:"tags"`
	// Slice of all the pasties on the paste
	Pasties []Pasty `json:"pasties"`
	// Slice of all edits
	Edits []Edit `json:"edits"`
}

// Pasty represents a single pasty, could also be perceived as a "file"
// on the PasteMyst website, contains language, code, and title.
type Pasty struct {
	// Id of the pasty
	Id string `json:"_id"`
	// Title of the pasty
	Title string `json:"title"`
	// Language of the pasty
	Language string `json:"language"`
	// Code of the pasty
	Code string `json:"code"`
}

// Edit holds information about a given edit based in 'id'.
//
// You can only edit pastes on your account, so you must provide the Authorization header.
type Edit struct {
	// Unique id of the edit
	Id string `json:"_id"`
	// Edit id, multiple edits can share the same id
	// to show that multiple properties were edited
	// at the same time
	EditId uint64 `json:"editId"`
	// Type of edit (incomplete)
	EditType uint64 `json:"editType"`
	// Various metadata, most used case - storing which pasty was edited
	Metadata []string `json:"metadata"`
	// The actual data of the edit, typically stores old data
	Edit string `json:"edit"`
	// Unix time of when the edit was executed
	EditedAt uint64 `json:"editedAt"`
}

// PastyCreateInfo represents the information needed to created a new pasty
type PastyCreateInfo struct {
	// Title represents the title of a pasty
	Title string `json:"title"`
	// Language represents the language of the pasty,
	// stores the name of the language, not the mode or MIME type.
	Language string `json:"language"`
	// Code represents the code of the pasty
	Code string `json:"code"`
}

// PasteCreateInfo represents the information needed to create a new paste
type PasteCreateInfo struct {
	// Title represents the title of the paste -- optional
	Title string `json:"title,omitempty"`
	// ExpiresIn represents when the paste will expire -- optional
	ExpiresIn string `json:"expiresIn,omitempty"`
	// IsPrivate represents if it is accessible by the owner -- optional
	IsPrivate bool `json:"isPrivate,omitempty"`
	// IsPublic represents if it is displayed on the owners public profile -- optional
	IsPublic bool `json:"isPublic,omitempty"`
	// Tags represents comma separated paste tags -- optional
	Tags string `json:"tags,omitempty"`
	// Pasties represents a slice of pasties -- mandatory
	Pasties []PastyCreateInfo `json:"pasties"`
}

// GetPaste gets a paste based on Id, a token is mandatory for accessing private pastes
//
// Params:
// 	(pasteId string)
//
// Returns:
//  (*Paste, error)
func (c *Client) GetPaste(pasteId string) (paste *Paste, err error) {
	url := EndpointPaste(pasteId)

	err = c.get(url, &paste)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

// TryGetPaste attempts to get a paste based on Id, a token is mandatory for accessing private pastes.
//
// Params:
// 	(pasteId string)
//
// Returns:
//  (*Paste, bool)
func (c *Client) TryGetPaste(pasteId string) (paste *Paste, ok bool) {
	paste, err := c.GetPaste(pasteId)
	if err != nil {
		return nil, false
	}

	return paste, true
}

// CreatePaste creates a new paste with the given PasteCreateInfo
//
// Posts new pastes to (https://paste.myst.rs/api/v2/paste)
//
// Params:
// 	(createInfo PasteCreateInfo)
//
// Returns:
//  (*Paste, error)
func (c *Client) CreatePaste(createInfo PasteCreateInfo) (paste *Paste, err error) {
	// There's no sense bothering with anything else if these checks fail
	// IsPrivate, IsPublic, and Tags are related to account features, if no token is passed
	// then these flags aren't allowed to be true.
	if (createInfo.IsPrivate || createInfo.IsPublic || createInfo.Tags != "") &&
		(*c.Token == "" || c.Token == nil) {
		return nil, err
	}

	// url for where the paste will go
	url := EndpointBase + "paste/"

	err = c.post(url, createInfo, &paste)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

func (c *Client) TryCreatePaste(createInfo PasteCreateInfo) (paste *Paste, ok bool) {
	paste, err := c.CreatePaste(createInfo)
	if err != nil {
		return nil, false
	}
	return paste, true
}

// DeletePaste deletes a paste with a specified account token -- mandatory
//
// You can only delete pastes on the account of the token that has been passed.
//
// A token is required for deleting a paste because this is an account feature.
//
// This action is irreversible.
//
// Params:
// 	(pasteId string)
//
// Returns:
//  (error)
func (c *Client) DeletePaste(pasteId string) error {
	_, ok := c.TryGetPaste(pasteId)
	if !ok {
		return errors.New("error: unable to locate the paste to delete, please ensure that you're attempting to delete a valid paste")
	}

	url := EndpointPaste(pasteId)

	ok, err := c.delete(url)
	if !ok || err != nil {
		return err
	}
	return nil
}

func (c *Client) TryDeletePaste(pasteId string) bool {
	err := c.DeletePaste(pasteId)
	if err != nil {
		return false
	}
	return true
}

// EditPaste edits a paste with a specified account token -- mandatory
//
// You can only edit pastes on the account of the token that has been passed.
//
// A token is required for editing a paste because this is an account feature.
//
// To edit values of a paste you must send back the exact same paste except with the
// adjusted values, you cannot edit expiration date, any result will have no effect.
//
// Params:
// 	(paste *Paste)
//
// Returns:
//  (*Paste, error)
func (c *Client) EditPaste(paste *Paste) (*Paste, error) {
	_, ok := c.TryGetPaste(paste.Id)

	if !ok {
		return nil, errors.New("error: unable to locate the paste to delete, please ensure that you're attempting to delete a valid paste")
	}

	url := EndpointPaste(paste.Id)

	err := c.patch(url, &paste)
	if err != nil {
		return nil, err
	}

	return paste, nil
}

func (c *Client) TryEditPaste(paste *Paste) (*Paste, bool) {
	editedPaste, err := c.EditPaste(paste)
	if err != nil {
		return nil, false
	}
	return editedPaste, true
}

func (c *Client) BulkDeletePastes(pastes []string) error {
	for _, paste := range pastes {
		err := c.DeletePaste(paste)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) TryBulkDeletePastes(pastes []string) bool {
	err := c.BulkDeletePastes(pastes)
	if err != nil {
		return false
	}
	return true
}

func (c *Client) PasteExists(pasteId string) bool {
	_, ok := c.TryGetPaste(pasteId)
	if !ok {
		return false
	}
	return true
}
