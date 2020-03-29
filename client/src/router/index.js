import Vue from "vue";
import VueRouter from "vue-router";
import store from "../store";
import Join from "../views/Join.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: Join
  },
  {
    path: "/join/:id",
    name: "join",
    component: Join
  },
  {
    path: "/join",
    redirect: "/"
  },
  {
    path: "/room/:id",
    name: "room",
    component: () => import(/* webpackChunkName: "room" */ "../views/Room.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.afterEach(() => {
  if (store.state.client.errorMessage != "") {
    // when going to a new page, clear the error.
    store.dispatch("client/error", { code: -1, reason: "" });
  }
});

export default router;
