
module.exports = class Compose {
  constructor() {
    this.canvasId = '#canvas-compose';
    this.canvas = $(this.canvasId)[0];
    this.ctx = this.canvas.getContext('2d');

    this.ctx.fillStyle = "#000000"

    this.width = 192;
    this.height = 64;

    this.resize();

    // Define and setup Pixel 2D array
    this.pixels = [];
    for (let x = 0; x < this.width; x++) {
      this.pixels.push([]);
      for (let y = 0; y < this.height; y++) {
        this.pixels[x].push(false);
      }
    }

    this.currentTool = 'pencil';
    this.currentSize = 'small';

    $(this.canvasId).on('pointerdown', this.onMouseDown.bind(this));

    $(this.canvasId).on('pointerup', this.onMouseUp.bind(this));

    $(this.canvasId).on('pointerenter', this.onMouseEnter.bind(this));

  }

  onMouseDown(e) {
    $(this.canvasId).on('pointermove', this.onMouseMove.bind(this)); // enable mousemove

    let pCoords = this.getPixel(e.offsetX, e.offsetY)

    this.lastPoint = pCoords; // save the pixel coords for interpolation.

    if (this.currentTool === 'pencil') {
      this.drawPixel(pCoords[0], pCoords[1]);
    } else if (this.currentTool === 'eraser') {
      this.clearPixel(pCoords[0], pCoords[1])
    }
  }

  onMouseUp(e) {
    $(this.canvasId).off('pointermove'); // disable mousemove
  }

  onMouseEnter(e) {
    this.lastPoint = this.getPixel(e.offsetX, e.offsetY)
  }

  onMouseMove(e) {
    let pCoords = this.getPixel(e.offsetX, e.offsetY)
    let delta = [pCoords[0] - this.lastPoint[0], pCoords[1] - this.lastPoint[1]]

    const inc = 1;
    let deltaLength = Math.sqrt((delta[0] * delta[0]) + (delta[1] * delta[1]))
    if (deltaLength < 0.5) { return; }

    for (let i = 0; i < deltaLength; i += inc) {
      let pc = i / deltaLength;
      let a = [delta[0] * pc, delta[1] * pc]
      let point = [Math.floor(a[0] + this.lastPoint[0]), Math.floor(a[1] + this.lastPoint[1])]

      if (this.currentTool === 'pencil') {
        this.drawPixel(point[0], point[1]);
      } else if (this.currentTool === 'eraser') {
        this.clearPixel(point[0], point[1]);
      }

    }

    this.lastPoint = pCoords;

  }

  clear() {

  }

  content() {

  }

  drawPixel(x, y) {
    let start = this.getPoint(Math.floor(x), Math.floor(y));
    let sizeX = Math.ceil(this.perX);
    let sizeY = Math.ceil(this.perY);
    this.ctx.fillRect(start[0], start[1], sizeX, sizeY);
    this.pixels[x][y] = true;
  }

  clearPixel(x, y) {
    let start = this.getPoint(Math.floor(x), Math.floor(y));
    let sizeX = Math.ceil(this.perX);
    let sizeY = Math.ceil(this.perY);
    this.ctx.clearRect(start[0], start[1], sizeX, sizeY);
    this.pixels[x][y] = false;
  }

  getPixel(x, y) {
    // Translate point coord to pixel coord
    let pX = Math.floor(x / this.perX);
    let pY = Math.floor(y / this.perY);
    return [pX, pY];
  }

  getPoint(x, y) {
    // Translate pixel coord to point coord
    let pX = Math.floor(x * this.perX);
    let pY = Math.floor(y * this.perY);
    return [pX, pY];
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

}
