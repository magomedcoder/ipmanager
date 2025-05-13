<script setup lang="ts">
import { RouterView } from 'vue-router'
import ru from 'element-plus/es/locale/lang/ru'

</script>

<template>
  <el-config-provider :locale="ru">
    <router-view />
  </el-config-provider>
</template>

<style scoped></style>
