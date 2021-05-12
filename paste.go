package pastemystgo

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Represents an enumeration of expiration values
// GetExpiresInString(expiresIn ExpiresIn)(string)
// will return the string format of expiration
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

type Paste struct {
	Id        string    `json:"_id"`
	OwnerId   string    `json:"ownerId"`
	Title     string    `json:"title"`
	CreatedAt uint64    `json:"createdAt"`
	ExpiresIn string    `json:"expiresIn"`
	DeletesAt uint64    `json:"deletesAt"`
	Stars     uint64    `json:"stars"`
	IsPrivate bool      `json:"isPrivate"`
	IsPublic  bool      `json:"isPublic"`
	Tags      []string  `json:"tags"`
	Pasties   []Pasty   `json:"pasties"`
	Edits     []Edit    `json:"edits"`
}

type Pasty struct {
	Id       string `json:"_id"`
	Language string `json:"language"`
	Title    string `json:"title"`
	Code     string `json:"code"`
}

type Edit struct {
	Id       string   `json:"_id"`
	EditId   string   `json:"editId"`
	EditType uint64   `json:"editType"`
	Metadata []string `json:"metadata"`
	Edit     string   `json:"edit"`
	EditedAt uint64   `json:"editedAt"`
}

/*type EditType int

const (
	Title EditType = iota
	PastyTitle
	PastyLanguage
	PastyContent
	PastyAdded
	PastyRemoved
)*/

type PastyCreateInfo struct { 
	Title    string `json:"title"`
	Language string `json:"language"`
	Code     string `json:"code"`
}

type PasteCreateInfo struct { 
	Title     string `json:"title"`
	ExpiresIn string `json:"expiresIn"`
	IsPrivate bool   `json:"isPrivate"`
	IsPublic  bool   `json:"isPublic"`
	Tags      string `json:"tags"`
	Pasties   []PastyCreateInfo `json:"pasties"`
}

// Token is allowed to be nil for when not passed
func GetPaste(id string, token string) (*Paste, error) {

	// Request the language from the API endpoint
	response, err := http.Get(PasteEndpoint + id)
	if err != nil {
		return nil, sadness("Error getting endpoint Response\n%v", err)
	}

	// Ensure that the status is found, otherwise there's no reason to continue
	if response.StatusCode == http.StatusNotFound {
		// Print out the type for a prettier error code view.
		return nil, sadness("Incorrect Status Code: (%+t)", errors.New(fmt.Sprintf("%T", response.StatusCode)))
	}

	if token != "" { 
		response.Request.Header.Add("Authorization", token)
	}

	// Read the responses body to get the raw text
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, sadness("Error reading Response Body\n%v", err)
	}

	var paste Paste
	//err = json.Unmarshal(bytes, &paste)
	err = deserializeJson(bytes, &paste)
	if err != nil {
		return nil, sadness("Error Deserializing the Response Body\n%v", err)
	}

	return &paste, nil
}

// Creates a new paste with the given PasteCreateInfo
// Posts new pastes to (https://paste.myst.rs/api/v2/paste)
// Returns Paste and an error if applicable
func CreatePaste(createInfo PasteCreateInfo, token string) (*Paste, error) { 	
	// There's no sense bothering with anything else if these checks fail
	// IsPrivate, IsPublic, and Tags are related to account features, if no token is passed
	// then these flags aren't allowed to be true. 
	if (createInfo.IsPrivate || createInfo.IsPublic || createInfo.Tags != "") && token == "" {
		return nil, errors.New("Error: Cannot use account features without a valid token.")
	}

	// url for where the paste will go
	url := BaseEndpoint + "paste/"
	 
	jsonBody := &bytes.Buffer{}
	json.NewEncoder(jsonBody).Encode(&createInfo)

	// The bytes are correct
	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody.Bytes()))
	if err != nil { 
		return nil, sadness("Unable to POST request.")
	}

	// Add headers 
	request.Header.Add("Content-Type", "application/json")
	if token != "" {
		request.Header.Add("Authorization", token)
	}

	// Finish request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil { 
		return nil, sadness("Unable to DO request.\n%v", err)
	}
	defer response.Body.Close()

	// Read the responses body to get the raw text
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, sadness("Error reading Response Body\n%v", err)
	}

	var paste Paste
	err = json.Unmarshal(bytes, &paste)
	if err != nil {
		return nil, sadness("Error Unmarshalling the Response Body\n%v", err)
	}

	return &paste, nil
}

func DeletePaste(id, token string) error { 
	url := PasteEndpoint + id
	request, err := http.NewRequest("DELETE", url, nil)
	if err != nil { 
		return sadness("Unable to delete given paste\n%v", err)
	}

	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Authorization", token)

	// Finish request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil { 
		sadness("Unable to DO request.\n%v", err)
	}
	//log.Fatal(response.Body)
	defer response.Body.Close()
	return nil
}

func EditPaste(paste Paste, token string) (*Paste, error) { 
	// url for where the paste will go
	url := PasteEndpoint + paste.Id
	
	jsonBytes := &bytes.Buffer{}
	json.NewEncoder(jsonBytes).Encode(&paste)

	// The bytes are correct
	request, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonBytes.Bytes()))
	if err != nil { 
		return nil, sadness("Error on POST\n%v", err)
	}

	// Finish request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil { 
		log.Fatal(err)
	}

	// Close body.
	defer response.Body.Close()
	return &paste, nil
}