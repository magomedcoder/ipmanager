import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { IpService } from '@/api/pb/ip_pb'
import { ConnectError } from '@connectrpc/connect'

const ipService = client(IpService)

export const useIpStore = defineStore('ip', {
  state: () => ({
    loading: false,
    subnetId: 0,
    total: 0,
    items: [],
    id: 0,
    item: {
      id: 0,
      ip: '',
      customerId: 0,
      customerName: '',
      serviceId: 0,
      serviceName: '',
      description: '',
    }
  }),
  actions: {
    setSubnetId(subnetId: number) {
      this.subnetId = subnetId
    },

    setId(id: number) {
      this.id = id
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

    async getIpById() {
      this.loading = true
      try {
        const res = await ipService.getIpById({ id: this.id })
        this.item = {
          ...res,
          id: Number(res.id),
          serviceId: Number(res.serviceId),
          customerId: Number(res.customerId),
        }
        this.loading = false
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        this.loading = false
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
    getLoading: state => state.loading,

    getTotal: state => state.total,

    getItems: state => state.items,

    getItem: state => state.item,
  }
})
