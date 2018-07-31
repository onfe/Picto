module.exports = class MessageBase {
  constructor(id) {
    this.id = id;
  }

  get html() {
    let html = `<div id="msg-${this.id}" class="msg"></div>`
    return html;
  }

  cleanup() {
    $('#msg-' + this.id).remove()
  }
}
