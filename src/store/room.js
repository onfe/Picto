import { Text } from "../assets/js/message.js";
import Vue from "vue";

const state = {
  id: null,
  name: null,
  showInfo: false,
  public: false,
  users: []
};

const getters = {
  title: state => (state.name ? state.name : state.id),
  invite: state => `${window.location.origin}/join/${state.id}`
};

const actions = {
  join: () => {
    // replaces client/init
    // TODO: refactor
  },
  leave: () => {
    // called with client/leave
    // TODO: refactor
  },
  updateUsers: () => {
    // triggered on user join/leave
    // TODO: refactor
  },
  rename: ({ commit, dispatch }, pl) => {
    // replaces client/renameRoom
    const user = pl.Payload.UserName;
    const name = pl.Payload.RoomName;

    var message = `${user} named the room '${name}'.`;
    if (name.length == 0) {
      message = `${user} removed the room name.`;
    }

    Vue.analytics.trackEvent("room", "rename");
    dispatch("messages/add", new Text(message, pl.Time), { root: true });
    commit("rename", name);
  }
};

const mutations = {
  rename: (state, name) => {
    state.name = name;
  }
};

const room = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default room;
