import Notepad from "./notepad.js";
import Camera from "./camera";

class Sketchpad {
  constructor(width, height, canvas, nameWidth) {
    this.width = width;
    this.height = height;
    this.notepad = new Notepad(width, height, canvas);

    this.pensize = 0;

    this.rainbowMode = false;
    this.colourIndex = 1;

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
    this.textMargin = 8;
    this.lineSpacing = 24;

    this.nameWidth = nameWidth || 0;

    this.cursorPos = [
      Math.round(this.width * this.nameWidth) + this.textMargin,
      this.textMargin
    ];

    /**typeLog contains charsLogged logs of text overlays in a circular queue.
     * It is drawn over imageData with this.overlayData() until it is baked,
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

    /**cameraEnabled keeps track of whether the camera has been enabled by the
     * user. It's automatically disabled when getSendableData() is called.
     */
    this.cameraEnabled = false;
    //this.enableCamera();
  }

  /**-------------------------------------------------- Utils */
  getSendableData() {
    if (this.cameraEnabled) {
      this.imageData = this.camera.bakeImage(this.imageData);
      this.disableCamera();
    }
    this.bakeText();

    //If the image is empty, we return null.
    if (this.imageData.data.reduce((a,b)=>a+b) == 0) {
      return null
    }

    var compressed = this.imageData;
    compressed.data = this.notepad.rleEncode(compressed.data);
    return compressed;
  }

  clear() {
    this.imageData = {
      span: this.width,
      data: new Array(this.width * this.height).fill(0)
    };
    this.resetCursorPos();
    this.resetTypeLog();
    this.refresh();
  }

  copy(data) {
    this.imageData = data
    this.notepad.ctx.clearRect(0, 0, this.width, this.height);
    this.notepad.placeImageData(data)
  }

  refresh() {
    if (this.cameraEnabled) {
      this.camera.loadFrame();
    } else {
      this.notepad.ctx.clearRect(0, 0, this.width, this.height);
    }

    this.overlayData(this.imageData);

    if (this.typeLog[this.typeLogHead] != undefined) {
      this.overlayData(this.typeLog[this.typeLogHead]);
    }
  }

  overlayData(data) {
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

  /**-------------------------------------------------- Button handlers */
  setPenMode() {
    if (this.colourIndex != 0) {
      this.rainbowMode = !this.rainbowMode;
    }
    if (this.rainbowMode) {
      this.colourIndex = 4;
    } else {
      this.colourIndex = 1;
    }
  }

  setEraserMode() {
    this.rainbowMode = false;
    this.colourIndex = 0;
  }

  setPenSize(newPenSize) {
    this.cursorPos[1] -= Math.round(
      (this.lineSpacing / 2) * (this.pensize + 1)
    );
    this.pensize = newPenSize;
    this.cursorPos[1] += Math.round(
      (this.lineSpacing / 2) * (this.pensize + 1)
    );
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
    this.cursorPos = [
      Math.round(this.width * this.nameWidth) + this.textMargin,
      this.textMargin
    ];
  }

  /**-------------------------------------------------- Mouse drawing */
  drawPix(event) {
    this.bakeText();

    var [x, y] = this.getMousePixPos(event.offsetX, event.offsetY);

    for (var xo = x - this.pensize; xo <= x + this.pensize; xo++) {
      for (var yo = y - this.pensize; yo <= y + this.pensize; yo++) {
        this.imageData["data"][yo * this.width + xo] = this.colourIndex;
      }
    }
    this.notepad.setPixel(x, y, this.colourIndex, this.pensize);

    if (this.rainbowMode) {
      this.colourIndex = ((this.colourIndex - 3) % 60) + 4;
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
        for (var xo = tempx - this.pensize; xo <= tempx + this.pensize; xo++) {
          for (
            var yo = tempy - this.pensize;
            yo <= tempy + this.pensize;
            yo++
          ) {
            this.imageData["data"][yo * this.width + xo] = this.colourIndex;
          }
        }
        this.notepad.setPixel(tempx, tempy, this.colourIndex, this.pensize);
      }
      if (this.rainbowMode) {
        this.colourIndex = ((this.colourIndex - 3) % 59) + 4;
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
    /*Ensuring char is of length 1.*/
    char = char.slice(0, 1);
    /*Setting up the styling for the text and ascertaining its size*/
    if (this.pensize == 0) {
      this.notepad.ctx.font = "32px 'pixel 5x7'";
    } else {
      this.notepad.ctx.font = "64px 'pixel 5x7'";
    }
    this.notepad.ctx.fillStyle = this.notepad.getColour(this.colourIndex);
    this.notepad.ctx.textBaseline = "alphabetic";
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
        this.cursorPos[1] + 2 * (this.lineSpacing * (this.pensize / 1.5 + 1)) <=
        this.height
      ) {
        this.cursorPos[1] += Math.round(
          this.lineSpacing * (this.pensize / 1.5 + 1)
        );
      } else {
        this.cursorPos = [
          Math.round(this.width * this.nameWidth) + this.textMargin,
          this.textMargin
        ];
      }
    }
    this.cursorPos[0] += charWidth;

    /*In order to do the thresholding to fix the antialiasing of the text in the
    canvas, we write the text to the canvas to get the image data, then return
    the canvas to its original state and use the image data to decide where to
    draw the text onto the canvas ourselves.*/
    var oldData = this.notepad.ctx.getImageData(0, 0, this.width, this.height);
    this.notepad.ctx.fillText(
      char,
      this.cursorPos[0],
      this.cursorPos[1] +
        Math.round((this.lineSpacing / 2) * (this.pensize + 1))
    );
    var newData = this.notepad.ctx.getImageData(0, 0, this.width, this.height);
    this.notepad.ctx.putImageData(oldData, 0, 0);

    /**newText isn't added to typeLog unless the character actually made a
     * difference on the canvas (e.g. wasn't a space)
     */
    var diff = false;
    for (var i = 0; i < newData["data"].length; i += 4) {
      /**The alpha channel is used for masking out the text */
      if (
        // [me]
        // >Hey, uh, I'd like to check, uh...
        // >if two arrays are equal?
        //        O/                  O    [js]
        //       \|                  /|/   >?you wanna what?
        (oldData["data"][i] != newData["data"][i] ||
          oldData["data"][i + 1] != newData["data"][i + 1] ||
          oldData["data"][i + 2] != newData["data"][i + 2] ||
          oldData["data"][i + 3] != newData["data"][i + 3]) &&
        newData["data"][i + 3] == 255
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

  /**-------------------------------------------------- Camera */
  enableCamera() {
    if (this.camera == undefined) {
      this.camera = new Camera(this.notepad);
    }
    this.cameraEnabled = true;
    this.cameraInterval = setInterval(
      function() {
        this.refresh();
      }.bind(this),
      1000 / 30
    );
  }

  disableCamera() {
    this.cameraEnabled = false;
    clearInterval(this.cameraInterval);
  }
}

export default Sketchpad;
