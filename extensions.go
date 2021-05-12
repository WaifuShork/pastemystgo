package pastemystgo

import (
	"encoding/json"
	"errors"
	"fmt"
)

// A helper function which wraps Unmarshal for readability
func deserializeJson(bytes []byte, v interface{}) (error) {
	return json.Unmarshal(bytes, &v)
}

// A helper function which wraps MarshalIndent and Marshal for readability
// providing the user with a choice to indent or not.
func serializeJson(v interface{}, isIndented bool) ([]byte, error) {
	if isIndented {
		return json.MarshalIndent(&v, "", "    ")
	} else { 
		return json.Marshal(&v)
	}
}

// Pretty much a joke function for returning errors
func sadness(message string, err ...error) error {
	errorMessage := fmt.Sprintf(message, err)
	return errors.New(errorMessage)
}