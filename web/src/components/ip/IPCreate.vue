<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { useIpStore } from '@/stores/ip'

const modelValue = defineModel()
const emit = defineEmits(['onReset'])

const ipStore = useIpStore()

interface IIP {
  ip: string
}

const ruleFormRef = ref<FormInstance>()

const form = reactive<IIP>({
  ip: '',
})

const rules = reactive<FormRules<IIP>>({
  ip: [
    {
      required: true,
      message: 'IP не должен быть пустым.',
      trigger: 'blur'
    },
    {
      min: 3,
      message: 'IP не должен быть короче 4 символов.',
      trigger: 'blur'
    }
  ],
})

const onSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid: boolean) => {
    if (valid) {
      ipStore.createIp({
        ip: form.ip,
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
    title="Добавить IP"
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
        label="IP"
        prop="ip"
      >
        <el-input v-model="form.ip" />
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
