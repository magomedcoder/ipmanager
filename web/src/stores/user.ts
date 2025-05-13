import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { UserService } from '@/api/pb/user_pb'
import { ConnectError } from '@connectrpc/connect'
import { setAccessToken, removeAccessToken } from '@/utils/cookie'
import { useJwt } from '@vueuse/integrations/useJwt'

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
        const { payload } = useJwt(accessToken)
        setAccessToken(accessToken, payload.value?.exp)
        success()
        loading(false)
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
    },

    createUser(user: { username: string, password: string, name: string, surname: string }): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await userService.createUser({
            username: user.username,
            password: user.password,
            name: user.name,
            surname: user.surname
          })
          resolve(res)
        } catch (err) {
          if (err instanceof ConnectError) {
            console.log(err.message)
          }
          reject()
        }
      })
    },

    getUsers(page: number, pageSize: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await userService.getUsers({
            page: page,
            pageSize: pageSize
          })
          resolve(res)
        } catch (err) {
          if (err instanceof ConnectError) {
            console.log(err.message)
          }
          reject()
        }
      })
    },

    getUser(id: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await userService.getUserById({ id: id })
          resolve(res)
        } catch (err) {
          if (err instanceof ConnectError) {
            console.log(err.message)
          }
          reject()
        }
      })
    }
  },
  getters: {

  }
})
