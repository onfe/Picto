import RunlengthEncoder from "./runlengthEncoder.js";

class Core {
  constructor(type, time) {
    this.type = type;
    this.time = time || Date.now();
    this.hash = Math.random()
      .toString(36)
      .substring(2, 15);
  }
}

class Message extends Core {
  constructor(data, span, author, colour, time) {
    super("Message", time);
    this._data = data;
    this.span = span;
    this.author = author;
    this.colour = colour;
    this.hidden = false;
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
    this.colour = "#000";
  }
}

class Text extends Core {
  constructor(text, time) {
    super("Text", time);
    this.text = text;
    this.colour = "#808080";
  }
}

export { Message, Announcement, Text };
