package pastemystgo

// A collection of extension tools used throughout the pastemystgo library
import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// A helper function which wraps json.Unmarshal for readability
//
// Params:
// 	(bytes []byte, v interface{})
//
// Returns:
//  (error)
func (c *Client) deserializeJson(bytes []byte, v interface{}) error {
	return json.Unmarshal(bytes, &v)
}

// A helper function which wraps json.MarshalIndent and json.Marshal for readability
// providing the user with a choice to indent or not.
//
// Params:
// 	(v interface{}, isIndented bool)
//
// Returns:
//  ([]byte, error)
func (c *Client) serializeJson(v interface{}, isIndented bool) ([]byte, error) {
	if isIndented {
		return json.MarshalIndent(&v, "", "    ")
	} else { 
		return json.Marshal(&v)
	}
}

// Converts the response.Body into a logical Golang struct to use
// accepts a reference to a pattern to deserialize based upon.
// 
// Params:
// 	(response *http.Response, pattern interface{})
//
// Returns:
//  (error)
func (c *Client) bodyToJson(response *http.Response, pattern interface{}) error {
	// Read the responses body to get the raw text
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err//newErrorf("error reading response body\n%v", err)
	}

	err = c.deserializeJson(bytes, &pattern)
	if err != nil {
		return err//newErrorf("error deserializing the Response Body\n%v", err)
	}
	return nil
}

// Executes a request using the http.Client, with a provided pattern to 'json-ify' into a Golang
// struct just in case the user needs to get the response body back in json form.
//
// Params:
// 	(client http.Client, request *http.Request, pattern interface{})
// 
// Returns:
//  (error)
func (c *Client) postBodyToJson(client http.Client, request *http.Request, pattern interface{}) error {
	// Post the actual request
	response, err := client.Do(request)
	if err != nil { 
		return err//newErrorf("unable to do request.\n%v", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			//newErrorf("error attempting to close the body reader\n%v", err)
		}
	} (response.Body)

	return c.bodyToJson(response, pattern)
}

// Wraps errors.New(error) so you can easily throw a new error without 
// dealing with formatting a message before feeding it to the method.
//
// Params:
// 	(message string, err ...interface{})
//
// Returns:
//  (error)
func newErrorf(message string, err ...interface{}) error {
	return errors.New(fmt.Sprintln(message, err))
}

// Wraps newErrorf for simplicity. Instead of being forced to string
// format on every call, you can simply pass the error message.  
// Properly prints the value of the error with default formatting as well.  
//
// Params:
// 	(err error)
//
// Returns:
//  (error)
func newError(err error) error {
	return newErrorf("%v", err)
}