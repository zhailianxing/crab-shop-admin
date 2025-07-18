<template>

    <div class="header">
        <el-button type="primary" @click="$emit('addEmit')">新增</el-button>
        <el-button type="primary" text @click="$emit('refreshEmit')">
            <el-icon size="24px">
                <Refresh />
            </el-icon>
        </el-button>
    </div>

</template>

<script setup>

defineEmits(["addEmit", "refreshEmit"])

</script>
<style lang="scss" scoped>
.header {
    display: flex;
    justify-content: space-between;
    align-items: center;
}
</style>