"use strict";

var jwt = require('jsonwebtoken');

const uuidV4 = require('uuid/v4');
const secret = uuidV4();

module.exports = class Token {
  static create(n, r) {
    return jwt.sign({ name: n, room: r }, secret, {
        expiresIn: 86400 // expires in 24 hours
      })
  }

  static verify(t, n, r) {
    var result = jwt.verify(t, secret, function(err, decoded) {
      if (err) {
        console.log(err)
        return false;
      }

      if (decoded.name === n && decoded.room === r) {
        // the token matches the user name, neat.
        return true;
      }

      return false;

    });
    return result;
  }
}
