import router from "../router";
import colour from "../assets/js/colours.js";
import Vue from "vue";

const state = {
  username: null,
  colour: colour(0),
  status: "idle",
  errorMessage: "",
  swStatus: "normal"
};

const actions = {
  join: ({ commit, dispatch }, { name, room }) => {
    commit("updateStatus", "connecting");
    dispatch("socket/connect", { name, room }, { root: true })
      .then(() => {
        Vue.analytics.trackEvent("client", "connect", "success");
        commit("updateStatus", "connected");
      })
      .catch(() => {
        Vue.analytics.trackEvent("client", "connect", "failure");
        commit("updateStatus", "fail");
        commit("updateError", "Couldn't connect to Picto.");
        // eslint-disable-next-line no-console
        console.log("Failed to connect to Picto");
      });
  },
  leave: ({ commit, dispatch }) => {
    dispatch("socket/disconnect", {}, { root: true });
    dispatch("messages/reset", {}, { root: true });
    dispatch("room/leave", {}, { root: true });
    Vue.analytics.trackEvent("client", "disconnect");
    commit("leave");
    if (router.currentRoute.name == "room") {
      router.replace(`/join/${router.currentRoute.params.id}`);
    }
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
  leave: state => {
    state.username = null;
    state.colour = colour(0);
    state.status = state.status == "connected" ? "idle" : state.status;
  },
  joined: (state, pl) => {
    state.username = pl.username;
    state.colour = pl.colour;
  },
  updateError: (state, error) => {
    state.errorMessage = error;
  },
  updateStatus: (state, payload) => {
    state.status = payload;
  },
  swUpdate: (state, status) => {
    state.swStatus = status;
  }
};

const client = {
  namespaced: true,
  state: state,
  actions: actions,
  mutations: mutations
};

export default client;
