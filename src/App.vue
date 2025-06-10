<script setup>
// 因为App.vue被main.js引用，App.vue使用helloWorld组件会被加载出来，所以删掉使用hellworld组件。
</script>

<template>
  <!-- 路由视图是用来显示 当前的路由，即route中的path对应的component -->
  <router-view></router-view>
</template>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
