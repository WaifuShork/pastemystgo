package pastemystgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// const baseEndpoint = "https://paste.myst.rs"

const (
	BaseEndpoint string =  `https://paste.myst.rs/api/v2/`
	DataEndpoint string =  BaseEndpoint + `data/`
	TimeEndpoint string =  BaseEndpoint + `time/`
	UserEndpoint string =  BaseEndpoint + `user/`
	PasteEndpoint string = BaseEndpoint + `paste/`

	DataLanguageByName string = DataEndpoint + `language?name=`
	DataLanguageByExt string = DataEndpoint + `languageExt?extension=`
	
	TimeExpiresInToUnix string = TimeEndpoint + `expiresInToUnixTime`
)

func getLanguageByName(endpoint, value string) (*Language, error) {
	
	// Request the language from the API endpoint
	response, err := http.Get(endpoint + value)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	// Read the responses body to get the raw text 
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}
	
	var language Language
	err = json.Unmarshal(bytes, &language)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	return &language, nil
}

func getLanguageByExtension(extension string) (*Language, error) { 

	response, err := http.Get(DataLanguageByExt + extension)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	var language Language
	err = json.Unmarshal(bytes, &language)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	return &language, nil
}