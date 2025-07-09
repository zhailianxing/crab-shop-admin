<template>

    <el-container class="image-layout">
        <!-- 1. 头部 -->
        <el-header class="image-header">
            <el-button type="primary" @click="handleAdd()">新增图片分类</el-button>
            <el-button type="success">上传图片</el-button>
            <!-- 新增图片的侧边drawer -->
            <FormDrawer ref="formDrawerRef" :title="formDrawerTitle" @submitEmit="handleSubmit()">
                <el-form :model="form" ref="formRef" :rules="rules" label-width="auto">
                    <!-- 一旦rules匹配上，label前会有红色星号标志 -->
                    <el-form-item label="图片分类名" prop="name">
                        <el-input v-model="form.name"></el-input>
                    </el-form-item>
                    <el-form-item label="order" prop="order">
                        <el-input-number v-model="form.order" :min="50" :max="1000" />
                    </el-form-item>
                </el-form>
            </FormDrawer>

        </el-header>
        <el-container class="image-layout-internal">
            <!-- 2. aside 侧边栏 -->
            <el-aside class="image-aside" v-loading="loading">
                <div class="aside-top">
                    <ImageAsideList :active="activeId == item.id" v-for="(item, index) in imageList" :key="index"
                        @edit="asideItemEdit(item)" @delete="asideItemDelete(item)" @click="changeActiveId(item.id)">
                        {{ item.name }}
                    </ImageAsideList>
                </div>
                <div class="aside-bottom">
                    <!-- 必须使用v-model双向绑定； ：只是单向绑定，从父组件到子组件传递 -->
                    <el-pagination background layout="prev, next" :total="totalCount" v-model:page-size="pageSize"
                        v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />
                </div>

            </el-aside>
            <!-- 3. Main 部分 -->
            <ImageMain ref="imageMainRef"></ImageMain>
        </el-container>

    </el-container>
</template>

<script setup>
import ImageAsideList from '~/components/ImageAsideList.vue'
import ImageMain from '~/components/ImageMain.vue'

import FormDrawer from '~/components/FormDrawer.vue'
import { getImageCategoryList, addImageCategory, updateImageCategory, deleteImageCategory } from '~/api/imageManger.js'
import { showSuccessMessage } from '~/common/util.js'
import { ref, onBeforeMount, reactive, computed } from 'vue';

// 1.获取图片分类列表
const imageList = ref([])
const activeId = ref(0)
const loading = ref(false)
onBeforeMount(() => {
    // 获取第一页的图片分类列表
    getData(1)
})

// 获取指定页码的数据
const getData = (page) => {
    if (typeof page == 'number') {
        currentPage.value = page
    }
    loading.value = true
    getImageCategoryList(currentPage.value)
        .then(res => {
            imageList.value = res.data.list
            totalCount.value = res.data.totalCount
            console.log("totalCount:", totalCount.value)
            if (imageList.value.length > 0 && imageList.value[0]) {
                activeId.value = imageList.value[0].id
            }
        })
        .finally(() => {
            loading.value = false
        })
}

// 2.分页功能
const pageSize = ref(10)
const totalCount = ref(1000)
const currentPage = ref(1)
// current-page 改变时触发（即页码发生改变）
const handleChangeCurrentChange = (newPage) => {
    // console.log("page:", page)
    getData(newPage)
}


// 3. 新增图片分类
const formDrawerRef = ref(null)
const handleAdd = () => {
    // 清空表单（防止上一次填写表单后，点击取消，表单数据还存在）
    form.name = ""
    form.order = 50

    // 调用子组件formDrawerRef暴露的open方法
    formDrawerRef.value.open()
}
// 表单相关
const form = reactive({
    name: "",
    order: 50
})

// 验证规则
const formRef = ref(null)
const rules = {
    name: [
        { required: true, message: '图片分类名不能为空', trigger: 'blur' },
    ],
}

const handleSubmit = () => {
    // 表单验证
    formRef.value.validate((vaild) => {
        if (!vaild) {
            console.log("form validate fail")
            return
        }
        formDrawerRef.value.showLoading()
        console.log("editId", editId.value)
        // 根据editID决定，使用 新增函数 还是编辑函数，返回的都是一个promise对象
        const fun = editId.value > 0 ? updateImageCategory(editId.value, form) : addImageCategory(form)
        fun.then((res) => {
            let showTitle = editId.value > 0 ? "编辑成功" : "新增成功"
            showSuccessMessage(showTitle)
            // 编辑的话刷新当前页面，新增的话就加载首页
            let newPage = editId.value > 0 ? currentPage.value : 1
            getData(newPage)
            // 关闭
            formDrawerRef.value.close()
        })
            .finally(() => {
                formDrawerRef.value.hideLoading()
            })
    })
}

// 4.编辑功能
const editId = ref(0)
const formDrawerTitle = computed(() => {
    return editId.value > 0 ? "编辑" : "新增"
})

const asideItemEdit = (item) => {
    // 设置 编辑id
    editId.value = item.id
    // 让fromDrawer上的表单显示 待编辑的值 
    form.name = item.name
    form.order = item.order
    // 显示fromDrawer
    formDrawerRef.value.open()
}


// 5.删除功能
const asideItemDelete = (item) => {
    console.log("edit delete")
    deleteImageCategory(item.id).then((res) => {
        showSuccessMessage("删除成功")
        getData(currentPage.value)
    })
}


// 6.切换分类时, 将选中的id传给 <ImageMain>, 从而加载图片列表
const imageMainRef = ref(null)
const changeActiveId = (imageCategoryId) => {
    activeId.value = imageCategoryId
    imageMainRef.value.changeImageCategoryId(imageCategoryId)
}
</script>


<style lang="scss" scoped>
// 隐藏el-aside的滚动条。 不然下面的flex：1会让aside产生滚动条
// 也就是说： el-aside本来就有滚动条机制,没有必要自己写了，将bottom设置absoulte的布局至于aside底部即可，其他元素过高就会出现滚动条，也不会覆盖bottom元素
.el-aside {
    overflow: hidden;
}

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
            height: 100%;
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
                margin-top: 10px;
                height: 30px;
            }
        }
    }
}
</style>