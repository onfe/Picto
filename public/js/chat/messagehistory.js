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
    console.log(msg)
    this.messages.push(msg);
    $("#msg-hist").prepend(msg.html);
    msg.draw();

    this.trim();

  }

  addStatus(id, content) {
    let stat = new Status(id, content);
    this.messages.push(stat);
    $("#msg-hist").prepend(stat.html);

    this.trim();
  }

  addSysMsg(id, content) {
    let sys = new SysMessage(id, content);
    this.messages.push(sys);
    $("#msg-hist").prepend(sys.html);

    this.trim();
  }

  trim() {
    if (this.messages.length > this.maxMessages) {
      this.messages[0].cleanup()
      this.messages.splice(0, 1);
    }
  }
}
