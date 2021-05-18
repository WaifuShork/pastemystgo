package pastemystgo

import (
	"net/url"
)

// GetLanguageByName gets a language based on its pretty name:
//
// Uses url encoding to convert it into a url friendly value
// returns a Language and error if applicable.
//
// Language will be nil if error is returned.
//
// Returns:
//  (*Language, error)
// BUG(r): Some languages will not return properly, and will error out.
func (c *Client) GetLanguageByName(endpoint, value string) (*Language, error) {
	// Request the language from the API endpoint
	endpointUrl := endpoint + url.QueryEscape(value)
	var language Language

	err := c.get(endpointUrl, &language)
	if err != nil {
		return nil, newError(err)
	}
	
	return &language, nil
}

// GetLanguageByExtension gets a language based on its extension
//
// Returns:
//  (*Language, error)
// BUG(r): Some languages will not return properly, and will error out.
func (c *Client) GetLanguageByExtension(extension string) (*Language, error) {
	var language Language	
	err := c.get(DataLanguageByExt + extension, &language)
	if err != nil { 
		return nil, newError(err)
	}

	return &language, err
}