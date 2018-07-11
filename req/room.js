"use strict"; // Force ES6 and strict mode.

var Utils = require('./utils')

var maxClients = 8;
var maxNameLength = 12;

module.exports = class Room {
  constructor(id) {
    this._id = id;
    this._clients = [];
    this.created = new Date();
    this.lastupdate = new Date();
    this._colours = ['orange', 'green', 'yellow', 'purple', 'blue'];
  }

  get clients() { return this._clients; }

  get id() { return this._id; }

  get full() {
    return !(this.clients.length < maxClients)
  }

  refresh() {
    this.lastupdate = new Date();
    return this.lastupdate;
  }

  addClient(c) {
    this.clients.push(c);
    var randomNum = Utils.getRandomInt(0, this._colours.length - 1)
    c.colour = this._colours[randomNum];
    this._colours.splice(randomNum, 1);
    // the client has been added, send a status update.
    this.sendStatus();
    this.broadcast('clientJoined', {
      name: c.name
    })
    this.refresh();
  }

  broadcast(type, pl) {
    this.refresh()
    for (var c = 0; c < this.clients.length; c++) {
      var client = this.clients[c]

      if (client.isAlive) {
        client.send(type, pl)
      }
    }
  }

  sendStatus() {
    this.broadcast('status', {
      numClients: this.clients.length
    });
  }

  cleanDeadClients() {
    for (var i = 0; i < this.clients.length; i++) {
      var client = this.clients[i];
      if (!client.isAlive) {
        var client = this.clients[i]
        this._colours.push[client.colour]
        this.clients.splice(i, 1);
        this.broadcast('clientLeft', {
          name: client.name
        });
        this.sendStatus();
      }
    }
  }

  findClient(name) {
    function finder(client) {
      if (client.name === this) {
        return client
      }
    }
    return this.clients.find(finder, name)
  }

  checkUsername(name) {
    var client = this.findClient(name)

    var available = true;
    var reason = '';

    if (client) { // client already exists with this name.
      available = false;
      reason = 'nametaken';
    }

    if (name.length > maxNameLength) {
      available = false;
      reason = 'namelength';
    }

    var msg = {
      room: this.id,
      name: name,
      available: available,
      reason: reason,
    }

    return msg;
  }

  join() {
    // move join check and stuff from main.js here
  }

}
