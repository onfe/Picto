
module.exports = class Compose {
  constructor() {
    this.canvasId = '#canvas-compose';
    this.canvas = $(this.canvasId)[0];
    this.ctx = this.canvas.getContext('2d');

    this.ctx.fillStyle = "#000000"

    this.width = 192;
    this.height = 64;

    // Define and setup Pixel 2D array
    this.pixels = [];
    for (let x = 0; x < this.width; x++) {
      this.pixels.push([]);
      for (let y = 0; y < this.height; y++) {
        this.pixels[x].push(false);
      }
    }

    this.resize();

    this.currentTool = 'pencil';
    this.currentSize = 'small';

    $(this.canvasId).on('pointerdown', this.onMouseDown.bind(this));

    $(this.canvasId).on('pointerup', this.onMouseUp.bind(this));

    $(this.canvasId).on('pointerenter', this.onMouseEnter.bind(this));

    $(window).on('resize', this.resize.bind(this));

  }

  onMouseDown(e) {
    $(this.canvasId).on('pointermove', this.onMouseMove.bind(this)); // enable mousemove

    e = this.extractOffset(e)

    let pCoords = this.getPixel(e.offsetX, e.offsetY)

    this.lastPoint = pCoords; // save the pixel coords for interpolation.

    this.draw(pCoords[0], pCoords[1]);
  }

  onMouseUp(e) {
    e = this.extractOffset(e)
    $(this.canvasId).off('pointermove'); // disable mousemove
  }

  onMouseEnter(e) {
    e = this.extractOffset(e)
    this.lastPoint = this.getPixel(e.offsetX, e.offsetY)
  }

  onMouseMove(e) {
    e = this.extractOffset(e)
    let pCoords = this.getPixel(e.offsetX, e.offsetY)
    let delta = [pCoords[0] - this.lastPoint[0], pCoords[1] - this.lastPoint[1]]

    const inc = 1;
    let deltaLength = Math.sqrt((delta[0] * delta[0]) + (delta[1] * delta[1]))
    if (deltaLength < 0.5) { return; }

    for (let i = 0; i < deltaLength; i += inc) {
      let pc = i / deltaLength;
      let a = [delta[0] * pc, delta[1] * pc]
      let point = [Math.floor(a[0] + this.lastPoint[0]), Math.floor(a[1] + this.lastPoint[1])]

      this.draw(point[0], point[1]);
    }

    this.lastPoint = pCoords;

  }

  clear() {
    for (var x = 0; x < this.pixels.length; x++) {
      for (var y = 0; y < this.pixels[x].length; y++) {
        this.clearPixel(x, y)
      }
    }
  }

  getContent() {
    var outData = '';
    for (var x = 0; x < this.pixels.length; x++) {
      for (var y = 0; y < this.pixels[x].length; y++) {
        if (this.pixels[x][y]) {
          outData += '1';
        } else {
          outData += '0';
        }
      }
    }
    return outData;
  }

  draw(x, y) {
    if (this.currentTool === 'pencil') {
      var toolFunc = this.drawPixel.bind(this)
    } else if (this.currentTool === 'eraser') {
      var toolFunc  = this.clearPixel.bind(this)
    }

    if (this.currentSize === 'small') {
      toolFunc(x, y)
    } else if (this.currentSize === 'big') {
      this.wide(x, y, toolFunc)
    }

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

  wide(x, y, dFunc) { // widen a stroke by drawing multiple pixels from seed pixel.
    var stroke = 3;
    var start = Math.floor(stroke / 2);
    for (let pX = 0; pX < stroke; pX++) {
      for (let pY = 0; pY < stroke; pY++) {
        dFunc((x - start) + pX, (y - start) + pY);
      }
    }
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


    this.redraw();
  }

  extractOffset(e) {
    var or = e.target.getBoundingClientRect()
    e.offsetX = e.clientX - or.left;
    e.offsetY = e.clientY - or.top;
    return e;
  }

  redraw() {
    for (var x = 0; x < this.pixels.length; x++) {
      for (var y = 0; y < this.pixels[x].length; y++) {
        if (this.pixels[x][y]) {
          this.drawPixel(x, y)
        } else {
          this.clearPixel(x, y)
        }
      }
    }
  }

}
