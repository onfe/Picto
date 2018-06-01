var ws = new WebSocket('ws://127.0.0.1:40510');
var username = 'onfe'

ws.onopen = function () {
  console.log('Connected to Picto! Listening for Messages...')
  sendJoin();
}

function sendJoin() {
  var pl = {}
  send('joinrequest', pl)
}

function sendMessage(msg) {
  var pl = {
    'msgcont': msg,
  }
  send('message', pl);
}

function send(type, payload) {
  var msg = {
    'type': type,
    'time': new Date(),
    'server': 'fakeserverID',
    'user': username,
    'auth': false,
    'payload': payload
  }
  var out = JSON.stringify(msg);
  ws.send(out);
}
