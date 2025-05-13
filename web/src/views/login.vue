<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import router from '@/router'
import { useUserStore } from '@/stores/user'

const userStore = useUserStore()

interface ILogin {
  username: string
  password: string
}

const loading = ref<boolean>(false)

const ruleFormRef = ref<FormInstance>()

const form = reactive<ILogin>({
  username: '',
  password: ''
})

const rules = reactive<FormRules<ILogin>>({
  username: [
    {
      required: true,
      message: 'Поле «Имя пользователя» должно быть заполнено',
      trigger: 'blur'
    },
    {
      min: 3,
      message: 'Должна быть не менее 3 символов',
      trigger: 'blur'
    }
  ],
  password: [
    {
      required: true,
      message: 'Поле «Пароль» должно быть заполнено',
      trigger: 'blur'
    },
    {
      min: 8,
      message: 'Длина пароля должна быть не менее 8',
      trigger: 'blur'
    }
  ]
})

const onSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid: boolean) => {
    if (valid) {
      loading.value = true
      userStore.login(form, (e: boolean) => (loading.value = e), () => router.push('/'))
    }
  })
}
</script>

<template>
  <div class="flex flex-col items-center justify-center px-6 py-8 mx-auto md:h-screen lg:py-0">
    <el-card class="w-full rounded-lg dark:border sm:max-w-md dark:bg-gray-800 dark:border-gray-700">
      <div class="md:space-y-6 sm:p-5">
        <h1 class="font-bold md:text-2xl dark:text-white text-center">
          IP Manager
        </h1>
        <el-form
          ref="ruleFormRef"
          :model="form"
          :rules="rules"
          label-position="top"
        >
          <el-form-item
            label="Имя пользователя"
            prop="username"
            class="font-medium text-gray-900 dark:text-white"
          >
            <el-input
              v-model="form.username"
              autocomplete="username"
            />
          </el-form-item>
          <el-form-item
            label="Пароль"
            prop="password"
            class="block mb-2 text-sm font-medium text-gray-900 dark:text-white"
          >
            <el-input
              type="password"
              v-model="form.password"
              autocomplete="password"
            />
          </el-form-item>
          <el-button
            type="primary"
            :loading="loading"
            class="w-full"
            @click="onSubmit(ruleFormRef)"
          >
            Войти
          </el-button>
        </el-form>
      </div>
    </el-card>
  </div>
</template>

<style scoped></style>
