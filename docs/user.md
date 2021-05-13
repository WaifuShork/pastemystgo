<h1 align="center">User</h1>

User contains 2 functions related to ensuring a user exists, and getting a user via username.

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
```
```go
user, err := client.GetUser("waifushork")
if err != nil {
    panic(err)
}
```

Get user with a username:
```go
func (c *Client) TryGetUser(username string) (*User, bool, error)
```
```go
func (c *Client) GetUser(username string) (*User, error)
```
```go
user, err := client.GetUser("waifushork")
if err != nil {
    panic(err)
}
```

<h2>Important Remarks</h2>

A user must have a public account to be located or retrieved.