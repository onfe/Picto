// import axios from "../services/axios";

const state = {
  socket: null,
  username: "" || sessionStorage.getItem("client/username")
};

const getters = {};

const actions = {
  join: ({ commit }, { name, room }) => {
    const here = window.location.host;
    const roomarg = room ? `&room=${room}` : "";
    const sock = new WebSocket(`ws://${here}/ws?name=${name}${roomarg}`);
    commit("join", sock);
    window._sock = sock;
    sock.onopen = () => {
      console.log("Connected to picto!");
    };
  }
};

const mutations = {
  join: (state, sock) => {
    state.socket = sock;
  }
};

const authentication = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default authentication;
