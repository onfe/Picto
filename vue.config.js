process.env.VUE_APP_VERSION = require("./package.json").version;

module.exports = {
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
    port: 8090
  },
  css: {
    loaderOptions: {
      sass: {
        // @/ is an alias to src/
        // so this assumes you have a file named `src/variables.scss`
        prependData: `@import "~@/assets/scss/mixins";`
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
    },
    workboxPluginMode: "InjectManifest",
    workboxOptions: {
      swSrc: "./src/sw.js",
      swDest: "service-worker.js"
    }
  }
};
