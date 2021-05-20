---
title: Intro
sidebar_position: 1
---

Each module consists of related helper functions for accessing the PasteMyst API.
The PasteMystGo API wrapper consists of four main modules, data, time, user, and paste. To get started with a specific module pick one from the list, or continue reading this page.

- [Client](02-client.md)
- [Data](03-data.md)
- [Paste](04-paste.md)
- [Time](05-time.md)
- [User](06-user.md)

## Preamble

When you initially create a PasteMyst session, you'll have the option to provide a "token", the following features are available to use without a token:

- Creating a public paste
- Getting a public paste
- Getting when a paste expires

The following features are unavailable to use without a token:

- Creating a private paste
- Getting a private paste
- Editing a private/public paste
- Deleting a public/private paste

You can get a token for the API here.

:::info Important: 

API access is restricted to 5 requests per second. When you exceed this limit, you will get "StatusCode 429 (too many requests)". Contact [CodeMyst](https://github.com/CodeMyst) to request more.
:::

For more information on the underlying API itself, please refer to the [PasteMyst website](https://paste.myst.rs/api-docs/index).

