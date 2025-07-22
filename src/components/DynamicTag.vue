<template>
    <div class="tagGroup">
        <el-tag v-for="tag in dynamicTags" :key="tag" closable :disable-transitions="false" @close="handleClose(tag)">
            {{ tag }}
        </el-tag>
        <el-input v-if="inputVisible" ref="InputRef" v-model="inputValue" size="small" @keyup.enter="handleInputConfirm"
            @blur="handleInputConfirm" />
        <el-button v-else class="button-new-tag" size="small" @click="showInput">
            + 输入规格
        </el-button>
    </div>
</template>

<script setup>
import { computed, nextTick, ref } from 'vue'


// 在子组件中访问父组件的modelValue，使用props.modelValue, 不需要再加.value
const inputValue = ref('')
const dynamicTags = computed(() => {
    if (props.modelValue == "") {
        return []
    }
    // console.log("dynamicTags init value:", props.modelValue.split(","))
    return props.modelValue.split(",")
})
const inputVisible = ref(false)
const InputRef = ref()

const handleClose = (tag) => {
    dynamicTags.value.splice(dynamicTags.value.indexOf(tag), 1)
}

const showInput = () => {
    inputVisible.value = true
    nextTick(() => {
        InputRef.value.input.focus()
    })
}

const handleInputConfirm = () => {
    if (inputValue.value) {
        dynamicTags.value.push(inputValue.value)
    }
    inputVisible.value = false
    inputValue.value = ''
    emit("update:modelValue", dynamicTags.value.join(","))
}

// 1.定义props
const props = defineProps({
    modelValue: {
        type: String
    }
})

// 2.定义emit
const emit = defineEmits(["update:modelValue"])

</script>

<style lang="scss">
.tagGroup {
    display: flex;
    gap: 2px;
}
</style>