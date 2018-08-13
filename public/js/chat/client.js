var Socket = require('./socket')
var Sound = require('./sound')
var Compose = require('./compose')
var Base64 = require('./base64')
var MessageHistory = require('./messagehistory')

module.exports = class Client {
  constructor() {

    // Get client username
    this.username = Cookies.get("user");

    var path = location.pathname.split('/')
    this.roomcode = path[path.length - 2]; // extract roomcode from page url. need to find better way to do this.
    this.joinLink = `${window.location.host}/join/${this.roomcode}/`;

    if (!this.username) { // check the username was correctly set.
      location.href = this.joinLink;
      return;
    }

    this.sock = new Socket(this.username, this.roomcode, this.recieve.bind(this))
    this.compose = new Compose();
    this.msgHist = new MessageHistory();

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

    $('.control.send').on('click', this.messageSend.bind(this));
    $('.control.clear').on('click', this.messageClear.bind(this));
    $('.control.get').on('click', this.messageGet.bind(this));

    // Useradd Overlay

    $('.useradd').on('click', function(e) {
      $('.useradd').addClass('show');
      $('#join-link').focus();
      $('#join-link').select();
      e.stopPropagation();
    });

    $(document).on('click', function(e) {
      $('.useradd').removeClass('show');
    })

  }

  recieve(pl) {
    console.log(pl)
    switch (pl.type) {
      case "joinresponse":
        this.load(pl);
        break;
      case "message":
        this.recieveMessage(pl);
        break;
      case "sent":
        // handle sent confirmation
        break;
      case "status":
        this.recieveStatus(pl);
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
    this.msgHist.addSysMsg('begin', 'Welcome to Picto!');
    $('#join-link').val(this.joinLink)

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

  recieveMessage(p) {
    this.Sound.message.play();
    var pl = p.payload
    this.msgHist.addMessage(pl.msgID, pl.sender, pl.colour, pl.msgCont)
  }

  recieveStatus(p) {
    var pl = p.payload;
    this.msgHist.addStatus(pl.id, pl.text)
    // also update the 'users tally';
  }

  messageSend() {
    var cont = this.compose.getContent();
    cont = Base64.encode(cont)
    this.sock.send('message', {msgCont: cont});
    this.messageClear();
  }

  messageClear() {
    this.compose.clear();
  }

  messageGet() {
    let srcMsg = this.msgHist.getRecent();
    if (srcMsg.data) {
      this.compose.load(srcMsg.databin)
    }
  }

}
