<template>
    <div>
        <h1>首页</h1>
        {{ $store.state.user }}



        <el-button @click="logoutBtn">退出登录</el-button>
    </div>
</template>


<script setup>

import { showModal, showSuccessMessage } from '../common/util';
import { logout } from '../api/api';
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'

const router = useRouter()
const store = useStore()

const logoutBtn = () => {
    showModal("确定要退出登录吗？")
        .then(res => {
            //调用退出登录接口,不管成功失败,都走finally逻辑
            logout().finally(() => {
                //1. 清除cookie、清除用户状态
                // store.dispatch本身是同步的，它会触发store中action，action通常是异步的，但是也可以放同步函数.
                store.dispatch("logout")
                // 3. 跳转到首页
                router.push("/login")
                // 4. 提示退出成功
                showSuccessMessage("退出登录成功")
            })
        })
        .catch(() => {
            // 点击"取消",啥也不干
        })
}



</script>