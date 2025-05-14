import type { Ref } from 'vue'

export const validateForm = (formRef: Ref): Promise<boolean> => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value?.validate((valid: boolean) => {
        if (valid) {
          resolve(true)
        } else {
          reject(new Error('Ошибка валидации формы'))
        }
      }).then()
    } else {
      reject(new Error('Ссылка на форму не определена'))
    }
  })
}
