<template>

    <el-drawer v-model="showDrawer" :title="title" :close-on-click-modal="false" :size="size"
        :destroy-on-close="destroyOnClose" :v-loading="loading">

        <div class="main">
            <div class="content">
                <slot></slot>
            </div>

            <div class="footer">
                <el-button type="primary" @click="submit">{{ submitText }}</el-button>
                <el-button type="primary" @click="close()">{{ cancelText }}</el-button>
            </div>
        </div>

    </el-drawer>

</template>

<script setup>

import { ref } from 'vue'
const showDrawer = ref(false)

//1. 暴露属性
const prop = defineProps({
    title: String,
    size: {
        type: String,
        default: "45%"
    },
    submitText: {
        type: String,
        default: "提交"
    },
    cancelText: {
        type: String,
        default: "取消"
    },
    // prop不支持 短横线， 即 destroy-on-close不支持
    destroyOnClose: {
        type: Boolean,
        default: true
    }
})


// 2. 暴露方法
const open = () => {
    showDrawer.value = true
}
const close = () => {
    showDrawer.value = false
}

const loading = ref(false)
const showLoading = () => {
    loading.value = true
}
const hideLoading = () => {
    loading.value = false
}

// 向调用者(即父组件)暴露方法
defineExpose({
    open,
    close,
    showLoading,
    hideLoading
})



// 3. 向父组件传递事件。
// 传递的事件名是submitEmit,在父组件中需要用：@submitEmit="myFunc"，来实现监听事件、执行myFunc
const emit = defineEmits(["submitEmit"])
const submit = () => {
    console.log("send emit")
    emit("submitEmit")
}


</script>


<style lang="scss" scoped>
.main {
    height: 100%;
    width: 100%;
    // background-color: red;
    display: flex;
    flex-direction: column;
}

.content {
    // flex: 1,即自动占据父容器剩余的所有空间。 前提是父容器一定要是flex布局
    // 注： 容器剩余的，也就是父容器下的其他所有子元素空间去掉后，这里就是去掉footer占据的空间
    flex: 1;
    // 元素在y方向超过容器长度，则显示滚动条
    overflow-y: auto;
}

.footer {
    height: 50px;
}
</style>
