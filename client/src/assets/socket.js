export default class Socket {
  constructor(username, room, handle) {
    const HOST = window.location.host;
    this.ws = new WebSocket(`ws://${HOST}`);

    this.ws.onopen = this.onOpen.bind(this);
    this.ws.onmessage = handle;

    this.username = username;
    this.room = rc;
    this.verified = false;
  }

  send(type, payload) {
    const msg = {
      type,
      // time: Date.now(),
      payload
    }
    this.ws.send(JSON.stringify(msg));
  }

  onOpen() {
    console.log('Connected to Picto...');
    this.join();
    this.send('join-request', {});
  }
}
