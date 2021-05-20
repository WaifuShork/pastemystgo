<h1 align="center">User</h1>

User contains several functions related to ensuring a user exists, getting a user via username, and getting a collection of the users pastes or paste ids.

> Note: 
> 
> Ensure that you have an active client. 

Golang struct representation of a User:
```go
type User struct {
	Id              string 
	Username        string 
	AvatarUrl       string 
	DefaultLang     string 
	PublicProfile   bool   
	SupporterLength uint64 
	IsContributor   bool   
}
```

Ensure user exists:
```go
func (c *Client) UserExists(username string) (bool, error)
```
```go
ok, err := client.UserExists("waifushork")
if err != nil || !ok { 
    panic(err)
}
```

Get user with a username:
```go
func (c *Client) GetUser(username string) (*User, error)
func (c *Client) TryGetUser(username string) (*User, bool)
```
```go
user, err := client.GetUser("waifushork")
if err != nil {
    panic(err)
}

user, ok := client.TryGetUser("waifushork")
if !ok { 
    fmt.Errorf("unable to get user")
}
```
Get self user: 
```go
func (c *Client) GetSelfUser() (*User, error)
func (c *Client) TryGetSelfUser() (*User, bool)
```
```go
user, err := client.GetSelfUser()
if err != nil {
	panic(err)
}

user, ok := client.TryGetSelfUser()
if !ok {
	fmt.Errorf("unable to get user")
}
```

Get self user pastes:
```go
func (c *Client) GetSelfPasteIds() ([]string, error)
func (c *Client) TryGetSelfPasteIds() ([]string, bool)

func (c *Client) GetSelfPastes() ([]*Paste, error)
func (c *Client) TryGetSelfPastes() ([]*Paste, book)

func (c *Client) GetSelfPastesByAmount(amount uint) ([]*Paste, error)
func (c *Client) TryGetSelfPastesByAmount(amount uint) ([]*Paste, bool)

func (c *Client) GetSelfPasteIdsByAmount(amount uint) ([]string, error)
func (c *Client) TryGetSelfPasteIdsByAmount(amount uint) ([]string, bool)
```
```go
pastes, err := client.GetSelfPasteIds()
if err != nil { 
	panic(err)
}
pastes, ok := client.TryGetPasteIds()
if !ok { 
	fmt.Errorf("unable to get paste ids")
}

pastes, err := client.GetSelfPastes()
if err != nil { 
	panic(err)
}
pastes, ok := client.TryGetSelfPastes()
if !ok { 
	fmt.Errorf("unable to get pastes")
}

pastes, err := client.GetSelfPastesByAmount(10) 
if err != nil { 
	panic(err)
}
pastes, ok := client.TryGetSelfPastesByAmount(10)
if !ok { 
	fmt.Errorf("unable to get selected pastes by amount")
}

pastes, err := client.GetSelfPasteIdsByAmount(10) 
if err != nil { 
	panic(err)
}
pastes, ok := client.TryGetSelfPasteIdsByAmount(10)
if !ok { 
	fmt.Errorf("unable to get selected paste ids by amount")
}
```


<h2>Important Remarks</h2>

A user must have a public account to be located or retrieved.


| [data](data.md) | [paste](paste.md) | [time](time.md)