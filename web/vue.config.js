const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/adapter-service": {
        target: 'http://192.168.10.159:8060',
        changeOrigin: true,
      },
      "/sqp/api": {
        target: 'http://adapter.xmig6.cn',
        changeOrigin: true,
      },
    },
  },
})
