package pastemystgo

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BaseEndpoint  string =  `https://paste.myst.rs/api/v2/`
	DataEndpoint  string =  BaseEndpoint + `data/`
	TimeEndpoint  string =  BaseEndpoint + `time/`
	UserEndpoint  string =  BaseEndpoint + `user/`
	PasteEndpoint string =  BaseEndpoint + `paste/`

	DataLanguageByName string = DataEndpoint + `language?name=`
	DataLanguageByExt  string = DataEndpoint + `languageExt?extension=`
	
	TimeExpiresInToUnix string = TimeEndpoint + `expiresInToUnixTime`
)

// Gets a language based on its pretty name:
//
// Uses url encoding to convert it into a url friendly value
// returns a Language and error if applicable. 
// 
// Language will be nil if error is returned. 
//
// Returns:
//  (*Language, error)
// BUG(r): Some languages will not return properly, and will error out.
func GetLanguageByName(endpoint, value string) (*Language, error) {
	// Request the language from the API endpoint
	response, err := http.Get(endpoint + url.QueryEscape(value))
	if err != nil { 
		return nil, err
	}
	var language Language

	// Read the responses body to get the raw text 
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		return nil, err
	}

	err = DeserializeJson(bytes, &language)
	// err = json.Unmarshal(bytes, &language)
	if err != nil { 
		return nil, err
	}

	return &language, nil
}

// Gets a language based on its extension:
//
// Wraps GetLanguageByName() to get the given language based on any
// extension that is applicable to the language.
//
// Returns:
//  (*Language, error)
// BUG(r): Some languages will not return properly, and will error out.
func GetLanguageByExtension(extension string) (*Language, error) { 

	return GetLanguageByName(DataLanguageByExt, extension)
}