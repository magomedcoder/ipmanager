import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { IpService } from '@/api/pb/ip_pb'
import { ConnectError } from '@connectrpc/connect'

const ipService = client(IpService)

export const useIpStore = defineStore('ip', {
  state: () => ({

  }),
  actions: {
    createIp(ipData: { ip: string }): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await ipService.createIp({
            ip: ipData.ip,
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

    getIps(page: number, pageSize: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await ipService.getIps({
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

    getIpById(id: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await ipService.getIpById({ id: id })
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
