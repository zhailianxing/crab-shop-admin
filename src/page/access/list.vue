<template>
    <el-card shadow="never" :body-style="{ border: '0' }">
        <Header @addEmit="handleHeaderAdd" @refreshEmit="handleHeaderRefresh" />
    </el-card>

</template>

<script setup>

import Header from '~/components/Header.vue'

const handleHeaderAdd = () => {

}

const handleHeaderRefresh = () => {

}


</script>