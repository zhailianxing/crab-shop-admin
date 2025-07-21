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
                            <el-button type="primary" size="small" @click="handleEdit(scope.$index, scope.row)" text>
                                配置权限
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

        <FormDrawer ref="editFormDrawerRef" :title="权限配置" @submitEmit="handleSubmit()">
            <!-- el-tree-v2 比 el-tree性能强点，数据多的话，不会卡顿 -->
            <el-tree-v2 ref="elTreeRef" style="max-width: 600px" :data="allPermissionList"
                :props="{ 'label': 'name', 'children': 'child' }" show-checkbox :height="700"
                :default-expanded-keys="defaultExpandNode" node-key="id" @check="handleChecked">
                <template #default="{ node, data }">

                    <div style="display: flex; align-items: center;">
                        <el-tag :type="data.menu == 1 ? 'primary' : 'info'">{{ data.menu == 1 ? '菜单' : '权限' }}</el-tag>
                        <!-- data.name 或者 node.label是一样的。 data是permissionList每个对象的值，对应到tree上就是node -->
                        <span style="margin-left: 2px; font-size: 16px;">{{ data.name }}</span>
                    </div>

                </template>
            </el-tree-v2>
        </FormDrawer>
    </div>


</template>

<script setup>
import FormDrawer from '~/components/FormDrawer.vue'
import { getPermissionList, getRoleList, addRole, modifyRole, deleteRole, changeRoleStatus, setRolePermission } from '~/api/manager.js'
import { showSuccessMessage } from "~/common/util.js"

import { computed, onBeforeMount, reactive, ref } from 'vue'

const loading = ref(false)



// 0. 分页功能
const currentPage = ref(1)
const totalCount = ref(100)
const pageSize = ref(10) // 这里写死了默认每页10个

const handleChangeCurrentChange = (newCurPage) => {
    getData(newCurPage)
}

// 角色列表返回的数据如下。 rules中的id是这个角色拥有的权限
// "list": [
//     {
//         "id": 643,
//         "status": 1,
//         "create_time": "2025-07-18 17:35:13",
//         "update_time": "2025-07-18 17:35:13",
//         "name": "运营",
//         "desc": "运营角色",
//         "rules": [
//             {
//                 "id": 5,
//                 "pivot": {
//                     "id": 26313,
//                     "role_id": 643,
//                     "rule_id": 5
//                 }
//             },
//             {
//                 "id": 10,
//                 "pivot": {
//                     "id": 26314,
//                     "role_id": 643,
//                     "rule_id": 10
//                 }
//             }
//         ]
//     }]
const allPermissionList = ref([])
const defaultExpandNode = ref([])

// 1. 列表功能
const tableData = ref([])
onBeforeMount(() => {
    getData()
    // 获取 所有权限列表，为了在新增/修改角色时，挂载不同的权限. [这里的权限其实就是每个http接口]
    getPermissionList().then(res => {
        allPermissionList.value = res.data.list
        // 默认展开一级菜单
        defaultExpandNode.value = allPermissionList.value.map(o => o.id)
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

// 2. 编辑功能
const editFormDrawerRef = ref(null)
const currentRoleId = ref(0)
const currentRoleRuleIds = ref([])
const elTreeRef = ref(null)

// 打开编辑框
const checkStrictly = ref(false)
const handleEdit = (index, row) => {
    currentRoleId.value = row.id
    editFormDrawerRef.value.open()
    // 获取当前角色下的所有权限id
    currentRoleRuleIds.value = row.rules.map(o => o.id)
    // 自动勾选上已有的权限
    // formDrawerRef.value.open()执行后，其下的elTree还没有立即被渲染出来,即elTreeRef是null,所以设置了一个timeout延迟执行
    setTimeout(() => {
        elTreeRef.value.setCheckedKeys(currentRoleRuleIds.value)
    }, 150);
}


// 当复选框被点击的时候触发	
/*
data是被点击的节点（选中状态或未选中状态）
info是额外的信息, checkedKeys 是所有子节点id,  halfCheckedKeys 是所有半选父节点的id。 即当某个节点下所有的子节点都被选时，halfCheckedKeys是空； checkedKeys不是所有子节点时, halfCheckedKeys才会有值
*/
const selectedKeys = ref([])
const handleChecked = (data, info) => {
    let { checkedKeys, halfCheckedKeys } = info
    console.log("checkedKeys:", checkedKeys, ", halfCheckedKeys:", halfCheckedKeys)
    // 拼接 父节点和子节点的id
    // selectedKeys.value = [...checkedKeys, ...halfCheckedKeys]
    // 由于check-strictly默认是false，即不断绝父子关系，有父必有子，有子必有父。 所以只要提交checkedKeys即可
    selectedKeys.value = [...checkedKeys]
}

// 编辑下的 提交 功能
const handleSubmit = () => {
    setRolePermission(currentRoleId.value, selectedKeys.value).then((res) => {
        showSuccessMessage("修改成功")
        getData()
        editFormDrawerRef.value.close()
    })
}

// 3. 角色列表中的删除功能
const handleDelete = (index, row) => {
    deleteRole(row.id).then(res => {
        getData()
    })
}

const handleStatusChange = (val, data) => {
    changeRoleStatus(data.id, val).then(res => {
        showSuccessMessage("修改状态成功")
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