<script lang="ts" setup>
import { ref } from 'vue'
import { useVlanStore } from '@/stores/vlan'

const modelValue = defineModel()
const props = defineProps(['id'])

const vlanStore = useVlanStore()

const onClose = () => modelValue.value = false

interface IItem {
  id: number
  name: string
}

const form = ref<IItem>({
  id: 0,
  name: ''
})

const load = async () => {
  vlanStore.getVlanById(props.id)
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
    title="Карточка"
    width="500"
    align-center
  >
    <el-descriptions :column="1">
      <el-descriptions-item label="Название:">{{form.name}}</el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
