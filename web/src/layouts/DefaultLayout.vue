<script lang="ts" setup>
import { ref, h } from 'vue'
import { ElMessageBox, ElMessage, ElInput, ElForm, ElFormItem } from 'element-plus'
import type { FormInstance, FormRules } from 'element-plus'
import router from '@/router'
import { useUserStore } from '@/stores/user'
import { User, Operation } from '@element-plus/icons-vue'

const userStore = useUserStore()

const onLogout = async () => userStore.logout(() => router.push('/login'))

const oldPassword = ref<string>('')
const newPassword = ref<string>('')
const confirmPassword = ref<string>('')
const formRef = ref<FormInstance | null>(null)

const validateForm = (): Promise<boolean> => {
  return new Promise((resolve, reject) => {
    if (formRef.value) {
      formRef.value.validate((valid: boolean) => {
        if (valid) {
          resolve(true)
        } else {
          reject(false)
        }
      })
    } else {
      reject(false)
    }
  })
}

const rules: FormRules = {
  oldPassword: [
    {
      required: true,
      message: 'Введите старый пароль',
      trigger: 'blur'
    }
  ],
  newPassword: [
    {
      required: true,
      message: 'Введите новый пароль',
      trigger: 'blur'
    },
    {
      min: 8,
      message: 'Пароль должен быть не менее 8 символов',
      trigger: 'blur'
    }
  ],
  confirmPassword: [
    {
      required: true,
      message: 'Подтвердите новый пароль',
      trigger: 'blur'
    },
    {
      min: 8,
      message: 'Пароль должен быть не менее 8 символов',
      trigger: 'blur'
    }
  ]
}

const showChangePasswordBox = () => {
  ElMessageBox({
    title: 'Смена пароля',
    showCancelButton: true,
    confirmButtonText: 'Изменить',
    message: () => h(ElForm, {
      ref: formRef,
      model: {
        oldPassword: oldPassword.value,
        newPassword: newPassword.value,
        confirmPassword: confirmPassword.value
      },
      rules,
      labelPosition: 'top'
    }, [
      h(ElFormItem, {
        label: 'Старый пароль',
        prop: 'oldPassword'
      }, [
        h(ElInput, {
          type: 'password',
          modelValue: oldPassword.value,
          'onUpdate:modelValue': (value: string) => oldPassword.value = value,
          placeholder: 'Старый пароль'
        })
      ]),
      h(ElFormItem, {
        label: 'Новый пароль',
        prop: 'newPassword'
      }, [
        h(ElInput, {
          type: 'password',
          modelValue: newPassword.value,
          'onUpdate:modelValue': (value: string) => newPassword.value = value,
          placeholder: 'Новый пароль'
        })
      ]),
      h(ElFormItem, {
        label: 'Подтвердите новый пароль',
        prop: 'confirmPassword'
      }, [
        h(ElInput, {
          type: 'password',
          modelValue: confirmPassword.value,
          'onUpdate:modelValue': (value: string) => confirmPassword.value = value,
          placeholder: 'Подтвердите новый пароль',
        })
      ])
    ]),
    beforeClose: (action, instance, done) => {
      if (action === 'confirm') {
        validateForm().then(() => {
          if (newPassword.value !== confirmPassword.value) {
            ElMessage({
              type: 'error',
              message: 'Новые пароли не совпадают!'
            })
          } else {
            userStore.password({
              oldPassword: oldPassword.value,
              password: confirmPassword.value
            }, (success: boolean) => {
              if (success) {
                ElMessage({
                  type: 'success',
                  message: 'Пароль успешно изменен!'
                })
                done()
              } else {
                ElMessage({
                  type: 'error',
                  message: 'Неверно введен пароль!'
                })
              }
            })
          }
        }).catch(() => { })
      } else {
        done()
      }
    },
  })
    .catch(() => { })
}

</script>

<template>
  <el-container>
    <el-header class="fixed w-full">
      <div class="container mx-auto">
        <el-menu
          mode="horizontal"
          text-color="white"
          :router="true"
          :ellipsis="false"
        >
          <el-menu-item :route="{ path: '/' }">
            <h3 class="header__logo">IP Manager</h3>
          </el-menu-item>
          <el-menu-item
            :route="{ path: '/' }"
            index="1"
          >
            IP
          </el-menu-item>
          <el-menu-item
            :route="{ path: '/vlans' }"
            index="2"
          >
            VLAN
          </el-menu-item>
          <el-menu-item
            :route="{ path: '/customers' }"
            index="3"
          >
            Клиенты
          </el-menu-item>
          <el-menu-item
            :route="{ path: '/users' }"
            index="4"
          >
            Пользователи
          </el-menu-item>
          <div class="flex-grow" />
          <div class="toolbar">
            <el-dropdown>
              <el-icon size="20"><User /></el-icon>
              <template #dropdown>
                <el-dropdown-menu>
                  <el-dropdown-item @click="showChangePasswordBox()">Сменить пароль</el-dropdown-item>
                  <el-dropdown-item @click="onLogout()" divided>Выйти</el-dropdown-item>
                </el-dropdown-menu>
              </template>
            </el-dropdown>
          </div>
        </el-menu>
      </div>
    </el-header>
    <el-main class="!pt-20">
      <slot/>
    </el-main>
  </el-container>
</template>

<style lang="scss" scoped>
header {
  display: flex;
  background-color: #262D34;
  height: 60px;
  border-bottom: 1px solid #374151;

  .header__logo {
    display: flex;
    align-items: center;
    font-size: 24px;
    color: whitesmoke;
    margin: 0 20px;
  }

  .toolbar {
    display: inline-flex;
    align-items: center;
    justify-content: center;
  }
}
</style>
