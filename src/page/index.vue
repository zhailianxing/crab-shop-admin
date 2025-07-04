<template>
    <div>
        <div class="statistics">
            <!-- 第一块： 统计信息 -->
            <el-row>
                <el-col v-if="panels.length == 0" v-for="i in 4" :key="i" :span="6">
                    <el-skeleton animated>
                        <template #template>
                            <!-- 将真正dom的el-card拷贝到这里，然后用el-skeleton-item代替 实际dom -->
                            <el-card class="card">
                                <template #header>
                                    <div class="card-header">
                                        <el-skeleton-item variant="text" style="margin-right: 16px" />
                                        <el-skeleton-item variant="text" style="margin-right: 16px" />
                                    </div>
                                </template>
                                <el-skeleton-item variant="p" style="margin-right: 16px" />
                                <template #footer>
                                    <div class="card-footer">
                                        <el-skeleton-item variant="text" style="margin-right: 16px" />
                                        <el-skeleton-item variant="text" style="margin-right: 16px" />
                                    </div>
                                </template>
                            </el-card>
                        </template>
                    </el-skeleton>
                </el-col>

                <el-col v-for="(item, index) in panels" :key="index" :span="6">
                    <el-card class="card">
                        <template #header>
                            <div class="card-header">
                                <span style="font-size: 16px;">{{ item.title }}</span>
                                <el-tag :type="item.unitColor">{{ item.unit }}</el-tag>
                            </div>
                        </template>
                        <p style="font-size: 32px; font-weight: bold;">{{ item.value }}</p>
                        <template #footer>
                            <div class="card-footer">
                                <span style="font-size: 16px; color: #6e6e77;">{{ item.subTitle }}</span>
                                <span>{{ item.subValue }}</span>
                            </div>
                        </template>
                    </el-card>
                </el-col>
            </el-row>


            <!-- 第二块： 功能图标 -->
            <icon-tab></icon-tab>

            <!-- 第三块:  echart图表展示统计信息 -->
            <el-row>
                <el-col :span="12" :offset="0">
                    <el-card class="echart-layout">
                        <template #header>
                            <div class="echart-layout-header">
                                <span>订单统计</span>
                                <div class="flex gap-2">
                                    <el-check-tag :checked="currentTab == item.value" v-for="(item, index) in echarTabs"
                                        @click="handleClickTab(item.value)">
                                        {{ item.label }}
                                    </el-check-tag>
                                </div>
                            </div>
                        </template>
                        <!-- card body -->
                    </el-card>
                </el-col>
                <el-col :span="12" :offset="0"></el-col>
            </el-row>

        </div>



    </div>
</template>


<script setup>
import { onBeforeMount, ref } from "vue";
import { getStatistics1 } from "~/api/api.js"

import IconTab from "~/components/IconTab.vue"

// 统计面板
const panels = ref([])
onBeforeMount(async () => {
    let res = await getStatistics1()
    if (res.msg == "ok") {
        panels.value = res.data.panels
    }
})



// echar 图表

// 选项卡
const currentTab = ref("week")
const echarTabs = [
    {
        label: "近1个月",
        value: "month"
    }, {
        label: "近1周",
        value: "week"
    }, {
        label: "近24小时",
        value: "day"
    },
]
// 选项卡切换 
const handleClickTab = (selectedValue) => {
    currentTab.value = selectedValue
}

</script>


<style lang="scss" scoped>
.card {
    max-width: 480px;
    margin: 0 10px;

    .card-header {
        display: flex;
        justify-content: space-between;
    }

    .card-footer {
        display: flex;
        justify-content: space-between;
    }
}


.echart-layout {
    width: 100%;
    margin-top: 10px;

    .echart-layout-header {
        display: flex;
        justify-content: space-between;
    }
}
</style>