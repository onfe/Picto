var express = require('express')
var app = express()
var path = require('path')

var WebSocketServer = require('ws').Server
var wss = new WebSocketServer({port: 40510})

app.use(express.static('public'))

app.get('/api/', function (req, res) {
  return
})

app.listen(8000, function () {
  console.log('Example app listening on port 8000!')
})

wss.on('connection', function (ws) {
  ws.on('message', function (message) {
    console.log('received: %s', message)
  })
})
