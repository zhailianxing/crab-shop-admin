import { createStore } from "vuex"

import {removeCookie} from '~/common/cookie'

const store = createStore({
    state:{
        user: null // 有效数据的格式是对象. 初始化为null，表示没有登录，方便判断（否则要判断为一个非空对象）
    },
    // mutations是同步操作,使用者调用store.commit函数触发
    mutations: {
        setUserInfo(state, user) {
            state.user = user
        }

    },
    // actions内是放异步操作，使用者通过store.dispatch函数触发. 但是action可以放同步的操作.
    // 异步操作 指的是需要等待结果返回的非阻塞操作，例如发送 HTTP 请求
    actions: {

        logout({commit}) {
            // 清除cookie
            removeCookie()
            // 清除用户状态
            commit("setUserInfo", {})
        }
    }

})


export default store