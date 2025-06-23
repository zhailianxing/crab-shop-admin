<template>
    <el-container>
        <el-header><global-header></global-header></el-header>
        <el-container>
            <!-- 设置宽度为auto，则宽度会根据子内容自适应. el-aside原先默认的宽度就是无效了 -->
            <el-aside width="auto"><global-aside></global-aside></el-aside>
            <el-main>
                <global-tag-list></global-tag-list>
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
