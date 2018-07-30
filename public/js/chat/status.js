module.exports = class Status {
  constructor(id, text) {
    this.text = text;
    this.id = id;
  }

  get html() {
    let html = `<div id="status-${this.id}" class="status">${this.text}</div>`
    return html;
  }

  cleanup() {
    $('#status-' + this.id).remove()
  }
}
