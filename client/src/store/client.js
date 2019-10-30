import router from "../router";

const axios = require("axios");

const state = {
  socket: null,
  username: '' || sessionStorage.getItem('client/username'),
};

const getters = {};

const actions = {};

const mutations = {};

const authentication = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default authentication;
