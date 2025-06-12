// 导入 创建哈希路由模式 和 创建路由 函数
import {createWebHashHistory, createRouter} from 'vue-router'

import myLoginView from '../page/login.vue' // 创建自己的页面
import myIndexView from '../page/index.vue' // 创建自己的页面
import NotFound from '../page/404.vue'

// 第一种方式报错： Cannot read properties of undefined (reading 'state'). 原因： useStore() 是一个组合式 API 钩子函数，只能在setup中使用， 不能在模块顶层使用(因为此时vuex还没有初始化)
// import {useStore} from 'vuex'
// const store = useStore()

// 第二种方式报错：The requested module '/src/store/index.js' does not provide an export named 'store'。 用{store}这种方式是命名导出。而vue项目是用默认导出 (default export) 
// import {store} from '../store'

// 第三种方式不报错
import store from '../store'

import {getToken} from '~/common/cookie.js'


import nprogress from 'nprogress'

import layoutIndex from '../layout/layoutIndex.vue'
// 定义路由表
const routes = [
    {
        name: "login",
        path: "/login", // 因为是哈希模式。所以用host:port/#/login访问
        component: myLoginView,
        meta: {
            title: "登录页"
        }
    } ,
    {
        name: "layoutIndex",
        path: "/", 
        component: layoutIndex,
        children: [{
            name: "index",
            path: "/", // 因为是哈希模式。所以用host:port/#/login访问
            component: myIndexView,
            meta: {
                title: "首页"
            }
        }]
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



router.beforeEach((to, from) => {
    nprogress.start()
    // 保存状态管理的数据一直存在(状态管理数据放在内存中,因此访问速度快,同时状态管理数据也是响应式的)
    let user = (store.state.user)  || JSON.parse(localStorage.getItem('user'))
    if (user && !store.state.user) {
        store.commit('setUserInfo', user)
    }
    
    const token = getToken()
    /* bug:
    // if (!token && from.path !== "/login") {  
    //     return {name: '/login'} 
    // }

    第一次导航：
    用户从"/"想去"/login"
    from = "/"，to = "/login"
    检查条件：!token && from.path !== "/login"
    假设没有token，条件成立
    守卫返回：{name: '/login'}，表示"重新导航到登录页"

    第二次导航（由第一次的重定向触发）：
    这是关键：此时Router认为"第一次导航已取消，开始新导航"
    新导航的目标仍然是"/login"
    但因为第一次导航被取消了，所以当前页面仍然是"/"
    所以在新一轮导航中：from = "/" (当前页面)，to = "/login" (目标页面)
    检查条件：!token && from.path !== "/login"
    条件仍然成立！
    守卫又返回：{name: '/login'}
    如此循环往复，永不结束

    更直观的理解
    想象这个过程：
    用户在首页("/")，点击"登录"按钮，想去登录页("/login")
    导航守卫说："等等，你需要先去登录页"，返回{name: '/login'}
    路由器说："好的，我取消去登录页的导航，改为去登录页"
    用户仍然在首页("/")，路由器尝试导航到登录页("/login")
    导航守卫又说："等等，你需要先去登录页"
    循环继续...

    总结：在导航守卫中，from总是表示"当前页面"，to表示"目标页面"。如果守卫返回一个新位置，会取消当前导航并开始新导航，但此时用户仍在原页面，所以在新导航中from仍然是原始页面。
在您的场景中，判断用户是否需要登录应该检查"他想去哪里(to)"，而不是"他从哪里来(from)"，这样才能避免无限循环。

    */

    // 1.token不存在 并且 目标不是登录页,则跳转登录页.  一定要检查to.path,防止无限重定向
    if (!token && to.path !== "/login") {  
        return {path: '/login'}
    }
    // 2.token存在，用户还想到登录页登录，则取消跳转
    if (token && to.path === "/login") {
        console.log("token存在，不让跳登录页")
        return false
    }
    // 3.返回undefine或者true,表示正常跳转到to上

    //设置 路由页面的 标题
    document.title =  (to.meta.title || "") + "- 兴化红膏大闸蟹"
})

router.afterEach((to, from) => {
    nprogress.done()  // 关不全局进度条
})


// 导出路由对象/实例
export default router

// 错误写法。下面是导出一个router对象的属性
// export default {
//     router // router:router的简写
// }

