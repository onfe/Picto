import Vue from "vue";
import Vuex from "vuex";
import client from "./client";
import socket from "./socket";
import messages from "./messages";
import compose from "./compose";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    client,
    socket,
    messages,
    compose
  }
});
