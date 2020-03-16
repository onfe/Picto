import RunlengthEncoder from "./runlengthEncoder.js";

class Core {
  constructor(type, time) {
    this.type = type;
    this.time = time || Date.now();
  }
}

class Message extends Core {
  constructor(data, span, author, colour, time) {
    super("Message", time);
    this._data = data;
    this.span = span;
    this.author = author;
    this.colour = colour;
  }

  get data() {
    return [...this._data];
  }

  raw() {
    return {
      data: this.data,
      span: this.span
    };
  }

  encoded() {
    return {
      data: RunlengthEncoder.encode(this.data),
      span: this.span
    };
  }
}

class Announcement extends Core {
  constructor(text, time) {
    super("Announcement", time);
    this.text = text;
  }
}

class Text extends Core {
  constructor(text, time) {
    super("Text", time);
    this.text = text;
  }
}

export { Message, Announcement, Text };
