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
}

const form = ref<IItem>({
  id: 0,
  ip: '',
})

const load = async () => {
  ipStore.getIp(props.id)
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
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
