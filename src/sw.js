/* eslint-disable no-console */

// Listen for a SW message 'perform-update'
self.addEventListener("message", e => {
  if (!e.data) {
    return;
  }

  if (e.data === "perform-update") {
    console.log("Switching to new content, please stand by...");
    self.skipWaiting().then(() => {
      console.log("Done. Refreshing...");
      setTimeout(() => {
        window.location.reload();
      }, 100);
    });
  }
});

// Claim the page on firstload (default is only claim after refresh).
// workbox.core.clientsClaim(); // disabled for now

// WIP Pre-caching

// self.__precacheManifest = [].concat(self.__precacheManifest || []);
// workbox.precaching.precacheAndRoute(self.__precacheManifest, {});
