<template>

    <div class="search">
        <div class="left">
            <el-form-item label="关键词搜索">
                <el-input v-model="searchName" style="width: 240px" placeholder="管理员昵称" />
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
                    <el-table-column label="管理员" width="180">
                        <template #default="scope">
                            <div class="admin">
                                <div class="avatar">
                                    <el-avatar size="small" :src="scope.row.avatar" />
                                </div>
                                <div class="adminInfo">
                                    <span>{{ scope.row.username }}</span>
                                    <span>ID: {{ scope.row.id }}</span>
                                </div>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column label="所属管理员" width="360">
                        <!-- 直接结构出row，省的写scope.row.xxx -->
                        <template #default="{ row }">
                            {{ row.role ? row.role.name : "" }}
                        </template>
                    </el-table-column>
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
                <el-form-item label="用户名">
                    <el-input v-model="form.username"></el-input>
                </el-form-item>
                <el-form-item label="密码">
                    <el-input v-model="form.password" type="password"></el-input>
                </el-form-item>
                <el-form-item label="头像">
                    <!-- 内部自定义实现 v-model  -->
                    <ChooseImage v-model="form.avatar" />
                </el-form-item>
                <el-form-item label="所属角色">
                    <el-select v-model="form.role_id" placeholder="" clearable filterable @change="">
                        <el-option v-for="item in rolesList" :key="item.id" :label="item.name" :value="item.id">
                        </el-option>
                    </el-select>
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
import { getManagerList, changeManagerStatus, delManager, addManager, modifyManager } from '~/api/manager.js'
import { showSuccessMessage } from '~/common/util.js'


import { computed, onBeforeMount, reactive, ref } from 'vue'
import ChooseImage from '~/components/ChooseImage.vue';

const loading = ref(false)

//  搜索功能
const searchName = ref("")
const handleSearch = () => {
    getData()
}
const handleReset = () => {
    searchName.value = ""
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
    let queryParam = { limit: pageSize.value, keyword: searchName.value }
    getManagerList(currentPage.value, queryParam).then((res) => {
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
    "username": "",
    "password": "",
    "role_id": "",
    "status": "",
    "avatar": ""
})
const title = computed(() => {
    return editId.value > 0 ? "编辑" : "新增"
})

const formDrawerRef = ref(null)
const formRef = ref(null)
const rules = {
    username: [
        { required: true, message: '名字不能为空', trigger: 'blur' }
    ],
    password: [
        { required: true, message: '密码不能为空', trigger: 'blur' }
    ],
    role_id: [
        { required: true, message: '角色不能为空', trigger: 'blur' }
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
        let fun = editId.value > 0 ? modifyManager(editId.value, form) : addManager(form)
        fun.then(res => {
            formDrawerRef.value.close()
            getData()
        }).finally(() => {
            formDrawerRef.value.hideLoading()
        })
    })
}

const handleAdd = () => {
    form.username = ""
    form.password = ""
    form.role_id = ""
    form.status = ""
    form.avatar = ""
    formDrawerRef.value.open()
}

const handleEdit = (index, row) => {
    editId.value = row.id
    form.username = row.username
    form.password = row.password
    form.role_id = row.role_id
    form.status = row.status
    form.avatar = row.avatar
    console.log("handlerEdit, row:", row)
    formDrawerRef.value.open()
}

const handleDelete = (index, row) => {
    delManager(row.id).then(res => {
        getData()
    })
}

// 更改 管理员状态
const handleStatusChange = (val, row) => {
    row.switchLoading = true
    changeManagerStatus(row.id, val).then(res => {
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