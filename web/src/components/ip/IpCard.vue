<script lang="ts" setup>
import { computed } from 'vue'
import { useIpStore } from '@/stores/ip'
import { showEditCustomerBox } from './customer-edit'
import { showEditServiceBox } from './service-edit'
import { showEditDescriptionBox } from './description-edit'
import type { IIp } from '@/types/ip'

const modelValue = defineModel()
const props = defineProps(['id'])

const form = computed<IIp>(() => ipStore.getItem)

const ipStore = useIpStore()

const load = async () => {
  ipStore.setId(props.id)
  await ipStore.getIpById()
}

const onClose = () => ipStore.getIps().then(() => {})

load()
</script>

<template>
  <el-dialog
    v-model="modelValue"
    :title="`Карточка IP (${form.ip})`"
    width="500"
    align-center
    @close="onClose"
  >
    <el-descriptions v-loading="ipStore.getLoading" :column="1">
      <el-descriptions-item label="IP:">{{form.ip}}</el-descriptions-item>
      <el-descriptions-item label="Клиент:">
        {{form.customerName}}
        <el-button
          type="info"
          size="small"
          link
          @click="showEditCustomerBox(form)"
        >
          Ред
        </el-button>
      </el-descriptions-item>
      <el-descriptions-item label="Сервис:">
        {{form.serviceName}}
        <el-button
          type="info"
          size="small"
          link
          @click="showEditServiceBox(form)"
        >
          Ред
        </el-button>
      </el-descriptions-item>
      <el-descriptions-item label="Описание:">
        {{form.description}}
        <el-button
          type="info"
          size="small"
          link
          @click="showEditDescriptionBox(form)"
        >
          Ред
        </el-button>
      </el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
