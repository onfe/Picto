import router from "../router";
import COLOURS from "../assets/js/colours.js";

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
  leave: ({ commit, dispatch }) => {
    dispatch("socket/disconnect", {}, { root: true });
    dispatch("messages/reset", {}, { root: true });
    commit("leave");
  },
  init: ({ commit }, payload) => {
    commit("init", payload);
    router.push(`/room/${payload.Payload.RoomID}`);
  },
  updateUser: ({ commit, state, dispatch }, d) => {
    const pl = d.Payload;
    if (pl.Users[pl.UserIndex] != "") {
      dispatch("messages/join", pl.UserName, { root: true });
    } else {
      dispatch("messages/leave", pl.UserName, { root: true });
    }
    if (d.Time > state.joinTime) {
      commit("updateUser", pl);
    }
  }
};

const mutations = {
  init: (state, d) => {
    state.room = d.Payload.RoomID;
    state.index = d.Payload.UserIndex;
    state.users = d.Payload.Users;
    state.colour = COLOURS[d.Payload.UserIndex];
    state.joinTime = d.Time;
  },
  updateUser: (state, payload) => {
    state.users = payload.Users;
  },
  leave: state => {
    state.room = null;
    state.index = -1;
    state.users = [];
    state.colour = COLOURS[0];
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
