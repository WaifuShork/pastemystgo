<h1 align="center">Data</h1>

Data contains several crucial functions related to fetching languages. You can get a language by either it's name, or extension. You can get the full list of languages and extensions [here](https://github.com/CodeMyst/pastemyst/blob/main/data/languages.json). 

> Note: 
> 
> Ensure that you have an active client. 

Golang struct representation of a Language:
```go
type Language struct {
	Name       string   
	Mode       string   
	Mimes      []string 
	Extensions []string 
	Color      string   
}
```

Getting a language by its name.
```go
func (c *Client) GetLanguageByName(name string) (*Language, error)
func (c *Client) TryGetLanguageByName(name string) (*Language, bool)
```
```go
language, err := client.GetLanguageByName("Go")
if err != nil {
    panic(err)
}

language, ok := client.TryGetLanguageByName("Go")
if !ok { 
    fmt.Errorf("unable to get language by name")
}
```

Getting a language by its extension.
```go
func (c *Client) GetLanguageByExtension(extension string) (*Language, error)
func (c *Client) TryGetLanguageByExtension(extension string) (*Language, bool)
```
```go
language, err := client.GetLanguageByExtension("go")
if err != nil { 
    panic(err)
}

language, ok := client.TryGetLanguageByExtension("go")
if !ok {
    fmt.Errorf("unable to get language by extension")
}
```

<h2>Important Remarks</h2>

Please note that getting an extension by name must be done by its "pretty name", such as "C#" or "C++", and extension must be in lowercase. 

> Known Bugs:
> 
> Some languages will not be retrievable.

| [paste](paste.md) | [time](time.md) | [user](user.md) |