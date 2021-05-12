package pastemystgo

// Language represents a language request
type Language struct {
	// Represents the name of the language
	Name       string   `json:"name"`
	// Language mode for the online editor (codemirror)
	Mode       string   `json:"mode"`
	// All supported mimes in a slice
	Mimes      []string `json:"mimes"`
	// All extensions for a language with a given name
	Extensions []string `json:"ext"`
	// Color of the language, not gauranteed for every language,
	// Default will be #FFFFFF if the language doesn't have one.
	Color      string   `json:"color"`
}