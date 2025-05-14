import { ref, h } from 'vue'
import { ElMessageBox, ElSelect, ElOption, ElForm, ElFormItem } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import { useCustomerStore } from '@/stores/customer'
import type { IIp } from '@/types/ip'
import type { ICustomer } from '@/types/customer'
import { validateForm } from '@/utils/validate'
import { useIpStore } from '@/stores/ip'

const customerStore = useCustomerStore()
const ipStore = useIpStore()

const formRef = ref<FormInstance>()
const id = ref<number>()
const customerList = ref<ICustomer[]>([])
const selected = ref<number>()

const rules: FormRules = {
  customer: [
    {
      required: true,
      message: 'Клиент не должен быть пустым.',
      trigger: 'blur'
    }
  ]
}

export const showEditCustomerBox = (form: IIp) => {
  id.value = form.id
  if (form.customerId !=0) {
    selected.value = form.customerId
  }

  customerStore.getCustomers(1, 100)
    .then(async (res: { total: number; items: ICustomer[] }) => {
      customerList.value = res.items.map((item) => {
        return {
          ...item,
          id: Number(item.id)
        }
      })

    }).finally(() => {})

  ElMessageBox({
    title: 'Редактирование клиент',
    showCancelButton: true,
    confirmButtonText: 'Сохранить',
    message: () => h('div', [
      h(ElForm, {
        ref: formRef,
        model: {
          customer: selected.value
        },
        rules,
        labelPosition: 'top'
      }, [
        h(ElFormItem, { prop: 'customer' }, [
          h(ElSelect, {
            modelValue: selected.value,
            'onUpdate:modelValue': (value: number) => selected.value = value,
            placeholder: 'Выберите клиент'
          }, customerList.value.map(customer =>
            h(ElOption, {
              label: customer.name,
              value: customer.id
            })
          ))
        ])
      ])
    ]),
    beforeClose: (action, instance, done) => {
      if (action === 'confirm') {
        validateForm(formRef).then(() => {
          ipStore.editCustomerById(id.value, selected.value).then(() => done())
        })
      } else {
        done()
      }
    },
  })
    .catch(() => { })
}
