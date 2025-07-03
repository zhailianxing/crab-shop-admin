<template>
    <div>
        <div class="statistics">
            <el-row>
                <el-col v-if="panels.length == 0" v-for="i in 4" :key="i" :span="6">
                    <el-skeleton animated>
                        <template #template>
                            <!-- 将真正dom的el-card拷贝到这里，然后用el-skeleton-item代替 实际dom -->
                            <el-card>
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
                    <el-card>
                        <template #header>
                            <div class="card-header">
                                <span style="font-size: 16px;">{{ item.title }}</span>
                                <el-tag :type="item.unitColor">{{ item.unit }}</el-tag>
                            </div>
                        </template>
                        <p style="font-size: 32px; font-weight: bold;">{{ item.value }}</p>
                        <template #footer>
                            <div class="card-footer">
                                <span style="font-size: 16px; color: #8a8a94;">{{ item.subTitle }}</span>
                                <span>{{ item.subValue }}</span>
                            </div>
                        </template>
                    </el-card>
                </el-col>
            </el-row>
        </div>


        <h1>首页</h1>
        {{ $store.state.user }}

    </div>
</template>


<script setup>
import { onBeforeMount, ref } from "vue";
import { getStatistics1 } from "~/api/api.js"

// 统计面板
const panels = ref([])
onBeforeMount(async () => {
    let res = await getStatistics1()
    if (res.msg == "ok") {
        panels.value = res.data.panels
    }
})
</script>


<style lang="scss" scoped>
.el-card {
    max-width: 480px;
    margin: 0 10px;
}

.card-header {
    display: flex;
    justify-content: space-between;
}

.card-footer {
    display: flex;
    justify-content: space-between;
}
</style>