<template>
  <div class="tabList" :style="{ left: $store.state.asideCurWidth }">
    <el-tabs v-model="avtiveTabsValue" type="card" @tab-change="handleTabChange" @tab-remove="removeTab">
      <!--可以每个tab选项卡分别设置是否可以关闭。 后台首页不可关闭，其他都可以  -->
      <el-tab-pane v-for="item in tabList" :closable="item.path != '/'" :key="item.path" :label="item.title"
        :name="item.path">
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
          <el-dropdown-item @click="closeOtherTab">关闭其它</el-dropdown-item>
          <el-dropdown-item @click="closeAll">关闭所有</el-dropdown-item>
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

const COOKIE_KEY = "tabList"
const tabList = ref([
  {
    title: '后台首页',
    path: '/',
  }
])

// 联动1: 刷新时/刚进入页面时， 浏览器的path路径 和 tab选项卡 绑定；  从cookie 中加载保存过的tablist
const avtiveTabsValue = ref()
const init = () => {
  let savedTabListVal = getCookie(COOKIE_KEY)
  console.log("savedTabList:", savedTabListVal)
  if (savedTabListVal) {
    tabList.value = JSON.parse(savedTabListVal)
  }
  const route = useRoute()
  avtiveTabsValue.value = route.path
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
    setCookie(COOKIE_KEY, JSON.stringify(tabList.value))
  }
  // 2.设置tab激活选中的项
  avtiveTabsValue.value = to.path
})


// 点击tab，显示tab 路由path内容
const handleTabChange = (tabPaneName) => {
  console.log(tabPaneName)
  router.push(tabPaneName)
}

// 关闭tab选项卡
// 功能： 关闭的是最后一个，选中上一个；否则，选中下一个
const removeTab = function (tabPaneName) {
  // 0. 后台首页 不可关闭
  if (tabPaneName == "/") {
    return
  }
  // 1. 更新 新的激活的选项卡
  let index = tabList.value.findIndex(obj => obj.path == tabPaneName)
  let nextActiveTab = ""
  if (index == tabList.value.length - 1) {
    nextActiveTab = tabList.value[tabList.value.length - 2].path
  } else {
    nextActiveTab = tabList.value[tabList.value.length - 1].path
  }
  avtiveTabsValue.value = nextActiveTab

  // 2. 删除关闭的选项卡
  tabList.value = tabList.value.filter(obj => obj.path != tabPaneName)
  // 3. 更新cookie
  setCookie(COOKIE_KEY, JSON.stringify(tabList.value))
  // 4. 显示新的激活的选项卡
  router.push(avtiveTabsValue.value)
}

// 关闭其他选项卡
const closeOtherTab = () => {
  tabList.value = tabList.value.filter(obj => obj.path == "/" || obj.path == avtiveTabsValue.value)
  setCookie(COOKIE_KEY, JSON.stringify(tabList.value))
}

// 关闭所有选项卡
const closeAll = () => {
  // 默认只有一个 后台首页 不能关闭
  tabList.value = [{
    title: "后台首页",
    path: "/"
  }]
  setCookie(COOKIE_KEY, JSON.stringify(tabList.value))
  router.push(tabList.value[0].path)
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