<script lang="ts" setup>
import { ref } from 'vue'
import { useUserStore } from '@/stores/user'

const modelValue = defineModel()
const props = defineProps(['id'])

const userStore = useUserStore()

interface IUser {
  id: number
  username: string
  name: string
  surname: string
}

const form = ref<IUser>({
  id: 0,
  username: '',
  name: '',
  surname: '',
})

const load = async () => {
  userStore.getUser(props.id)
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
    title="Карточка пользователя"
    width="500"
    align-center
  >
    <el-descriptions :column="1">
      <el-descriptions-item label="Имя пользователя:">@ {{form.username}}</el-descriptions-item>
      <el-descriptions-item label="Имя:">{{form.name}}</el-descriptions-item>
      <el-descriptions-item label="Фамилия:">{{form.surname}}</el-descriptions-item>
    </el-descriptions>
  </el-dialog>
</template>

<style scoped></style>
