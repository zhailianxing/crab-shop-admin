<template>


    <el-card shadow="never">
        <template #header>
            <div class="header">
                <el-button type="primary"> 新增</el-button>
                <el-button type="primary" text>
                    <el-icon size="24px">
                        <Refresh />
                    </el-icon>
                </el-button>
            </div>
        </template>
        <div class="body">
            <el-table :data="tableData" stripe style="width: 100%">
                <el-table-column prop="date" label="Date" width="180" />
                <el-table-column prop="name" label="Name" width="180" />
                <el-table-column prop="address" label="Address" />

                <el-table-column label="操作">
                    <template #default="scope">
                        <el-button size="small" @click="handleEdit(scope.$index, scope.row)">
                            编辑
                        </el-button>
                        <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">
                            删除
                        </el-button>
                    </template>
                </el-table-column>
            </el-table>

            <el-pagination background layout="prev, pager, next" :total="totalCount" v-model:page-size="pageSize"
                v-model:current-page="currentPage" @current-change="handleChangeCurrentChange" />

        </div>
    </el-card>

</template>

<script setup>

import { ref } from 'vue'
const tableData = [
    {
        date: '2016-05-03',
        name: 'Tom',
        address: 'No. 189, Grove St, Los Angeles',
    },
    {
        date: '2016-05-02',
        name: 'Tom',
        address: 'No. 189, Grove St, Los Angeles',
    },
    {
        date: '2016-05-04',
        name: 'Tom',
        address: 'No. 189, Grove St, Los Angeles',
    },
    {
        date: '2016-05-01',
        name: 'Tom',
        address: 'No. 189, Grove St, Los Angeles',
    },
]


const currentPage = ref(1)
const totalCount = ref(100)
const pageSize = ref(5)

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