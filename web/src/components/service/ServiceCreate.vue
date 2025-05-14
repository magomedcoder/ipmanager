<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useServiceStore } from '@/stores/service'

const modelValue = defineModel()
const emit = defineEmits(['onReset'])

const serviceStore = useServiceStore()

interface IService {
  name: string
}

const ruleFormRef = ref<FormInstance>()

const form = reactive<IService>({
  name: '',
})

const rules = reactive<FormRules<IService>>({
  name: [
    {
      required: true,
      message: 'Service не должен быть пустым.',
      trigger: 'blur'
    },
    {
      min: 3,
      message: 'Service не должен быть короче 4 символов.',
      trigger: 'blur'
    }
  ],
})

const onSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid: boolean) => {
    if (valid) {
      serviceStore.createService({
        name: form.name,
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
        label="Название"
        prop="name"
      >
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item
        label="VLAN"
        prop="name"
      >
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item
        label="IP"
        prop="name"
      >
        <el-input v-model="form.name" />
      </el-form-item>
      <el-form-item
        label="Клиент"
        prop="name"
      >
        <el-input v-model="form.name" />
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
