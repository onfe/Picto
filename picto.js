var express = require('express')
var app = express()
var path = require('path')

const MongoClient = require('mongodb').MongoClient;
const assert = require('assert');



var WebSocketServer = require('ws').Server
var wss = new WebSocketServer({port: 40510})

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

wss.on('connection', function (ws) {
  ws.on('message', function (message) {
    console.log('received: %s', message)
  })
})
