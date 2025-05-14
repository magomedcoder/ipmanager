<script lang="ts" setup>
import { ref } from 'vue'
import { useIpStore } from '@/stores/ip'
import { showEditCustomerBox } from './customer-edit'
import { showEditServiceBox } from './service-edit'
import { showEditDescriptionBox } from './description-edit'
import type { IIp } from '@/types/ip'

const modelValue = defineModel()
const props = defineProps(['id'])

const ipStore = useIpStore()

const onClose = () => modelValue.value = false

const form = ref<IIp>({
  id: 0,
  ip: '',
  customerId: 0,
  customerName: '',
  serviceId: 0,
  serviceName: '',
  description: '',
})

const load = async () => {
  ipStore.getIpById(props.id)
    .then(async (res: any) => {
      form.value = {
      ...res,
        id: Number(res.id),
        serviceId: Number(res.serviceId),
        customerId: Number(res.customerId),
      }
    }).finally(() => {
  })
}

load()
</script>

<template>
  <el-dialog
    v-model="modelValue"
    :title="`Карточка IP (${form.ip})`"
    width="500"
    align-center
  >
    <el-descriptions :column="1">
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
    <el-button
      type="primary"
      @click=""
    >
      Зарезервировать
    </el-button>
    <el-button
      type="primary"
      @click=""
    >
      Освободить
    </el-button>
  </el-dialog>
</template>

<style scoped></style>
