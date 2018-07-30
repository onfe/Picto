var Base64 = require('./base64')

module.exports = class Message {
  constructor(id, data, author, colour) {
    this.id = id;
    this.data = data;
    this.databin = Base64.decode(this.data)
    this.colour = colour;
    this.auth = author;

    this.width = 192;
    this.height = 64;
  }

  get html() {
    let html = `<div id="msg-${this.id}" class="msg ${this.colour}">
    <canvas resize id="canv-${this.id}"></canvas>
    <span class="msg-auth">${this.auth}</span>`;
    return html;
  }

  draw() {
    this.canvas = $('#canv-' + this.id)[0];
    this.ctx = this.canvas.getContext('2d');
    this.ctx.fillStyle = "#000000"

    this.resize(); // resize the canvas before drawing.

    for (let y = 0; y < this.height; y++) {
      for (let x = 0; x < this.width; x++) {
        let dataIndex = (x * this.height) + y;
        if (this.databin[dataIndex] == '1') {
          this.drawPixel(x, y)
        }
      }
    }
  }

  drawPixel(x, y) {
    let start = [Math.floor(x * this.perX), Math.floor(y * this.perY)];
    let sizeX = Math.ceil(this.perX);
    let sizeY = Math.ceil(this.perY);
    this.ctx.fillRect(start[0], start[1], sizeX, sizeY);
  }

  resize() {

    let width = this.canvas.clientWidth;
    let height = this.canvas.clientHeight;

    if (this.canvas.width !== width || this.canvas.height !== height) {
      this.canvas.width = width;
      this.canvas.height = height;
    }
    this.perX = width / this.width;
    this.perY = height / this.height;

  }

  cleanup() {
    $('#msg-' + this.id).remove()
  }

}
