import Vue from "vue";
import App from "./App.vue";
import VueAnalytics from "vue-ua";
import "./registerServiceWorker";
import store from "./store";
import router from "./router";
import "./assets/scss/style.scss";
import "./plugins/fontawesome-vue";

Vue.config.productionTip = false;

Vue.use(VueAnalytics, {
  appName: "Picto",
  appVersion: "0.0.1",
  trackingId: "UA-108088302-4",
  vueRouter: router, // Pass the router instance to automatically sync with router (optional)
  trackPage: true, // Whether you want page changes to be recorded as pageviews (website) or screenviews (app), default: false
  createOptions: {
    // Optional, Option when creating GA tracker, ref: https://developers.google.com/analytics/devguides/collection/analyticsjs/field-reference
    siteSpeedSampleRate: 10
  }
});

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");
