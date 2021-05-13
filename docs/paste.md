<h1 align="center">Paste</h1>

Paste contains 4 crucial functions related to creating, deleting, editing, and getting pastes. If you wish to create/edit a private paste, or delete a private/public paste, you need to be using an API key as these are restricted to account features.

> Note: 
> 
> Ensure that you have an active client. 

Golang struct representation of a Paste:
```go
type Paste struct {
	Id          string    
	OwnerId     string   
	Title       string    
	CreatedAt   uint64    
	ExpiresIn   string    
	DeletesAt   uint64    
	Stars       uint64   
	IsPrivate   bool      
	IsPublic    bool     
	IsEncrypted bool     
	Tags        []string  
	Pasties     []Pasty   
	Edits       []Edit    
}
```

Getting a paste using an Id:
```go
func (c *Client) GetPaste(id string) (*Paste, error)
```
```go
paste, err := client.GetPaste("sewevxee")
if err != nil {
    panic(err)
}
```

Creating a paste from scratch:
```go
func (c *Client) CreatePaste(createInfo PasteCreateInfo) (*Paste, error)
```
```go
pastyCreateInfo := []pastemystgo.PastyCreateInfo{
    {
        Title: "pasty1",
        Language: "plain text",
        Code: "asd asd asd",
    },
}

createInfo := pastemystgo.PasteCreateInfo{
    Title:     "api test paste",
    ExpiresIn: "never",
    IsPrivate: false,
    IsPublic:  false,
    Tags:      "",
    Pasties: pastyCreateInfo,
}

paste, err := client.CreatePaste(createInfo)
if err != nil { 
    panic(err)
}
```

Deleting a paste:
```go
func (c *Client) DeletePaste(id string) (error)
```
```go
paste, err := client.GetPaste("sewevxee")
if err != nil { 
    panic(err)
}

err := client.DeletePaste(paste.Id)
if err != nil { 
    panic(err)
}
```

Editing a paste:
```go
func (c *Client) EditPaste(paste Paste) (*Paste, error)
```
```go
paste, err := client.GetPaste("sewevxee")
if err != nil { 
    panic(err)
}

paste.Title = "edited title"

editedPaste, err := client.EditPaste(*paste)
if err != nil {
    panic(err)
}
```