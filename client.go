package pastemystgo

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)


type Client struct {
	Token *string
}

type RequestMethod int
const (
	Get RequestMethod = iota
	Post 
	Patch
	Delete
)

func (c *Client) Get(url string, pattern interface{}) (interface{}, error) {
	// A specific type to get must be specified
	if pattern == nil { 
		return nil, nil
	}

	// This won't work
	if c.Token == nil { 
		return nil, nil
	}

	// Request the language from the API endpoint
	response, err := http.Get(url)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	// Ensure StatusCode 200
	if response.StatusCode != http.StatusOK {
		return nil, err
	}

	// Get the body of the page
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	// Deserialize the specific pattern from the page contents 
	err = json.Unmarshal(bytes, &pattern)
	if err != nil { 
		log.Fatalf("%v", err)
		return nil, err
	}

	// Cowabunga idk what else to check
	return &pattern, nil
}

func (c *Client) MakeRequest(url string, method RequestMethod, body interface{}) (*http.Response, interface{}) { 
	// request, err := http.NewRequest(getRequestType(method), url, nil)
	return nil, nil
}

func getRequestType(rt RequestMethod) string { 
	switch rt { 
	case Get:
		return "GET"
	case Post:
		return "POST"
	case Patch: 
		return "PATCH"
	case Delete:
		return "DELETE"
	default: 
		return ""
	}
}