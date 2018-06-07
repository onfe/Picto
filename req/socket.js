"use strict"; // Force ES6 and strict mode.

var WebSocket = require('ws')

module.exports = class Socket {
  constructor(ws) {
    this._ws = ws;
  }

  send(type, payload) {
    var msg = {
      type: type,
      time: new Date(),
      payload: payload
    }
    if (this.isAlive) {
      try {
        this._ws.send(JSON.stringify(msg))
      } catch (err) {
        console.log(err)
      }
    }

    console.log('SENT: ', msg)
  }

  get isAlive() {
    return (this._ws.readyState === WebSocket.OPEN)
  }
}
