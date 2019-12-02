const state = {
  tool: "pencil",
  size: "small"
};

const getters = {};

const actions = {
  send: ({ dispatch }) => {
    const data = window._sketch.getSendableData();
    const pl = {
      Event: "message",
      Message: data
    };
    dispatch("messages/addSelf", pl, { root: true });
    dispatch("socket/send", pl, { root: true });
  },
  clear: () => {
    window._sketch.clear();
  },
  pencil: ({ commit }) => {
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
  }
};

const mutations = {
  pencil: state => {
    state.tool = "pencil";
  },
  eraser: state => {
    state.tool = "eraser";
  },
  large: state => {
    state.size = "large";
  },
  small: state => {
    state.size = "small";
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
