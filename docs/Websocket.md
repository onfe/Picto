# Websocket Protocol

This document details the expected data structure of information sent down the websocket. 



## Contents

- [Wrapper](#wrapper)
- [Payloads](#Payloads)
- [Error Codes](#ErrorCodes)
- [Message Encoding](#Message Encoding)



## Wrapper

Every message across the WebSocket must be contained in the following wrapper's `payload` field:

Client -> server

```JSON
{
  "event": "",
  "payload": {event data}
}
```

Server -> client

```json
{
  "event": "",
  "time": 1582128345655,
  "payload": {event data}
}
```

The server will be in charge of dating events as it receives them.

`time` is in UNIX time to millisecond precision.



## Payloads

### `init` - Joining a room

```JSON
{
  "RoomID": "id",
  "RoomName": "default",
  "Static": true,
  "UserIndex": 1,
  "Users": ["Eddie", null, "Josho", null, null, "Martin", "Elle", null],
}
```
`Users` has length equal to the max number of users in a room. Colours are assigned using the index of the user. You have the index `UserIndex` in the array. The `RoomID` is the value for `/room/RoomID` and serves as the link for inviting friends to the room.



### `user` - User join/leave

User join:
```JSON
{
  "UserIndex": 1,
  "UserName":"Jordie",
  "Users": ["Eddie", "Jordie", "Josho", null, null, "Martin", "Elle", null],
}
```

User leave:
```JSON
{
  "UserIndex": 1,
  "UserName":"Jordie",
  "Users": ["Eddie", null, "Josho", null, null, "Martin", "Elle", null],
}
```

Including `UserName` may seem a bit redundant but it is necessary in the case of cached join/leave events.



### `message` - Client Message

Client -> Server:

```JSON
{
  "Data": [255, 255, ...],
  "Span": 192
}
```

Server -> Client:

```JSON
{
  "ColourIndex": 2,
  "Sender": "Josh",
  "Data": "///...",
  "Span": 192
}
```

`ColourIndex` was the index of the user that sent the message in the `users` array when they sent it, it differs from `UserIndex` in that there may or may not be the same user that sent the message in that index in the `Users` array. 

See the [Message Encoding](#Message-Encoding) section for a description of the `Data` field.



### `announcement` - message from server

Server -> Client:
```JSON
{
  "Announcement": "Welcome to Picto!",
}
```



### `rename` - Rename Room

Client -> server
```JSON
{
  "RoomName": "Denver Airport"
}
```
Server -> client

```json
{
  "UserName": "Josj",
  "RoomName": "Denver Airport"
}
```

`rename` events are ignored in static rooms.



## Error Codes

| Code | Description                                                  |
| ---- | ------------------------------------------------------------ |
| 4400 | Invalid username provided (too long/empty string etc).       |
| 4404 | The server can't find the room specified.                    |
| 4409 | The server can't add the client to the room (the name is already taken or the room is full). |
| 4503 | The server can't create a room (has reached maximum rooms capacity). |
| 4666 | If this happens, an error isn't being handled appropriately by the server. |



## Message Encoding

A byte is used per pixel.

| Range | Use                                       |
| ----- | ----------------------------------------- |
| 0     | Transparent                               |
| 1     | Black                                     |
| 2-3   | Greyscale (2 = light grey, 3 = dark grey) |
| 4-62  | Rainbow colours                           |
| 63    | RLE encoding start character              |

RLE encoding is `255 [counts] 0 [value]` where the total count is the sum of `counts`, plus 4 (if four or less characters are repeated they're not RLE'd as it'd be less efficient, and 0 is an illegal character in `[counts]`, so we know there's at least 5). The sum of `[counts]` is used as opposed to a product as to avoid having to complicate message checking for illegally large images.

The RLE decoder will cut off as soon as the current count sums to more than the canvas size, or the image length becomes longer than the canvas size.