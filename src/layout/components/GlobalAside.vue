<template>
    <!-- router: 启用该模式会在激活导航时以 index 作为 path 进行路由跳转 使用 default-active 来设置加载时的激活项。 -->
    <el-menu default-active="1" class="my-menu" @open="handleOpen" @close="handleClose" router>
        <!-- 加了一层tempalte作为for循环 -->
        <template v-for="(item, index) in menuData">
            <el-sub-menu v-if="item.child && item.child.length > 0" :index="'' + item.id">
                <template #title>
                    <el-icon>
                        <component :is="item.icon"></component>
                    </el-icon>
                    <span>{{ item.name }}</span>
                </template>
                <el-menu-item v-for="(secondItem, index) in item.child" :index="secondItem.path">
                    <template #title>
                        <el-icon>
                            <component :is="secondItem.icon"></component>
                        </el-icon>
                        <span>{{ secondItem.name }}</span>
                    </template>
                </el-menu-item>
            </el-sub-menu>
            <el-sub-menu v-else :index="item.path">
                <template #title>
                    <el-icon>
                        <component :is="item.icon"></component>
                    </el-icon>
                    <span>{{ item.name }}</span>
                </template>
            </el-sub-menu>
        </template>
    </el-menu>
</template>


<script setup>


const menuData = [
    {
        name: "仪表盘",
        path: "/",
        icon: "HomeFilled",
        id: 1
    },
    {
        name: "商品管理",
        path: "",
        icon: "HomeFilled",
        id: 2,
        child: [
            {
                name: "商品列表",
                path: "/goods/list",
                icon: "Goods",
                id: 3
            }
        ]
    }
]

</script>

<style lang="scss" scoped>
.my-menu {
    width: 300px;
    border: 0;
}
</style>