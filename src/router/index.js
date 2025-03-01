// 导入 创建哈希路由模式 和 创建路由 函数
import {createWebHashHistory, createRouter} from 'vue-router'

import myLoginView from '../page/login.vue' // 创建自己的页面
import NotFound from '../page/404.vue'
// 定义路由表
const routes = [
    {
        path: "/login", // 因为是哈希模式。所以用host:port/#/login访问
        component: myLoginView,
    } ,
    { 
        path: '/:pathMatch(.*)*', 
        name: 'NotFound', 
        component: NotFound 
    },

]

// 创建路由
const router = createRouter({
    history: createWebHashHistory(), // 函数调用
    routes
})

// 导出路由对象/实例
export default router

// 错误写法。下面是导出一个router对象的属性
// export default {
//     router // router:router的简写
// }

