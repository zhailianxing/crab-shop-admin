<template>
    <!-- 规格管理 -->
    <div>
        <el-card shadow="never">
            <Header @addEmit="handleAdd" @refreshEmit="handleRefresh" />

            <div class="body">
                <el-table :data="tableData" stripe style="width: 100%" v-loading="loading">
                    <el-table-column prop="name" label="商品规格" width="180" />
                    <el-table-column prop="default" label="规格值" width="360" />
                    <el-table-column prop="order" label="排序" width="360" />
                    <el-table-column label="状态" width="360">
                        <template #default="scope">
                            <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
                                v-loading="switchLoading" @change="(val) => handleStatusChange(val, scope.row)" />
                        </template>
                    </el-table-column>

                    <el-table-column label="操作">
                        <template #default="scope">
                            <el-button size="small" @click="handleEdit(scope.$index, scope.row)">
                                编辑
                            </el-button>

                            <el-popconfirm title="确认删除吗?" confirm-button-text="确认" cancel-button-text="取消"
                                @confirm="handleDelete(scope.$index, scope.row)">
                                <template #reference>
                                    <el-button size="small" type="danger">
                                        删除
                                    </el-button>
                                </template>
                            </el-popconfirm>


                        </template>
                    </el-table-column>
                </el-table>

                <el-pagination background layout="prev, pager, next" :total="totalCount" v-model:page-size="pageSize"
                    v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />
            </div>
        </el-card>

        <FormDrawer ref="formDrawerRef" :title="title" @submitEmit="handleSubmit()">
            <el-form :model="form" ref="formRef" :rules="rules" label-width="80px" :inline="false">
                <el-form-item label="商品规格">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="排序">
                    <el-input-number v-model="form.order" :min="1" :max="1000" />
                </el-form-item>
                <el-form-item label="规格值">
                    <!-- <el-input v-model="form.default"></el-input> -->
                    <DynamicTag v-model="form.default"></DynamicTag>
                </el-form-item>
            </el-form>

        </FormDrawer>
    </div>


</template>

<script setup>
import Header from '~/components/Header.vue'
import FormDrawer from '~/components/FormDrawer.vue'
import { getSkusList, addSkus, modifySkus, deleteSkus, changeSkusStatus } from '~/api/skus.js'
import { showSuccessMessage } from "~/common/util.js"
import DynamicTag from '~/components/DynamicTag.vue'

import { computed, onBeforeMount, reactive, ref } from 'vue'

const loading = ref(false)



// 0. 分页功能
const currentPage = ref(1)
const totalCount = ref(100)
const pageSize = ref(10) // 这里写死了默认每页10个

const handleChangeCurrentChange = (newCurPage) => {
    getData(newCurPage)
}


// 1. 列表功能
const tableData = ref([])
onBeforeMount(() => {
    getData()
})

const getData = (page) => {
    if (typeof (page) === "number") {
        currentPage.value = page
    }
    getSkusList(currentPage.value).then((res) => {
        tableData.value = res.data.list
        totalCount.value = res.data.totalCount
    })
}

// 2. 新增或编辑功能
const editId = ref(0)
const form = reactive({
    "name": "",
    "order": 50,
    "default": "",
    "status": 1
})
const title = computed(() => {
    return editId.value > 0 ? "编辑" : "新增"
})

const formRef = ref(null)
const rules = {
    name: [
        { required: true, message: '规格名不能为空', trigger: 'blur' }
    ]
}
// 新增或编辑 提交
const handleSubmit = () => {
    formRef.value.validate(valid => {
        if (!valid) {
            console.log("validate fail")
            return
        }
        loading.value = true
        let fun = editId.value > 0 ? modifySkus(editId.value, form) : addSkus(form)
        fun.then(res => {
            formDrawerRef.value.close()
            loading.value = false
            getData()
        }).finally(() => {
            loading.value = false
        })
    })
}

const formDrawerRef = ref(null)
const handleAdd = () => {
    form.name = ""
    form.order = 1
    form.default = ""

    formDrawerRef.value.open()
}

const handleEdit = (index, row) => {
    editId.value = row.id
    form.name = row.name
    form.order = row.order
    form.default = row.default
    formDrawerRef.value.open()
}

const handleDelete = (index, row) => {
    deleteSkus(row.id).then(res => {
        showSuccessMessage("删除成功")
        getData()
    })
}

// 更改规格状态
const switchLoading = ref(false)
const handleStatusChange = (val, row) => {
    row.switchLoading = true
    changeSkusStatus(row.id, val).then(res => {
        showSuccessMessage("修改状态成功")
    }).finally(() => {
        row.switchLoading = false
    })
}



</script>

<style scoped lang="scss">
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}

.body {
    display: flex;
    // flex-direction: column → 主轴是 垂直方向（上下）。此时 align-items 控制的是 交叉轴，也就是 水平方向（左右）
    flex-direction: column;
    align-items: center;

    .el-pagination {
        // margin: 0 auto;
        margin-top: 16px;
    }
}
</style>