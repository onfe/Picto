class Sketchpad {
  constructor(width, height, canvas) {
    this.width = width;
    this.height = height;
    this.notepad = new Notepad(width, height, canvas);

    this.rainbowMode = false;
    this.colourIndex = 1;
    this.saturation = 255; /*Saturation should be in range 0-255 inclusive*/

    /**last coords of mouse, -1 when the mouse is outside the canvas bounds. */
    this.lastMousePos = [-1, -1];

    /**These events monitor if the mouse is up/down */
    this.mousedown = false;
    this.notepad.canvas.addEventListener(
      "pointerdown",
      () => (this.mousedown = true)
    );
    this.notepad.canvas.addEventListener(
      "pointerup",
      () => (this.mousedown = false)
    );

    /**Drawing events */
    this.notepad.canvas.addEventListener("pointermove", this.drawTo.bind(this));
    this.notepad.canvas.addEventListener(
      "pointerdown",
      this.drawPix.bind(this)
    );
    this.notepad.canvas.addEventListener(
      "pointerleave",
      this.resetMousePos.bind(this)
    );

    /**imageData is all the drawn pixels and baked text */
    this.imageData = {
      span: this.width,
      data: new Array(this.width * this.height).fill(0)
    };

    /*Text styling */
    this.textMargin = 1;
    this.lineSpacing = 10;

    this.cursorPos = [this.textMargin, this.textMargin];

    /**typeLog contains charsLogged logs of text overlays in a circular queue.
     * It is drawn over imageData with this.overlayText() until it is baked,
     * at which point it becomes part of imageData.
     */
    this.charsLogged = 16;
    this.typeLog = new Array(this.charsLogged);
    this.typeLogHead = 0;
    this.typeLog[this.typeLogHead] = {
      data: new Array(this.width * this.height).fill(0),
      span: this.width,
      cursor: this.cursorPos.slice(0)
    };
  }

  /**-------------------------------------------------- Utils */
  getSendableData() {
    this.bakeText();
    return this.imageData;
  }

  clear() {
    this.imageData = {
      span: this.width,
      data: new Array(this.width * this.height).fill(0)
    };
    this.resetTypeLog();
    this.refresh();
  }

  refresh() {
    this.notepad.ctx.clearRect(0, 0, this.width, this.height);
    this.notepad.loadImageData(this.imageData);
    if (this.typeLog[this.typeLogHead] != undefined) {
      this.overlayText(this.typeLog[this.typeLogHead]);
    }
  }

  toggleRainbowMode() {
    this.rainbowMode = !this.rainbowMode;
    if (this.rainbowMode) {
      this.colourIndex = 2;
    } else {
      this.colourIndex = 1;
    }
  }

  /**-------------------------------------------------- Resets */
  resetMousePos() {
    this.lastMousePos = [-1, -1];
  }

  resetTypeLog() {
    this.typeLog = new Array(this.charsLogged);
    this.typeLogHead = 0;
    this.typeLog[this.typeLogHead] = {
      data: new Array(this.width * this.height).fill(0),
      span: this.width,
      cursor: this.cursorPos.slice(0)
    };
  }

  resetCursorPos() {
    this.cursorPos = [this.textMargin, this.textMargin];
  }

  /**-------------------------------------------------- Mouse drawing */
  drawPix(event) {
    this.bakeText();

    var [x, y] = this.getMousePixPos(event.offsetX, event.offsetY);

    this.imageData["data"][y * this.width + x] = this.colourIndex;
    this.notepad.setPixel(x, y, this.colourIndex);

    if (this.rainbowMode) {
      this.colourIndex = ((this.colourIndex + 1) % 254) + 2;
    }

    [this.lastMousePos[0], this.lastMousePos[1]] = [x, y];
    this.cursorPos = this.lastMousePos.slice(0);
  }

  drawTo(event) {
    var [x, y] = this.getMousePixPos(event.offsetX, event.offsetY);

    if (
      this.mousedown &&
      !(this.lastMousePos[0] == -1 && this.lastMousePos[1] == -1)
    ) {
      var deltas = [x - this.lastMousePos[0], y - this.lastMousePos[1]];
      var dist = Math.sqrt(deltas[0] * deltas[0] + deltas[1] * deltas[1]);

      for (var i = 0; i < dist; i += 0.5) {
        var tempx = Math.round(x - deltas[0] * (i / dist));
        var tempy = Math.round(y - deltas[1] * (i / dist));
        this.imageData["data"][tempy * this.width + tempx] = this.colourIndex;
        this.notepad.setPixel(tempx, tempy, this.colourIndex);
      }
      if (this.rainbowMode) {
        this.colourIndex = ((this.colourIndex + 1) % 254) + 2;
      }
    }

    [this.lastMousePos[0], this.lastMousePos[1]] = [x, y];
  }

  getMousePixPos(offsetX, offsetY) {
    return [
      Math.round((offsetX / this.notepad.canvas.clientWidth) * this.width),
      Math.round((offsetY / this.notepad.canvas.clientHeight) * this.height)
    ];
  }

  /**-------------------------------------------------- Text drawing */
  drawChar(char) {
    /*Setting up the styling for the text and ascertaining its size*/
    this.notepad.ctx.font = "16px 'Generic Pixel Font 5x7 Neue'";
    this.notepad.ctx.fillStyle = this.notepad.getColour(this.colourIndex);
    this.notepad.ctx.textBaseline = "hanging";
    this.notepad.ctx.textAlign = "end";
    var charWidth = Math.round(this.notepad.ctx.measureText(char).width);

    /**newText will hold the all the unbaked text on the sketchpad*/
    var newText = {
      data: new Array(this.width * this.height).fill(0),
      span: this.imageData["span"],
      cursor: this.cursorPos.slice(0)
    };
    if (this.typeLog[this.typeLogHead] != undefined) {
      newText["data"] = this.typeLog[this.typeLogHead]["data"].slice(0);
    }

    /**Adjusting the cursor position */
    if (this.cursorPos[0] + charWidth + this.textMargin > this.width) {
      this.cursorPos[0] = this.textMargin;
      if (
        this.cursorPos[1] + 2 * this.lineSpacing + this.textMargin <=
        this.height
      ) {
        this.cursorPos[1] += this.lineSpacing;
      } else {
        this.cursorPos[1] = this.textMargin;
      }
    }
    this.cursorPos[0] += charWidth;

    /*In order to do the thresholding to fix the antialiasing of the text in the
    canvas, we write the text to the canvas to get the image data, then return
    the canvas to its original state and use the image data to decide where to
    draw the text onto the canvas ourselves.*/
    var oldData = this.notepad.ctx.getImageData(0, 0, this.width, this.height);
    this.notepad.ctx.fillText(char, this.cursorPos[0], this.cursorPos[1]);
    var newData = this.notepad.ctx.getImageData(0, 0, this.width, this.height);
    this.notepad.ctx.putImageData(oldData, 0, 0);

    /**newText isn't added to typeLog unless the character actually made a
     * difference on the canvas (e.g. wasn't a space)
     */
    var diff = false;
    for (var i = 0; i < newData["data"].length; i += 4) {
      /**The alpha channel is used for masking out the text */
      if (
        oldData["data"][i + 3] != newData["data"][i + 3] &&
        newData["data"][i + 3] > 128
      ) {
        newText["data"][Math.floor(i / 4)] = this.colourIndex;
        diff = true;
      }
    }
    if (diff) {
      this.typeLogHead = (this.typeLogHead + 1) % this.charsLogged;
      this.typeLog[this.typeLogHead] = newText;
      this.refresh();
    }

    if (this.rainbowMode) {
      this.colourIndex = ((this.colourIndex + Math.round(254 / 16)) % 254) + 2;
    }
  }

  backspace() {
    var nextTypeLogHead =
      (this.typeLogHead + this.charsLogged - 1) % this.charsLogged;
    if (this.typeLog[nextTypeLogHead] != undefined) {
      this.cursorPos = this.typeLog[this.typeLogHead]["cursor"].slice(0);
      this.typeLog[this.typeLogHead] = undefined;
      this.typeLogHead = nextTypeLogHead;
      this.refresh();
      return;
    }
  }

  overlayText(data) {
    for (var i = 0; i < data["data"].length; i++) {
      if (data["data"][i] != 0) {
        this.notepad.setPixel(
          i % data["span"],
          Math.floor(i / data["span"]),
          data["data"][i]
        );
      }
    }
  }

  /**bakeText takes the head of the textLog and writes it onto the imageData */
  bakeText() {
    var lastLog = this.typeLog[this.typeLogHead];
    if (lastLog != undefined) {
      for (var i = 0; i < lastLog["data"].length; i++) {
        if (lastLog["data"][i] > 0) {
          this.imageData["data"][i] = lastLog["data"][i];
        }
      }
      this.resetCursorPos();
      this.resetTypeLog();
      this.refresh();
    }
  }
}

export default Sketchpad;
