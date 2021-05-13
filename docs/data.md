<h1 align="center">Data</h1>

Data contains 2 crucial functions related to fetching languages. You can get a language by either it's name, or extension. You can get the full list of languages and extensions [here](https://github.com/CodeMyst/pastemyst/blob/main/data/languages.json). 

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
func (c *Client) GetLanguageByName(endpoint, value string) (*Language, error)
```
```go
language, err := client.GetLanguageByName("Go")
if err != nil {
    panic(err)
}
```

Getting a language by its extension.
```go
func (c *Client) GetLanguageByExtension(extension string) (*Language, error)
```
```go
language, err := client.GetLanguageByExtension("go")
if err != nil { 
    panic(err)
}
```

<h2>Important Remarks</h2>

Please note that getting an extension by name must be done by it's "pretty name", such as "C#" or "C++", and extension must be in lowercase. 

> Known Bugs:
> 
> Some languages will not be fetchable.