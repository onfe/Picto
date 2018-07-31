const MessageBase = require('./base')

module.exports = class Status extends MessageBase {
  constructor(id, text) {
    super(id);
    this.text = text;
  }

  get html() {
    let html = `<div id="msg-${this.id}" class="status">${this.text}</div>`
    return html;
  }
}
