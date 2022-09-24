const { defineConfig } = require('@vue/cli-service')

process.env.VUE_APP_VERSION = require("./package.json").version;

module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/ws": {
        target: "ws://localhost:8080",
        ws: true
      },
      "/api": {
        target: "http://localhost:8080"
      }
    },
    port: 8090,

    // webpack uses /ws for it's hot reload socket by default, but that's what we
    // use for our endpoints, so override this.
    client: {
      webSocketURL: "ws://localhost:8090/wp-ws"
    },
    webSocketServer: {
      type: "ws",
      options: {
        path: "/wp-ws"
      }
    }
  },
  css: {
    loaderOptions: {
      sass: {
        additionalData: `@import "~@/assets/scss/mixins";`
      }
    }
  },
  pwa: {
    themeColor: "#4d4d4d",
    icons: [
      {
        src: "img/icons/maskable-1024x1024.png",
        sizes: "1024x1024",
        type: "image/png",
        purpose: "maskable"
      }
    ],
    iconPaths: {
      favicon32: "img/icons/favicon-32x32.png",
      favicon16: "img/icons/favicon-16x16.png",
      maskIcon: "img/icons/favicon.svg"
    }
  }
})