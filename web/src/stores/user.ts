import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { UserService } from '@/api/pb/user_pb'
import { ConnectError } from '@connectrpc/connect'

const userService = client(UserService)

interface ILogin {
  username: string
  password: string
}

interface IPassword {
  oldPassword: string
  password: string
}

export const useUserStore = defineStore('user', {
  state: () => ({
    accessToken: ''
  }),
  actions: {
    async login(
      form: ILogin,
      loading: (value: boolean) => void,
      success: () => void,
    ) {
      loading(true)
      try {
        const { accessToken } = await userService.login({
          username: form.username,
          password: form.password
        })
        success()
        loading(false)
        console.log(accessToken)
      } catch (err) {
        loading(false)
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
      }
    },

    async logout(loading: (value: boolean) => void) {
      loading(true)
      try {
        const { success } = await userService.logout()
        loading(false)
        console.log(success)
      } catch (err) {
        loading(false)
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
      }
    },

    async password(
      form: IPassword,
      success: (value: boolean) => void
    ) {
      try {
        const { success } = await userService.password({
          oldPassword: form.oldPassword,
          password: form.password
        })
        success(success)
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
      }
    }
  },
  getters: {

  }
})
