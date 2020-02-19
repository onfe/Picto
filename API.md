# API and WebSocket Protocol



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
  "Message": "NPXkOU8..."
}
```

Server -> Client:

```JSON
{
  "ColourIndex": 2,
  "Sender": "Josh",
  "Message": "NPXkOU8..."
}
```

`ColourIndex` was the index of the user that sent the message in the `users` array when they sent it, it differs from `UserIndex` in that there may or may not be the same user that sent the message in that index in the `Users` array. 



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
  "UserIndex": 2,
  "RoomName": "Denver Airport"
}
```

`UserIndex` is the index of the user in the users array who changed the room's name.

`rename` events are not cached, so we don't need to worry about `UserIndex` becoming incorrect on user join/leaves.



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