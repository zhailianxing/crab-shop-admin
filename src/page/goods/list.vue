<template>
    <div>
        <!-- 标签页 -->
        <el-tabs v-model="searchForm.tab" @tab-change="handleTabChange">
            <!-- searchForm.tab 是 选中选项卡的 name，默认值是第一个 tab 的 name -->
            <el-tab-pane label="全部" name="all"></el-tab-pane>
            <el-tab-pane label="审核中" name="checking"></el-tab-pane>
            <el-tab-pane label="出售中" name="saling"></el-tab-pane>
            <el-tab-pane label="已下架" name="off"></el-tab-pane>
            <el-tab-pane label="库存预警" name="min_stock"></el-tab-pane>
            <el-tab-pane label="回收站" name="delete"></el-tab-pane>
        </el-tabs>
        <!-- 搜索 -->
        <div>

            <el-row>
                <el-col :span="8" :offset="0">
                    <el-form-item label="关键词搜索">
                        <el-input v-model="searchForm.title" style="width: 240px" placeholder="商品名" />
                    </el-form-item>
                </el-col>
                <el-col :span="8" :offset="0" v-if="isShowCategory">
                    <el-form-item label="商品分类">
                        <el-select v-model="searchForm.category_id" placeholder="请选择" clearable filterable @change="">
                            <el-option v-for="item in categoryList" :key="item.id" :label="item.name" :value="item.id">
                            </el-option>
                        </el-select>
                    </el-form-item>
                </el-col>
                <el-col :span="8" :offset="isShowCategory ? 0 : 8" style="text-align: right;">
                    <el-button type="primary" @click="handleSearch()">搜索</el-button>
                    <el-button type="primary" @click="handleReset()">重置</el-button>
                    <el-button type="primary" @click="isShowCategory = !isShowCategory">
                        {{ isShowCategory ? '收起' : "展开" }}
                        <el-icon>
                            <ArrowDown v-if="isShowCategory" />
                            <ArrowUp v-else />
                        </el-icon>

                    </el-button>
                </el-col>
            </el-row>



        </div>
        <!-- 列表内容 -->
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
                        <el-table-column prop="title" label="商品名" width="360">
                            <template #default="{ row }">
                                <div class="title">
                                    <el-image :src="row.cover" fit="cover"
                                        style="width: 100px; height: 100px; margin-right: 5px;"></el-image>
                                    <div class="title_right">
                                        <span>{{ row.title }}</span>
                                        <div>
                                            <span class="span1">￥{{ row.min_price }}</span>
                                            <span class="span2">|</span>
                                            <span class="span3">￥ {{ row.min_oprice }} </span>
                                        </div>
                                        <span class="span_bottom">分类: {{ row.category ? row.category.name : "未分类"
                                        }}</span>
                                        <span class="span_bottom">创建时间： {{ row.create_time }}</span>
                                    </div>
                                </div>
                            </template>
                        </el-table-column>
                        <el-table-column prop="sale_count" label="实际销量" width="100" />
                        <!-- 回收站tab签下，没有 审核状态 按钮 -->
                        <el-table-column label="商品状态" width="150" align="center" v-if="searchForm.tab != 'delete'">
                            <template #default="scope">
                                <div v-if="scope.row.ischeck == 0" style="display: flex; flex-direction: column;">
                                    <el-button type="primary" size="small" @click="" plain
                                        style="padding: 0;">审核通过</el-button>
                                    <el-button type="danger" size="small" @click="" plain
                                        style="padding: 0; margin: 10px 0;">审核拒绝</el-button>
                                </div>
                                <span v-else>{{ scope.row.ischeck == 1 ? "通过" : "拒绝" }}</span>
                            </template>
                        </el-table-column>
                        <el-table-column label="商品状态" width="100">
                            <template #default="scope">
                                <el-tag v-if="scope.row.status == 1" type="success">已上架</el-tag>
                                <el-tag v-else type="danger">在仓库</el-tag>
                            </template>
                        </el-table-column>
                        <el-table-column prop="stock" label="总库存" width="100" />
                        <el-table-column label="操作" align="center">
                            <template #default="scope">
                                <!-- 回收站tab签下，没有 操作 按钮 -->
                                <div v-if="searchForm.tab != 'delete'">
                                    <el-button type="primary" text size="small" @click="">
                                        修改
                                    </el-button>
                                    <el-button type="primary" text size="small" @click="">
                                        商品规格
                                    </el-button>
                                    <el-button type="primary" text size="small" @click="">
                                        设置轮播图
                                    </el-button>
                                    <el-button type="primary" text size="small" @click="">
                                        商品详情
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
                                <span v-else>暂无操作</span>
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
    </div>

</template>

<script setup>
import FormDrawer from '~/components/FormDrawer.vue'
import { getGoodsList, changeGoodsStatus, delGoods, addGoods, modifyGoods } from '~/api/goods.js'
import { getCategoryList, changeCategoryStatus, addCategory, modifyCategory } from '~/api/category.js'

import { showSuccessMessage } from '~/common/util.js'


import { computed, onBeforeMount, reactive, ref } from 'vue'

const loading = ref(false)

//  搜索功能
const searchForm = reactive({
    tab: "all",
    title: "",
    category_id: null
})
const handleSearch = () => {
    getData()
}
const handleReset = () => {
    searchForm.title = ""
    searchForm.category_id = null
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
    getGoodsList(currentPage.value, searchForm).then((res) => {
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

// 3.更改 商品状态
const handleStatusChange = (val, row) => {
    row.switchLoading = true
    changeGoodsStatus(row.id, val).then(res => {
        showSuccessMessage("修改状态成功")
    }).finally(() => {
        row.switchLoading = false
    })
}

// 4.标签页发生改变
const handleTabChange = (tabName) => {
    // console.log("tabName:", tabName)
    searchForm.tab = tabName
    getData()
}

// 5. 商品分类
const categoryList = ref([])
onBeforeMount(() => {
    getCategoryList().then(res => {
        categoryList.value = res.data
    })
})
const isShowCategory = ref(false)

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

    .body {
        .title {
            display: flex;
            align-items: center;

            .title_right {
                flex: 1;
                display: flex;
                flex-direction: column;

                .span1 {
                    color: red;
                }

                .span2 {
                    margin: 0 5px;
                }

                .span3 {
                    color: gray;
                }

                .span_bottom {
                    font-size: 12px;
                    color: gray;
                }

            }
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