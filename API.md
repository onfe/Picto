# API & WebSocket Protocol



## WebSocket Protocol

Every message across the WebSocket must be contained in the following wrapper's `payload` field:

Client -> server

```JSON
{
  "event": "",
  "payload": {event data}
}
```

Server ​​-> client

```json
{
  "event": "",
  "time": 1582128345655,
  "payload": {event data}
}
```

The server will be in charge of dating events as it receives them.

`time` is in UNIX time to millisecond precision.



### `init` - Joining a room

```JSON
{
  "RoomID": "id",
  "RoomName": "default",
  "UserIndex": 1,
  "Users": ["Eddie", null, "Josho", null, null, "Martin", "Elle", null],
}
```
`Users` has length equal to the max number of users in a room. Colours are assigned using the index of the user. You have the index `UserIndex` in the array. The `RoomID` is the value for `/room/code` and serves as the link for inviting friends to the room.



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

`UserIndex` is the index of the user in the users array who changed the room's name.

`rename` events are not cached, so we don't need to worry about `UserIndex` becoming incorrect on user join/leaves.



## WebSocket Errors

| Code | Description                                                  |
| ---- | ------------------------------------------------------------ |
| 4400 | Invalid username provided (too long/empty string etc).       |
| 4404 | The server can't find the room specified.                    |
| 4409 | The server can't add the client to the room (the name is already taken or the room is full). |
| 4503 | The server can't create a room (has reached maximum rooms capacity). |
| 4666 | If this happens, an error isn't being handled appropriately by the server. |



## API - Public

### room_exists

`/api/?method=room_exists&room_id=ROOM_ID`

If `ROOM_ID` exists, returns `true`. Otherwise `false`.



## API - Private

### get_state

`/api/?token=API_TOKEN&method=get_state`
Returns the state of the entire server by marshalling the roomManager. Not wise to use in prod.

### get_room_ids

`/api/?token=API_TOKEN&method=get_room_ids`
Returns a list of all the ids of the currently open rooms.

### get_room_state

`/api/?token=API_TOKEN&method=get_room_state&room_id=ROOM_ID`
Returns the state of the `ROOM_ID` specified by marshalling the room object.

### announce

`/api/?token=API_TOKEN&method=announce&message=MESSAGE`
Announces `MESSAGE` to ALL ROOMS.

`/api/?token=API_TOKEN&method=announce&message=MESSAGE&room_id=ROOM_ID`
Announces `MESSAGE` to the `ROOM_ID` specified.

### close_room

`/api/?token=API_TOKEN&method=close_room&room_id=ROOM_ID&reason=REASON`
Closes `ROOM_ID` and announces message `REASON` beforehand.

### create_static_room

`/api/?token=API_TOKEN&method=create_static_room&room_name=ROOM_NAME&room_size=ROOM_SIZE`
Creates a static room (continues to exist when there are no clients connected) with name `ROOM_NAME` and a max clients of `ROOM_SIZE`.



# Message Encoding

A byte is used per pixel.

| Range | Use                                       |
| ----- | ----------------------------------------- |
| 0     | Transparent                               |
| 1     | Black                                     |
| 2-3   | Greyscale (2 = light grey, 3 = dark grey) |
| 4-62  | Rainbow colours                           |
| 63    | RLE encoding start character              |

RLE encoding is `255 [counts] 0 [value]` where the total count is the sum of `counts`, plus 4 (if four or less characters are repeated they're not RLE'd as it'd be less efficient, and 0 is an illegal character in `[counts]`, so we know there's at least 5). The sum of `[counts]` is used as opposed to a product as to avoid having to complicate message checking for illegally large images