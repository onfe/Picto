"use strict"; // Force ES6 and strict mode.

var express = require('express')
var app = express()
var path = require('path')

const uuidV4 = require('uuid/v4');
const secret = uuidV4();


var jwt = require('jsonwebtoken');

var WebSocket = require('ws')
var wss = new WebSocket.Server({port: 40510})


var Client = require('./req/client')
var Room = require('./req/room')
var Socket = require('./req/socket')



app.use(express.static('public'))

app.get('/room/:roomcode', function (req, res) {
  return res.send(req.params)
})

app.get('/api/', function (req, res) {
  console.log(req.query)
  if (req.query.roomCode == 'abc') {
    var avail = true
  } else {
    var avail = false
  }
  return res.send({roomCode: req.query.roomCode, available: avail})
})


app.listen(8000, function () {
  console.log('Picto listening on port 8000')
})









function createUserToken(n, r) {
  return jwt.sign({ name: n, room: r }, secret, {
      expiresIn: 86400 // expires in 24 hours
    })
}

function verifyUserToken(t, n, r) {
  var result = jwt.verify(t, secret, function(err, decoded) {
    if (err) {
      console.log(err)
      return false;
    }

    if (decoded.name === n && decoded.room === r) {
      // the token matches the user name, neat.
      return true;
    }

    console.log(n, r, decoded.name, decoded.room)

    return false;

  });
  return result;
}

var rooms = [];
rooms.push(new Room('fakeRoomID'));

wss.on('connection', function (ws) {
  ws.on('message', function (msg) {
    var pl = JSON.parse(msg)

    console.log('RECIEVED: ', pl)

    if (pl.type === 'joinrequest') {
      var auth = createUserToken(pl.name, pl.room)
      var cli = new Client(pl.name, ws, auth);
      rooms[0].addClient( cli ); // FORCE ROOM 0
      cli.send('joinresponse', {auth: cli.auth})
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

function messageHandler(pl) {

  var senderName = pl.name;
  var senderAuth = pl.auth;
  var senderRoom = pl.room;

  var result = verifyUserToken(senderAuth, senderName, senderRoom);
  if (!result) {
    console.log('unauthorised')
    return
  }



  var plOut = {
    msgCont: pl.payload.msgCont,
    sender: senderName,
    colour: 'orange',
    msgID: randomHex()
  }

  var roomID = pl.room;
  var room = findRoomByID(roomID);
  room.broadcast('message', plOut);
  room.cleanDeadClients();
}
