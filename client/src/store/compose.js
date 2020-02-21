import RunlengthEncoder from "../assets/js/runlengthEncoder.js";

const state = {
  tool: "pencil",
  size: "small",
  rainbow: false
};

const getters = {};

const actions = {
  send: ({ dispatch }) => {
    const msg = window._sketch.getBakedImageData();
    if (msg != null) {
      //We don't send empty messages.
      const pl = {
        Event: "message",
        Time: 1000,
        Payload: {
          Message: {
            data: msg.data,
            span: msg.span
          }
        }
      };
      const socket_pl = {
        Event: "message",
        Time: 1000,
        Payload: {
          Message: {
            data: RunlengthEncoder.encode(msg.data),
            span: msg.span
          }
        }
      };

      dispatch("clear");
      dispatch("messages/addSelf", pl.Payload, { root: true });
      dispatch("socket/send", socket_pl, { root: true });
    }
  },
  clear: () => {
    window._sketch.clear();
  },
  copy: ({ rootState }, id) => {
    if (id != null) {
      // var msg = console.log('id');
    } else {
      console.log(rootState);
      var msg = rootState.messages.history.sort((a, b) => {
        a.id - b.id;
      })[0];
      console.log(msg.data.data);
      window._sketch.loadImageData(msg.data);
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
