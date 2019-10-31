module.exports = {
  devServer: {
    proxy: {
      '/ws': {
         target: 'ws://localhost:8080',
         ws: true
      },
      '/api': {
         target: 'http://localhost:8080',
      }
    },
    port: 8090
  }
}
