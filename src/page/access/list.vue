<template>
    <el-card shadow="never" :body-style="{ border: '0' }">
        <Header @addEmit="handleHeaderAdd" @refreshEmit="handleHeaderRefresh" />

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
                        <el-switch v-model="data.status" :active-value="1" :inactive-value="0" />
                        <el-button type="primary" link @click="append(data)" text>
                            修改
                        </el-button>
                        <el-button type="primary" link @click="append(data)" text>
                            增加
                        </el-button>
                        <el-button style="margin-left: 4px" type="danger" link @click="remove(node, data)">
                            删除
                        </el-button>
                    </div>
                </div>
            </template>
        </el-tree>

    </el-card>

</template>

<script setup>

import { onMounted, ref } from 'vue';
import Header from '~/components/Header.vue'
import { getMenus } from '~/api/api.js'

const handleHeaderAdd = () => {

}

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
    getMenus().then((res) => {
        console.log("menu:", res.data.menus)
        dataSource.value = res.data.menus
        // node-key设置的是"id"，所以默认展开第一级菜单的话，只获取第一层的id
        defaultExpandedKeys.value = dataSource.value.map(o => o.id)
    })
})

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