<script lang="ts" setup>
import 'element-plus/es/components/empty/style/css'
defineProps(['loading', 'columns', 'items', 'rowEventHandlers'])
const emit = defineEmits(['load'])
</script>

<template>
  <el-auto-resizer>
    <template #default="{ height, width }">
      <el-table-v2
        v-loading="loading"
        :columns="columns"
        :data="items"
        :width="width"
        :height="height"
        fixed
        :scrollbar-always-on="true"
        row-class="cursor-pointer"
        :row-event-handlers="rowEventHandlers"
      >
        <template #empty>
          <div class="el-empty">
            <div class="el-empty__description">
              <p>Нет данных</p>
            </div>
            <div class="el-empty__bottom">
              <el-button
                type="primary"
                @click="emit('load')"
              >
                Обновить
              </el-button>
            </div>
          </div>
        </template>
      </el-table-v2>
    </template>
  </el-auto-resizer>
</template>

<style scoped></style>
