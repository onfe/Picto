"use strict"; // Force ES6 and strict mode.

// Require express for static files, etc.
var express = require('express')
var app = express()
var path = require('path')

// Require and setup WebSockets
var WebSocket = require('ws')
var wss = new WebSocket.Server({port: 40510})

// Require custom modules
var Client = require('./req/client')
var Room = require('./req/room')
var Socket = require('./req/socket')

var Token = require('./req/token')

var maxClients = 8;
var maxNameLength = 12;

var maxRooms = 8;



app.use(express.static('public'))

app.get('/room/:roomcode/', function (req, res) {
  var roomCheck = checkRoom(req.params.roomcode)
  if (roomCheck.available) {
    if (!roomCheck.full) {
      res.sendFile(__dirname + "/public/room.html")
    } else {
      res.redirect('/#room-full')
    }
  } else {
    res.redirect('/#room-not-available')
  }
})

app.get('/join/:roomcode/', function (req, res) {
  res.sendFile(__dirname + "/public/join.html")
});

app.get('/api/:type/', function (req, res) {
  switch (req.params.type) {

    case "room":
      var msg = checkRoom(req.query.room)
      res.send(msg)
      break;

    case "username":
      var room = findRoomByID(req.query.room);
      if (!room) {
        res.status(400).send();
      } else {
        var msg = checkUsername(room, req.query.name)
        res.send(msg);
      }
      break;

    case "createroom":
      if (rooms.length < maxRooms) {
        var newroom = new Room(randomHex());
        var msg = {
          created: true,
          room: newroom.id,
        }
        rooms.push(newroom)
        console.log(rooms);
      } else {
        msg = {
          created: false,
        }
      }

      res.send(msg);
      break;

    default:
      res.status(400).send();
      break;
  }
})

app.listen(8000, function () {
  console.log('Picto listening on port 8000');
})


// CREATE THE ROOMS LIST
var rooms = [];
// VERY IMPORTANT - DO. NOT. ACCIDENTLY. DELETE.
// (AND THEN SPEND HALF AN HOUR WORRYING ABOUT ALL THE ERRORS)

wss.on('connection', function (ws) {
  ws.on('message', function (msg) {
    var pl = JSON.parse(msg)

    console.log('RECIEVED: ', pl)

    if (pl.type === 'joinrequest') {
      var roomOK = checkRoom(pl.room)
      if (!roomOK.available || roomOK.full) {
        var cli = new Client(pl.name, ws, false)
        var msg = {auth: false, room: false};
      } else {
        var room = findRoomByID(pl.room)
        var userOK = checkUsername(room, pl.name)
        if (userOK.available) {
          // everything is good, let's join!
          var auth = Token.create(pl.name, pl.room)
          var cli = new Client(pl.name, ws, auth);
          room.addClient(cli);

          var msg = {auth: cli.auth, room: room.id, colour: cli.colour}
        }
      }

      cli.send('joinresponse', msg)
    }

    if (pl.type === 'message') {
      messageHandler(pl)
    }

  })
})


function randomHex() {
  return Math.floor(Math.random()*16777215).toString(16);
}

function findRoomByID(id) {
  function finder(room) {
    if (room.id === this) {
      return room
    }
  }
  return rooms.find(finder, id)
}

function findClientByName(room, name) {
  function finder(client) {
    if (client.name === this) {
      return client
    }
  }
  console.log(room)
  return room.clients.find(finder, name)
}

function checkRoom(roomid) {
  var room = findRoomByID(roomid)
  if (room) {

    let full = !(room.clients.length < maxClients)

    var msg = {
      room: room.id,
      numCli: room.clients.length,
      available: true,
      full: full,
    }

  } else {
    var msg = {
      room: roomid,
      available: false,
    }
  }
  return msg;
}

function checkUsername(room, name) {
  var client = findClientByName(room, name)

  var available = true;
  var reason = '';

  if (client) { // client already exists with this name.
    available = false;
    reason = 'nametaken';
  }

  if (name.length > maxNameLength) {
    available = false;
    reason = 'namelength';
  }

  var msg = {
    room: room.id,
    name: name,
    available: available,
    reason: reason,
  }

  return msg;
}

function messageHandler(pl) {

  var verified = Token.verify(pl.auth, pl.name, pl.room);
  if (!verified) {
    console.log('unauthorised')
    return;
  }

  var room = findRoomByID(pl.room);
  if (!room) {
    console.log('roomfailure')
    return;
  };
  var client = findClientByName(room, pl.name)
  if (!client) {
    console.log('clientfailure')
    return;
  }

  var plOut = {
    msgCont: pl.payload.msgCont,
    sender: pl.name,
    colour: client.colour,
    msgID: randomHex()
  }

  client.send('sent', { msgID: plOut.msgID }) // message sent response
  room.broadcast('message', plOut);
  room.cleanDeadClients();
}
