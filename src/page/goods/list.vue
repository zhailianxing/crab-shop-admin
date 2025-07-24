<template>

    <div class="search">
        <div class="left">
            <el-form-item label="关键词搜索">
                <el-input v-model="searchKey.title" style="width: 240px" placeholder="商品名" />
            </el-form-item>
        </div>
        <div class="right">
            <el-button type="primary" @click="handleSearch()">搜索</el-button>
            <el-button type="primary" @click="handleReset()">重置</el-button>
        </div>

    </div>
    <div class="middle">
        <el-card shadow="never">
            <template #header>
                <div class="header">
                    <el-button type="primary" @click="handleAdd()">新增</el-button>
                    <el-button type="primary" text>
                        <el-icon size="24px">
                            <Refresh />
                        </el-icon>
                    </el-button>
                </div>
            </template>
            <div class="body">
                <el-table :data="tableData" stripe style="width: 100%" v-loading="loading">
                    <el-table-column prop="title" label="商品名" width="180" />
                    <el-table-column label="状态" width="360">
                        <template #default="scope">
                            <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
                                :loading="scope.row.switchLoading" :disabled="scope.row.super == 1"
                                @change="(val) => handleStatusChange(val, scope.row)" />
                        </template>
                    </el-table-column>
                    <el-table-column label="操作">
                        <template #default="scope">
                            <div v-if="scope.row.super == 1">
                                禁止 操作
                            </div>
                            <div v-else>
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
                            </div>

                        </template>
                    </el-table-column>
                </el-table>
            </div>
        </el-card>

        <div class="bottom">
            <el-pagination background layout="prev, pager, next" :total="totalCount" v-model:page-size="pageSize"
                v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />
        </div>

        <FormDrawer ref="formDrawerRef" :title="title" @submitEmit="handleSubmit()">
            <el-form :model="form" ref="formRef" :rules="rules" label-width="80px" :inline="false">
                <el-form-item label="商品名">
                    <el-input v-model="form.title"></el-input>
                </el-form-item>
                <el-form-item label="状态">
                    <el-switch v-model="form.status" :active-value="0" :inactive-value="1" />
                </el-form-item>
            </el-form>
        </FormDrawer>
    </div>


</template>

<script setup>
import FormDrawer from '~/components/FormDrawer.vue'
import { getGoodsList, changeGoodsStatus, delGoods, addGoods, modifyGoods } from '~/api/goods.js'
import { showSuccessMessage } from '~/common/util.js'


import { computed, onBeforeMount, reactive, ref } from 'vue'

const loading = ref(false)

//  搜索功能
const searchKey = reactive({
    tab: "all",
    title: "",
    category_id: null
})
const handleSearch = () => {
    getData()
}
const handleReset = () => {
    searchKey.title = ""
    getData()
}


// 0. 分页功能
const currentPage = ref(1)
const totalCount = ref(100)
const pageSize = ref(10) // 这里写死了默认每页10个

const handleChangeCurrentChange = (newCurPage) => {
    getData(newCurPage)
}


// 1. 列表功能
const rolesList = ref([]) // 角色列表
const tableData = ref([])
onBeforeMount(() => {
    getData()
})

const getData = (page) => {
    if (typeof (page) === "number") {
        currentPage.value = page
    }
    loading.value = true
    getGoodsList(currentPage.value, searchKey).then((res) => {
        rolesList.value = res.data.roles
        tableData.value = res.data.list.map(obj => {
            // 新增一个switchLoading变量，用于控制每个switch的loading显示
            obj.switchLoading = false
            return obj
        })
        totalCount.value = res.data.totalCount
    }).finally(() => {
        loading.value = false
    })
}

// 2. 新增或编辑功能
const editId = ref(0)
const form = reactive({
    "title": "",
})
const title = computed(() => {
    return editId.value > 0 ? "编辑" : "新增"
})

const formDrawerRef = ref(null)
const formRef = ref(null)
const rules = {
    title: [
        { required: true, message: '商品名不能为空', trigger: 'blur' }
    ]
}
// 新增或编辑 提交
const handleSubmit = () => {
    formRef.value.validate(valid => {
        if (!valid) {
            console.log("validate fail")
            return
        }
        formDrawerRef.value.showLoading()
        let fun = editId.value > 0 ? modifyGoods(editId.value, form) : addGoods(form)
        fun.then(res => {
            formDrawerRef.value.close()
            getData()
        }).finally(() => {
            formDrawerRef.value.hideLoading()
        })
    })
}

const handleAdd = () => {
    form.title = ""
    formDrawerRef.value.open()
}

const handleEdit = (index, row) => {
    editId.value = row.id
    form.title = row.title
    formDrawerRef.value.open()
}

const handleDelete = (index, row) => {
    delGoods(row.id).then(res => {
        getData()
    })
}

// 更改 商品状态
const handleStatusChange = (val, row) => {
    row.switchLoading = true
    changeGoodsStatus(row.id, val).then(res => {
        showSuccessMessage("修改状态成功")
    }).finally(() => {
        row.switchLoading = false
    })
}


</script>

<style scoped lang="scss">
.search {
    display: flex;
    justify-content: space-between;
    align-items: center;

    .left {
        margin-left: 20px;
        font-size: 14px;
    }
}

.middle {
    .el-card {
        border: 0;
    }

    .header {
        display: flex;
        justify-content: space-between;
        align-items: center;
    }

    .admin {
        display: flex;
        align-items: center;
        gap: 10px;

        .adminInfo {
            display: flex;
            flex-direction: column;
        }
    }
}


.bottom {
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 20px;
}
</style>