"use strict"; // Force ES6+ and strict mode.

/* -----------------------------------------------------------------------------
To build use browserify to produce picto-bundle.js
browserify main.js -o picto-bundle.js
------------------------------------------------------------------------------*/

var Client = require('./client')

// Create the client instance.
window.cli = new Client();

// hotkey detection
$(document).keypress(function(e) {
  if (e.which == 13) {
    cli.messageSend();
  }
})
