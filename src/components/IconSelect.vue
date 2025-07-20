<template>
    <!-- 对于实现自己的v-model，这个展示部分可加可不加 -->
    <div style="display: flex; align-items: center; gap: 10px;">
        <el-icon v-if="modelValue">
            <component :is="modelValue" />
        </el-icon>
        <el-select :modelValue="modelValue" placeholder="选择图标" filterable @change="handleChange">
            <el-option v-for="iconName in options" :key="iconName" :label="iconName" :value="iconName">
                <div style="display: flex; align-items: center; justify-content: space-between;">
                    <el-icon>
                        <component :is="iconName" />
                    </el-icon>
                    <span style="float: left">{{ iconName }}</span>
                </div>
            </el-option>
        </el-select>
    </div>
</template>


<script setup>

import * as Icons from '@element-plus/icons-vue'
import { ref, onMounted } from 'vue'
// 获取所有图标的名称
const options = ref([])
onMounted(() => {
    options.value = Object.keys(Icons)
})

defineProps({
    modelValue: {
        type: String,
        default: ''
    }
})

const emit = defineEmits(['update:modelValue'])
// 3. modelValue发生变化时,通知父组件即可
const handleChange = (value) => {
    emit('update:modelValue', value)
}

</script>