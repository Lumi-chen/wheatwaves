<!-- filename: CustomLayout.vue -->
<script setup>
import WSide from './WSide.vue'
import WHeader from './WHeader.vue'
</script>

<template>
  <el-container>
    <el-aside width="250px">
      <WSide />
    </el-aside>
    <el-container>
      <el-header>
        <WHeader />
      </el-header>
      <el-main>
        <KeepAlive>
          <RouterView />
        </KeepAlive>
      </el-main>
    </el-container>
  </el-container>
</template>

<style lang="scss" scoped>
.el-container {
  width: 100%;
  height: 100%;

  .el-header {
    border-bottom: 1px solid #e7e7e7;
  }

  .el-aside {
    border-right: 1px solid #e7e7e7;
  }

  .el-main {
    padding: 1rem;
  }
}
</style>