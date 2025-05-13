import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { ServiceService } from '@/api/pb/service_pb'
import { ConnectError } from '@connectrpc/connect'

const serviceService = client(ServiceService)

export const useServiceStore = defineStore('service', {
  state: () => ({

  }),
  actions: {
    createService(serviceData: { name: string }): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await serviceService.createService({
            name: serviceData.name,
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

    getServices(page: number, pageSize: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await serviceService.getServices({
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

    getServiceById(id: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await serviceService.getServiceById({ id: id })
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
