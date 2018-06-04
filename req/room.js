"use strict"; // Force ES6 and strict mode.

module.exports = class Room {
  constructor(id) {
    this._id = id;
    this._clients = [];
    this.created = Date.now();
    this._lastUpdate = Date.now();
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
    this._clients.push(c);
  }

  broadcast(type, pl) {
    for (var c = 0; c < this._clients.length; c++) {
      var client = this._clients[c]
      console.log(client)

      if (client.isAlive) {
        client.send(type, pl)
        console.log('msg')
      }
    }
    this.refresh()
  }

  cleanDeadClients() {
    for (var i = 0; i < this._clients.length; i++) {
      var client = this._clients[i];
      if (!client.isAlive) {
        this._clients.splice(i, 1);
      }
    }
  }

}
