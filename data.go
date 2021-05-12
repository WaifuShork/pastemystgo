package pastemystgo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	BaseEndpoint  string =  `https://paste.myst.rs/api/v2/`
	DataEndpoint  string =  BaseEndpoint + `data/`
	TimeEndpoint  string =  BaseEndpoint + `time/`
	UserEndpoint  string =  BaseEndpoint + `user/`
	PasteEndpoint string = BaseEndpoint + `paste/`

	DataLanguageByName string = DataEndpoint + `language?name=`
	DataLanguageByExt  string = DataEndpoint + `languageExt?extension=`
	
	TimeExpiresInToUnix string = TimeEndpoint + `expiresInToUnixTime`
)

func GetLanguageByName(endpoint, value string) (*Language, error) {
	// Request the language from the API endpoint
	response, err := http.Get(endpoint + url.QueryEscape(value))
	if err != nil { 
		return nil, err
	}

	// Read the responses body to get the raw text 
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		return nil, err
	}

	var language Language
	err = json.Unmarshal(bytes, &language)
	if err != nil { 
		return nil, err
	}

	return &language, nil
}

func GetLanguageByExtension(extension string) (*Language, error) { 

	return GetLanguageByName(DataLanguageByExt, extension)
}