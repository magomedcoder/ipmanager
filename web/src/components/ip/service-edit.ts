import { ref, h } from 'vue'
import { ElMessageBox, ElSelect, ElOption, ElForm, ElFormItem } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import type { IIp } from '@/types/ip'
import type { IService } from '@/types/service'
import { validateForm } from '@/utils/validate'
import { useIpStore } from '@/stores/ip'
import { useServiceStore } from '@/stores/service'

const ipStore = useIpStore()
const serviceStore = useServiceStore()

const formRef = ref<FormInstance>()
const id = ref<number>()
const serviceList = ref<IService[]>([])
const selected = ref<number | null>(null)

const rules: FormRules = {
  service: [
    {
      required: false,
      message: 'Сервис не должен быть пустым.',
      trigger: 'blur'
    }
  ]
}

export const showEditServiceBox = (form: IIp) => {
  id.value = form.id
  selected.value = (form.serviceId !=0) ? form.serviceId : null

  serviceStore.getServices(1, 100)
    .then(async (res: { total: number; items: IService[] }) => {
      serviceList.value = res.items.map((item) => {
        return {
          ...item,
          id: Number(item.id)
        }
      })

    }).finally(() => {})

  ElMessageBox({
    title: 'Редактирование сервиса',
    showCancelButton: true,
    confirmButtonText: 'Сохранить',
    message: () => h('div', [
      h(ElForm, {
        ref: formRef,
        model: {
          service: selected.value
        },
        rules,
        labelPosition: 'top'
      }, [
        h(ElFormItem, { prop: 'service' }, [
          h(ElSelect, {
            modelValue: selected.value,
            'onUpdate:modelValue': (value: number) => selected.value = value,
            placeholder: 'Выберите cервис',
            clearable: true
          }, serviceList.value.map(service =>
            h(ElOption, {
              label: service.name,
              value: service.id
            })
          ))
        ])
      ])
    ]),
    beforeClose: (action, instance, done) => {
      if (action === 'confirm') {
        validateForm(formRef).then(() => {
          ipStore.editServiceById(id.value, selected.value).then(() => done())
          ipStore.getIpById()
        })
      } else {
        done()
      }
    },
  })
    .catch(() => { })
}
