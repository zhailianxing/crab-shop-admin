<template>

    <el-container class="image-layout">
        <el-header class="image-header">
            <el-button type="primary">新增图片分类</el-button>
            <el-button type="success">上传图片</el-button>
        </el-header>
        <el-container class="image-layout-internal">
            <el-aside class="image-aside" v-loading="loading">
                <div class="aside-top">
                    <AsideList :active="activeId == item.id" v-for="(item, index) in imageList" :key="index"
                        @edit="asideItemEdit()" @delete="asideItemDelete()"> {{ item.name }} </AsideList>
                </div>
                <div class="aside-bottom">
                    <el-icon>
                        <ArrowLeft />
                    </el-icon>
                    <el-icon>
                        <ArrowRight />
                    </el-icon>
                </div>
            </el-aside>
            <el-main>Main</el-main>
        </el-container>
    </el-container>
</template>

<script setup>
import AsideList from '~/components/AsideList.vue'

import { getImageCategoryList } from '~/api/imageManger.js'
import { ref, onBeforeMount } from 'vue';

const imageList = ref([])
const activeId = ref(0)
const loading = ref(false)
onBeforeMount(() => {
    loading.value = true
    // 获取第一页的图片分类列表
    getImageCategoryList(1)
        .then(res => {
            imageList.value = res.data.list
            if (imageList.value.length > 0 && imageList.value[0]) {
                activeId.value = imageList.value[0].id
            }
        })
        .finally(() => {
            loading.value = false
        })
})

const asideItemEdit = () => {
    console.log("edit happen")
}

const asideItemDelete = () => {
    console.log("edit delete")
}


</script>


<style lang="scss" scoped>
.image-layout {
    background-color: white;
    width: 100%;
    // 嵌套er-container
    // 继承的是layoutIndex.vue的中el-main的height，而el-main中有一个tagList的height为40px，所以要减掉
    // 目的： 一开始空内容不应该让父容器layoutIndex.vue产生滚动条，而是真正超出内容时才有
    height: calc(100% - 40px);

    .image-header {
        display: flex;
        align-items: center;
    }

    .image-layout-internal {
        width: 100%;
        //  100%是指继承父亲100%的高度。上面还有个class="image-header"的元素，高度为60px，所以要减60px
        height: calc(100% - 60px);

        .image-aside {
            width: 200px;
            display: flex;
            flex-direction: column;
            justify-content: center;

            .aside-top {
                flex: 1;
                overflow-y: auto;

            }

            .aside-bottom {
                display: flex;
                align-items: center;
                justify-content: center;
                height: 30px;
            }
        }
    }
}
</style>