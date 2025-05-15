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
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
      } finally {
        loading(false)
      }
    },

    async logout(loading: (value: boolean) => void) {
      loading(true)
      try {
        const { success } = await userService.logout()
        console.log(success)
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
      } finally {
        loading(false)
      }
    },

    async password(
      form: IPassword,
      success: (value: boolean) => void
    ) {
      try {
        const res = await userService.password({
          oldPassword: form.oldPassword,
          newPassword: form.password
        })
        success(res.success)
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
      }
    },

    async createUser(user: { username: string, password: string, name: string, surname: string }) {
      try {
        return await userService.createUser({
          username: user.username,
          password: user.password,
          name: user.name,
          surname: user.surname
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async getUsers(page: number, pageSize: number) {
      try {
        return await userService.getUsers({
          page: page,
          pageSize: pageSize
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async getUser(id: number) {
      try {
        return await userService.getUserById({ id: id })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    }
  },
  getters: {}
})
