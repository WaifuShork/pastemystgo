---
title: Client
sidebar_position: 2
---

The Client is a back-end system that PasteMystGo utilizes to ensure a session keeps a configuration, namely an authorization token.  

## The Backend Client

PasteMystGo is powered by a back-end client/runner, this system is responsible for ensuring safety among API access.

```go
client := pastemystgo.NewClient("API-Token")
```

Once you've registered a client, you can access all the PasteMystGo API features available without worry. Passing an empty string for the token means you won't get access to the account related features as mentioned above. It's best to declare your client in the global scope to ensure you only ever have one client active. To mark an active client for deletion, you may do the following:

```go
pastemystgo.DeleteClient(client)
```
It is recommended to delete a client when you're no longer using it. If you create a client while inside a Goroutine, there's no need to delete it manually, since the Goroutine will handle deletion.