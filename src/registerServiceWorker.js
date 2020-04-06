/* eslint-disable no-console */

import { register } from "register-service-worker";

if (process.env.NODE_ENV === "production") {
  register(`${process.env.BASE_URL}service-worker.js`, {
    ready() {
      console.log(
        "App is being served from cache by a service worker.\n" +
          "For more details, visit https://goo.gl/AFskqB"
      );
    },
    registered() {
      console.log("Service worker has been registered.");
    },
    cached() {
      console.log("Content has been cached for offline use.");
    },
    updatefound(r) {
      console.log("New content is downloading.");
      window._worker = r;
      document.dispatchEvent(
        new CustomEvent("sw-status", { detail: "update-preparing" })
      );
    },
    updated(r) {
      console.log("New content is available.");
      window._worker = r;
      document.dispatchEvent(
        new CustomEvent("sw-status", { detail: "update-ready" })
      );
    },
    offline() {
      console.log(
        "No internet connection found. App is running in offline mode."
      );
      document.dispatchEvent(
        new CustomEvent("sw-status", { detail: "offline" })
      );
    },
    error(error) {
      console.error("Error during service worker registration:", error);
    }
  });
}

window.document.addEventListener("sw-perform-update", () => {
  if (window._worker) {
    window._worker.waiting.postMessage({ type: "SKIP_WAITING" });

    navigator.serviceWorker.addEventListener("controllerchange", () => {
      console.log("rf");
      window.location.reload();
    });
  }
});
