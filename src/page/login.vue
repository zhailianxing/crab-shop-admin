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
                <el-form :model="form" label-width="auto" style="max-width: 600px; margin: 1em auto;">
                    <el-form-item label="用户名">
                        <el-input v-model="form.name" />
                    </el-form-item>
                    <el-form-item label="密码">
                        <el-input v-model="form.pwd" />
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
import { reactive } from 'vue'

// do not use same name with ref
const form = reactive({
    name: '',
    pwd: '',
})

const onSubmit = () => {
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
        // 先设置父容器flex布局为columnp排列 并且justify-content:center(主轴方向居中)，这样内容(即子元素）是垂直居中了
        display: flex;
        flex-direction: column;
        justify-content: center;
        color: white;

        .leftWelcome {
            display: flex;
            justify-content: center;
            align-items: center;
            font-size: 2rem;
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