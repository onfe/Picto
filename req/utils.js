"use strict"; // Force ES6 and strict mode.

module.exports.getRandomInt = function (min, max) {
  min = Math.ceil(min);
  max = Math.floor(max);
  return Math.floor(Math.random() * (max - min)) + min; //The maximum is exclusive and the minimum is inclusive
}

module.exports.randomHex = function() {
  return Math.floor(Math.random()*16777215).toString(16);
}

module.exports.shortHex = function() {
  return Math.floor(Math.random()*65535).toString(16);
}
