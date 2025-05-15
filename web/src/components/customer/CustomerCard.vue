<script lang="ts" setup>
import { ref } from 'vue'
import { useCustomerStore } from '@/stores/customer'

const modelValue = defineModel()
const props = defineProps(['id'])

const customerStore = useCustomerStore()

interface IItem {
  id: number
  name: string
  surname: string
}

const form = ref<IItem>({
  id: 0,
  name: '',
  surname: ''
})

const load = async () => {
  customerStore.getCustomerById(props.id)
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
    title="Карточка клиента"
    width="500"
    align-center
  >
    <el-descriptions :column="1">
      <el-descriptions-item label="Имя:">{{form.name}}</el-descriptions-item>
      <el-descriptions-item label="Фамилия:">{{form.surname}}</el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
