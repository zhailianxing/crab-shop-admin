<template>


    <div class="top-list" :class="{ 'active': active }">
        <span class=" title">
            <slot />
        </span>
        <span>
            <!-- <el-button type="primary" :icon="Edit" circle /> -->
            <el-button text size="small" style="margin-right: 1px; padding: 0;" @click="$emit('edit')">
                <el-icon size="16">
                    <Edit />
                </el-icon>
            </el-button>

            <el-button text size="small" style="margin-right: 2px; padding: 0;" @click="$emit('delete')">
                <el-icon size="16">
                    <Close />
                </el-icon>
            </el-button>
        </span>
    </div>

</template>

<script setup>
import {
    Edit,
    Close,
} from '@element-plus/icons-vue'


defineProps({
    active: {
        type: Boolean,
        default: false
    }
})

const emits = defineEmits(["edit", "delete"])
</script>


<style lang="scss" scoped>
.top-list {
    display: flex;
    align-items: center;
    justify-content: space-between;

    .title {
        margin-left: 5px;
        font-size: 0.875rem;
        color: #4B5563;
    }

    // & 不能少
    &:hover {
        background-color: #EFF6FF;
        cursor: pointer;
    }
}

.top-list.active {
    background-color: #EFF6FF;
    cursor: pointer;
}
</style>