"use strict";

module.exports = class Socket {
  constructor(un, rc, handler) {
    this.ws = new WebSocket(`ws://${window.location.hostname}:40510`)

    this.ws.onopen = this.wsOpened.bind(this);
    this.ws.onmessage = this.messageparse.bind(this);

    this.handler = handler;

    this.auth = false;

    this.username = un;
    this.roomcode = rc;

  }

  send(type, payload) {
    let msg = {
      type: type,
      time: new Date(),
      room: this.roomcode,
      name: this.username,
      auth: this.auth,
      payload: payload
    }
    let out = JSON.stringify(msg);
    this.ws.send(out);
  }

  wsOpened() {
    console.log('Successfully opened WebSocket. Connected to Picto servers...');
    this.join();
  }

  messageparse(msg) {
    let pl = JSON.parse(msg.data)
    // the 'joinresponse' is important to the socket too, and if authorisation failed, redirect the user.
    if (pl.type === "joinresponse") {
      if (!pl.payload.auth) {
        location.href = '/';
        return;
      } else {
        this.auth = pl.payload.auth;
      }
    }
    this.handler(pl)
  }

  join() {
    var pl = {}
    this.send('joinrequest', pl)
  }
}
