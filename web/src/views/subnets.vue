<script setup lang="ts">
import { ref } from 'vue'
import SubnetCreate from '@/components/subnet/SubnetCreate.vue'
import BaseTable from '@/components/base/BaseTable.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import { useSubnetStore } from '@/stores/subnet'
import { useRouter } from 'vue-router'

const subnetStore = useSubnetStore()
const router = useRouter()

interface IColumn {
  key: string
  dataKey: string
  title: string
  width: number
}

interface IItem {
  id: number
  ip: string
  vlanName: string
  description: string
}

const loading = ref<boolean>(false)
const total = ref<number>(0)
const page = ref<number>(1)
const pageSize = ref<number>(15)
const items = ref<IItem[]>([])
const showCreation = ref<boolean>(false)

const columns = ref<IColumn[]>([
  {
    key: "ip",
    dataKey: "ip",
    title: "IP",
    width: 350
  },
  {
    key: "vlanName",
    dataKey: "vlanName",
    title: "VLAN",
    width: 350
  },
  {
    key: "description",
    dataKey: "description",
    title: "Описание",
    width: 350
  },
])

const load = async (_page: number) => {
  loading.value = true
  subnetStore.getSubnets(page.value, pageSize.value)
    .then((res: { total: number; items: IItem[] }) => {
      total.value = Number(res.total)
      items.value = res.items.map((item) => {
        return {
          ...item,
          id: Number(item.id)
        }
      })
    }).finally(() => {
    loading.value = false
  })
}

const rowEventHandlers = {
  onClick: (e: any) => {
    router.push({
      name: 'IPView',
      params: {
        id: e.rowData.id
      }
    })
  }
}

load(-1)
</script>

<template>
  <default-layout>
    <el-card class="container mx-auto">
      <template #header>
        <div class="flex justify-between">
          <h1 class="sm:text-2xl pt-1 font-extrabold dark:text-white tracking-tight">Подсети</h1>
          <el-button
            type="primary"
            @click="showCreation = true"
            class="ml-auto !flex"
          >
            Добавить
          </el-button>
        </div>
      </template>
      <div style="height: 690px">
        <base-table
          v-loading="loading"
          :columns="columns"
          :items="items"
          @load="load(1)"
          :row-event-handlers="rowEventHandlers"
        />
      </div>
      <div class="p-3 lg:px-6 border-gray-100 dark:border-slate-800 flex">
        <el-pagination
          v-model:current-page="page"
          :disabled="loading"
          background
          layout="total, prev, pager, next"
          :total="total"
          :page-size="pageSize"
          @current-change="load"
          class="ml-auto"
        />
      </div>
    </el-card>
    <subnet-create
      v-if="showCreation"
      v-model="showCreation"
      @on-reset="load"
    />
  </default-layout>
</template>

<style scoped>

</style>
