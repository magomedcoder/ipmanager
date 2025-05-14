import { ref, h } from 'vue'
import { ElMessageBox, ElSelect, ElOption, ElForm, ElFormItem } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useVlanStore } from '@/stores/vlan'
import type { ISubnet } from '@/types/subnet'
import type { IVlan } from '@/types/vlan'
import { validateForm } from '@/utils/validate'
import { useSubnetStore } from '@/stores/subnet'

const subnetStore = useSubnetStore()
const vlanStore = useVlanStore()

const formRef = ref<FormInstance>()
const id = ref<number>()
const vlanList = ref<IVlan[]>([])
const selectedVlan = ref<number | null>(null)

const rules: FormRules = {
  vlan: [
    {
      required: true,
      message: 'VLAN не должен быть пустым.',
      trigger: 'blur'
    }
  ]
}

export const showEditVlanBox = (form: ISubnet) => {
  id.value = form.id
  selectedVlan.value = form.vlanId

  vlanStore.getVlans(1, 100)
    .then(async (res: { total: number; items: IVlan[] }) => {
      vlanList.value = res.items.map((item) => {
        return {
          ...item,
          id: Number(item.id)
        }
      })
    }).finally(() => {})

  ElMessageBox({
    title: 'Редактирование VLAN',
    showCancelButton: true,
    confirmButtonText: 'Сохранить',
    message: () => h('div', [
      h(ElForm, {
        ref: formRef,
        model: {
          vlan: selectedVlan.value
        },
        rules,
        labelPosition: 'top'
      }, [
        h(ElFormItem, { prop: 'vlan' }, [
          h(ElSelect, {
            modelValue: selectedVlan.value,
            'onUpdate:modelValue': (value: number) => selectedVlan.value = value,
            placeholder: 'Выберите VLAN'
          }, vlanList.value.map(vlan =>
            h(ElOption, {
              label: vlan.name,
              value: vlan.id
            })
          ))
        ])
      ])
    ]),
    beforeClose: (action, instance, done) => {
      if (action === 'confirm') {
        validateForm(formRef).then(() => {
          subnetStore.editVlanById(id.value, selectedVlan.value).then(() => done())
        })
      } else {
        done()
      }
    },
  })
    .catch(() => { })
}
