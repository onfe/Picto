import router from "../router";
import colour from "../assets/js/colours.js";
import { Announcement, Text } from "../assets/js/message.js";
import Vue from "vue";

const state = {
  index: -1,
  colour: colour(0),
  room: null,
  roomName: "",
  status: "idle",
  users: [],
  showInfo: false,
  errorMessage: "",
  static: false,
  swStatus: "normal"
};

const getters = {
  username: state => state.users[state.index] || "",
  roomTitle: state => (state.roomName.length > 0 ? state.roomName : state.room),
  userColours: state =>
    state.users.map((k, i) => [k, colour(i)]).filter(e => e[0]),
  inviteLink: state => `${window.location.origin}/join/${state.room}`
};

const actions = {
  join: ({ commit, dispatch }, { name, room }) => {
    commit("updateStatus", "connecting");
    dispatch("socket/connect", { name, room }, { root: true })
      .then(() => {
        Vue.analytics.trackEvent("room", "connect", "success");
        commit("updateStatus", "connected");
      })
      .catch(() => {
        Vue.analytics.trackEvent("room", "connect", "failure");
        commit("updateStatus", "fail");
        commit("updateError", "Couldn't connect to Picto.");
        // eslint-disable-next-line no-console
        console.log("Failed to connect to Picto");
      });
  },
  leave: ({ commit, dispatch }) => {
    dispatch("socket/disconnect", {}, { root: true });
    dispatch("messages/reset", {}, { root: true });
    Vue.analytics.trackEvent("room", "leave");
    commit("leave");
    if (router.currentRoute.name == "room") {
      router.replace(`/join/${router.currentRoute.params.id}`);
    }
  },
  init: ({ commit, dispatch }, payload) => {
    commit("init", payload);
    Vue.analytics.trackEvent("room", "join", "success");
    // eslint-disable-next-line no-console
    console.log("Connected to Picto.");
    dispatch("messages/add", new Announcement("Welcome to Picto!"), {
      root: true
    });
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
  },
  toggleInfo: ({ commit }) => {
    commit("toggleInfo");
  },
  renameRoom: ({ commit, dispatch }, pl) => {
    const user = pl.Payload.UserName;
    const name = pl.Payload.RoomName;
    var message = `${user} named the room '${name}'.`;
    if (name.length == 0) {
      message = `${user} removed the room name.`;
    }
    Vue.analytics.trackEvent("room", "rename");
    dispatch("messages/add", new Text(message, pl.Time), { root: true });
    commit("renameRoom", name);
  },
  error: ({ commit }, error) => {
    Vue.analytics.trackEvent("room", "join", "error", error.code);
    commit("updateError", error.reason);
    commit("updateStatus", "fail");
  },
  clearError: ({ commit }) => {
    commit("updateError", "");
  },
  swUpdate: ({ commit }, type) => {
    if (type == "offline") {
      commit("updateError", "Picto is offline.");
    }
    commit("swUpdate", type);
  }
};

const mutations = {
  init: (state, d) => {
    state.room = d.Payload.RoomID;
    state.index = d.Payload.UserIndex;
    state.users = d.Payload.Users;
    state.colour = colour(d.Payload.UserIndex);
    state.roomName = d.Payload.RoomName;
    state.joinTime = d.Time;
    state.showInfo = false;
    state.errorMessage = "";
    state.errorCode = -1;
    state.static = d.Payload.Static;
  },
  updateUser: (state, payload) => {
    state.users = payload.Users;
  },
  leave: state => {
    state.room = null;
    state.index = -1;
    state.users = [];
    state.colour = colour(0);
    state.status = state.status == "connected" ? "idle" : state.status;
  },
  updateError: (state, error) => {
    state.errorMessage = error;
  },
  updateStatus: (state, payload) => {
    state.status = payload;
  },
  toggleInfo: state => {
    state.showInfo = !state.showInfo;
  },
  renameRoom: (state, name) => {
    state.roomName = name;
  },
  swUpdate: (state, status) => {
    state.swStatus = status;
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
