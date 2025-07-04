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
            <el-row style="margin-top: 10px;" gutter="20">
                <el-col :span="12" :offset="0">
                    <el-card class="echart-layout">
                        <!-- card header -->
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
                        <div ref="el" id="echartID" style="width: 100%; height:400px;"></div>
                    </el-card>
                </el-col>
                <el-col :span="12" :offset="0">
                    <IndexShopCard title="店铺及商品提示" tips="店铺及商品展示" :data="goods" />
                    <IndexShopCard title="交易提示" tips="需要立即处理的交易订单" :data="order" />
                </el-col>
            </el-row>

        </div>



    </div>
</template>


<script setup>
import { onBeforeMount, onBeforeUnmount, onMounted, ref } from "vue";
import { getStatistics1, getEchartData, getShopsEchartData } from "~/api/api.js"
// 导入所有的echart表。后续如果只用到柱状图，可以按需导入
import * as echarts from 'echarts';
import { useResizeObserver } from "@vueuse/core";


import IconTab from "~/components/IconTab.vue"
import IndexShopCard from "~/components/IndexShopCard.vue"


/*
    第一部分： 统计面板
*/
const panels = ref([])
const goods = ref([])
const order = ref([])
onBeforeMount(async () => {
    let res = await getStatistics1()
    if (res.msg == "ok") {
        panels.value = res.data.panels
    }

    let res2 = await getShopsEchartData()
    if (res2.msg == "ok") {
        goods.value = res2.data.goods
        order.value = res2.data.order
    }


})



/* 
    第三部分： echar 图表
*/
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
        value: "hour"
    },
]
// 选项卡切换 
const handleClickTab = (selectedValue) => {
    currentTab.value = selectedValue
    // 切换选项卡，加载不同数据，显示图表
    getDataAndShowChart(currentTab.value)
}

var myChart;
onMounted(() => {
    myChart = echarts.init(document.getElementById('echartID'));
    getDataAndShowChart(currentTab.value)
})

const getDataAndShowChart = (type) => {
    // 自带的loading加载框
    myChart.showLoading()
    getEchartData(type).then(res => {
        // let xData = [1, 2]
        // let yData = [100, 200]
        if (res.msg != "ok") {
            console.log("getEchartData error")
        }
        let xData = res.data.x
        let yData = res.data.y

        let option = {
            title: {
                // text: 'ECharts 入门示例'
            },
            tooltip: {},
            xAxis: {
                data: xData
            },
            yAxis: {},
            series: [
                {
                    name: '销量',
                    type: 'bar',
                    data: yData
                }
            ]
        }
        myChart.setOption(option);
    }).finally(() => {
        myChart.hideLoading()
    })
}

// 每当绑定的 DOM 元素（即放置 ECharts 图表的 div）尺寸变化时，自动调用 myChart.resize() 让图表自适应宽和高，以便完整展示图表
// useResizeObserverh函数： 监听某个元素(el绑定)尺寸变化（宽高）。
const el = ref(null)
useResizeObserver(el, (entries) => {
    const entry = entries[0]
    const { width, height } = entry.contentRect
    console.log("width: ", width, ", height:", height)
    myChart.resize()
})


// 页面即将销毁时，销毁echart实例。 防止页面卡顿等异常情况
onBeforeUnmount(() => {
    if (myChart) {
        echarts.dispose(myChart)
    }
})

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

    .echart-layout-header {
        display: flex;
        justify-content: space-between;
    }
}
</style>