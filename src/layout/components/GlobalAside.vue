<template>
    <div :style="appendStyle">
        <!-- router: 启用该模式会在激活导航时以 index 作为 path 进行路由跳转 使用 default-active 来设置加载时的激活项。 -->
        <el-menu default-active="1"  @open="handleOpen" @close="handleClose" router :collapse="collapse"  :collapse-transition="false" >
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
    </div>
</template>


<script setup>

import { computed, ref } from 'vue'

import { useStore } from 'vuex';
const store = useStore()

const collapse = computed(() => {
    console.log("store.state.asideCurWidth :" , store.state.asideCurWidth )
    return !(store.state.asideCurWidth == "300px")
})

const appendStyle = computed(() => {
    return {
        width: store.state.asideCurWidth,
        transition: 'all 0.2s',
        // 隐藏水平滚动条
        // 在动态计算css时，overflow-x则需要用驼峰命名法
        overflowX: 'hidden'
    }
})

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
.el-menu {
    border: 0;
}
</style>