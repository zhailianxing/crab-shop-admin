<template>

    <el-card shadow="always">
        <template #header>
            <div class="card-header">
                <span style="font-size: 16px;">{{ title }}</span>
                <el-tag type="danger">{{ tips }}</el-tag>
            </div>
        </template>
        <el-row :gutter="10">
            <el-col :span="24 / data.length" v-for="(item, index) in data" :key="index">
                <el-card>
                    <div class="card-body">
                        <span class="my-digit">{{ item.value }}</span>
                        <span class="my-label">{{ item.label }}</span>
                    </div>
                </el-card>

            </el-col>
        </el-row>

    </el-card>




</template>


<script setup>

const props = defineProps({
    title: String,
    tips: String,
    data: Array
})


</script>

<style lang="scss" scoped>
.card-header {
    display: flex;
    justify-content: space-between;
}

.card-body {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    .my-digit {
        font-size: 20px;
    }

    .my-label {
        font-size: 16px;
        color: #6e6e77;
    }
}
</style>