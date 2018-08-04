module.exports = class Sound {
  constructor() {
    this.join = new Audio('/snd/bell.mp3');
    this.button = new Audio('/snd/click.mp3');
    this.message = new Audio('/snd/message.mp3');
  }
}
