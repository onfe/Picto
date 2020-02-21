import COLOURS from "../assets/js/colours.js";
import RunlengthEncoder from "../assets/js/runlengthEncoder.js";

const state = {
  history: [],
  iter: 0
};

const getters = {};

const actions = {
  add: ({ state, commit }, pl) => {
    pl.Message.data = RunlengthEncoder.decode(pl.Message.data)
    const message = {
      type: "normal",
      author: pl.Sender,
      colour: COLOURS[pl.ColourIndex],
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
      data: pl.Message,
      id: state.iter
    };
    commit("add", message);
  },
  announce: ({ commit }, pl) => {
    const message = {
      type: "announcement",
      text: pl.Announcement
    };
    commit("add", message);
  },
  join: ({ commit }, pl) => {
    const message = {
      text: `${pl} joined.`
    };
    commit("add", message);
  },
  leave: ({ commit }, pl) => {
    const message = {
      text: `${pl} left.`
    };
    commit("add", message);
  },
  reset: ({ commit }) => {
    commit("reset");
  }
};

const mutations = {
  add: (state, message) => {
    state.history.unshift(message);
    state.iter += 1;
  },
  reset: state => {
    state.history = [];
    state.iter = 0;
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
