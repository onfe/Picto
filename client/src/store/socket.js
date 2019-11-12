const state = {
  _socket: null
};

const getters = {
  open: state => state._socket.readyState === WebSocket.OPEN
};

const actions = {
  connect: ({ commit, dispatch }, { name, room }) => {
    return new Promise((res, rej) => {
      const here = window.location.host;
      const roomarg = room ? `&room=${room}` : "";
      const sock = new WebSocket(`ws://${here}/ws?name=${name}${roomarg}`);
      commit("create", sock);
      window._sock = sock;
      sock.onmessage = m => dispatch("_onMessage", m);
      sock.onopen = res;
      sock.onerror = rej;
    });
  },
  _onMessage: (s, pl) => {
    const now = new Date();
    // eslint-disable-next-line no-console
    console.log(
      `[SOCK] (${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}): ${pl}`
    );
  },
  send: ({ state }, pl) => {
    // TODO: check if connected, if not, dispatch socket/reconnect
    state._socket.send(pl);
    const now = new Date();
    // eslint-disable-next-line no-console
    console.log(
      `[SOCK] (${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}): ${pl}`
    );
  },
  reconnect: () => {}
};

const mutations = {
  create: (state, sock) => {
    state._socket = sock;
  }
};

const client = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default client;
