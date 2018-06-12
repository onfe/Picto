"use strict"; // Force ES6 and strict mode.

var Utils = require('./utils')

module.exports = class Room {
  constructor(id) {
    this._id = id;
    this._clients = [];
    this.created = Date.now();
    this._lastUpdate = Date.now();
    this._colours = ['orange', 'green', 'yellow', 'purple', 'blue'];
  }

  get clients() {
    return this._clients;
  }

  get id() {
    return this._id;
  }

  refresh() {
    this._lastUpdate = Date.now();
    return this._lastUpdate;
  }

  addClient(c) {
    this.clients.push(c);
    var randomNum = Utils.getRandomInt(0, this._colours.length - 1)
    c.colour = this._colours[randomNum];
    this._colours.splice(randomNum, 1);
  }

  broadcast(type, pl) {
    for (var c = 0; c < this.clients.length; c++) {
      var client = this.clients[c]

      if (client.isAlive) {
        client.send(type, pl)
      }
    }
    this.refresh()
  }

  cleanDeadClients() {
    for (var i = 0; i < this.clients.length; i++) {
      var client = this.clients[i];
      if (!client.isAlive) {
        var client = this.clients[i]
        this._colours.push[client.colour]
        this.clients.splice(i, 1);
      }
    }
  }

}
