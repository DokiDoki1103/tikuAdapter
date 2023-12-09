const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/adapter-service": {
        target: 'http://adapter.xmig6.cn',
        changeOrigin: true
      }
    }
  }
})
