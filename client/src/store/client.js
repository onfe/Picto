// import axios from "../services/axios";

const state = {
  socket: null,
  username: "" || sessionStorage.getItem("client/username")
};

const getters = {};

const actions = {
  join: ({ commit }, name) => {
    const sock = new WebSocket(`ws://${window.location.host}/ws?name=${name}`);
    commit("setSock", sock);
    window._sock = sock;
    sock.onopen = () => {
      setTimeout(() => sock.send("aiii"), 1000);
    };
  }
};

const mutations = {
  setSock: (state, sock) => {
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
