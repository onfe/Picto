import axios from "../services/axios";

const state = {
  socket: null,
  username: "" || sessionStorage.getItem("client/username")
};

const getters = {};

const actions = {
  join: (s, name) => {
    axios({ url: "/api/ws", data: name, method: "POST" })
  }
};

const mutations = {};

const authentication = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default authentication;
