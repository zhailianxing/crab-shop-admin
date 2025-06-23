<template>
    <div class="global-header">
        <div class="global-header-left">
            <el-icon class="g-h-l-icon">
                <House />
            </el-icon>
            <span class="g-h-l-title">兴化大闸蟹</span>
            <el-icon class="g-h-l-icon g-h-l-fold" @click="modifyMenuWidth">
                <Expand v-if="isCollapse" />
                <Fold v-else />
            </el-icon>
            <el-icon class="g-h-l-icon" @click="handleRefresh">
                <Refresh />
            </el-icon>
        </div>
        <div class="global-header-right">
            <el-icon class="g-h-r-icon">
                <VideoCamera />
            </el-icon>
            <el-icon class="g-h-r-icon" @click="handleFullscreen">
                <!-- 目前是非全屏状态，则显示全屏图标，否则显示已在全屏的图标 -->
                <FullScreen v-if="!isFullscreen" />
                <Aim v-else />
            </el-icon>
            <el-dropdown>
                <span class="el-dropdown-link">
                    <!--  开始： 这里的内容自由替换  -->
                    <el-avatar :size="25" :src="avatar" />
                    <!-- store也是响应式变量,如果退出登录时，将store.state.user设置为null，那这里机会报错. 所以用 计算属性 优雅判断下 -->
                    <!-- {{ $store.state.user.avatar }} -->
                    {{ username }}
                    <!--  结束： 这里的内容自由替换  -->
                    <el-icon class="el-icon--right">
                        <arrow-down />
                    </el-icon>
                </span>
                <template #dropdown>
                    <el-dropdown-menu>
                        <el-dropdown-item @click="handleLogout">退出登录</el-dropdown-item>
                        <el-dropdown-item @click="handleModifyPassword">修改密码</el-dropdown-item>
                    </el-dropdown-menu>
                </template>
            </el-dropdown>
        </div>
    </div>

    <!-- 接受到子组件发送的submitEmit事件后，去执行onUpatePwdSubmit函数 -->
    <form-drawer ref="formDrawerRef" title="修改密码呀" size="45%" @submitEmit="onUpatePwdSubmit" submitText="确认修改">
        <el-form :model="form" ref="formRef" :rules="rules" label-width="auto"
            style="max-width: 600px; margin: 1em auto;">
            <el-form-item label="旧密码" prop="oldPwd">
                <el-input v-model="form.oldPwd"></el-input>
            </el-form-item>
            <el-form-item label="新密码" prop="newPwd">
                <el-input v-model="form.newPwd" type="password" />
            </el-form-item>
            <el-form-item label="确认新密码" prop="reNewPwd">
                <el-input v-model="form.reNewPwd" type="password" />
            </el-form-item>
        </el-form>
    </form-drawer>
</template>

<script setup>
//TODO: 优化，导入所有图标，不再一个一个导入
import { House, VideoCamera, FullScreen, Aim, Fold, Refresh } from '@element-plus/icons-vue'

import { showModal, showSuccessMessage } from '~/common/util';
import { logout, modifyPwd } from '~/api/api';
import { useStore } from 'vuex'
const store = useStore()
import { useRouter } from 'vue-router'
const router = useRouter()

import { useFullscreen } from '@vueuse/core'
import { computed } from 'vue';
// isFullscreen判断是否全屏。 toggle是自动切换函数，如果处于非全屏，调用toggle就会进入全屏。如果是全屏状态，调用toggle进入非全屏.
const { isFullscreen, toggle } = useFullscreen()

import { ref, reactive, onMounted, onUnmounted } from 'vue'

import FormDrawer from '~/components/formDrawer.vue';

const avatar = computed(() => {
    if (store.state.user) {
        return store.state.user.avatar
    }
    return ""
})

const username = computed(() => {
    if (store.state.user) {
        return store.state.user.username
    }
    return ""
})

// 刷新
const handleRefresh = () => {
    window.location.reload()
}

// 全屏切换
const handleFullscreen = () => {
    toggle()
}

// 退出登录
const handleLogout = () => {
    showModal("是否确定退出登录")
        .then(res => {
            //调用退出登录接口,不管成功失败,都走finally逻辑
            logout().finally(() => {
                //1. 清除cookie、清除用户状态
                // store.dispatch本身是同步的，它会触发store中action，action通常是异步的，但是也可以放同步函数.
                store.dispatch("logout")
                // 2. 跳转到首页
                router.push("/login")
                // 3. 提示退出成功
                showSuccessMessage("退出登录成功")
            })
        })
        .catch(() => {  // 走到这里的时机：1. 点击了"取消"; 2.then回调内异常
            console.log("取消退出")
        })
}

// 修改密码
const formDrawerRef = ref(null)
const handleModifyPassword = () => {
    formDrawerRef.value.open()
}

// 修改密码 - 确认修改
const formRef = ref(null)

const form = reactive({
    oldPwd: '',
    newPwd: '',
    reNewPwd: '',
})

const rules = {
    // key 要和 const from =reactive(...)定义的属性名一样
    oldPwd: [
        // trigger: 'blur'，失去焦点
        { required: true, message: '旧密码不能为空', trigger: 'blur' },
    ],
    newPwd: [
        { required: true, message: '新密码不能为空', trigger: 'blur' },
    ],
    reNewPwd: [
        { required: true, message: '确认新密码不能为空', trigger: 'blur' },
    ]
}

const onUpatePwdSubmit = () => {
    formRef.value.validate(boolResult => {
        if (!boolResult) {
            console.log("验证失败")
            return
        }
        const data = { oldPwd: form.oldPwd, newPwd: form.newPwd, reNewPwd: form.reNewPwd }
        modifyPwd(data)
            .then(res => {
                // 提示成功
                showSuccessMessage("修改成功")
                // 清空
                store.dispatch("logout")
                // 跳转到首页
                router.push('/login')
            })
            .finally(() => {
                showDrawer.value = false
            })
    })
}

//显示 展开/折叠 图标
const isCollapse = computed(() => {
    return !(store.state.asideCurWidth == "300px")
})
const modifyMenuWidth = () => {
    console.log("click 1")
    store.commit("changeAsideWidth")
}

</script>

<style lang="scss" scoped>
.global-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    background-color: #2a93cf;
    height: 60px;
    position: fixed;
    top: 0;
    left: 0;
    right: 0;

    .global-header-left {
        display: flex;
        align-items: center;
        gap: 10px; //gap: 10px 用于设置 Flex 子项之间的间距，避免手动给每个子项加 margin-right 或类似的样式，更方便、语义更清晰。

        .g-h-l-icon {
            font-size: 20px;
            margin-left: 20px;
            height: 60px;
            width: 60px;
            cursor: pointer; // 鼠标悬停时显示手型

            &:hover {
                // 鼠标悬停时，背景颜色变为深蓝色，会将width和height区域的背景颜色改掉，所以设置了height
                background-color: #3a8fc0;
            }

        }

        .g-h-l-title {
            font-size: 20px;
        }

        .g-h-l-fold {
            font-size: 20px;
            margin-left: 60px;
            ;
        }
    }

    .global-header-right {
        display: flex;
        align-items: center;
        gap: 20px;

        .g-h-r-icon {
            font-size: 20px;
            height: 60px;
            width: 60px;
            cursor: pointer; // 鼠标悬停时显示手型

            &:hover {
                // 鼠标悬停时，背景颜色变为深蓝色，会将width和height区域的背景颜色改掉，所以设置了height
                background-color: #3a8fc0;
            }
        }

        .el-dropdown-link {
            display: flex;
            align-items: center;
            gap: 5px;
        }
    }
}
</style>
