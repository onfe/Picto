const state = {};

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
  }
};

const mutations = {};

const client = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default client;
