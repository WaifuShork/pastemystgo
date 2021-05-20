---
title: Paste
sidebar_position: 4
---

Paste contains several crucial functions related to creating, deleting, editing, and getting pastes. If you wish to create/edit a private paste, or delete a private/public paste, you need to be using an API key as these are restricted to account features.

:::info 
Ensure that you have an active [Client](02-client.md). 
:::

### Paste struct
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

### Getting a paste
```go
func (c *Client) GetPaste(id string) (*Paste, error)
func (c *Client) TryGetPaste(id string) (*Paste, book)
```
```go
paste, err := client.GetPaste("sewevxee")
if err != nil {
    panic(err)
}

paste, ok := client.TryGetPaste("sewevxee")
if !ok { 
    fmt.Errorf("unable to get paste")
}
```

### Creating a paste
```go
func (c *Client) CreatePaste(createInfo PasteCreateInfo) (*Paste, error)
func (c *Client) TryCreatePaste(createInfo PasteCreateInfo) (*Paste, bool)
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

paste, ok := client.TryCreatePaste(createInfo)
if !ok {
    fmt.Errorf("unable to create paste")
}
```

### Deleting a paste
```go
func (c *Client) DeletePaste(id string) error
func (c *Client) TryDeletePaste(id string) bool
```
```go
err := client.DeletePaste("sewevxee")
if err != nil { 
    panic(err)
}

ok := client.TryDeletePaste("sewevxee")
if !ok { 
    fmt.Errorf("unable to delete paste")
}
```

### Bulk deleting pastes
```go
func (c *Client) BulkDeletePastes(pastes []string) error
func (c *Client) TryBulkDeletePastes(pastes []string) book
```
```go
err := client.BulkDeletePastes("sewevxee", "i3dcx8ab", "g36wu5to")
if err != nil {
    panic(err)
}

ok := client.TryBulkDeletePastes("sewevxee", "i3dcx8ab", "g36wu5to")
if !ok {
    fmt.Errorf("unable to delete pastes")
}
```

### Editing a paste
```go
func (c *Client) EditPaste(paste Paste) (*Paste, error)
func (c *Client) TryEditPaste(paste Paste) (*Paste, book)
```
```go
paste, err := client.GetPaste("sewevxee")
if err != nil { 
    panic(err)
}

paste.Title = "edited title"

editedPaste, err := client.EditPaste(paste)
if err != nil {
    panic(err)
}

editedPaste, ok := client.TryEditPaste(paste)
if !ok { 
    fmt.Errorf("unable to edit paste")
}
```

### ExpiresIn enum
```go
// ExpiresIn represents all the possible time formats
// for when a paste will expire.
type ExpiresIn int
const (
	Never ExpiresIn = iota // string form -> "never"
	OneHour  // string form -> "1h"
	TwoHours // string form -> "2h"
	TenHours // string form -> "10h"
	OneDay   // string form -> "1d"
	TwoDays  // string form -> "2d"
	OneWeek  // string form -> "1w"
	OneMonth // string form -> "1m"
	OneYear  // string form -> "1y"
)
```