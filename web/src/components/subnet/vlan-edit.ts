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
const vlans = ref<IVlan[]>([])
const selected = ref<number | null>(null)

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
  selected.value = (form.vlanId !=0) ? Number(form.vlanId) : null

  vlanStore.getVlans(1, 100)
    .then((res: { total: number; items: IVlan[] }) => {
      vlans.value = res.items.map((item) => {
        return {
          id: Number(item.id),
          name: item.name
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
          vlan: selected.value
        },
        rules,
        labelPosition: 'top'
      }, [
        h(ElFormItem, { prop: 'vlan' }, [
          h(ElSelect, {
            modelValue: selected.value,
            'onUpdate:modelValue': (value: number) => selected.value = value,
            placeholder: 'Выберите VLAN'
          }, vlans.value.map(vlan =>
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
          subnetStore.editVlanById(id.value, selected.value).then(() => done())
          subnetStore.getSubnetById()
        })
      } else {
        done()
      }
    },
  })
    .catch(() => { })
}
