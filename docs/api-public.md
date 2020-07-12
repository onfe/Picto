# API - Public

All of the methods in this document can be called by anyone.



## room_exists

`/api/?method=room_exists&id=ROOM_ID`

If `ROOM_ID` exists, returns `true`. Otherwise `false`.



## get_public_rooms

`api/?method=get_public_rooms`

Returns a list of public rooms including their description, population and capacity as follows:

```json
[
	{
		"Name":"Pets",
        "Desc":"This room is for drawing pictures of your pets! You can send 1 picto to this room per day.",
		"Cap":64,
		"Pop":12
	}
]
```

The response is ordered and should be displayed in the order it is received.