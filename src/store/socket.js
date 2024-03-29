const state = {
  _socket: null
};

const getters = {
  open: state =>
    state._socket !== null && state._socket.readyState === WebSocket.OPEN
};

const actions = {
  connect: ({ commit, dispatch }, { name, room }) => {
    return new Promise((res, rej) => {
      const proto = window.location.protocol == "https:" ? "wss" : "ws";
      const here = window.location.host;
      const roomarg = room ? `&room=${room}` : "";
      var sock = new WebSocket(`${proto}://${here}/ws?name=${name}${roomarg}`);

      commit("create", sock);
      window._sock = sock;

      sock.onmessage = m => dispatch("_onMessage", m);
      sock.onopen = () => {
        res();
      };
      sock.onerror = e => rej(e);
      sock.onclose = e => dispatch("_onClose", e);
    });
  },
  disconnect: ({ state, commit }) => {
    if (state._socket !== null) {
      state._socket.close();
    }
    commit("destroy");
  },
  _onMessage: ({ dispatch }, pl) => {
    pl = JSON.parse(pl.data);
    if (!pl.Event) {
      throw "Payload does not contain event field.";
    }
    const now = new Date();
    // eslint-disable-next-line no-console
    console.log(
      `[SOCK][R] (${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}):`,
      pl
    );

    switch (pl.Event) {
      case "message":
        dispatch("messages/message", pl, { root: true });
        break;
      case "init":
        dispatch("room/join", pl, { root: true });
        break;
      case "user":
        dispatch("room/updateUsers", pl, { root: true });
        break;
      case "announcement":
        dispatch("messages/announce", pl, { root: true });
        break;
      case "rename":
        dispatch("room/rename", pl, { root: true });
        break;
      default:
        // eslint-disable-next-line no-console
        console.log(pl);
    }
  },
  _onClose: ({ commit, dispatch }, e) => {
    commit("destroy");
    if (4000 <= e.code && e.code < 5000) {
      // If 4000 <= code < 5000, it's a custom code to display an error
      dispatch("client/error", e, { root: true });
    }
    dispatch("client/leave", {}, { root: true });
  },
  send: ({ state }, { event, payload }) => {
    const packet = {
      Time: Date.now(),
      Event: event,
      Payload: payload
    };

    state._socket.send(JSON.stringify(packet));
    const now = new Date();
    // eslint-disable-next-line no-console
    console.log(
      `[SOCK][S] (${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}):`,
      packet
    );
  },
  reconnect: () => {}
};

const mutations = {
  create: (state, sock) => {
    state._socket = sock;
  },
  destroy: state => {
    state._socket = null;
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
