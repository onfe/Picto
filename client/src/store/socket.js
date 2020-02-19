import router from "../router";

const state = {
  _socket: null
};

const getters = {
  open: state => state._socket.readyState === WebSocket.OPEN
};

const actions = {
  connect: ({ commit, dispatch }, { name, room }) => {
    return new Promise((res, rej) => {
      const proto = window.location.protocol == "https:" ? "wss" : "ws";
      const here = window.location.host;
      const roomarg = room ? `&room=${room}` : "";
      const sock = new WebSocket(
        `${proto}://${here}/ws?name=${name}${roomarg}`
      );

      commit("create", sock);
      window._sock = sock;

      sock.onmessage = m => dispatch("_onMessage", m);
      sock.onopen = res;
      sock.onerror = rej;
      sock.onclose = () => dispatch("_onClose");
    });
  },
  disconnect: ({ state }) => {
    state._socket.close();
  },
  _onMessage: ({ dispatch }, pl) => {
    pl = JSON.parse(pl.data);
    console.log(pl);
    if (!pl.Event) {
      throw "Payload does not contain event field.";
    }
    const now = new Date();
    // eslint-disable-next-line no-console
    console.log(
      `[SOCK] (${now.getHours()}:${now.getMinutes()}:${now.getSeconds()}): ${
        pl.Event
      }`
    );

    switch (pl.Event) {
      case "message":
        dispatch("messages/add", pl.Payload, { root: true });
        break;
      case "init":
        dispatch("client/init", pl, { root: true });
        break;
      case "user":
        dispatch("client/updateUser", pl, { root: true });
        break;
      case "announcement":
        dispatch("messages/announce", pl.Payload, { root: true });
        break;
      case "rename":
        // TODO: ADD Rename event.
        break;
      default:
        // eslint-disable-next-line no-console
        console.log(pl);
    }
  },
  _onClose: ({ commit }) => {
    if (router.currentRoute.name == "room") {
      router.replace(`/join/${router.currentRoute.params.id}`);
    }
    commit("destroy");
  },
  send: ({ state }, pl) => {
    // TODO: check if connected, if not, dispatch socket/reconnect
    if (!pl.Event) {
      throw "Payload does not contain event field.";
    }

    state._socket.send(JSON.stringify(pl));
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
