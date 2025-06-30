<template>
  <div class="tabList" :style="{ left: $store.state.asideCurWidth }">
    <el-tabs v-model="curTabsValue" type="card" closable @tab-change="handleTabChange">
      <el-tab-pane v-for="item in tabList" :key="item.path" :label="item.title" :name="item.path">
      </el-tab-pane>
    </el-tabs>

    <el-dropdown class="tab-btn">
      <span class="el-dropdown-link">
        <el-icon class="el-icon--right">
          <arrow-down />
        </el-icon>
      </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item>Action 1</el-dropdown-item>
          <el-dropdown-item>Action 2</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
  </div>

</template>

<script lang="ts" setup>
import { ref } from 'vue'
import { setCookie, getCookie } from "~/common/cookie.js"
import { onBeforeRouteUpdate, useRoute, useRouter } from 'vue-router'
const router = useRouter()

const cookieKey = "tabList"
const tabList = ref([
  {
    title: '后台首页',
    path: '/',
  }
])

// 联动1: 刷新时/刚进入页面时， 浏览器的path路径 和 tab选项卡 绑定；  从cookie 中加载保存过的tablist
const curTabsValue = ref()
const init = () => {
  let savedTabListVal = getCookie(cookieKey)
  console.log("savedTabList:", savedTabListVal)
  if (savedTabListVal) {
    tabList.value = JSON.parse(savedTabListVal)
  }
  const route = useRoute()
  curTabsValue.value = route.path
}
init()


// 联动2： 点击 菜单栏时， tab也对应改变
// onBeforeRouteUpdate：路由跳转之前 执行的函数
onBeforeRouteUpdate((to, from) => {
  //1. 不存在则追加
  const notFind = tabList.value.findIndex(obj => obj.path == to.path)
  if (notFind == -1) {
    const newTab = {
      title: to.meta.title,
      path: to.path
    }
    tabList.value.push(newTab)
    setCookie(cookieKey, JSON.stringify(tabList.value))
  }
  // 2.设置tab激活选中的项
  curTabsValue.value = to.path
})




// 点击tab，显示tab 路由path内容
const handleTabChange = (tabPaneName) => {
  // console.log(tabPaneName)
  router.push(tabPaneName)
}


</script>

<style scoped>
.tabList {
  display: flex;
  position: fixed;
  top: 60px;
  /* right:0，表示宽度占满到最右边*/
  right: 0;
  /* 设置z-index和 background-color是为了在滑动主区域内容时， 标签栏不会被遮挡 */
  /* 添加 z-index 确保在其他元素上方 */
  z-index: 100;
  /* 添加背景白色 */
  background-color: #fff;

}


/* 去掉border */
:deep(.el-tabs--card>.el-tabs__header .el-tabs__item) {
  border: 0;
  margin-left: 10px;
}

:deep(.el-tabs--card>.el-tabs__header .el-tabs__nav) {
  border: 0;
}

.tab-btn {
  margin-left: auto;
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
}

/* :deep(.el-tabs) {
  max-width: 900px;  //必须指定max-width，才会显示 滚动箭头
} */
/*替换上面max-width的。 */
/*flex：1，表示宽度占满到最右边, min-width: 0，防止flex子项溢出,在必要时收缩到比其内容更小的宽度*/
:deep(.el-tabs) {
  flex: 1;
  min-width: 0;
  /*这里也可以min-width: 20px等等*/
}

/*滚到最前或者最后时, 鼠标放在滚动箭头时, 鼠标样式为 禁止*/
:deep(.el-tabs__nav-next.is-disabled) {
  cursor: not-allowed;
}

:deep(.el-tabs__nav-prev.is-disabled) {
  cursor: not-allowed;
}
</style>