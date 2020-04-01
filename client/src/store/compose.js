import { Message } from "../assets/js/message.js";
import Vue from "vue";

const state = {
  tool: "pencil",
  size: "small",
  rainbow: false
};

const getters = {};

const actions = {
  send: ({ dispatch, rootState, rootGetters }) => {
    const raw = window._sketch.getBakedImageData();
    if (raw == null) {
      // Don't send empty messages.
      return;
    }

    const msg = new Message(
      raw.data,
      raw.span,
      rootGetters["client/username"],
      rootState.client.colour
    );

    dispatch("clear");
    Vue.analytics.trackEvent("message", "send");
    dispatch("messages/add", msg, { root: true });
    dispatch(
      "socket/send",
      { event: "message", payload: msg.encoded() },
      { root: true }
    );
  },
  clear: () => {
    window._sketch.clear();
  },
  copy: ({ rootState }, msg) => {
    if (!msg) {
      var msgs = rootState.messages.history
        .sort((a, b) => {
          a.id - b.id;
        })
        .filter(a => a.type === "Message");
      if (msgs.length < 1) {
        // eslint-disable-next-line no-console
        console.error("No messages to copy!");
        return;
      }
      msg = msgs[msgs.length-1];
      window._sketch.loadImageData(msg.raw());
    } else {
      window._sketch.loadImageData(msg.raw());
    }
  },
  pencil: ({ commit, state }) => {
    commit("rainbow", state.tool == "pencil" && !state.rainbow);
    window._sketch.setPenMode();
    commit("pencil");
  },
  eraser: ({ commit }) => {
    window._sketch.setEraserMode();
    commit("eraser");
  },
  large: ({ commit }) => {
    window._sketch.pensize = 1;
    commit("large");
  },
  small: ({ commit }) => {
    window._sketch.pensize = 0;
    commit("small");
  },
  write: (_, chr) => {
    window._sketch.drawChar(chr);
  },
  backspace: () => {
    window._sketch.backspace();
  }
};

const mutations = {
  pencil: state => {
    state.tool = "pencil";
  },
  eraser: state => {
    state.tool = "eraser";
    state.rainbow = false;
  },
  large: state => {
    state.size = "large";
  },
  small: state => {
    state.size = "small";
  },
  rainbow: (state, to) => {
    state.rainbow = to;
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
