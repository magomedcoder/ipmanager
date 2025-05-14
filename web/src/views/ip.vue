<script setup lang="ts">
import { reactive, ref, computed } from 'vue'
import IpCard from '@/components/ip/IpCard.vue'
import BaseTable from '@/components/base/BaseTable.vue'
import DefaultLayout from '@/layouts/DefaultLayout.vue'
import { useRoute } from 'vue-router'
import { useIpStore } from '@/stores/ip'
import { useSubnetStore } from '@/stores/subnet'
import BaseChart from '@/components/base/BaseChart.vue'
import type { IIp } from '@/types/ip'
import { showEditVlanBox } from '@/components/subnet/vlan-edit'
import { showEditDescriptionBox } from '@/components/subnet/description-edit'
import { ISubnet } from '@/types/subnet'
import { getLastOctet } from '@/utils/util'

const route: any = useRoute()
const { params } = route

const ipStore = useIpStore()
const subnetStore = useSubnetStore()

interface IColumn {
  key: string
  dataKey: string
  title: string
  width: number
}

const loading = ref<boolean>(false)

const subnet = ref<ISubnet>({
  id: 0,
  ip: null,
  vlanId: null,
  vlanName: null,
  description: null
})

const total = ref<number>(0)
const items = computed<IIp[]>(() => ipStore.getItems)
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

const chart = reactive({
  labels: ['Занято', 'Свободно'],
  data: []
})

const getDetail = async () => {
  subnetStore.getSubnetById(params.id)
    .then((res: any) => {
      subnet.value = res
      chart.data = [
        {
          data: res.charts.map((item) => Number(item)),
          backgroundColor: ['#ff0000', '#1E252B']
        }
      ]
    }).finally(() => {
  })
}

const load = async (_page: number) => {
  loading.value = true
  ipStore.setSubnetId(params.id)
  ipStore.getIps().finally(() => {
    loading.value = false
  })
}

const onClick = (rowId: number) => {
  id.value = rowId
  showCard.value = true
}

const rowEventHandlers = {
  onClick: (e: any) => {
    onClick(e.rowData.id)
  }
}

load(-1)
getDetail()

</script>

<template>
  <default-layout>
    <el-card class="container mx-auto">
      <template #header>
        <div class="flex justify-between">
          <h1 class="sm:text-2xl pt-1 font-extrabold dark:text-white tracking-tight">{{ subnet.ip }}</h1>
        </div>
      </template>
      <el-row>
        <el-col :span="6">
        <el-card class="h-100">
          <el-descriptions :column="1">
            <el-descriptions-item label="VLAN:">
              {{ subnet.vlanName }}
              <el-button
                type="info"
                size="small"
                link
                @click="showEditVlanBox(subnet)"
              >
                Ред
              </el-button>
            </el-descriptions-item>
            <el-descriptions-item label="Описание:">
              {{ subnet.description }}
              <el-button
                type="info"
                size="small"
                link
                @click="showEditDescriptionBox(subnet)"
              >
                Ред
              </el-button>
            </el-descriptions-item>
          </el-descriptions>
          <el-divider />
          <base-chart
            v-if="chart.data.length > 0"
            :labels="chart.labels"
            :data="chart.data"
          />
          </el-card>
        </el-col>
        <el-col :span="18">
          <el-card class="h-100">
            <h3 class="mb-3 font-extrabold dark:text-white tracking-tight">IP-сетка</h3>
            <el-row :gutter="20">
              <el-col
                v-for="(item, index) in items"
                :key="index"
                :span="1"
                :style="{ padding: '5px' }"
              >
                <el-button
                  @click="onClick(item.id)"
                  class="ip-table"
                  :class="{ active: item.busy }"
                >
                  {{ getLastOctet(item.ip) }}
                </el-button>
              </el-col>
            </el-row>
          </el-card>
        </el-col>
      </el-row>
      <el-card>
        <h3 class="mb-3 font-extrabold dark:text-white tracking-tight">IP-таблица</h3>
        <div style="height: 690px">
          <base-table
            v-loading="loading"
            :columns="columns"
            :items="items"
            @load="load(1)"
            :row-event-handlers="rowEventHandlers"
          />
        </div>
      </el-card>
    </el-card>
    <ip-card
      v-if="showCard"
      v-model="showCard"
      :id="id"
      @on-reset="load"
    />
  </default-layout>
</template>

<style lang="scss" scoped>
.h-100 {
  height: 100%;
}

.ip-table {
  width: 28px;
  height: 22px;

  &.active {
    background: #af0000;
  }

  &.active:hover {
    background: #c1d4e3;
  }
}
</style>
