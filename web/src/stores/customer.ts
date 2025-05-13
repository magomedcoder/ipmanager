import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { CustomerService } from '@/api/pb/customer_pb'
import { ConnectError } from '@connectrpc/connect'

const customerService = client(CustomerService)

export const useCustomerStore = defineStore('customer', {
  state: () => ({

  }),
  actions: {
    createCustomer(customerData: { name: string; surname: string }): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await customerService.createCustomer({
            name: customerData.name,
            surname: customerData.surname,
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

    getCustomers(page: number, pageSize: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await customerService.getCustomers({
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

    getCustomerById(id: number): Promise<any> {
      return new Promise(async (resolve, reject) => {
        try {
          const res = await customerService.getCustomerById({ id: id })
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
