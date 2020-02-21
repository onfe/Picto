const itobs =
  "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/";

class RunlengthEncoder {
    static encode(d) {
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

  static decode(d) {
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

export default RunlengthEncoder;