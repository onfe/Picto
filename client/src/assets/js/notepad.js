class Notepad {
  constructor(width, height, canvas) {
    this.canvas = canvas;
    this.ctx = this.canvas.getContext("2d");
    this.width = width;
    this.height = height;
    this.canvas.width = this.width;
    this.canvas.height = this.height;
  }

  loadImageData(data) {
    this.ctx.clearRect(0, 0, this.width, this.height);
    for (var i = 0; i < data["data"].length; i++) {
      if (data["data"][i] != 0) {
        this.setPixel(
          Math.round(i % data["span"]),
          Math.floor(i / data["span"]),
          data["data"][i]
        );
      }
    }
  }

  setPixel(x, y, i) {
    this.ctx.fillStyle = this.getColour(i);
    this.ctx.clearRect(x, y, 1, 1);
    this.ctx.fillRect(x, y, 1, 1);
  }

  getColour(i) {
    switch (i) {
      case 0:
        return "rgba(255,255,255,0)";
      case 1:
        return "rgba(0,0,0,1)";
      default:
        /**This is my (Josh's) nice way of generating rainbows.
         * No touchy. >:c
         * It goes
         * ___      ___      ___      ___      ___
         *    \    /|  \    /...\    /  |\    /   \    /
         *     \  /.|   \  /|...|\  /   |.\  /     \  /
         *      \/..|    \/ |...| \/    |..\/       \/
         *         ^up        ^flat      ^down
         */
        var toCode = (r, g, b) => {
          return "rgba(" + r + "," + g + "," + b + ",1)";
        };

        var up = i =>
          Math.round(255 - this.saturation + this.saturation * (i % 1));
        var flat = () => 255;
        var down = i =>
          Math.round(255 - this.saturation + this.saturation * (1 - (i % 1)));

        var ci = ((i - 2) / 254) * 3;

        switch (Math.floor(ci)) {
          case 0:
            return toCode(flat(ci), up(ci), down(ci));
          case 1:
            return toCode(down(ci), flat(ci), up(ci));
          case 2:
            return toCode(up(ci), down(ci), flat(ci));
        }
    }
  }
}

export default Notepad;
