<script setup lang="ts">
import { ref } from 'vue'
import IpCreate from '@/components/ip/IPCreate.vue'
import IpCard from '@/components/ip/IPCard.vue'
import BaseTable from '@/components/base/BaseTable.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import { useIpStore } from '@/stores/ip'

const ipStore = useIpStore()

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
  customerName: string
  description: string
}

const loading = ref<boolean>(false)
const total = ref<number>(0)
const page = ref<number>(1)
const pageSize = ref<number>(15)
const items = ref<IItem[]>([])
const showCreation = ref<boolean>(false)
const showCard = ref<boolean>(false)
const id = ref<number>()

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
    title: "Vlan",
    width: 350
  },
  {
    key: "customerName",
    dataKey: "customerName",
    title: "Клиент",
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
  ipStore.getIps(page.value, pageSize.value)
    .then(async (res: { total: number; items: IItem[] }) => {
      total.value = Number(res.total)
      items.value = res.items
    }).finally(() => {
    loading.value = false
  })
}

const rowEventHandlers = {
  onClick: (e: any) => {
    id.value = e.rowData.id
    showCard.value = true
  }
}

load(-1)
</script>

<template>
  <DefaultLayout>
    <div class="container mx-auto">
      <el-card>
        <template #header>
          <div class="flex justify-between">
            <h1 class="sm:text-2xl pt-1 font-extrabold dark:text-white tracking-tight">IP</h1>
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
    </div>
    <ip-create
      v-if="showCreation"
      v-model="showCreation"
      @on-reset="load"
    />
    <ip-card
      v-if="showCard"
      v-model="showCard"
      :id="id"
      @on-reset="load"
    />
  </DefaultLayout>
</template>

<style scoped>

</style>
