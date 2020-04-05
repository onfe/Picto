import colour from "../assets/js/colours.js";
import RunlengthEncoder from "../assets/js/runlengthEncoder.js";
import { Message, Announcement, Text } from "../assets/js/message.js";

const state = {
  history: [],
  unread: 0
};

const getters = {};

const actions = {
  add: ({ commit }, message) => {
    commit("add", message);
  },
  message: ({ commit }, d) => {
    const pl = d.Payload;
    pl.Data = RunlengthEncoder.decode(pl.Data);
    const message = new Message(
      pl.Data,
      pl.Span,
      pl.Sender,
      colour(pl.ColourIndex),
      d.Time
    );
    commit("add", message);
  },
  announce: ({ commit }, d) => {
    const announce = new Announcement(d.Payload.Announcement, d.Time);
    commit("add", announce);
  },
  join: ({ commit }, pl) => {
    const text = new Text(`${pl.name} joined.`, pl.time);
    commit("add", text);
  },
  leave: ({ commit }, pl) => {
    const text = new Text(`${pl.name} left.`, pl.time);
    commit("add", text);
  },
  reset: ({ commit }) => {
    commit("reset");
  },
  toggleHidden: ({ commit }, msg) => {
    commit("toggleHidden", msg);
  },
  read: ({ commit }) => {
    commit("read");
  }
};

const mutations = {
  add: (state, message) => {
    if (document.visibilityState == "hidden") {
      state.unread++;
    }
    state.history.unshift(message);
    state.history.sort((a, b) => b.time - a.time);
    state.history.length = Math.min(state.history.length, 32);
  },
  toggleHidden: (state, message) => {
    for (var i = 0; i < state.history.length; i++) {
      if (state.history[i].hash == message.hash) {
        state.history[i].hidden = !state.history[i].hidden;
        return;
      }
    }
  },
  reset: state => {
    state.history = [];
  },
  read: state => {
    state.unread = 0;
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
