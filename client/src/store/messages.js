import COLOURS from "../assets/js/colours.js";
import RunlengthEncoder from "../assets/js/runlengthEncoder.js";

const state = {
  history: []
};

const getters = {};

const actions = {
  add: ({ commit }, d) => {
    const pl = d.Payload;
    pl.Message.data = RunlengthEncoder.decode(pl.Message.data);
    const message = {
      type: "normal",
      author: pl.Sender,
      colour: COLOURS[pl.ColourIndex],
      data: pl.Message,
      id: d.Time
    };
    commit("add", message);
  },
  addSelf: ({ rootState, commit }, pl) => {
    console.log(pl);
    const message = {
      type: "normal",
      author: rootState.client.users[rootState.client.index],
      colour: COLOURS[rootState.client.index],
      data: pl,
      id: Date.now()
    };
    commit("add", message);
  },
  announce: ({ commit }, d) => {
    const pl = d.Payload;
    const message = {
      type: "announcement",
      text: pl.Announcement,
      id: d.Time
    };
    commit("add", message);
  },
  join: ({ commit }, pl) => {
    const message = {
      text: `${pl.name} joined.`,
      id: pl.time
    };
    commit("add", message);
  },
  leave: ({ commit }, pl) => {
    const message = {
      text: `${pl.name} left.`,
      id: pl.Time
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
    state.history.sort((a, b) => b.id - a.id);
  },
  reset: state => {
    state.history = [];
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
