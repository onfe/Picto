# API - Public

All of the methods in this document can be called by anyone.



## room_exists

`/api/?method=room_exists&id=ROOM_ID`

If `ROOM_ID` exists, returns `true`. Otherwise `false`.



## get_public_rooms

`api/?method=get_public_rooms`

Returns a list of public rooms including their population and capacity, as follows:

```json
[
	{
		"Name":"Parlor",
		"Cap":64,
		"Pop":12
	},
	{
		"Name":"Library",
		"Cap":64,
		"Pop":43
	},
	{
		"Name":"Garden",
		"Cap":64,
		"Pop":27
	}
]
```

The response is ordered and should be displayed in the order it is received.