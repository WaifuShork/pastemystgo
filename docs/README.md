<h1 align="center">Welcome to PasteMystGo</h1>

Each module consists of related helper functions for accessing the PasteMyst API.
The PasteMystGo API wrapper consists of four main modules, data, time, user, and paste. To get started with a specific module pick one from the list, or continue reading this page.

* [Data](data.md)
* [Time](time.md)
* [User](user.md)
* [Paste](paste.md)

<h2>Preamble</h2>

When you initially create a pastemyst session, you'll have the option to provide a "token", the following features are available to use without a token: 
* Creating a public paste
* Getting a public paste
* Getting when a paste expires 

The following features are unavailable to use without a token:
* Creating a private paste
* Getting a private paste
* Editing a private/public paste
* Deleting a public/private paste

You can get a token for the API [here](https://paste.myst.rs/user/settings). 

> Important Note: API access is restricted to 5 requests per second. When you exceed this limit, you will get "StatusCode 429 (too many requests)". Contact [CodeMyst](https://github.com/CodeMyst) to request more.  

For more information on the underlying API itself, please refer to the [pastemyst website](https://paste.myst.rs/api-docs/index). 

<h2>The Backend Client</h2>

PasteMystGo is powered by a back-end client/runner, this system is responsible for ensuring safety among API access. 

```go
client := pastemystgo.NewClient("API-Token")
```

Once you've registered a client, you can access all the PasteMystGo API features available without worry. Passing an empty string for the token means you won't get access to the account related features as mentioned above. It's best to declare your client in the global scope to ensure you only ever have one client active. To mark an active client for deletion, you may do the following:

```go
client.DeleteClient()
```

It is recommended to delete a client when you're no longer using it. If you create a client while inside a Goroutine, there's no need to delete it manually, since the Goroutine will handle deletion.

