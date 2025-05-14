<script lang="ts" setup>
import { reactive, ref } from 'vue'
import type { FormInstance, FormRules } from 'element-plus'
import { ElSelectV2 } from 'element-plus'
import { useSubnetStore } from '@/stores/subnet'
import { useVlanStore } from '@/stores/vlan'
import { IVlan } from '@/types/vlan'
import { ISubnet } from '@/types/subnet'

const modelValue = defineModel()
const emit = defineEmits(['onReset'])

const subnetStore = useSubnetStore()
const vlanStore = useVlanStore()

const ruleFormRef = ref<FormInstance>()
const loading = ref<boolean>(false)

const form = reactive<ISubnet>({
  id: 0,
  ip: '',
  vlanId: null,
  vlanName: null,
  description: ''
})

const rules = reactive<FormRules<ISubnet>>({
  ip: [
    {
      required: true,
      message: 'Подсеть не должен быть пустым.',
      trigger: 'blur'
    },
    {
      min: 3,
      message: 'Подсеть не должен быть короче 4 символов.',
      trigger: 'blur'
    }
  ],
  vlanId: [],
  description: [],
})

interface IVlanItem {
  value: number
  label: string
}

const vlans = ref<IVlanItem[]>([])

const loadVlan = async () => {
  loading.value = true
  vlanStore.getVlans(1, 100)
    .then((res: { total: number; items: IVlan[] }) => {
      vlans.value = res.items.map((item) => {
        return {
          value: Number(item.id),
          label: item.name
        }
      })
    }).finally(() => {
    loading.value = false
  })
}

const onSubmit = async (formEl: FormInstance | undefined) => {
  if (!formEl) return
  await formEl.validate((valid: boolean) => {
    if (valid) {
      subnetStore.createSubnet(form.ip, form.vlanId, form.description)
        .then(async (res: any) => {
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

loadVlan()
</script>

<template>
  <el-dialog
    v-model="modelValue"
    title="Добавить подсеть"
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
      <el-form-item
        label="VLAN"
        prop="vlan"
      >
        <el-select-v2
          v-model="form.vlanId"
          class="w-100"
          :options="vlans"
        />
      </el-form-item>
      <el-form-item
        label="Описание"
        prop="description"
      >
        <el-input
          type="textarea"
          v-model="form.description"
        />
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
