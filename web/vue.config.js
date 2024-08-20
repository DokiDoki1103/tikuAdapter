const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true,
  devServer: {
    proxy: {
      "/adapter-service": {
        target: 'http://127.0.0.1:8060',
        changeOrigin: true,
      },
      "/sqp/api": {
        target: 'http:/127.0.0.1:8060',
        changeOrigin: true,
      },
    },
  },
})
