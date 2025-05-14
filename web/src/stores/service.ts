import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { ServiceService } from '@/api/pb/service_pb'
import { ConnectError } from '@connectrpc/connect'

const serviceService = client(ServiceService)

export const useServiceStore = defineStore('service', {
  state: () => ({}),
  actions: {
    async createService(serviceData: { name: string }) {
      try {
        return await serviceService.createService({
          name: serviceData.name,
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async getServices(page: number, pageSize: number) {
      try {
        return await serviceService.getServices({
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

    async getServiceById(id: number) {
      try {
        return await serviceService.getServiceById({ id: id })
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
