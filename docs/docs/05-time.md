---
title: Time
sidebar_position: 5
---

Time contains a single function for converting the expiration time to the unix time format.

:::info 
Ensure that you have an active [Client](02-client.md). 
:::

### Get time until expiration in unix time
```go
func (c *Client) ExpiresInToUnixTime(createdAt uint64, expires ExpiresIn) (uint64, error)
```
```go
time, err := client.ExpiresInToUnixTime(1588441258, pastemyst.OneWeek)
if err != nil {
    panic(err)
}
```