package pastemystgo

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type RequestMethod int 
const (
	GET RequestMethod = iota
	POST
	PATCH
	DELETE
)

// Client represents a backend client with a token
// used for registration of a new context
type Client struct {
	Token *string
}

// NewClient registers a new Client to use in the backend for API operations.
//
// Returns:
//  (*Client)
func NewClient(tok string) *Client {
	var token *string
	token = &tok

	return &Client{
		Token: token,
	}
}

// DeleteClient marks a Client for deletion by assigning it to nil,
// allowing it to be handled by the Garbage Collector.
//
// Params:
//  (client *Client)
//
// Remarks: This function should be called when you're done 
// using the client, ensure that cleanup is executed.
func DeleteClient(client *Client) {
	client = nil
}

func (c *Client) get(url string, pattern interface{}) error {
	response, err := c.makeRequest(url, GET, nil, &pattern)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK { 
		return err
	}

	return nil
}

func (c *Client) post(url string, body interface{}, pattern interface{}) error {
	response, err := c.makeRequest(url, POST, body, &pattern)
	if err != nil { 
		return err
	}
	
	if response.StatusCode != http.StatusOK { 
		return err
	}

	return nil
}

func (c *Client) patch(url string, body interface{}) error {
	response, err := c.makeRequest(url, PATCH, body, nil)
	if err != nil { 
		return err
	}

	if response.StatusCode != http.StatusOK { 
		return err
	}

	return nil
}

func (c *Client) delete(url string) (bool, error) {
	response, err := c.makeRequest(url, DELETE, nil, nil)
	if err != nil {
		return false, err
	}

	return response.StatusCode == http.StatusOK, nil
}

// makeRequest is the back-end for all client actions
//
// Params:
// 	(url string, method RequestMethod, body interface{}, outPattern interface{})
//
// Returns:
//  (*http.Response, error)
func (c *Client) makeRequest(url string, method RequestMethod, body interface{}, outPattern interface{}) (*http.Response, error) {
	reqMethod := c.getRequestMethod(method)

	// Converts the `body interface{}` into a buffer for feeding NewRequest the bytes.
	jsonBody := &bytes.Buffer{}
	if body != nil {
		err := json.NewEncoder(jsonBody).Encode(&body)
		if err != nil {
			return nil, err
		}
	}

	// It's possible body to be nil, considering not everything requires a pattern body.
	request, err := http.NewRequest(reqMethod, url, bytes.NewBuffer(jsonBody.Bytes()))
	if err != nil { 
		return nil, err
	}

	// Apply headers
	request.Header.Add("Content-Type", "application/json")
	if *c.Token != "" || c.Token != nil {
		request.Header.Add("Authorization", *c.Token)
	}

	// Execute actual request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil { 
		return nil, err
	}

	if body != nil || outPattern != nil {
		err = c.bodyToJson(response, &outPattern)
		if err != nil {
			return nil, err
		}
	}

	// P A N I C
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	} (response.Body)

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

func (c *Client) IsAuthorized() bool {
	if *c.Token == "" || c.Token == nil {
		return false
	}
	return true
}