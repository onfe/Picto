import RunlengthEncoder from "./runlengthEncoder.js";

class Message {
  constructor(data, span) {
    this._data = data;
    this._span = span;
  }

  get data() {
    return [...this._data];
  }

  get span() {
    return this._span;
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

export default Message;
