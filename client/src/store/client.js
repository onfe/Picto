const state = {
  socket: null,
  username: ""
};

const getters = {};

const actions = {
  join: ({ dispatch }, { name, room }) => {
    dispatch("socket/connect", { name, room }, { root: true })
    .then(() => {
      console.log("Connected to Picto.");
    })
    .catch(() => {
      console.log("Failed to connect to Picto");
    })
  }
};

const mutations = {
  join: (state, sock) => {
    state.socket = sock;
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
