<template>

    <div>
        <el-card shadow="never">
            <Header layout="add, refresh" @addEmit="handleAdd" />

            <div class="body">
                <el-table :data="tableData" stripe style="width: 100%" v-loading="loading">
                    <el-table-column prop="name" label="优惠劵名字" width="360">
                        <template #default="scope">
                            <div class="couponColumn">
                                <span class="title">{{ scope.row.name }}</span>
                                <span class="time"> {{ scope.row.start_time }} - {{ scope.row.end_time }}</span>
                            </div>
                        </template>
                    </el-table-column>
                    <el-table-column prop="statusText" label="状态" width="180" />
                    <el-table-column label="优惠">
                        <template #default="scope">
                            {{ scope.row.type == 0 ? "满减" : "打折" }} {{ scope.row.type == 0 ? "￥" + scope.row.value :
                                scope.row.value + "折" }}
                        </template>
                    </el-table-column>
                    <el-table-column prop="total" label="发放数量" />
                    <el-table-column prop="used" label="已使用" />
                    <el-table-column label="操作">
                        <template #default="scope">
                            <el-popconfirm title="确认失效吗?" confirm-button-text="确认" cancel-button-text="取消"
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
            <el-form :model="form" ref="formRef" :rules="rules" label-width="100px" :inline="false">
                <el-form-item label="优惠劵名字">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="类型">
                    <el-radio-group v-model="form.type">
                        <el-radio :value="0">满减</el-radio>
                        <el-radio :value="1">折扣</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="面值">
                    <el-input v-model="form.value" style="width: 30%">
                        <template #append>{{ form.type == 0 ? '元' : "折" }}</template>
                    </el-input>
                </el-form-item>

                <el-form-item label="发放数量">
                    <el-input-number v-model="form.total" :min="1" :max="10000"></el-input-number>
                </el-form-item>
                <el-form-item label="最低使用价格">
                    <el-input-number v-model="form.min_price" :min="1" :max="10"></el-input-number>
                </el-form-item>
                <el-form-item label="排序">
                    <el-input-number v-model="form.order" :min="50" :max="1000"></el-input-number>
                </el-form-item>
                <el-form-item label="时间">
                    <!-- v-model绑定的值是一个数组：[start_time, end_time]，时间格式可以是时间戳、Date对象、字符串等 -->
                    <!-- format格式是组件上肉眼看到的格式。  value-format是绑定到变量的值格式，默认是Date对象-->
                    <el-date-picker v-model="form.pickedTimeArr" type="datetimerange" range-separator="至"
                        start-placeholder="开始日期" end-placeholder="结束日期" @change="handleSelectDate"
                        format="YYYY-MM-DD HH:mm:ss" value-format="YYYY-MM-DD HH:mm:ss" />
                </el-form-item>
                <el-form-item label="描述">
                    <el-input v-model="form.desc" type="textarea"></el-input>
                </el-form-item>
            </el-form>
        </FormDrawer>
    </div>


</template>

<script setup>
import FormDrawer from '~/components/FormDrawer.vue'
import Header from '~/components/Header.vue'
import { getCouponList, addCoupon, modifyCoupon, deleteCoupon, changeCouponStatus } from '~/api/coupon.js'
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
    getCouponList(currentPage.value).then((res) => {
        tableData.value = res.data.list.map(obj => {
            // 增加自定义的'状态'字段
            obj.statusText = getStatusText(obj)
            return obj
        })
        totalCount.value = res.data.totalCount
    })
}

function getStatusText(row) {
    let statusText = "领取中"
    let now = (new Date()).getTime()
    let start_time = (new Date(row.start_time)).getTime()
    let end_time = (new Date(row.end_time)).getTime()
    if (now < start_time) {
        statusText = "未开始"
    } else if (end_time < now) {
        statusText = "已结束"
    } else if (row.status == 0) {
        statusText = "已失效"
    }
    return statusText
}
// 2. 新增或编辑功能
const editId = ref(0)
const form = reactive({
    "name": "",
    "type": 1,
    "value": "",
    "min_price": 1,
    "total": 0,
    "used": 0,
    "order": 50,
    "pickedTimeArr": [], // 临时变量，用于存放[start_time, end_time]
    "start_time": "",
    "end_time": "",
})
const title = computed(() => {
    return editId.value > 0 ? "编辑" : "新增"
})

const formRef = ref(null)
const rules = {
    name: [
        { required: true, message: '优惠劵名字不能为空', trigger: 'blur' }
    ],

}
// 新增或编辑 提交
const handleSubmit = () => {
    formRef.value.validate(valid => {
        if (!valid) {
            console.log("validate fail")
            return
        }
        loading.value = true
        // 真是服了. form.start_time要求是毫秒时间戳，通过列表拿到的又是YYYY-MM—DD HH:mm:ss格式
        console.log("form.start_time:", form.start_time)
        form.start_time = (new Date(form.start_time))
        form.end_time = (new Date(form.end_time)).getTime()
        console.log("form.start_time:", form.start_time)

        let fun = editId.value > 0 ? modifyCoupon(editId.value, form) : addCoupon(form)
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
    form.type = 1
    form.value = ""
    form.min_price = 1
    form.total = 0
    form.used = 0
    form.order = 50
    form.pickedTimeArr = []
    formDrawerRef.value.open()
}

const handleEdit = (index, row) => {
    editId.value = row.id
    form.name = row.name
    form.type = row.type
    form.value = row.value
    form.total = row.total
    form.min_price = row.min_price
    form.start_time = row.start_time
    form.end_time = row.end_time
    form.pickedTimeArr = [form.start_time, form.end_time]
    formDrawerRef.value.open()
}

const handleDelete = (index, row) => {
    deleteCoupon(row.id).then(res => {
        getData()
    })
}

// 报错： 当form.start_time改变后，与computed的只读属性冲突。 其他的解决办法在computed写set、get方法
// // 初始化：开始和结束时间
// const pickedTimeArr = computed(() => {
//     if (form.start_time && form.end_time) {
//         return [form.start_time, form.end_time]
//     }
//     return []
// })

// 在el-date-picker上选择时间后，触发的回调
const handleSelectDate = (val) => {
    if (val) {
        form.start_time = val[0]
        form.end_time = val[1]
    }
    // 用户选择了cancel
    return
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
        margin-top: 16px;
    }

    .couponColumn {
        display: flex;
        flex-direction: column;
        background-color: #ffeceb;

        .title {
            margin: 5px;
            font-size: 18px;
            font-weight: 600;
            color: rgb(209, 40, 68);
        }

        .time {
            margin: 5px;
            color: rgb(173, 54, 74);
        }
    }
}
</style>