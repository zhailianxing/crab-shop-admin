<template>

    <el-container class="image-layout">
        <el-header>Header</el-header>
        <el-container class="image-layout-internal">
            <el-aside width="200px">Aside</el-aside>
            <el-main style="background-color: black;">Main</el-main>
        </el-container>
    </el-container>


</template>



<style lang="scss" scoped>
.image-layout {
    background-color: pink;
    width: 100%;
    // 不知道为什么使用了height:100%就出出现滑动条，所以设置了一个保守的高度
    height: calc(100vh - 60px - 55px - 80px);

    .image-layout-internal {
        width: 100%;
        height: 100%;
    }
}
</style>