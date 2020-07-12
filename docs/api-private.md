# Private API

All of the methods in this document require a token to be called.



## get_state

`/api/?token=API_TOKEN&method=get_state`

Returns the state of the entire server by marshalling the `roomManager`. Not permitted in prod.



## get_room_ids

`/api/?token=API_TOKEN&method=get_room_ids`

Returns a list of all the ids of the currently open rooms.



## get_room_state

`/api/?token=API_TOKEN&method=get_room_state&id=ROOM_ID`

Returns the state of the `ROOM_ID` specified by marshalling the room object.



## get_static_rooms

`/api/?token=API_TOKEN&method=get_static_rooms`

Returns a list of static rooms with their population and capacity as follows:

```json
[
	{
		"Name":"StaticRoom",
		"Cap":16,
		"Pop":0
	}
]
```



## get_moderated_rooms

`/api/?token=API_TOKEN&method=get_moderated_rooms`

Returns a list of moderated rooms with their population, capacity and the number of (visible) messages in that room, as follows:

```json
[
	{
		"Name":"ModeratedRoom",
        "Desc":"This is a moderated room.",
		"Cap":64,
		"Pop":5,
        "Visible":10,
        "Invisible":34
	}
]
```



## get_moderated_messages

`api/?token=API_TOKEN&method=get_moderated_messages&room_id=ROOM_ID`

Returns all the messages in the moderationCache of the `ROOM_ID` specified, as follows:

```json
[
    {
        "ID":"127.0.0.1-8-April",
     	"Addr":"127.0.0.1:36262",
        "Message":
 	       {
               "ColourIndex":0,
               "Sender":"Josh",
               "Data":"//////////gAABBBB/KAAB///rAABAAAB/JAAB///rAABAAABABBBBAABAAABABBBB///oAABBBBAABAAABABAAABABAAAB///nAABAAABAB/BAABAAABABAAAB///nAABAAABAB/BAABAAABABAAAB///nAABBBBAAB/CAABBBBABAAAB///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////wAA",
               "Span":192
           }
    }
]
```

The messages are ordered such that the oldest message is first.



## set_moderated_message_state

`/api/?token=API_TOKEN&method=set_moderated_message_state&room_id=ROOM_ID&message_id=MESSAGE_ID&state=STATE`

Sets the state of message `MESSAGE_ID` in `ROOM_ID` to `STATE`.

Valid `STATE` values are: "invisible" and "visible".

If a message's `STATE` is set to "visible" and there are already `MaxVisibleMessages` visible messages, then the oldest one is discarded.



## delete_message

`/api/?token=API_TOKEN&method=delete_message&room_id=ROOM_ID&message_id=MESSAGE_ID&offensive=OFFENSIVE`

Deletes the message `MESSAGE_ID` in `ROOM_ID`.

`OFFENSIVE` must be `true` or `false` exactly (it is case sensitive).

If `OFFENSIVE` is `true`, the client who sent the message will be ignored for `ClientIgnoreTime` (set in `config.go`) from when this endpoint is called. Otherwise, the message is simply discarded.



## announce

`/api/?token=API_TOKEN&method=announce&message=MESSAGE&id=ROOM_ID`

Announces `MESSAGE` to the `ROOM_ID` specified.

If no `ROOM_ID` is supplied, the announcement is made to **ALL ROOMS** - be careful!



## close_room

`/api/?token=API_TOKEN&method=close_room&id=ROOM_ID&reason=REASON&time=TIME`

Closes `ROOM_ID` and announces message `REASON` before closing the room after `TIME` seconds. 

If `REASON` is not supplied then it defaults to "`This room is being closed by the server.`".

If `TIME` is not supplied it defaults to the `DefaultCloseTime` set in `config.go`.

This time can be overwritten by making another call to this endpoint, but you will need to supply a second reason. Once called it cannot be cancelled.



## create_static_room

`/api/?token=API_TOKEN&method=create_static_room&name=ROOM_NAME&size=ROOM_SIZE`

Creates a static room (continues to exist when there are no clients connected) with name `ROOM_NAME` and a max clients of `ROOM_SIZE`. 

If `ROOM_SIZE` is not supplied it defaults to the `DefaultRoomSize` set in `config.go`.



## create_moderated_room

`/api/?token=API_TOKEN&method=create_moderated_room&name=ROOM_NAME&desc=DESCRIPTION&size=ROOM_SIZE`

Creates a moderated room (appears on front page, is public and moderated) with name `ROOM_NAME` and description `DESCRIPTION` that can hold `ROOM_SIZE` clients.

Unlike `create_static_room`, a `ROOM_SIZE` must be specified.



