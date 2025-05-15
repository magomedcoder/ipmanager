import { ref, h } from 'vue'
import { ElMessageBox, ElInput, ElForm, ElFormItem } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import type { IIp } from '@/types/ip'
import { validateForm } from '@/utils/validate'
import { useSubnetStore } from '@/stores/subnet'

const subnetStore = useSubnetStore()

const formRef = ref<FormInstance>()
const id = ref<number>()
const description = ref<string>('')

const rules: FormRules = {
  description: [
    {
      required: true,
      message: 'Описание не должно быть пустым.',
      trigger: 'blur'
    }
  ]
}

export const showEditDescriptionBox = (form: IIp) => {
  id.value = form.id
  description.value = form.description

  ElMessageBox({
    title: 'Редактирование описание',
    showCancelButton: true,
    confirmButtonText: 'Сохранить',
    message: () => h('div', [
      h(ElForm, {
        ref: formRef,
        model: { description: description.value },
        rules,
        labelPosition: 'top'
      }, {
        default: () => [
          h(ElFormItem, { prop: 'description' }, {
            default: () => [
              h(ElInput, {
                type: "textarea",
                modelValue: description.value,
                'onUpdate:modelValue': (value: string) => description.value = value,
                placeholder: 'Описание'
              })
            ]
          })
        ]
      })
    ]),
    beforeClose: (action, instance, done) => {
      if (action === 'confirm') {
        validateForm(formRef).then(() => {
          subnetStore.editDescriptionById(id.value, description.value).then(() => done())
          subnetStore.getSubnetById()
        })
      } else {
        done()
      }
    },
  })
    .catch(() => { })
}
