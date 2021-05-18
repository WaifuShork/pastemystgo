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