<template>

    <el-main :v-loading="loading">
        <div class="main">
            <div class="top">
                <div v-for="(item, index) in list" :key="index"> {{ item.url }}</div>
            </div>
            <div class="bottom">
                <el-pagination background layout="prev, pager, next" :total="totalCount" v-model:page-size="pageSize"
                    v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />
            </div>
        </div>

    </el-main>


</template>


<script setup>

import { changeGlobalNodesTarget } from 'element-plus/es/utils/index.mjs';
import { ref } from 'vue'
import { getImagesByCategoryId } from '~/api/imageManger.js'

const loading = ref(false)
const list = ref([])
const ImageCategoryId = ref(0)

// 分页参数
const pageSize = ref(10)
const totalCount = ref(1000)
const currentPage = ref(1)

// 获取指定页码的数据
const getData = (page) => {
    if (typeof page == 'number') {
        currentPage.value = page
    }
    loading.value = true
    getImagesByCategoryId(ImageCategoryId.value, currentPage.value)
        .then(res => {
            list.value = res.data.list
            totalCount.value = res.data.totalCount
        })
        .finally(() => {
            loading.value = false
        })
}

// current-page 改变时触发（即页码发生改变）
const handleChangeCurrentChange = (newPage) => {
    getData(newPage)
}

//暴露设置 ImageCategoryId 的方法
const changeImageCategoryId = (id) => {
    ImageCategoryId.value = id
    getData(1)
}
defineExpose({
    changeImageCategoryId
})
</script>



<style lang="scss" scoped>
.main {
    position: relative;
    height: 100%; 
    width: 100%;

    .bottom {
        position: absolute;
        // 首先：将元素的左边缘放在父容器宽度的中点
        left: 50%;
        // 接着：将元素向左移动自身宽度的一半，实现完美居中
        transform: translateX(-50%) ;
        bottom: 0px;
        margin-top: 10px;
        height: 30px;

    }

}
</style>