<template>
    <div class="common-layout">
        <el-container class="containder-index">
            <el-header class="header">Header</el-header>
            <el-container>
                <el-aside width="200px" class="aside">Aside</el-aside>
                <el-container>
                    <!-- 这里不一定是el-header,用个div都行 -->
                    <el-header class="tab-header">
                        <span>tab1</span>
                        <span>tab2</span>
                    </el-header>
                    <el-main class="main">
                        <p v-for="i in 100">hahaha.{{ i }}</p>
                    </el-main>
                </el-container>
            </el-container>
        </el-container>
    </div>
</template>


<!-- 
与mycontainer.vue的唯一区别： 在el-main上加了一个tab-header
-->

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
    position: relative;
    background-color: black;
    width: 100%;
    height: 60px;
}

.aside {
    background-color: pink;
    height: 100%;
}


.tab-header {
    height: 50px;
    background-color: orange;
}

.main {
    overflow-y: auto;
    // 减去header1 ，再减去tabHeader
    height: calc(100vh - 60px - 50px);
}
</style>