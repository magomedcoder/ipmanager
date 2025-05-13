<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useCustomerStore } from '@/stores/customer'

const modelValue = defineModel()
const emit = defineEmits(['onReset'])

const customerStore = useCustomerStore()

interface ICustomer {
  name: string
  surname: string
}

const ruleFormRef = ref<FormInstance>()

const form = reactive<ICustomer>({
  name: '',
  surname: ''
})

const rules = reactive<FormRules<ICustomer>>({
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
      customerStore.createCustomer({
        name: form.name,
        surname: form.surname,
      }).then(async (res: any) => {
        if (res.id !=0) {
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
    title="Добавить"
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
