<template>
    <el-card shadow="never" :body-style="{ border: '0' }">
        <Header @addEmit="handleAdd" @refreshEmit="handleRefresh" />

        <!-- node-key是唯一标识 -->
        <!-- 必须设置prop属性，label是设置显示的名称，children是子tree。 对应的value(即'name'、'children')是dataSource里字段 -->
        <el-tree :data="dataSource" :props="{ label: 'name', children: 'child' }" node-key="id"
            :default-expanded-keys="defaultExpandedKeys">
            <template #default="{ node, data }">
                <div class="custom-tree-node">
                    <!-- data就是dataSource中的每个item对象 -->
                    <div class="left">
                        <el-tag size="small">{{ data.child && data.child.length != 0 ? "菜单" : "权限" }}</el-tag>
                        <el-icon v-if="data.icon" :size="16">
                            <component :is="data.icon"></component>
                        </el-icon>
                        <span>{{ data.name }}</span>
                    </div>
                    <div @click.stop class="right">
                        <el-switch v-model="data.status" :active-value="1" :inactive-value="0"
                            @change="(val) => handleStatusChange(val, data)" />
                        <el-button type="primary" link @click="handleEdit(data)" text>
                            修改
                        </el-button>
                        <el-button style="margin-left: 4px" type="danger" link @click="remove(node, data)">
                            删除
                        </el-button>
                    </div>
                </div>
            </template>
        </el-tree>

        <FormDrawer ref="formDrawerRef" :title="title" @submitEmit="handleSubmit()">
            <el-form :model="form" ref="formRef" :rules="rules" label-width="80px" :inline="false" v-loading="loading">
                <el-form-item label="上级菜单">
                    <el-cascader v-model="form.rule_id" :options="dataSource"
                        :props="{ 'label': 'name', 'children': 'child', 'checkStrictly': true, 'value': 'id' }"
                        clearable @change="handleCascader()" :emitPath="false" />
                </el-form-item>
                <el-form-item label="菜单/规则">
                    <el-radio-group v-model="form.menu">
                        <el-radio :value="1" border>菜单</el-radio>
                        <el-radio :value="0" border>权限</el-radio>
                    </el-radio-group>
                </el-form-item>
                <el-form-item label="名称">
                    <el-input v-model="form.name"></el-input>
                </el-form-item>
                <el-form-item label="后端规则" v-if="form.menu == 0">
                    <!-- 提供的接口别名,比如：getGoodsList (暂时用不到) -->
                    <el-input v-model="form.condition"></el-input>
                </el-form-item>
                <el-form-item label="请求方法" v-if="form.menu == 0">
                    <el-input v-model="form.method"></el-input>
                </el-form-item>
                <el-form-item label="排序">
                    <el-input v-model="form.order"></el-input>
                </el-form-item>
                <el-form-item label="图标" v-if="form.menu == 1">
                    <IconSelect v-model="form.icon"></IconSelect>
                </el-form-item>
                <el-form-item label="前端路由" v-if="form.menu == 1 && form.rule_id > 0">
                    <el-input v-model="form.frontpath"></el-input>
                </el-form-item>
            </el-form>
        </FormDrawer>

    </el-card>

</template>

<script setup>

import { onMounted, ref, reactive, computed } from 'vue';
import Header from '~/components/Header.vue'
import FormDrawer from '~/components/FormDrawer.vue'
import IconSelect from '~/components/IconSelect.vue'

import { changeManagerMenuStatus, addManagerMenu, modifyManagerMenu, deleteManagerMenu, getPermissionList } from '~/api/manager.js'

const handleHeaderRefresh = () => {

}

// dataSource的格式如下。 重点关注id、name、child字段
//  [
//     {
//         "id": 5,
//         "rule_id": 0,
//         "status": 1,
//         "create_time": "2019-08-11 13:36:09",
//         "update_time": "2021-12-21 19:31:11",
//         "name": "后台面板",
//         "desc": "index",
//         "frontpath": null,
//         "condition": null,
//         "menu": 1,
//         "order": 1,
//         "icon": "help",
//         "method": "GET",
//         "child": [
//             {
//                 "id": 10,
//                 "rule_id": 5,
//                 "status": 1,
//                 "create_time": "2019-08-11 13:37:02",
//                 "update_time": "2021-12-21 20:21:23",
//                 "name": "主控台",
//                 "desc": "index",
//                 "frontpath": "\/",
//                 "condition": null,
//                 "menu": 1,
//                 "order": 20,
//                 "icon": "home-filled",
//                 "method": "GET",
//                 "child": []
//             }
//         ]
//     }
// ]
const dataSource = ref([])
const defaultExpandedKeys = ref([])
onMounted(() => {
    getPermissionList().then((res) => {
        dataSource.value = res.data.list
        // node-key设置的是"id"，所以默认展开第一级菜单的话，只获取第一层的id
        defaultExpandedKeys.value = dataSource.value.map(o => o.id)
    })
})


//新增或修改
const loading = ref(false)
const editId = ref(0)
const form = reactive({
    "rule_id": 0,
    "menu": 0,
    "name": "",
    "condition": "",
    "method": "",
    "status": 1, // 1:启用 0:禁用
    "order": 50,
    "icon": "",
    "frontpath": ""
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
        let fun = editId.value > 0 ? modifyManagerMenu(editId.value, form) : addManagerMenu(form)
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
    form.rule_id = 0
    form.menu = 0
    form.name = ""
    form.condition = ""
    form.method = ""
    form.status = 1
    form.order = 50
    form.icon = ""
    form.frontpath = ""
    formDrawerRef.value.open()
}

const handleEdit = (data) => {
    editId.value = data.id
    form.rule_id = data.rule_id
    form.menu = data.menu
    form.name = data.name
    form.condition = data.condition
    form.method = data.method
    form.status = data.status
    form.order = data.order
    form.icon = data.icon
    form.frontpath = data.frontpath // 前端路由 
    formDrawerRef.value.open()
}

const handleStatusChange = (val, data) => {
    changeManagerMenuStatus(data.id, val).then(res => {
        showSuccessMessage("修改状态成功")
    })
}

// 级联选择器，值发生变化
const handleCascader = (value) => {
    console.log("handleCascader value:", value)
    console.log("handleCascader form.rule_id:", form.rule_id)
}





</script>


<style lang="css">
.custom-tree-node {
    flex: 1;
    display: flex;
    align-items: center;
    justify-content: space-between;
    font-size: 16px;
    padding-right: 8px;

    .left {
        display: flex;
        align-items: center;
        gap: 5px;
    }

    .right {
        display: flex;
        gap: 5px
    }
}

.el-tree-node__content {
    padding: 15px;
}
</style>