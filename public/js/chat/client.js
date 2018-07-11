var Socket = require('./socket')
var Sound = require('./sound')
var Compose = require('./compose')

module.exports = class Client {
  constructor() {

    // Get client username
    this.username = Cookies.get("user");

    var path = location.pathname.split('/')
    this.roomcode = path[path.length - 2]; // extract roomcode from page url. need to find better way to do this.

    if (!this.username) { // check the username was correctly set.
      location.href = "/";
      return;
    }

    this.sock = new Socket(this.username, this.roomcode, this.recieve.bind(this))
    this.compose = new Compose();

    // ------------------------
    // SETUP UI EVENT LISTENERS
    // ------------------------

    $('.tool.draw').on('click', function (e) {
      $('.tool.draw').removeClass('selected');
      $(e.target).addClass('selected');
      this.compose.currentTool = $(e.target).attr('id')
    }.bind(this));

    $('.tool.size').on('click', function (e) {
      $('.tool.size').removeClass('selected');
      $(e.target).addClass('selected');
      this.compose.currentSize = $(e.target).attr('id')
    }.bind(this));

  }

  recieve(pl) {
    console.log(pl)
    switch (pl.type) {
      case "joinresponse":
        this.load(pl);
        break;
      case "message":
        console.log('message')
        this.message(pl);
        break;
      case "sent":
        // handle sent confirmation
        break;
      case "status":
        // handle status update
        break;
      default:
        break;
    }
  }

  set username(un) {
    this._username = un
    $('.msg-auth.self').html(un);
  }

  get username() {
    return this._username;
  }

  set colour(col) {
    if (this._colour) { return; } // colour cannot be changed once set.
    this._colour = col;
    $('.msg-current').addClass(col)
  }

  get colour() {
    return this._colour;
  }

  load(pl) {
    this.colour = pl.payload.colour; // set the user's colour
    this.Sound = new Sound(); // Load all sounds async.

    // Check that all non-js assets are also loaded before calling ready.
    if (document.readyState === "complete") {
      this.ready()
      return;
    } else {
      var start = new Date();
      console.log('Waiting for assets...')
      $(window).on('load', function() {
        var delta = new Date() - start;
        console.log(`Assets loaded, took ${delta}ms`);


        this.ready()
      }.bind(this));

      return;
    }
  }

  ready() {
    $('.loader').addClass('hide'); // Hide the loader
    this.Sound.join.play(); // Play the join sound.
  }

  message() {
    this.Sound.message.play();
  }

}
