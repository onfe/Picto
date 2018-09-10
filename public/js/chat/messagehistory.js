var Message = require('./message/message')
var Status = require('./message/status')
var SysMessage = require('./message/system')

module.exports = class MessageHistory {
  constructor() {
    this.maxMessages = 32;
    this.msgHistID = "#msg-hist";
    this.messages = [];
  }

  addMessage(id, sender, colour, content) {
    // Recieve a b64 encoded message, add it to the message history and draw it.
    let msg = new Message(id, content, sender, colour);
    this.add(msg)
    $('#msg-breadcrumbs').prepend(`<a class="breadcrumb ${colour}"></a>`)

    msg.draw(); // draw the message
  }

  addStatus(id, content) {
    let stat = new Status(id, content);
    this.add(stat)
    $('#msg-breadcrumbs').prepend(`<a class="breadcrumb"></a>`)
  }

  addSysMsg(id, content) {
    let sys = new SysMessage(id, content);
    this.add(sys)
    $('#msg-breadcrumbs').prepend(`<a class="breadcrumb system"></a>`)
  }

  add(msg) {
    this.messages.push(msg)
    this.trim()

    var willScroll = false;
    var scrollBefore = $('#msg-hist').scrollTop()
    var histHeight = $('#msg-hist').height()
    var scrollHeight = $("#msg-hist")[0].scrollHeight;
    if ((scrollHeight - (scrollBefore + histHeight)) <= 10) { // if the message history is not focused on the most recent messages (+ tolerance), do not scroll.
      // console.log(scrollHeight, scrollBefore, scrollHeight)
      var willScroll = true;
    }
    // console.log(willScroll);

    $("#msg-hist").prepend(msg.html); // add message html

    var newScroll = $("#msg-hist")[0].scrollHeight - $('#msg-hist').height();

    $('#msg-hist').scrollTop(scrollBefore)
    if (willScroll) {
      $('#msg-hist').animate({scrollTop: newScroll}, 300);
    }

  }

  getRecent() {
    if (this.messages.length == 0) {
      return;
    }
    return this.messages[this.messages.length - 1]
  }

  trim() {
    if (this.messages.length > this.maxMessages) {
      this.messages[0].cleanup()
      this.messages.splice(0, 1);
    }
    if ($('.breadcrumb').length > this.maxMessages) {
      $('.breadcrumb').last().remove();
    }
  }
}
