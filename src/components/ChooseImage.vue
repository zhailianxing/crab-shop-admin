<template>
    <div style="display: flex; align-items: center; flex-wrap: wrap;">
        <!-- 显示选择的照片 -->
        <template v-if="typeof showImages == 'string' && showImages != ''">
            <el-image :src="showImages" fit="fill" style="width: 100px; height: 100px; margin-right: 5px;"></el-image>
        </template>
        <template v-else>
            <div v-for="(url, index) in showImages" :key="index" style="position: relative; display: inline-flex; ">
                <el-image  :src="url" fit="fill" style="width: 100px; height: 100px; margin-right: 5px;">
                </el-image>
                <el-icon style="display: absolute; top: 5px; right: 20px;" @click="handleDelete(url)"><CircleClose /></el-icon>
            </div>
        </template>
        <!-- 将chooseImage按钮直接放在与图片同级的位置, 这样在同一个flex布局中 -->
        <div class="chooseImage">
            <el-icon size="25" @click="handleOpen()">
                <Plus />
            </el-icon>

            <el-dialog v-model="dialogVisible" title="选择图片" width="60vw" top="5vh">
                <!-- 内容从image/list.vue中复制 -->
                <el-container class="image-layout">
                    <el-container class="image-layout-internal">
                        <!-- 2. aside 侧边栏 -->
                        <el-aside class="image-aside" v-loading="loading">
                            <div class="aside-top">
                                <ImageAsideList :active="activeId == item.id" v-for="(item, index) in imageCategoryList"
                                    :key="index" @edit="asideItemEdit(item)" @delete="asideItemDelete(item)"
                                    @click="changeActiveId(item.id)">
                                    {{ item.name }}
                                </ImageAsideList>
                            </div>
                            <div class="aside-bottom">
                                <!-- 必须使用v-model双向绑定； ：只是单向绑定，从父组件到子组件传递 -->
                                <el-pagination background layout="prev, next" :total="totalCount"
                                    v-model:page-size="pageSize" v-model:current-page="currentPage"
                                    @current-change="handleChangeCurrentChange" />
                            </div>

                        </el-aside>
                        <!-- 3. Main 部分 -->
                        <ImageMain ref="imageMainRef" hasCheckBox @chooseImageEvent="HandleChooseImageEvent">
                        </ImageMain>
                    </el-container>
                </el-container>

                <template #footer>
                    <div class="dialog-footer">
                        <el-button @click="dialogVisible = false">取消</el-button>
                        <el-button type="primary" @click="confirmChooseImage">
                            确认
                        </el-button>
                    </div>
                </template>

            </el-dialog>

        </div>

    </div>


</template>

<script setup>

import { ref } from 'vue'
const dialogVisible = ref(false)
//打开 Dialog对话框
const handleOpen = () => {
    dialogVisible.value = true
}

import ImageAsideList from '~/components/ImageAsideList.vue'
import ImageMain from '~/components/ImageMain.vue'
import Upload from '~/components/Upload.vue'

import FormDrawer from '~/components/FormDrawer.vue'
import { getImageCategoryList, addImageCategory, updateImageCategory, deleteImageCategory } from '~/api/imageManger.js'
import { showSuccessMessage } from '~/common/util.js'
import { onBeforeMount, reactive, computed } from 'vue';

// 1.获取图片分类列表
const imageCategoryList = ref([])
const activeId = ref(0)
const loading = ref(false)
const showImages = ref(null)
onBeforeMount(() => {
    showImages.value = props.modelValue
    // 获取第一页的图片分类列表
    getData(1)
})

// 获取指定页码的分类数据数据
const getData = (page) => {
    if (typeof page == 'number') {
        currentPage.value = page
    }
    loading.value = true
    getImageCategoryList(currentPage.value)
        .then(res => {
            imageCategoryList.value = res.data.list
            totalCount.value = res.data.totalCount
            console.log("totalCount:", totalCount.value)
            if (imageCategoryList.value.length > 0 && imageCategoryList.value[0]) {
                // 既更改分类列表的分类id，又去加载该分类下的图片列表
                changeActiveId(imageCategoryList.value[0].id)
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

// 3.1 上传图片
const uploadRef = ref(null)
const handleUpload = () => {
    uploadRef.value.open()
}
// data是上传文件发送的额外数据.
// 文件实际数据被组件自动添加到body里key-value中，key通过组件:name属性指令。 value是上传的文件路径等其他其他数据
const data = computed(() => {
    return {
        image_class_id: activeId.value
    }
})
const uploadSuccess = (response) => {
    console.log("父组件, resp", response)
    showSuccessMessage("文件上传成功")
    uploadRef.value.close()
    getData(currentPage.value)
}


// 3.2 新增图片分类
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
    if (imageMainRef.value) {
        activeId.value = imageCategoryId
        imageMainRef.value.changeImageCategoryId(imageCategoryId)
    }
}

// 7.选择图片
const selectedImage = ref("")
const HandleChooseImageEvent = (urls) => {
    if (urls.length > 0) {
        selectedImage.value = urls[0].url
    }
}
// 7.1确认选择的图片
const confirmChooseImage = () => {
    dialogVisible.value = false
    console.log("confirmChooseImage:", selectedImage.value)
    if (typeof showImages.value == 'string') { // 只支持单选
        emit("update:modelValue", selectedImage.value)
    } else { // 支持多选
        showImages.value.push(selectedImage.value)
        emit("update:modelValue", showImages.value)
    }
}

// 7.2 v-model 实现: 定义prop、定义update:modelValue的emit 并且发送emit事件
const props = defineProps({
    modelValue: {
        type: [String, Array]
    }
})
const emit = defineEmits(["update:modelValue"])

// 7.3 删除图片
const handleDelete = (url) => {
    console.log("handleDelete:", url)
    showImages.value = showImages.value.filter(item => item != url)
    emit("update:modelValue", showImages.value)
}

</script>

<style lang="scss">
.chooseImage {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100px;
    height: 100px;
    cursor: pointer;
    background-color: lightgray;
}


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
            width: 235px;
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