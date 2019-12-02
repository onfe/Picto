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
  "Event": "init",
  "RoomID": "id",
  "RoomName": "default",
  "UserIndex": 1,
  "Users": ["Eddie", null, "Josho", null, null, "Martin", "Elle", null],
  "NumUsers": 4
}
```
`Users` has length equal to the max number of users in a room. Colours are
assigned using the modulo of the index of the user. You have the index 
`UserIndex` in the array. The `RoomID` is the value for `/room/code` and serves
as the link for inviting friends to the room.

### User join/leave

User join:
```JSON
{
  "Event":"user",
  "UserIndex": 1,
  "Users": ["Eddie", "Jordie", "Josho", null, null, "Martin", "Elle", null],
  "NumUsers": 5
}
```

User leave:
```JSON
{
  "Event":"user",
  "UserIndex": 1,
  "Users": ["Eddie", null, "Josho", null, null, "Martin", "Elle", null],
  "NumUsers": 4
}
```

### Message

Server -> Client:
```JSON
{
  "Event": "message",
  "UserIndex": 2,
  "Message": "NPXkOU8..."
}
```
Client -> Server:
```JSON
{
  "Event": "message",
  "Message": "NPXkOU8..."
}
```

### Announcement

Server -> Client:
```JSON
{
  "Event": "announcement",
  "Announcement": "Welcome to Picto!",
}
```

### Rename Room

Client -> Server -> Client(s)
```JSON
{
  "Event": "rename",
  "UserIndex": 2,
  "RoomName": "Denver Airport"
}
```
`UserIndex` is the index of the user in the users array who changed the room's 
name.

---

## API - Public

### room_exists

`/API/?method=room_exists&room_id=ROOM_ID`

If `ROOM_ID` exists, returns `true`. Otherwise `false`.

---

## API - Private

### get_state

`/API/?token=API_TOKEN&method=get_state`
Returns the state of the entire server by marshalling the roomManager. Not wise to use in prod.

### get_room_ids

`/API/?token=API_TOKEN&method=get_room_ids`
Returns a list of all the ids of the currently open rooms.

### get_room_state

`/API/?token=API_TOKEN&method=get_room_state&room_id=ROOM_ID`
Returns the state of the `ROOM_ID` specified by marshalling the room object.

### announce

`/API/?token=API_TOKEN&method=announce&message=MESSAGE`
Announces `MESSAGE` to ALL ROOMS.

`/API/?token=API_TOKEN&method=announce&message=MESSAGE&room_id=ROOM_ID`
Announces `MESSAGE` to the `ROOM_ID` specified.