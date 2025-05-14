import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { IpService } from '@/api/pb/ip_pb'
import { ConnectError } from '@connectrpc/connect'

const ipService = client(IpService)

export const useIpStore = defineStore('ip', {
  state: () => ({
    subnetId: 0,
    total: 0,
    items: []
  }),
  actions: {
    setSubnetId(subnetId: number) {
      this.subnetId = subnetId
    },

    async getIps() {
      try {
        const { total, items } = await ipService.getIps({
          subnetId: this.subnetId
        })
        this.total = total
        this.items = items?.map((item) => {
          return {
            ...item,
            id: Number(item.id),
            customerId: Number(item.customerId)
          }
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async getIpById(id: number) {
      try {
        return await ipService.getIpById({ id: id })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async editCustomerById(id: number, customerId: number) {
      try {
        return await ipService.editIpCustomer({
          id: id,
          customerId: customerId
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async editServiceById(id: number, serviceId: number) {
      try {
        return await ipService.editIpService({
          id: id,
          serviceId: serviceId
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async editDescriptionById(id: number, description: string) {
      try {
        return await ipService.editIpDescription({
          id: id,
          description: description
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    }
  },
  getters: {
    getTotal: state => state.total,

    getItems: state => state.items
  }
})
