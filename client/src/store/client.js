import router from "../router";
import COLOURS from "../assets/js/colours.js"

const state = {
  index: -1,
  colour: COLOURS[0],
  room: null,
  users: []
};

const getters = {
  username: state => state.users[state.index] || ""
};

const actions = {
  join: ({ dispatch }, { name, room }) => {
    dispatch("socket/connect", { name, room }, { root: true })
      .then(() => {
        // eslint-disable-next-line no-console
        console.log("Connected to Picto.");
      })
      .catch(() => {
        // eslint-disable-next-line no-console
        console.log("Failed to connect to Picto");
      });
  },
  init: ({ commit }, payload) => {
    commit("init", payload);
    router.push(`/room/${payload.RoomID}`);
  }
};

const mutations = {
  init: (state, payload) => {
    state.room = payload.RoomID;
    state.index = payload.UserIndex;
    state.users = payload.Users;
    state.colour = COLOURS[payload.UserIndex];
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
