package pastemystgo

// Language represents a language request
type Language struct {
	// Name represents the name of the language
	Name       string   `json:"name"`
	// Language represents the language mode for the online editor (codemirror)
	Mode       string   `json:"mode"`
	// Mimes represents all supported mimes in a slice
	Mimes      []string `json:"mimes"`
	// Extensions represents all extensions for a language with a given name
	Extensions []string `json:"ext"`
	// Color represents the color language, not guaranteed for every language,
	// Default will be #FFFFFF if the language doesn't have one.
	Color      string   `json:"color"`
}

// GetLanguageByName gets a language based on its pretty name:
//
// Uses url encoding to convert it into a url friendly value
// returns a Language and error if applicable.
//
// Language will be nil if error is returned.
//
// Params:
// 	(name string)
//
// Returns:
//  (*Language, error)
// BUG(r): Some languages will not return properly, and will error out.
func (c *Client) GetLanguageByName(name string) (*Language, error) {
	// Request the language from the API endpoint
	url := DataLanguageByName(name)
	var language Language

	err := c.get(url, &language)
	if err != nil {
		return nil,err// newError(err)
	}

	return &language, nil
}

// GetLanguageByExtension gets a language based on its extension
//
// Params:
// 	(extension string)
//
// Returns:
//  (*Language, error)
// BUG(r): Some languages will not return properly, and will error out.
func (c *Client) GetLanguageByExtension(extension string) (language *Language, err error) {
	err = c.get(DataLanguageByExt(extension), &language)
	if err != nil {
		return nil, err//newError(err)
	}

	return language, err
}