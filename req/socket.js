"use strict"; // Force ES6 and strict mode.

var WebSocket = require('ws')

module.exports = class Socket {
  constructor(ws) {
    this._ws = ws;
  }

  send(type, payload) {
    var msg = {
      type: type,
      time: Date.now(),
      payload: payload
    }
    this._ws.send(JSON.stringify(msg))
    console.log('SENT: ', msg)
  }

  get isAlive() {
    return (this._ws.readyState === WebSocket.OPEN)
  }
}
