import COLOURS from "../assets/js/colours.js";

const state = {
  history: [],
  iter: 0
};

const getters = {};

const actions = {
  add: ({ state, rootState, commit }, pl) => {
    const message = {
      type: "normal",
      author: rootState.client.users[pl.UserIndex],
      colour: COLOURS[pl.UserIndex],
      data: pl.Message,
      id: state.iter
    };
    commit("add", message);
  },
  addSelf: ({ rootState, commit }, pl) => {
    const message = {
      type: "normal",
      author: rootState.client.users[rootState.client.index],
      colour: COLOURS[rootState.client.index],
      data: pl.Message
    };
    commit("add", message);
  },
  announce: ({ commit }, pl) => {
    const message = {
      type: "announcement",
      text: pl.Announcement
    };
    commit("add", message);
  }
};

const mutations = {
  add: (state, message) => {
    state.history.unshift(message);
    state.iter += 1;
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
