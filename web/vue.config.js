const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/adapter-service": {
        target: 'http://192.168.0.108:8060',
        changeOrigin: true
      }
    }
  }
})
