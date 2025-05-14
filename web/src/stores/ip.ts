import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { IpService } from '@/api/pb/ip_pb'
import { ServiceService } from '@/api/pb/service_pb'
import { ConnectError } from '@connectrpc/connect'

const ipService = client(IpService)
const serviceService = client(ServiceService)

export const useIpStore = defineStore('ip', {
  state: () => ({}),
  actions: {
    async getIps(subnetId: number) {
      try {
        return await ipService.getIps({
          subnetId: subnetId
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
  getters: {}
})
