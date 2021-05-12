<p align="center">
    <a href="https://paste.myst.rs/">
        <img width="450" src="logo.svg" alt="Logo" /><br>
    </a>
</p>

<h1 align="center"> 
    <a href="https://github.com/WaifuShork/pastemyst-go/actions">
        <img alt="Build status" src="https://img.shields.io/github/workflow/status/WaifuShork/pastemyst-go/Go?label=BUILD%20STATUS&logo=github&style=for-the-badge">
    </a>
    <a href="https://github.com/WaifuShork/pastemyst-go/issues">
        <img alt="GitHub issues" src="https://img.shields.io/github/issues/WaifuShork/pastemyst-go?style=for-the-badge&logo=github">
    </a>
</h1>

<p align="center">
A small, fast, reliable pastemyst API wrapper written in Golang. </br>
Official pastemyst API docs found <a href="https://paste.myst.rs/api-docs/index">here</a>. If you'd like enhanced documentation please visit my <a href="https://paste.myst.rs/sewevxee">collection of pastes</a>
</p>

<h2> 
Data
</h2>

```go
// Get language by name
language, err := pastemystgo.GetLanguageByName(endpoint, "C#")
if err != nil {
    panic(err)
}
fmt.Println(language.Name)

// Get language by extension
language, err := pastemystgo.GetLanguageByExtension("cs")
if err != nil { 
    panic(err)
}

fmt.Println(language.Name)
```

<h2> 
Time
</h2>

```go
// Get the time till expiration in Unix time format
time, err := pastemystgo.ExpiresInToUnixTime(1_588_441_258, OneWeek)
if err != nil {
    panic(err)
}

fmt.Printf("%d", time)
```

<h2> 
User
</h2>

```go
// Check whether a user exists
ok, err := pastemystgo.UserExists("waifushork")
if err != nil || !ok { 
    panic(err)
}

fmt.Println("%t", ok)

// Get user from username
user, err := pastemystgo.GetUser("waifushork")
if err != nil {
    panic(err)
}

fmt.Println(user.Username)
```

<h2> 
Paste
</h2>

```go
// Getting a paste
paste, err := pastemystgo.GetPaste("sewevxee", "")
if err != nil {
    panic(err)
}
fmt.Println(paste.Title)

// Creating a paste
pastyCreateInfo := []PastyCreateInfo{
    {
        Title: "pasty1",
        Language: "plain text",
        Code: "asd asd asd",
    },
}

createInfo := PasteCreateInfo{
    Title:     "api test paste",
    ExpiresIn: "never",
    IsPrivate: false,
    IsPublic:  false,
    Tags:      "",
    Pasties: pastyCreateInfo,
}

// No token needs to be passed for a public paste
paste, err := CreatePaste(createInfo, "")
if err != nil { 
    panic(err)
}

fmt.Println(paste.Title)
```
```go
// Creating a private paste
pastyCreateInfo := []PastyCreateInfo{
    {
        Title: "pasty1",
        Language: "plain text",
        Code: "asd asd asd",
    },
}

createInfo := PasteCreateInfo{
    Title:     "api test paste",
    ExpiresIn: "never",
    IsPrivate: true,
    IsPublic:  false,
    Tags:      "",
    Pasties: pastyCreateInfo,
}
fmt.Println(paste.Title)
fmt.Println("%t", paste.IsPrivate)

// Token MUST be passed for private paste.
paste, err := CreatePaste(createInfo, token)
if err != nil { 
    panic(err)
}

fmt.Println(paste.Title)
```
```go
// Deleting a paste
paste, err := GetPaste("sewevxee", token)
if err != nil { 
    panic(err)
}

// A token must be passed because you can only delete your own pastes
err := DeletePaste(paste.Id, token)
if err != nil { 
    panic(err)
}
```
```go
// Editing a paste
paste, err := GetPaste("sewevxee", token)
if err != nil { 
    panic(err)
}

paste.Title = "edited title"

// A token must be passed because you can only edit your own pastes
editedPaste, err := EditPaste(*paste, token)
if err != nil {
    panic(err)
}

fmt.Println(editedPaste.Title)
```