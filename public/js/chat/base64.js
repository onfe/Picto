module.exports = class Base64 {

  static decode(st) {
    var out = '';
    var base64List = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=';

    for (var i = 0; i < st.length; i++) {
      var char = st[i];
      var num = base64List.indexOf(char);
      var bin = num.toString(2);
      var final = ('000000' + bin).slice(-6); // ensure is 6 chars
      out += final;
    }

    return out;
  }

  static encode(st) {
    var out = '';
    var base64List = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-=';
    var chunks = this.chunkString(st, 6);
    for (var c = 0; c < chunks.length; c++) {
      var num = parseInt(chunks[c], 2) // Binary string to decimal integer.
      var char = base64List[num];
      out += char.toString(2);

    }
    return out;
  }

  static chunkString(str, length) {
    return str.match(new RegExp('.{1,' + length + '}', 'g'));
  }
}
