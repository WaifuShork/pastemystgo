package pastemystgo

// A collection of extension tools used throughout the pastemystgo library
import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// A helper function which wraps json.Unmarshal for readability
//
// Returns:
//  (error)
func (c *Client) deserializeJson(bytes []byte, v interface{}) (error) {
	return json.Unmarshal(bytes, &v)
}

// A helper function which wraps json.MarshalIndent and json.Marshal for readability
// providing the user with a choice to indent or not.
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
// Reference via &pattern to alter the reference instead of throwing it away after the Goroutine
// completes its execution cycle.
//
// Returns:
//  (error)
func (c *Client) bodyToJson(response *http.Response, pattern interface{}) (error) { 
	// Read the responses body to get the raw text
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return sadness("Error reading Response Body\n%v", err)
	}

	err = c.deserializeJson(bytes, &pattern)
	if err != nil {
		return sadness("Error Deserializing the Response Body\n%v", err)
	}
	return nil
}

// Executes a request using the http.Client, with a provided pattern to 'json-ify' into a Golang
// struct just in case the user needs to get the response body back in json form.
//
// Reference via &pattern to alter the reference instead of throwing it away after the Goroutine
// completes its execution cycle.
// 
// Returns:
//  (error)
func (c *Client) postBodyToJson(client http.Client, request *http.Request, pattern interface{}) (error) { 
	// Post the actual request
	response, err := client.Do(request)
	if err != nil { 
		return sadness("Unable to DO request.\n%v", err)
	}

	defer response.Body.Close()
	return c.bodyToJson(response, pattern)
}

// Wraps errors.New(error) so you can easily throw a new error without 
// dealing with formatting a message before feeding it to the method.
//
// Returns:
//  (error)
func sadness(message string, err ...interface{}) error {
	return errors.New(fmt.Sprintln(message, err))
}