<template>
  <div class="w-header center">
    <el-breadcrumb separator="/">
      <el-breadcrumb-item :to="{ path: '/' }">homepage</el-breadcrumb-item>
      <el-breadcrumb-item>promotion list</el-breadcrumb-item>
    </el-breadcrumb>
    <div class="center">
      <el-switch v-model="mode" style="--el-switch-on-color: #313852" active-value="dark" inactive-value="day"
        @change="changeModeType">
        <template #active-action>
          <Moon />
        </template>
        <template #inactive-action>
          <Sunny color="black" />
        </template>
      </el-switch>
      <el-avatar :size="32" src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        style="margin: 0 0.5rem" />
      <span>NickName</span>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const mode = ref('day')
function changeModeType(value) {
  document.getElementById("app").className = value
}
</script>


<style lang="scss" scoped>
.w-header {
  height: 100%;
  justify-content: space-between;
}
</style>