import router from "../router";
import COLOURS from "../assets/js/colours.js";
import { Announcement } from "../assets/js/message.js";

const state = {
  index: -1,
  colour: COLOURS[0],
  room: null,
  status: "idle",
  users: []
};

const getters = {
  username: state => state.users[state.index] || ""
};

const actions = {
  join: ({ commit, dispatch }, { name, room }) => {
    commit("updateStatus", "connecting");
    dispatch("socket/connect", { name, room }, { root: true })
      .then(() => {
        commit("updateStatus", "connected");
        // eslint-disable-next-line no-console
        console.log("Connected to Picto.");
      })
      .catch(() => {
        commit("updateStatus", "fail");
        // eslint-disable-next-line no-console
        console.log("Failed to connect to Picto");
      });
  },
  leave: ({ commit, dispatch }) => {
    dispatch("socket/disconnect", {}, { root: true });
    dispatch("messages/reset", {}, { root: true });
    commit("leave");
    if (router.currentRoute.name == "room") {
      router.replace(`/join/${router.currentRoute.params.id}`);
    }
  },
  init: ({ commit, dispatch }, payload) => {
    commit("init", payload);
    dispatch("messages/add", new Announcement("Welcome to Picto!"), { root: true });
    router.push(`/room/${payload.Payload.RoomID}`);
  },
  updateUser: ({ commit, state, dispatch }, d) => {
    const pl = d.Payload;
    const data = {
      name: pl.UserName,
      time: d.Time
    };
    if (pl.Users[pl.UserIndex] != "") {
      dispatch("messages/join", data, { root: true });
    } else {
      dispatch("messages/leave", data, { root: true });
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
    state.status = "idle";
  },
  updateStatus: (state, payload) => {
    state.status = payload;
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
