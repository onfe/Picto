"use strict"; // Force ES6 and strict mode.

var Room = require('./room')
var Utils = require('./utils')

module.exports = class Picto {
  constructor() {
    this.rooms = []

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
    var id = Utils.randomHex()
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


}
