"use strict"; // Force ES6 and strict mode.

var Socket = require('./socket')

module.exports = class Client {
  constructor(name, socket, auth) {
    this.name = name;
    this.auth = auth;
    this.socket = new Socket(socket);
    this.colour = '';
  }

  send(type, payload) {
    console.log('msg at client')
    this.socket.send(type, payload);
  }

  get isAlive() {
    return this.socket.isAlive;
  }
}
