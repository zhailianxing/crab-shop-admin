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
})
