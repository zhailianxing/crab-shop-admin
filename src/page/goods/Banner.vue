<template>

    <el-drawer title="设置商品轮播图" v-model="drawerVisible" dsize="30%" destroy-on-close>
        <el-form :model="form" ref="formRef">
            <el-form-item label="图片">
                <ChooseImage v-model="form.dispalyBanners"></ChooseImage>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">确定设置</el-button>
                <el-button @click="close()">取消</el-button>
            </el-form-item>
        </el-form>

    </el-drawer>



</template>


<script setup>

import { reactive, ref } from 'vue'
import ChooseImage from '~/components/ChooseImage.vue'

const drawerVisible = ref(false)

const form = reactive({
    "dispalyBanners": []
})
const open = (banner) => {
    form.dispalyBanners = banner
    drawerVisible.value = true
}

const close = (banner) => {
    form.dispalyBanners = []
    drawerVisible.value = false
}

defineExpose({
    open,
    close
})

const emit = defineEmits(["confirmSetBannerEmit"])
const onSubmit = () => {
    emit("confirmSetBannerEmit", form.dispalyBanners)
}

</script>