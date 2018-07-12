var Message = require('./message')

module.exports = class MessageHistory {
  constructor() {
    this.maxMessages = 32;
    this.msgHistID = "#msg-hist";
    this.messages = [];
  }

  addMessage(id, sender, colour, content) {
    // Recieve a b64 encoded message, add it to the message history and draw it.
    var msg = new Message(id, content, sender, colour);
    console.log(msg)
    this.messages.push(msg);
    $("#msg-hist").prepend(msg.html);
    msg.draw();

  }

  addStatus() {
    // Recieve a status update and add that to the message history
  }

  trim() {
    // Remove any messages over the message limit (this.maxMessages)
    // Keeping the message history trimmed should reduce the load on clients from dealing with bloated DOMs.
  }
}
