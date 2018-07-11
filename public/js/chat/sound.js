module.exports = class Sound {
  constructor() {
    this.join = new Audio('/snd/bell.mp3');
    this.button = new Audio('/snd/button.wav');
    this.message = new Audio('/snd/message.wav');
  }
}
