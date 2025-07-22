<template>

    <div class="header">
        <div>
            <el-button v-if="layoutTag.includes('add')" type="primary" @click="$emit('addEmit')">新增</el-button>
            <el-popconfirm v-if="layoutTag.includes('delete')" title="确认删除选择规格吗?" confirm-button-text="确认"
                cancel-button-text="取消" @confirm="$emit('deleteEmit')">
                <template #reference>
                    <el-button type="danger">批量删除</el-button>
                </template>
            </el-popconfirm>

        </div>
        <el-button v-if="layoutTag.includes('refresh')" type="primary" text @click="$emit('refreshEmit')">
            <el-icon size="24px">
                <Refresh />
            </el-icon>
        </el-button>
    </div>

</template>

<script setup>
import { computed } from 'vue';


const layoutTag = computed(() => {
    // 去除空格并按逗号分割
    return props.layout.replace(/\s+/g, '').split(',');
})

const props = defineProps({
    layout: {
        type: String,
        default: "add,refresh"
    }
})
defineEmits(["addEmit", "refreshEmit"])

</script>
<style lang="scss" scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 20px;
}
</style>