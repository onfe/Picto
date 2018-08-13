"use strict"; // Force ES6 and strict mode.

// Require express for static files, etc.
var express = require('express')
var app = express()
var path = require('path')

const Server = require('http').createServer();
const Port = (process.env.PORT || 80)
Server.on('request', app);

// Require and setup WebSockets
var WebSocket = require('ws')
var wss = new WebSocket.Server({server: Server})

// Require custom modules
var Picto = require('./req/picto')
var Client = require('./req/client')
var Token = require('./req/token')

var maxRooms = 8;

// CREATE THE EVERYTHING CLASS
var picto = new Picto()
// VERY IMPORTANT - DO. NOT. ACCIDENTLY. DELETE.

app.use(express.static('public'))

// START API

app.get('/room/:roomcode/', function (req, res) {
  var roomCheck = picto.checkRoom(req.params.roomcode)
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
      var msg = picto.checkRoom(req.query.room)
      res.send(msg)
      break;

    case "username":
      var room = picto.findRoom(req.query.room);
      if (!room) {
        res.status(400).send();
      } else {
        var msg = room.checkUsername(req.query.name)
        res.send(msg);
      }
      break;

    case "createroom":
      if (picto.rooms.length < maxRooms) {
        var newroom = picto.createRoom();
        var msg = {
          created: true,
          room: newroom.id,
        }
        console.log(picto.rooms);
      } else {
        msg = {
          created: false,
        }
      }

      res.send(msg);
      break;

    case "stats":
      var num_cli = picto.numClients();
      var num_rooms = picto.rooms.length;
      var msg = {
        numClients: num_cli,
        numRooms: num_rooms,
      }
      res.send(msg);
      break;

    default:
      res.status(400).send();
      break;
  }
})

// END API

wss.on('connection', function (ws) {
  ws.on('message', function (msg) {
    var pl = JSON.parse(msg)

    console.log('RECIEVED: ', pl)

    if (pl.type === 'joinrequest') {
      // ----------------------
      // TODO: MOVE TO ROOM.JOIN
      // ----------------------
      var roomOK = picto.checkRoom(pl.room)
      if (!roomOK.available || roomOK.full) {
        var cli = new Client(pl.name, ws, false)
        var msg = {auth: false, room: false};
      } else {
        var room = picto.findRoom(pl.room)
        var userOK = room.checkUsername(pl.name)
        if (userOK.available) {
          // everything is good, let's join!
          var auth = Token.create(pl.name, pl.room)
          var cli = new Client(pl.name, ws, auth);
          room.addClient(cli);

          var msg = {auth: cli.auth, room: room.id, colour: cli.colour}
        } else {
          var cli = new Client(pl.name, ws, false)
          var msg = {auth: false, room: false};
        }
      }

      cli.send('joinresponse', msg)
    }

    if (pl.type === 'message') {
      picto.recieve(pl);
    }

  })
})

// messageHandler() moved to picto.recieve();

// Finally, start the server.
Server.listen(Port, function() {
  console.log(`Picto started on port ${Port}`);
})
