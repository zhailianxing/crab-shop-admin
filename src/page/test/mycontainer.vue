<template>
    <div class="common-layout">
        <el-container class="containder-index">
            <el-header class="header">Header</el-header>
            <el-container>
                <el-aside width="200px" class="aside">Aside</el-aside>
                <el-main class="main">
                    <p v-for="i in 100">test.{{ i }}</p>
                </el-main>
            </el-container>
        </el-container>
    </div>
</template>

<style lang="scss" scoped>
.common-layout {
    background-color: red;
    width: 100%;
    height: 100%;
}

.containder-index {
    background-color: yellow;
    height: 100%;
    width: 100%;
}

.header {
    background-color: black;
    width: 100%;
    height: 60px;
}

.aside {
    background-color: pink;
    height: 100%;
}

/*

解决方案：使用固定高度 + 内部滚动
核心思路： 不让整个页面滚动；,而是让 <el-main> 独立滚动；
<el-header> 保持固定即可。

实现效果：
1.页面头部 <el-header> 固定；
2.内容多的时候，只有 <el-main> 会滚动；
3.滚动时不会遮挡或影响头部

关键实现：
 common-layout、containder-index高度要有继承或者自己设置，不要依赖子元素高度。
 el-main 的高度计算方式为：height: calc(100vh - header高度 - footer高度);
 el-main的内容较多时，实现垂直滚动：overflow-y: auto;
*/
.main {
    overflow-y: auto;
    // 减号，前后要有空格
    height: calc(100vh - 60px);
}
</style>