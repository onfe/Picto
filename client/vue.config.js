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
    themeColor: '#4d4d4d'
  }
};
