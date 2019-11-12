# API and WebSocket Protocol

## WebSocket Protocol
Every message across the WebSocket must be a JSON Object, that contains the
`event` field.

```JSON
{
  "event": "",
  "...": ""
}
```

### Joining a room
```JSON
{
  "event": "init",
  "room": "code",
  "token": "base64rejoincode",
  "index": 1,
  "users": ["Eddie", null, "Josho", null, null, "Martin", "Elle", "Jordie"],
  "numUsers": 4
}
```
`users` has length equal to the max number of users in a room. Colours are
assigned using the modulo of the index of the user. You have the index `index`
in the array. The `room` is the value for `/room/code` and serves as the link
for inviting friends to the room.

The token is used in case of the socket connection being momentarily dropped.
The client

### Message

```JSON
{
  "event": "message",
  "from": 3,
  "data": "NPXkOU8..."
}
```
When sending from the client to the server, the `from` field is optional.
