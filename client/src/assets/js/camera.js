class Camera {
  constructor(notepad) {
    this.notepad = notepad;
    this.video = document.createElement("video");
    this.video.autoplay = true;

    this.imageData = {
      span: this.notepad.width,
      data: new Array(this.notepad.width * this.notepad.height).fill(0)
    };

    if (navigator.mediaDevices.getUserMedia) {
      navigator.mediaDevices.getUserMedia({ video: true }).then(
        function(stream) {
          this.video.srcObject = stream;
        }.bind(this)
      );
    }
  }

  bakeImage(drawingData) {
    this.loadFrame();
    for (var i = 0; i < drawingData["data"].length; i++) {
      if (drawingData["data"][i] == 0 && this.imageData["data"][i] != 0) {
        drawingData["data"][i] = this.imageData["data"][i];
      }
    }
    return drawingData;
  }

  loadFrame() {
    var source_height = Math.round(
      this.notepad.height / (this.notepad.width / this.video.videoWidth)
    );

    this.notepad.ctx.drawImage(
      this.video,
      0,
      (this.video.videoHeight - source_height) / 2,
      this.video.videoWidth,
      source_height,
      0,
      0,
      this.notepad.width,
      this.notepad.height
    );

    this.img = this.notepad.ctx.getImageData(
      0,
      0,
      this.notepad.width,
      this.notepad.height
    );

    var nearest = x => (x < 128 ? 0 : 255);

    var stride = this.notepad.width * 4;
    for (var i = 0; i < this.img.data.length - 1; i += 4) {
      var old =
        (this.img.data[i] + this.img.data[i + 1] + this.img.data[i + 2]) / 3;

      var n = nearest(old);
      var err = old - n;

      var colourCode = n == 0 ? 1 : 0;
      this.imageData["data"][Math.floor(i / 4)] = colourCode;
      this.notepad.setPixel(
        Math.floor(i / 4) % this.imageData["span"],
        Math.floor(i / 4 / this.imageData["span"]),
        colourCode
      );

      this.img.data[i + 4] += (err * 7) / 16;
      this.img.data[i + 5] += (err * 7) / 16;
      this.img.data[i + 6] += (err * 7) / 16;

      this.img.data[i + stride - 4] += (err * 3) / 16;
      this.img.data[i + stride - 3] += (err * 3) / 16;
      this.img.data[i + stride - 2] += (err * 3) / 16;

      this.img.data[i + stride] += (err * 5) / 16;
      this.img.data[i + stride + 1] += (err * 5) / 16;
      this.img.data[i + stride + 2] += (err * 5) / 16;

      this.img.data[i + stride + 4] += (err * 1) / 16;
      this.img.data[i + stride + 5] += (err * 1) / 16;
      this.img.data[i + stride + 6] += (err * 1) / 16;
    }
  }
}

export default Camera;
