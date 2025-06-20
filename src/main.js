import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import ElementPlus  from  'element-plus'
import 'element-plus/dist/index.css'
import router from './router/index.js'
import store from './store/index.js'

import * as ElementPlusIconsVue from '@element-plus/icons-vue'


import 'nprogress/nprogress.css' // 导入nprogress的样式; 显示颜色、高度

const app = createApp(App)

//导入所有图标
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
  app.component(key, component)
}

app.use(router)
app.use(ElementPlus)
app.use(store)
app.mount('#app')
