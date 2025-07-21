<template>

    <div>
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
                    <el-table-column prop="name" label="角色名称" width="180" />
                    <el-table-column prop="desc" label="角色描述" width="360" />
                    <el-table-column label="状态" width="360">
                        <template #default="scope">
                            <el-switch v-model="scope.row.status" :active-value="1" :inactive-value="0"
                                :loading="scope.row.switchLoading" :disabled="scope.row.super == 1"
                                @change="(val) => handleStatusChange(val, scope.row)" />
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
            <!-- el-tree-v2 比 el-tree性能强点，数据多的话，不会卡顿 -->
            <el-tree-v2 style="max-width: 600px" :data="menus" :props="{ 'label': 'name', 'children': 'child' }"
                show-checkbox :height="700" />

        </FormDrawer>
    </div>


</template>

<script setup>
import FormDrawer from '~/components/FormDrawer.vue'
import { getRoleList, addRole, modifyRole, deleteRole, changeRoleStatus } from '~/api/manager.js'
import { getMenus } from '~/api/api.js'

import { computed, onBeforeMount, reactive, ref } from 'vue'

const loading = ref(false)



// 0. 分页功能
const currentPage = ref(1)
const totalCount = ref(100)
const pageSize = ref(10) // 这里写死了默认每页10个

const handleChangeCurrentChange = (newCurPage) => {
    getData(newCurPage)
}

const menus = ref([])

// 1. 列表功能
const tableData = ref([])
onBeforeMount(() => {
    getData()
    // 获取 菜单列表，为了在新增/修改角色时，挂载不同的菜单
    getMenus().then(res => {
        menus.value = res.data.menus
    })
})

const getData = (page) => {
    if (typeof (page) === "number") {
        currentPage.value = page
    }
    getRoleList(currentPage.value).then((res) => {
        tableData.value = res.data.list
        totalCount.value = res.data.totalCount
    })
}

// 2. 新增或编辑功能
const editId = ref(0)
const form = reactive({
    "title": "",
    "content": ""
})
const title = computed(() => {
    return editId.value > 0 ? "编辑" : "新增"
})

const formRef = ref(null)
const rules = {
    title: [
        { required: true, message: '标题不能为空', trigger: 'blur' }
    ],
    content: [
        { required: true, message: '内容不能为空', trigger: 'blur' }
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
        let fun = editId.value > 0 ? modifyRole(editId.value, form) : addRole(form)
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
    form.title = ""
    form.content = ""
    formDrawerRef.value.open()
}

const handleEdit = (index, row) => {
    editId.value = row.id
    form.title = row.title
    form.content = row.content
    formDrawerRef.value.open()
}

const handleDelete = (index, row) => {
    deleteRole(row.id).then(res => {
        getData()
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