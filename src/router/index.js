import Vue from "vue";
import Router from "vue-router";
import Meta from "vue-meta";
import store from "../store";
import Join from "../views/Join.vue";

Vue.use(Router);
Vue.use(Meta, {
  // optional pluginOptions
  refreshOnceOnNavigation: true
});

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
  },
  {
    path: "/room",
    redirect: "/"
  }
];

const router = new Router({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

router.afterEach(() => {
  if (store.state.client.errorMessage != "") {
    // when going to a new page, clear the error.
    store.dispatch("client/clearError");
  }
});

export default router;
