const state = {
  tool: "pencil",
  size: "small",
  rainbow: false
};

const getters = {};

const actions = {
  send: ({ dispatch }) => {
    const data = window._sketch.getSendableData();
    const pl = {
      Event: "message",
      Message: data
    };
    dispatch("clear");
    dispatch("messages/addSelf", pl, { root: true });
    dispatch("socket/send", pl, { root: true });
  },
  clear: () => {
    window._sketch.clear();
  },
  copy: ({ rootState }, id) => {
    if (id != null) {
      // var msg = console.log('id');
    } else {
      var msg = rootState.messages.history.sort((a, b) => {
        a.id - b.id;
      })[0];
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
