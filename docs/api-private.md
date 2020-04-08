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



## get_submission_rooms

`/api/?token=API_TOKEN&method=get_submission_rooms`

Returns a list of submission rooms with their population, capacity and the number of (unpublished) submissions in that room, as follows:

```json
[
	{
		"Name":"SubmissionRoom",
		"Cap":64,
		"Pop":5,
        "Submissions":34
	}
]
```



## get_submissions

`api/?token=API_TOKEN&method=get_submissions&room_id=ROOM_ID`

Returns all the submissions in the submission cache of the `ROOM_ID` specified, as follows:

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

The submissions are ordered such that the oldest submission is first.



## publish_submission

`/api/?token=API_TOKEN&method=publish_submission&room_id=ROOM_ID&submission_id=SUBMISSION_ID`

Publishes the submission `SUBMISSION_ID` in `ROOM_ID`.



## reject_submission

`/api/?token=API_TOKEN&method=reject_submission&room_id=ROOM_ID&submission_id=SUBMISSION_ID`

Rejects the submission `SUBMISSION_ID` in `ROOM_ID`.



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



## create_submission_room

`/api/?token=API_TOKEN&method=create_submission_room&name=ROOM_NAME&desc=DESCRIPTION&size=ROOM_SIZE`

Creates a submission room (appears on front page, is public and moderated) with name `ROOM_NAME` and description `DESCRIPTION` that can hold `ROOM_SIZE` clients.

Unlike `create_static_room`, a `ROOM_SIZE` must be specified.



