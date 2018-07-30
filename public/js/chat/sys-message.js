module.exports = class SysMessage {
  constructor(id, text) {
    this.text = text;
    this.id = id;
  }

  get html() {
    let html = `<div id="msg-${this.id}" class="msg system">
    <div class='msg-text'>${this.text}</div>
    </div>`;
    return html;
  }

  cleanup() {
    $('#msg-' + this.id).remove()
  }
}
