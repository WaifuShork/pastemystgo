package pastemystgo

// A collection of extension tools used throughout the pastemystgo library
import (
	"encoding/json"
	"errors"
	"fmt"
)

// A helper function which wraps json.Unmarshal for readability
//
// Returns:
//  ([]byte, error)
func DeserializeJson(bytes []byte, v interface{}) (error) {
	return json.Unmarshal(bytes, &v)
}


// A helper function which wraps json.MarshalIndent and json.Marshal for readability
// providing the user with a choice to indent or not.
//
// Returns:
//  ([]byte, error)
func SerializeJson(v interface{}, isIndented bool) ([]byte, error) {
	if isIndented {
		return json.MarshalIndent(&v, "", "    ")
	} else { 
		return json.Marshal(&v)
	}
}

// Wraps errors.New(error) so you can easily throw a new error without 
// dealing with formatting a message before feeding it to the method.
//
// Returns:
//  error
func sadness(message string, err ...interface{}) error {
	errorMessage := fmt.Sprintf(message, err...)
	return errors.New(errorMessage)
}