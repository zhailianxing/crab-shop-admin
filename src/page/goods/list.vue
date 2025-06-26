<template>
    <div>
        商品列表啦
    </div>
    <div>
        商品列表啦
    </div>
    <div>
        商品列表啦
    </div>
    <div v-for="n in 100" :key="n">
        第 {{ n }} 项
    </div>
</template>