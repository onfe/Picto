"use strict"; // Force ES6 and strict mode.

var Room = require('./room')
var Utils = require('./utils')
const Token = require('./token')

module.exports = class Picto {
  constructor() {
    this.rooms = []
    setInterval(this.refresh.bind(this), 10000);

  }

  findRoom(id) {
    function finder(room) {
      if (room.id === this) {
        return room
      }
    }
    return this.rooms.find(finder, id)
  }

  createRoom() {
    var id = Utils.shortHex()
    var newroom = new Room(id)
    this.rooms.push(newroom)
    return newroom;
  }

  checkRoom(roomid) {
    var room = this.findRoom(roomid)
    if (room) {

      var msg = {
        room: room.id,
        numCli: room.clients.length,
        available: true,
        full: room.full,
      }

    } else {
      var msg = {
        room: roomid,
        available: false,
      }
    }
    return msg;
  }

  numClients() {
    var num = 0;
    for (var i = 0; i < this.rooms.length; i++) {
      num += this.rooms[i].clients.length
    }
    return num;
  }

  refresh() {
    // if a room's last update was over x time ago, and there are no connected clients, close the room.
    for (var i = 0; i < this.rooms.length; i++) {
      var room = this.rooms[i];
      // check for dead clients.
      room.cleanDeadClients();
      var now = new Date();
      var sinceUpdate = now - room.lastupdate;
      if (sinceUpdate > 30000 && room.clients.length == 0) {
        this.rooms.splice(i, 1); // delete the room
      }
    }
  }

  recieve(pl) {
    let verified = Token.verify(pl.auth, pl.name, pl.room);
    if (!verified) { return; };

    let room = this.findRoom(pl.room);
    if (!room) { return; };

    let client = room.findClient(pl.name);
    if (!client) { return; };

    var payload = {
      msgCont: pl.payload.msgCont,
      sender: pl.name,
      colour: client.colour,
      msgID: Utils.randomHex()
    }

    client.send('sent', { msgID: payload.msgID })
    room.broadcast('message', payload);
    room.cleanDeadClients();
  }
}
