import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { VlanService } from '@/api/pb/vlan_pb'
import { ConnectError } from '@connectrpc/connect'

const vlanService = client(VlanService)

export const useVlanStore = defineStore('vlan', {
  state: () => ({}),
  actions: {
    async createVlan(vlanData: { name: string }) {
      try {
        return await vlanService.createVlan({
          name: vlanData.name,
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err // Propagate error
      }
    },

    async getVlans(page: number, pageSize: number) {
      try {
        return await vlanService.getVlans({
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

    async getVlanById(id: number) {
      try {
        return await vlanService.getVlanById({ id: id })
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
