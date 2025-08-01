<template>
    <FormDrawer ref="formDrawerRef" title="设置商品详情" size="50%">
        <el-form :model="form" ref="formRef">
            <el-form-item label="">
                <Editor></Editor>
            </el-form-item>
        </el-form>
    </FormDrawer>

</template>


<script setup>

import FormDrawer from '~/components/FormDrawer.vue'
import Editor from '~/components/Editor.vue'
import { reactive, ref } from 'vue'

const formDrawerRef = ref(null)
const form = reactive({
})
const open = () => {
    formDrawerRef.value.open()
}

const close = (banner) => {
    formDrawerRef.value.close()

}

defineExpose({
    open,
    close
})

const emit = defineEmits(["confirmEmit"])
const onSubmit = () => {
}

</script>