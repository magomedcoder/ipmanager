import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { VlanService } from '@/api/pb/vlan_pb'
import { ConnectError } from '@connectrpc/connect'

const vlanService = client(VlanService)

export const useVlanStore = defineStore('vlan', {
  state: () => ({

  }),
  actions: {
    createVlan(vlanData: { name: string }): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await vlanService.createVlan({
            name: vlanData.name,
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

    getVlans(page: number, pageSize: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await vlanService.getVlans({
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

    getVlanById(id: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await vlanService.getVlanById({ id: id })
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
