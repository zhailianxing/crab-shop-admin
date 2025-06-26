<template>
    <el-container>
        <el-header><global-header></global-header></el-header>
        <el-container>
            <el-aside><global-aside></global-aside></el-aside>
            <el-main style="padding: 0;">
                <global-tag-list></global-tag-list>
                <!-- header、taglist都是fixed定位，脱离了工作流，主内容会被挡住，所以加了一个div占用位置，让<router-view>正常显示不被覆盖 -->
                <div style="height: 100px;"></div>
                <router-view></router-view>
            </el-main>
        </el-container>
    </el-container>
</template>

<script setup>
// 文件命名方式是PascalCase格式，
// 即：每个单词首字母大写。Vue 内部会把组件名 Fhead（PascalCase）转换为 fhead（小写）和 f-head（kebab-case）两种形式都进行匹配
// 所以在模板中可以使用<global-header>等标签.
import GlobalAside from './components/GlobalAside.vue';
import GlobalHeader from './components/GlobalHeader.vue';
import GlobalTagList from './components/GlobalTagList.vue';

</script>
