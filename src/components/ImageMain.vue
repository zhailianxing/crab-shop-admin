<template>

    <el-main :v-loading="loading">
        <div class="main">
            <div class="top">
                <el-row :gutter="20">
                    <el-col :span="6" v-for="(item, index) in list" :key="index">
                        <el-card shadow="hover">
                            <el-image :src="item.url" fit="fill" :lazy="true" :preview-src-list="[item.url]"></el-image>
                            <div class="name">{{ item.name }}</div>
                            <div class="action">
                                <el-checkbox v-if="hasCheckBox" v-model="item.checked"
                                    @change="(val) => handleCheck(val, item)" />
                                <el-button type="primary" size="small" @click="rename(item)" text>重命名</el-button>
                                <el-popconfirm title="确认删除图片吗？" confirm-button-text="确认" cancel-button-text="取消"
                                    @confirm="handleDeleteImage(item)">
                                    <template #reference>
                                        <el-button type="primary" size="small" text>删除</el-button>
                                    </template>
                                </el-popconfirm>
                            </div>
                        </el-card>
                    </el-col>
                </el-row>
            </div>
            <div class="bottom">
                <el-pagination background layout="prev, pager, next" :total="totalCount" v-model:page-size="pageSize"
                    v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />
            </div>
        </div>

    </el-main>


</template>


<script setup>

import { showModalInput, showSuccessMessage } from '~/common/util.js'
import { ref } from 'vue'
import { getImagesByCategoryId, modifyName, deleteImage } from '~/api/imageManger.js'
import { showErrorMessage } from '../common/util';

const loading = ref(false)
const list = ref([])
const imageCategoryId = ref(0)

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
    getImagesByCategoryId(imageCategoryId.value, currentPage.value)
        .then(res => {
            list.value = res.data.list.map(o => {
                o.checked = false
                return o
            })
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

//暴露设置 imageCategoryId 的方法
const changeImageCategoryId = (id) => {
    imageCategoryId.value = id
    getData(1)
}
defineExpose({
    changeImageCategoryId
})


// 重命名功能
const rename = (item) => {
    showModalInput()
        .then(({ value }) => { // value不能换，这里是解构，对象中key和value字段，都叫value
            modifyName(item.id, value)
                .then((res) => {
                    showSuccessMessage("重命名成功")
                    getData(currentPage.value)
                })
        })
        .catch(() => {
            ElMessage({
                type: 'info',
                message: 'Input canceled',
            })
        })
}

// 删除功能
const handleDeleteImage = (item) => {
    let imageId = item.id
    deleteImage([imageId])
        .then(res => {
            showSuccessMessage("删除成功")
            getData(currentPage.value)
        })
        .finally(() => { })
}

// 选择框功能(给manger/list使用)
defineProps({
    hasCheckBox: {
        type: Boolean,
        default: false
    }
})

let urls = []
const emit = defineEmits(["chooseImageEvent"])
const handleCheck = (val, imageItem) => {
    console.log("checkVal", val, ", obj:", imageItem.url)
    if (urls.length >= 1 && val) {
        urls.push(imageItem)
        showErrorMessage("最多只能选一张图片")
        return
    }
    // 取消选中
    if (!val) {
        urls = urls.filter(o => o.id != imageItem.id)
    } else {
        urls.push(imageItem)
    }
    emit("chooseImageEvent", urls)
}



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
        transform: translateX(-50%);
        bottom: 0px;
        margin-top: 10px;
        height: 30px;

    }
}

.el-card {
    position: relative;

    .el-image {
        width: 100%;
        height: 160px;
        /* 固定图片高度 */
        display: flex;
        justify-content: center;
        align-items: center;
        padding: 0;
    }

    .name {
        // 使用绝对定位，是相对上一个 relative定位父元素
        position: absolute;
        left: 0;
        right: 0;
        bottom: 25px;
        //等价于: background-color: #1f2937;  opacity: 0.5;
        background-color: rgba(31, 41, 55, 0.5);
    }

    .action {
        display: flex;
        justify-content: center;
        align-items: center;
    }
}


// 在浏览器查看时，发现有个padding
:deep(.el-card__body) {
    padding: 0;
}
</style>