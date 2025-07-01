<template>
    <el-row class="bg">
        <el-col :span="16" class="bgLeft">
            <div class="leftWelcome">
                <span>欢迎光临</span>
            </div>
            <div class="description">
                <span>此站点是兴化红膏大闸蟹的厚点演示地址</span>
            </div>
        </el-col>
        <el-col :span="8" class="bgRight">
            <div class="comeBack">
                欢迎回来
            </div>
            <div class="tips">
                <!-- TODO: 如何实现 ---- ？ -->
                <span>-------</span>
                <span>账号密码登录</span>
                <span>-------</span>
            </div>
            <div>
                <el-form :model="form" ref="formRef" :rules="rules" label-width="auto"
                    style="max-width: 600px; margin: 1em auto;">
                    <!-- prop="name"：这是 Element UI中 el-form-item 组件的一个属性，用途之一是表单验证功能。用于指定该表单项对应的字段（form.name）并与表单验证规则关联。-->
                    <el-form-item label="用户名" prop="name">
                        <el-input v-model="form.name">
                            <!-- 方法一: 使用插槽, #preifx表示放在input前面。 #suffix表示放在后面 -->
                            <template #prefix>
                                <el-icon class="el-input__icon">
                                    <search />
                                </el-icon>
                            </template>
                        </el-input>
                    </el-form-item>
                    <el-form-item label="密码" prop="pwd">
                        <!-- 方法二： 使用 : prefix-icon 定义图标 -->
                        <el-input v-model="form.pwd" :prefix-icon="Lock" type="password" show-password />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" :loading="loading" round @click="onSubmit"
                            style="margin:0 auto; width: 100px;">登录</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </el-col>
    </el-row>
</template>

<script setup>
import { ref, reactive, onMounted, onUnmounted } from 'vue'
import { Lock, Search } from '@element-plus/icons-vue'
import { LoginApi, getUserInfo } from '~/api/api.js'
import { ElNotification } from 'element-plus'
import { useRouter } from 'vue-router'
import { setToken } from '~/common/cookie.js'
import { showSuccessMessage } from '~/common/util.js'
import { useStore } from 'vuex'
const store = useStore()
const router = useRouter()

const formRef = ref(null)

const loading = ref(false)

// do not use same name with ref
const form = reactive({
    name: '',
    pwd: '',
})

function handlerKeyDown(e) {
    console.log("handlerKeyDown , e:", e)
    if (e.key == "Enter") {
        onSubmit()
    }
}

onMounted(() => {
    console.log('组件已经挂载到页面了（即DOM渲染完成了），可以操作DOM或发请求')
    document.addEventListener('keydown', handlerKeyDown)
})

onUnmounted(() => {
    // 页面销毁时，放移除，防止内存泄漏.
    // document表示本vue组件，windows表示整个窗口。只要在登录页面监听用户按回车后，执行登录操作就行
    document.removeEventListener('keydown', handlerKeyDown)
})




const rules = {
    // key 要和 const from =reactive(...)定义的属性名一样
    name: [
        // trigger: 'blur'，失去焦点
        { required: true, message: '用户名不能为空', trigger: 'blur' },
        { min: 3, max: 5, message: '用户名长度3-5位', trigger: 'blur' },
    ],
    pwd: [
        { min: 5, max: 10, message: '密码5到10位', trigger: 'blur' },

    ]
}


const onSubmit = () => {
    formRef.value.validate(boolResult => {
        if (!boolResult) {
            console.log("验证失败")
            return
        }
        // 登录时，显示loading,可以防止多次点击
        loading.value = true
        LoginApi(form.name, form.pwd)
            .then(async res => {
                // 提示成功
                console.log("登录成功")
                // 将token保存到cookie中
                setToken(res.data.token)
                // 提示成功
                showSuccessMessage("登录成功")
                // 获取用户信息
                const d =  await getUserInfo()
                let user = {
                    username: d.data.username,
                    avatar: d.data.avatar,  
                    id: d.data.id,
                }
                // 数据 保存 到 localStorage中
                localStorage.setItem('user', JSON.stringify(user)) //localStorage只能存字符串
                // 跳转到首页
                router.push('/')
            })
            .finally(() => { // 无论成功还是失败，都会执行finally，即最后都关闭loading
                loading.value = false
            })
        // 错误处理 放到axios拦截器里了
        // .catch(err => {})
    })
    console.log('submit!')
}
</script>

<style lang="scss">
.bg {
    min-height: 100vh;
    width: 100vw; // 屏幕占满.使用el-col分栏时，是将el-row的宽度分成24栏

    .bgLeft {
        background-color: rgba(99, 102, 241, 1);
        // justify-content是 主轴方向，align-items是 交叉轴方向. 当flex-direction:column时，主轴方向是垂直方向，交叉轴方向是水平方向
        // 知识点： justify-content:center 是主轴方向居中； align-items:center 是交叉轴方向居中
        // 先设置父容器flex布局为column排列 并且justify-content:center(主轴方向居中)，这样内容(即子元素）是垂直居中了
        display: flex;
        flex-direction: column;
        justify-content: center;
        color: white;

        .leftWelcome {
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 2rem;
            font-weight: bold;
        }

        .description {
            display: flex;
            justify-content: center;
            align-items: center;
        }
    }

    .bgRight {
        background-color: white;
        display: flex;
        flex-direction: column;
        justify-content: center;

        .comeBack {
            font-size: 2em;
            margin: 0.25em auto;
            font-weight: bold; // 标题加粗后，感觉好看些
        }

        .tips {
            color: gray;
            display: flex;
            justify-content: space-around;
            margin: 0.5em auto;

            span {
                padding: 0 0.1em;
            }
        }
    }
}
</style>