<script lang="ts" setup>
import { ref } from 'vue'
import { useServiceStore } from '@/stores/service'

const modelValue = defineModel()
const props = defineProps(['id'])

const serviceStore = useServiceStore()

interface IItem {
  id: number
  name: string
  ipName: string[]
}

const form = ref<IItem>({
  id: 0,
  name: '',
  ipName: []
})

const load = async () => {
  serviceStore.getServiceById(props.id)
    .then(async (res: any) => {
      form.value = {
        ...res,
        ipName: res.ips.map((ip) => ip.ip).join(', ')
      }
    }).finally(() => {
  })
}

load()
</script>

<template>
  <el-dialog
    v-model="modelValue"
    title="Карточка"
    width="500"
    align-center
  >
    <el-descriptions :column="1">
      <el-descriptions-item label="Название:">{{form.name}}</el-descriptions-item>
<!--      <el-descriptions-item label="VLAN:">{{form.name}}</el-descriptions-item>-->
      <el-descriptions-item label="IP:">{{form.ipName}}</el-descriptions-item>
<!--      <el-descriptions-item label="Клиент:">{{form.name}}</el-descriptions-item>-->
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
