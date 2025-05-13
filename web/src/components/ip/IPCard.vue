<script lang="ts" setup>
import { ref } from 'vue'
import { useIpStore } from '@/stores/ip'

const modelValue = defineModel()
const props = defineProps(['id'])

const ipStore = useIpStore()

const onClose = () => modelValue.value = false

interface IItem {
  id: number
  ip: string
  vlanName: string
  customerName: string
  description: string
}

const form = ref<IItem>({
  id: 0,
  ip: '',
  vlanName: '',
  customerName: '',
  description: '',
})

const load = async () => {
  ipStore.getIpById(props.id)
    .then(async (res: any) => {
      form.value = res
    }).finally(() => {
  })
}

load()
</script>

<template>
  <el-dialog
    v-model="modelValue"
    title="Карточка IP"
    width="500"
    align-center
  >
    <el-descriptions :column="1">
      <el-descriptions-item label="IP:">{{form.ip}}</el-descriptions-item>
      <el-descriptions-item label="Vlan:">{{form.vlanName}}</el-descriptions-item>
      <el-descriptions-item label="Клиент:">{{form.customerName}}</el-descriptions-item>
      <el-descriptions-item label="Описание:">{{form.description}}</el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
