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
                        <el-input v-model="form.pwd" :prefix-icon="Lock"  type="password" show-password />
                    </el-form-item>
                    <el-form-item>
                        <el-button type="primary" round @click="onSubmit"
                            style="margin:0 auto; width: 100px;">登录</el-button>
                    </el-form-item>
                </el-form>
            </div>
        </el-col>
    </el-row>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { Lock, Search } from '@element-plus/icons-vue'
import { LoginApi } from '~/api/api.js'
import { ElNotification } from 'element-plus'
import { useRouter } from 'vue-router'

const router = useRouter()
const formRef = ref(null)

// do not use same name with ref
const form = reactive({
    name: '',  
    pwd: '',
})

const rules = {
    // 要和 const from =reactive(...)定义的一样
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
        LoginApi(form.name, form.pwd)
            .then(res => {
                // 提示成功
                console.log("登录成功")
                console.log(res)
                // 保存用户信息(token、cookie等)

                // 跳转到首页
                router.push('/')
            })
            .catch(err => {
                // 提示错误信息
                console.log(err)
                ElNotification({
                    title: '错误',
                    message: (err.response && err.response.data && err.response.data.message) || '登录失败',
                    type: 'error',
                    duration: 2500
                })
            })
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