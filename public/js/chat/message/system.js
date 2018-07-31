const MessageBase = require('./base')

module.exports = class System extends MessageBase {
  constructor(id, text) {
    super(id);
    this.text = text;
  }

  get html() {
    let html = `<div id="msg-${this.id}" class="msg system">
    <div class='msg-text'>${this.text}</div>
    </div>`;
    return html;
  }
}
