module.exports = {
  devServer: {
    proxy: {
      '/api': 'http://localhost:8080/api'
    }
    port: 8090
  }
}
