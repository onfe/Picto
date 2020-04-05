# Private API

All of the methods in this document require a token to be called.



## get_state

`/api/?token=API_TOKEN&method=get_state`
Returns the state of the entire server by marshalling the roomManager. Not permitted in prod.



## get_room_ids

`/api/?token=API_TOKEN&method=get_room_ids`
Returns a list of all the ids of the currently open rooms.



## get_room_state

`/api/?token=API_TOKEN&method=get_room_state&id=ROOM_ID`
Returns the state of the `ROOM_ID` specified by marshalling the room object.



## get_static_rooms

`/api/?token=API_TOKEN&method=get_static_rooms`

Returns a list of static rooms including if they're public, their population and capacity as follows:

```json
[
	{
		"Name":"Parlor",
		"Public":true,
		"Cap":64,
		"Pop":0
	},
	{
		"Name":"Library",
		"Public":true,
		"Cap":64,
		"Pop":0
	},
	{
		"Name":"Garden",
		"Public":true,
		"Cap":64,
		"Pop":0
	},
	{
		"Name":"Hidden Static Room",
		"Public":false,
		"Cap":16,
		"Pop":0
	}
]
```



## announce

`/api/?token=API_TOKEN&method=announce&message=MESSAGE`
Announces `MESSAGE` to ALL ROOMS - careful!

`/api/?token=API_TOKEN&method=announce&message=MESSAGE&id=ROOM_ID`
Announces `MESSAGE` to the `ROOM_ID` specified.



## close_room

`/api/?token=API_TOKEN&method=close_room&id=ROOM_ID&reason=REASON&time=TIME`
Closes `ROOM_ID` and announces message `REASON` before closing the room after `TIME` seconds. This time can be overwritten by making another call to this endpoint, but you will need to supply a second reason. Once called it cannot be cancelled.



## create_static_room

`/api/?token=API_TOKEN&method=create_static_room&name=ROOM_NAME&size=ROOM_SIZE&public=PUBLIC`
Creates a static room (continues to exist when there are no clients connected) with name `ROOM_NAME` and a max clients of `ROOM_SIZE`. If `PUBLIC` is "true" then it is a public room that will be displayed in the front page, any other value will be interpreted as false and it is case sensitive.


