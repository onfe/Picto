var ws = new WebSocket('ws://' + window.location.hostname + ':40510');
var username = 'onfe'
var auth = false;

ws.onopen = function () {
  console.log('Connected to Picto! Listening for Messages...')
  sendJoin();
}

ws.onmessage = function (msg) {
  pl = JSON.parse(msg.data);

  console.log(pl);

  if (pl.type === 'joinresponse') {
    auth = pl.payload.auth;
  }

  if (pl.type === 'message') {
    messageRecieved(pl.payload); // back to main.js
  }

  if (pl.type === 'sent')
    msgSent(pl.payload) // back to main.js
}

function sendJoin() {
  var pl = {}
  send('joinrequest', pl)
}

function sendMessage(msg) {
  var pl = {
    msgCont: msg,
  }
  send('message', pl);
}

function send(type, payload) {
  var msg = {
    type: type,
    time: new Date(),
    room: 'fakeRoomID',
    name: username,
    auth: auth,
    payload: payload
  }
  var out = JSON.stringify(msg);
  ws.send(out);
  console.log(out);
}
