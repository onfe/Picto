import COLOURS from "../assets/js/colours.js";

const state = {
  history: [
    { author: "onfe", data: "ahsdald", colour: "#ffcced" },
    { author: "pedanticat", data: "ahsdald", colour: "#00ffdd" }
  ]
};

const getters = {};

const actions = {
  add: ({ rootState, commit }, pl) => {
    const message = {
      author: rootState.client.users[pl.UserIndex],
      colour: COLOURS[pl.UserIndex],
      data: pl.Message
    }
    commit("add", message)
  }
};

const mutations = {
  add: (state, message) => {
    state.history = [message, ...state.history];
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
