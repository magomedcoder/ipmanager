<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useUserStore } from '@/stores/user'

const modelValue = defineModel()
const emit = defineEmits(['onReset'])

const userStore = useUserStore()

interface ILogin {
  username: string
  password: string
  name: string
  surname: string
}

const ruleFormRef = ref<FormInstance>()

const form = reactive<ILogin>({
  username: '',
  password: '',
  name: '',
  surname: ''
})

const rules = reactive<FormRules<ILogin>>({
  username: [
    {
      required: true,
      message: 'Имя пользователя не должен быть пустым.',
      trigger: 'blur'
    },
    {
      min: 3,
      message: 'Имя пользователя не должен быть короче 4 символов.',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: 'Пароль не должен быть пустым.',
      trigger: 'blur'
    },
    {
      min: 8,
      message: 'Пароль не должен быть короче 8 символов.',
      trigger: 'blur'
    }
  ],
  name: [
    {
      required: true,
      message: 'Имя не должно быть пустым.',
      trigger: 'blur'
    },
    {
      min: 2,
      message: 'Имя не должен быть короче 3 символов.',
      trigger: 'blur'
    }
  ],
  surname: [
    {
      required: true,
      message: 'Фамилия не должно быть пустым.',
      trigger: 'blur'
    },
    {
      min: 2,
      message: 'Фамилия не должен быть короче 3 символов.',
      trigger: 'blur'
    }
  ]
})

const onSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid: boolean) => {
    if (valid) {
      userStore.CreateUser({
        username: form.username,
        password: form.password,
        name: form.name,
        surname: form.surname
      }).then(async (res: any) => {
        const { code, data } = res
        if (code == 200 && data.id) {
          onClose()
          emit('onReset')
        }
      }).finally(() => {

      })
    }
  })
}

const onClose = () => modelValue.value = false
</script>

<template>
  <el-dialog
    v-model="modelValue"
    title="Добавить пользователя"
    width="500"
    align-center
  >
    <el-form
      ref="ruleFormRef"
      :model="form"
      :rules="rules"
      label-position="top"
    >
      <el-form-item
        label="Имя пользователя"
        prop="username"
      >
        <el-input v-model="form.username" />
      </el-form-item>
      <el-form-item
        label="Пароль"
        prop="password"
      >
        <el-input
          v-model="form.password"
          type="password"
        />
      </el-form-item>
      <el-form-item
        label="Имя"
        prop="name"
      >
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item
        label="Фамилия"
        prop="surname"
      >
        <el-input v-model="form.surname" />
      </el-form-item>
      <el-button
        type="primary"
        class="w-full"
        @click="onSubmit(ruleFormRef)"
      >
        Добавить
      </el-button>
    </el-form>
  </el-dialog>
</template>

<style scoped></style>
