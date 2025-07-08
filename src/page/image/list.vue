<template>

    <el-container class="image-layout">
        <el-header class="image-header">
            <el-button type="primary" @click="handleAdd()">新增图片分类</el-button>
            <el-button type="success">上传图片</el-button>
            <!-- 新增图片的侧边drawer -->
            <FormDrawer ref="formDrawerRef" title="新增图片分类" @submitEmit="handleSubmit()">
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
            <el-aside class="image-aside" v-loading="loading">
                <div class="aside-top">
                    <AsideList :active="activeId == item.id" v-for="(item, index) in imageList" :key="index"
                        @edit="asideItemEdit()" @delete="asideItemDelete()"> {{ item.name }} </AsideList>
                </div>
                <div class="aside-bottom">
                    <!-- 必须使用v-model双向绑定； ：只是单向绑定，从父组件到子组件传递 -->
                    <el-pagination background layout="prev, next" :total="totalCount" v-model:page-size="pageSize"
                        v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />
                </div>
            </el-aside>
            <el-main>Main</el-main>
        </el-container>
    </el-container>
</template>

<script setup>
import AsideList from '~/components/AsideList.vue'
import FormDrawer from '~/components/FormDrawer.vue'
import { getImageCategoryList } from '~/api/imageManger.js'
import { ref, onBeforeMount, reactive } from 'vue';

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

const asideItemEdit = () => {
    console.log("edit happen")
}

const asideItemDelete = () => {
    console.log("edit delete")
}

// 3. 新增图片分类
const formDrawerRef = ref(null)
const handleAdd = () => {
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
        console.log("form validate success")
    })
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