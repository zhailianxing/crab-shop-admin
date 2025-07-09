<template>
    <el-drawer v-model="drawer" title="图片上传">
        <el-upload class="upload-demo" drag :action="UPLOAD_URL" :headers="{ token: token }" :data="data"
            :on-success="uploadSuccess" multiple name="img">
            <el-icon class="el-icon--upload"><upload-filled /></el-icon>
            <div class="el-upload__text">
                Drop file here or <em>click to upload</em>
            </div>
            <template #tip>
                <div class="el-upload__tip">
                    jpg/png files with a size less than 500kb
                </div>
            </template>
        </el-upload>
    </el-drawer>
</template>



<script setup>
import { UPLOAD_URL } from '~/api/imageManger.js'
import { ref } from 'vue'
import { getToken } from '~/common/cookie.js'

const token = getToken()
const drawer = ref(false)

const open = () => {
    drawer.value = true
}
const close = () => {
    drawer.value = false
}

defineExpose({
    open,
    close
})

// 父组件传递要上传图片post的数据
const props = defineProps({
    data: {
        type: Object
    }
})

const emit = defineEmits(['uploadSuccess'])
// 文件上传的回调
const uploadSuccess = (response, uploadFile, uploadFiles) => {
    console.log("upload sucess res:", response)
    // 通知父组件
    emit('uploadSuccess', response)
}

</script>