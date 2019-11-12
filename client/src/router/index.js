import Vue from "vue";
import VueRouter from "vue-router";
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

export default router;
