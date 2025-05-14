import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { CustomerService } from '@/api/pb/customer_pb'
import { ConnectError } from '@connectrpc/connect'

const customerService = client(CustomerService)

export const useCustomerStore = defineStore('customer', {
  state: () => ({}),
  actions: {
    async createCustomer(customerData: { name: string; surname: string }) {
      try {
        return await customerService.createCustomer({
          name: customerData.name,
          surname: customerData.surname,
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }

        throw err
      }
    },

    async getCustomers(page: number, pageSize: number) {
      try {
        return await customerService.getCustomers({
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

    async getCustomerById(id: number) {
      try {
        return await customerService.getCustomerById({ id: id })
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
