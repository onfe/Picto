import { Text, Announcement } from "../assets/js/message.js";
import colour from "../assets/js/colours.js";
import router from "../router";
import Vue from "vue";

const state = {
  id: null,
  name: null,
  showInfo: false,
  public: false,
  joinTime: -1,
  users: []
};

const getters = {
  title: state => (state.name ? state.name : state.id),
  invite: state => `${window.location.origin}/join/${state.id}`,
  colours: state => state.users.map((k, i) => [k, colour(i)]).filter(e => e[0])
};

const actions = {
  join: ({ commit, dispatch }, pl) => {
    // replaces client/init
    commit("join", pl);
    Vue.analytics.trackEvent("room", "join", "success");
    dispatch("messages/add", new Announcement("Welcome to Picto!"), {
      root: true
    });
    router.push(`/room/${pl.Payload.RoomID}`);
    // the client needs it's confirmed username and colour.
    commit(
      "client/joined",
      {
        username: pl.Payload.Users[pl.Payload.UserIndex],
        colour: colour(pl.Payload.UserIndex)
      },
      { root: true }
    );
  },
  leave: ({ commit }) => {
    // supplements client/leave
    commit("leave");
    Vue.analytics.trackEvent("room", "leave");
  },
  updateUsers: ({ commit, dispatch, state }, d) => {
    console.log("User trigger");
    // triggered on user join/leave
    const pl = d.Payload;
    const update = {
      name: pl.UserName,
      time: d.Time
    };

    // Add a note to the message history.
    if (pl.Users[pl.UserIndex] != "") {
      dispatch("messages/join", update, { root: true });
    } else {
      dispatch("messages/leave", update, { root: true });
    }

    console.log(state);

    // If the update is current (after we joined), mutate the user list.
    if (d.Time > state.joinTime) {
      commit("updateUsers", pl);
    }
  },
  rename: ({ commit, dispatch }, pl) => {
    // replaces client/renameRoom
    const user = pl.Payload.UserName;
    const name = pl.Payload.RoomName;

    var message = `${user} named the room '${name}'.`;
    if (name.length == 0) {
      message = `${user} removed the room name.`;
    }

    Vue.analytics.trackEvent("room", "rename");
    dispatch("messages/add", new Text(message, pl.Time), { root: true });
    commit("rename", name);
  },
  toggleInfo: ({ commit }) => {
    commit("toggleInfo");
  }
};

const mutations = {
  join: (state, d) => {
    const pl = d.Payload;
    state.id = pl.RoomID;
    state.users = pl.Users;
    state.name = pl.Name;
    state.public = pl.Public;
    state.showInfo = false;
    state.joinTime = Date.now();
  },
  leave: state => {
    state.id = null;
    state.name = null;
    state.users = [];
    state.public = false;
    state.showInfo = false;
    state.joinTime = -1;
  },
  updateUsers: (state, pl) => {
    state.users = pl.Users;
  },
  rename: (state, name) => {
    state.name = name;
  },
  toggleInfo: state => {
    state.showInfo = !state.showInfo;
  }
};

const room = {
  namespaced: true,
  state: state,
  getters: getters,
  actions: actions,
  mutations: mutations
};

export default room;
