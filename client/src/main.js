import Vue from "vue";
import App from "./App.vue";
import "./registerServiceWorker";
import store from "./store";
import router from "./router";
import "./assets/scss/style.scss";
import "./plugins/fontawesome-vue";

Vue.config.productionTip = false;

new Vue({
  store,
  router,
  render: h => h(App)
}).$mount("#app");

import VueAnalytics from "vue-ua";

Vue.use(VueAnalytics, {
  appName: "Picto",
  appVersion: "0.0.1",
  trackingId: "UA-108088302-3",
  debug: false,
  vueRouter: router, // Pass the router instance to automatically sync with router (optional)
  // ignoredViews: ['homepage'], // If router, you can exclude some routes name (case insensitive) (optional)
  trackPage: false, // Whether you want page changes to be recorded as pageviews (website) or screenviews (app), default: false
  createOptions: {
    // Optional, Option when creating GA tracker, ref: https://developers.google.com/analytics/devguides/collection/analyticsjs/field-reference
    siteSpeedSampleRate: 10
  }
  // globalDimensions: [ // Optional
  //   {dimension: 1, value: 'MyDimensionValue'},
  //   {dimension: 2, value: 'AnotherDimensionValue'}
  // ],
  // globalMetrics: [ // Optional
  //     {metric: 1, value: 'MyMetricValue'},
  //     {metric: 2, value: 'AnotherMetricValue'}
  //   ]
});
