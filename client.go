package pastemystgo

import (
	"bytes"
	"encoding/json"
	"net/http"
)


type RequestMethod int 
const (
	GET RequestMethod = iota
	POST
	PATCH
	DELETE
)
// Represents a backend client with a token
// used for registration of a new context
type Client struct {
	Token string
}

// Registers a new Client to use in the backend for API operations.
//
// Returns:
//  (*Client)
func NewClient(token string) *Client { 
	return &Client{
		Token: token,
	}
}

// Marks a Client for deletion by assigning it to nil,
// allowing it to be handled by the Garbage Collector.
//
// Remarks: This function should be called when you're done 
// using the client, ensure that cleanup is executed.
func (c *Client) DeleteClient() { 
	c = nil
}

func (c *Client) get(url string, pattern interface{}) (error) { 
	response, err := c.makeRequest(url, GET, nil, &pattern)
	if err != nil { 
		sadness("%v", err)
	}

	if response.StatusCode != http.StatusOK { 
		sadness("Error: Expected StatusOK\nGot: %v", response.StatusCode)
	}

	return nil	
}

func (c *Client) post(url string, body interface{}, pattern interface{}) error { 
	response, err := c.makeRequest(url, POST, body, &pattern)
	if err != nil { 
		return sadness("%v", err)
	}
	
	if response.StatusCode != http.StatusOK { 
		return sadness("Error: Expected StatusOK\nGot: %v", response.StatusCode)
	}

	return nil
}

func (c *Client) patch(url string, body interface{}) (error) {
	response, err := c.makeRequest(url, PATCH, body, nil)
	if err != nil { 
		return sadness("%v", err)
	}

	if response.StatusCode != http.StatusOK { 
		return sadness("StatusCode was not 200\nStatusCode: %v", response.StatusCode)
	}

	return nil
}

func (c *Client) delete(url string, pattern interface{}) (bool, error) { 
	response, err := c.makeRequest(url, DELETE, nil, nil)
	if err != nil {
		return false, sadness("%v", err)
	}

	return response.StatusCode == http.StatusOK, nil
}

func (c *Client) makeRequest(url string, method RequestMethod, body interface{}, outPattern interface{}) (*http.Response, error) { 
	reqMethod := c.getRequestMethod(method)
	// endpointUrl := BaseEndpoint + url

	jsonBody := &bytes.Buffer{}
	json.NewEncoder(jsonBody).Encode(&body)
	// It's possible body to be nil, considering not everything requires a pattern body. 
	request, err := http.NewRequest(reqMethod, url, bytes.NewBuffer(jsonBody.Bytes()))
	if err != nil { 
		return nil, sadness("%v", err)
	}

	request.Header.Add("Content-Type", "application/json")

	if c.Token != "" { 
		request.Header.Add("Authorization", c.Token)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil { 
		return nil, sadness("%v", err)
	}

	if body != nil || outPattern != nil { 
		err = c.bodyToJson(response, &outPattern)
		if err != nil {
			return nil, sadness("%v", err)
		}
	}

	defer response.Body.Close()
	return response, nil
}

func (c *Client) getRequestMethod(method RequestMethod) string {
	switch method { 
	case GET: 
		return "GET"
	case POST:
		return "POST"
	case PATCH:
		return "PATCH"
	case DELETE:
		return "DELETE"
	default: 
		// http.NewRequest() should default to "GET" but you should still,
		// specify a method for clarity on call.
		return ""
	}
}