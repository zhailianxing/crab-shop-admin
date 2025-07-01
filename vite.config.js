import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import path from 'path' // node中的path库
// https://vite.dev/config/
export default defineConfig({
  resolve: {
    alias: {
      //将xxx/xxx/src设置别名为~
      "~": path.resolve(__dirname, "src")
    }
  },
  plugins: [vue()],
  server: {
    proxy: {
      '/whatever': {
        target: 'http://ceshi13.dishait.cn', // 你后端服务的地址
        changeOrigin: true,              // 是否改变请求源头（伪装成目标地址）
        rewrite: (path) => path.replace(/^\/whatever/, ''), // 重写路径：去掉 /whatever 前缀
        // 添加调试日志
        // configure: (proxy, options) => {
        //   proxy.on('error', (err, req, res) => {
        //     console.error('Proxy error:', err);
        //   });
        //   proxy.on('proxyReq', (proxyReq, req, res) => {
        //     console.log('Sending request:', req.method, req.url);
        //   });
        //   proxy.on('proxyRes', (proxyRes, req, res) => {
        //     console.log('Received response:', proxyRes.statusCode, req.url);
        //   });
        // },
      },
    },
  },
})
