<template>

    <el-row>
        <!-- el-row被分成24列等份， 每个el-col占据 24/length  -->
        <el-col :span="24 / iconList.length" v-for="(item, index) in iconList" :offset="0"
            @click="$router.push(item.path)">
            <div class="item">
                <el-icon :size="20" :color="item.color">
                    <component :is="item.icon"></component>
                </el-icon>
                <span style="margin-top:8px;">{{ item.title }}</span>
            </div>

        </el-col>
    </el-row>




</template>


<script setup>

import { ref } from 'vue'

const iconList = ref([])

iconList.value = [
    {
        icon: "user",
        color: "deepskyblue",
        title: "用户",
        path: "/user/list"
    },
    {
        icon: "shopping-cart",
        color: "violet",
        title: "商品",
        path: "/goods/list"
    },
    {
        icon: "tickets",
        color: "fuchsia",
        title: "订单",
        path: "/order/list"
    },
    {
        icon: "chat-dot-square",
        color: "teal",
        title: "评价",
        path: "/comment/list"
    },
    {
        icon: "picture",
        color: "hotpink",
        title: "图库",
        path: "/image/list"
    },
    {
        icon: "bell",
        color: "mediumseagreen",
        title: "公告",
        path: "/notice/list"
    },
    {
        icon: "set-up",
        color: "#8a8a94",
        title: "配置",
        path: "/setting/base"
    },
    {
        icon: "files",
        color: "goldenrod",
        title: "优惠券",
        path: "/coupon/list"
    }
]


</script>



<style lang="scss">
.item {
    display: flex;
    flex-direction: column;
    margin-top: 10px;
    margin-left: 15px;
    align-items: center;
    justify-content: center;
    // 设置宽高更好看
    width: 80px;
    height: 100px;
    // gpt:圆角+阴影效果
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);

    cursor: pointer;
}
</style>