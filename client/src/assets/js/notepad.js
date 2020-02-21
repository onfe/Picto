const itobs =
  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

class Notepad {
  constructor(width, height, canvas) {
    this.canvas = canvas;
    this.ctx = this.canvas.getContext("2d");
    this.width = width;
    this.height = height;
    this.canvas.width = this.width;
    this.canvas.height = this.height;
    this.saturation = 255; /*Saturation should be in range 0-255 inclusive*/
  }

  placeImageData(data) {
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

  loadImageData(data) {
    this.ctx.clearRect(0, 0, this.width, this.height);
    data.data = this.rleDecode(data.data);
    this.placeImageData(data);
  }

  setPixel(x, y, i, s) {
    s = s || 0;
    this.ctx.fillStyle = this.getColour(i);
    this.ctx.clearRect(x - s, y - s, 1 + s * 2, 1 + s * 2);
    this.ctx.fillRect(x - s, y - s, 1 + s * 2, 1 + s * 2);
  }

  getColour(i) {
    switch (i) {
      case 0:
        return "rgba(255,255,255,0)";
      case 1:
        return "rgba(0,0,0,1)";
      case 2:
        return "rgba(85,85,85,1)";
      case 3:
        return "rgba(170,170,170,1)";
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

        var ci = ((i - 4) / 59) * 3;

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

  rleEncode(d) {
    var occurences = 0;
    var prev = null;
    var resArr = [];
    for (var i = 0; i <= d.length; i++) {
      if (prev != d[i] || i == d.length) {
        if (occurences > 4) {
          occurences -= 4;
          resArr.push(63);
          while (occurences > 63) {
            resArr.push(63);
            occurences -= 63;
          }
          resArr.push(occurences);
          resArr.push(0, prev);
        } else if (occurences > 0) {
          resArr.push(...Array(occurences).fill(prev));
        }
        occurences = 1;
        prev = d[i];
      } else {
        occurences += 1;
      }
    }
    var res = resArr.map(i => itobs[i]).join("");
    return res;
  }

  rleDecode(d) {
    d = d.split("").map(char => itobs.indexOf(char));
    var res = [];
    for (var i = 0; i < d.length; i++) {
      if (d[i] == 63) {
        var count = 4;
        i += 1;
        for (i; d[i] != 0; i++) {
          count += d[i];
        }
        res.push(...Array(count).fill(d[i + 1]));
        i += 1;
      } else {
        res.push(d[i]);
      }
    }
    return res;
  }
}

export default Notepad;
