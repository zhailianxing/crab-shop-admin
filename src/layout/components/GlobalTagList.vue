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
import { useRoute, useRouter } from 'vue-router'
const router = useRouter()

// 用当前路由的path值 初始化, 选中选项卡
const curTabsValue = ref()
const init = () => {
  const route = useRoute()
  curTabsValue.value = route.path
}
init()


const tabList = ref([
  {
    title: '商品列表',
    path: '/goods/list',  // 传给el-tab-pane的name属性
  },
  {
    title: '分类管理',
    path: '/goods/category',
  },
  {
    title: '分类管理222',
    path: '/goods/category2',
  },
  {
    title: '分类管理333',
    path: '/goods/category3',
  },
  {
    title: '分类管理444',
    path: '/goods/category4',
  },
  {
    title: '分类管理555',
    path: '/goods/category5',
  },
  {
    title: '分类管理666',
    path: '/goods/category6',
  },
  {
    title: '分类管理777',
    path: '/goods/category7',
  },  

  {
    title: '分类管理888',
    path: '/goods/category8',
  },

  {
    title: '分类管理999',
    path: '/goods/category9',
  },  

  {
    title: '分类管理1000',
    path: '/goods/category10',
  },
  {   
    title: '分类管理1111',
    path: '/goods/category11',
  },
  {   
    title: '分类管理1222',
    path: '/goods/category12',
  },
  {   
    title: '分类管理1333',
    path: '/goods/category13',
  }
  
])



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
  min-width: 0;  /*这里也可以min-width: 20px等等*/
}

/*滚到最前或者最后时, 鼠标放在滚动箭头时, 鼠标样式为 禁止*/
:deep(.el-tabs__nav-next.is-disabled) {
  cursor: not-allowed;
}

:deep(.el-tabs__nav-prev.is-disabled) {
  cursor: not-allowed;
}


</style>