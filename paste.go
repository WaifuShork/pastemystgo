package pastemystgo

import (
	"errors"
)

// Represents an enumeration of expiration values
// GetExpiresInString(expiresIn ExpiresIn)(string)
// will return the string format of expiration
// All possible expiration values
type ExpiresIn int
const (
	Never ExpiresIn = iota
	OneHour
	TwoHours
	TenHours
	OneDay
	TwoDays
	OneWeek
	OneMonth
	OneYear
)

// Represents a single paste, containing all edits and pasties attached
type Paste struct {
	// Paste Id
	Id          string    `json:"_id"`
	// Owner of the paste, if none then will be " "
	OwnerId     string    `json:"ownerId"`
	// Title of the paste
	Title       string    `json:"title"`
	// Date in unix time when the paste was created
	CreatedAt   uint64    `json:"createdAt"`
	// When the paste expires
	ExpiresIn   string    `json:"expiresIn"`
	// Date in unix time when the paste will be deleted
	DeletesAt   uint64    `json:"deletesAt"`
	// Amount of stars the paste has
	Stars       uint64    `json:"stars"`
	// Is the paste private?
	IsPrivate   bool      `json:"isPrivate"`
	// Is the paste public?
	IsPublic    bool      `json:"isPublic"`
	// Is the paste encrypted?
	IsEncrypted bool      `json:"encrypted"`
	// Slices of all tags for this paste
	Tags        []string  `json:"tags"`
	// Slice of all the pasties on the paste
	Pasties     []Pasty   `json:"pasties"`
	// Slice of all edits
	Edits       []Edit    `json:"edits"`
}
// Represents a single pasty, could also be perceived as a "file" 
// on the PasteMyst website, contains language, code, and title.
type Pasty struct {
	// Id of the pasty
	Id       string `json:"_id"`
	// Title of the pasty
	Title    string `json:"title"`
	// Language of the pasty
	Language string `json:"language"`
	// Code of the pasty
	Code     string `json:"code"`
}

// Holds information about a given edit based in 'id'
type Edit struct {
	// Unique id of the edit
	Id       string   `json:"_id"`
	// Edit id, multiple edits can share the same id
	// to show that multiple properties were edited
	// at the same time
	EditId   string   `json:"editId"`
	// Type of edit (incomplete)
	EditType uint64   `json:"editType"`
	// Various metadata, most used case - storing which pasty was edited
	Metadata []string `json:"metadata"`
	// The actual data of the edit, typically stores old data
	Edit     string   `json:"edit"`
	// Unix time of when the edit was executed
	EditedAt uint64   `json:"editedAt"`
}

// Information needed to created a new pasty
type PastyCreateInfo struct { 
	// Title of pasty
	Title    string `json:"title"`
	// Language of the pasty, stores the name of the language,
	// not the mode or MIME type.
	Language string `json:"language"`
	// Code of the pasty
	Code     string `json:"code"`
}

// Information needed to create a new paste
type PasteCreateInfo struct { 
	// Title -- optional
	Title     string            `json:"title"`
	// ExpiresIn -- optional
	ExpiresIn string            `json:"expiresIn"`
	// Is it accessible by the owner? -- optional
	IsPrivate bool              `json:"isPrivate"`
	// Is it displayed on the owners public profile? -- optional
	IsPublic  bool              `json:"isPublic"`
	// Tags, comma separated -- optional
	Tags      string            `json:"tags"`
	// List of pasties -- mandatory
	Pasties   []PastyCreateInfo `json:"pasties"`
}

// Gets a paste based on Id, a token is mandatory for accessing private pastes
//  
// Returns:
//  (*Paste, error)
func (c *Client) GetPaste(id string) (*Paste, error) {
	url := PasteEndpoint + id
	client := NewClient(c.Token) 

	var paste Paste
	err := client.get(url, &paste)
	if err != nil {
		return nil, sadness("%v", err)
	}

	return &paste, nil
}

// Creates a new paste with the given PasteCreateInfo
// 
// Posts new pastes to (https://paste.myst.rs/api/v2/paste)
//  
// Returns:
//  (*Paste, error)
func (c *Client) CreatePaste(createInfo PasteCreateInfo) (*Paste, error) { 	
	// There's no sense bothering with anything else if these checks fail
	// IsPrivate, IsPublic, and Tags are related to account features, if no token is passed
	// then these flags aren't allowed to be true. 
	if (createInfo.IsPrivate || createInfo.IsPublic || createInfo.Tags != "") && c.Token == "" {
		return nil, errors.New("Error: Cannot use account features without a valid token.")
	}

	// url for where the paste will go
	url := BaseEndpoint + "paste/"
	client := NewClient(c.Token) 

	var paste Paste
	err := client.post(url, createInfo, &paste)
	if err != nil { 
		return nil, sadness("%v", err)
	}
	
	return &paste, nil
}

// Deletes a paste with a specified account token -- mandatory
//
// You can only delete pastes on the account of the token that has been passed.
// 
// A token is required for deleting a paste because this is an account feature.
// 
// This action is irreversible.
//  
// Returns:
//  (error)
func (c *Client) DeletePaste(id string) error { 
	url := PasteEndpoint + id
	
	client := NewClient(c.Token) 
	ok, err := client.delete(url, &Paste{})

	if !ok || err != nil {
		return sadness("Unable to delete paste\n%v", err)
	}

	return nil
}

// Edits a paste with a specified account token -- mandatory
//
// You can only edit pastes on the account of the token that has been passed.
// 
// A token is required for editing a paste because this is an account feature.
// 
// To edit values of a paste you must send back the exact same paste except with the 
// adjusted values, you cannot edit expiration date, any result will have no effect.
//  
// Returns:
//  (*Paste, error)
func (c *Client) EditPaste(paste Paste) (*Paste, error) { 
	// url for where the paste will go
	url := PasteEndpoint + paste.Id
	client := NewClient(c.Token) 

	err := client.patch(url, &paste)
	if err != nil {
		return nil, sadness("%v", err)
	}

	return &paste, nil
}