import { defineStore } from 'pinia'
import { client } from '@/api/client'
import { SubnetService } from '@/api/pb/subnet_pb'
import { ConnectError } from '@connectrpc/connect'

const subnetService = client(SubnetService)

export const useSubnetStore = defineStore('subnet', {
  state: () => ({
    id: 0,
    item: {},
    chart: {}
  }),
  actions: {
    setId(id: number) {
      this.id = id
    },

    async createSubnet(ip: string, vlanId: number, description: string) {
      try {
        return await subnetService.createSubnet({
          ip: ip,
          vlanId: vlanId,
          description: description,
        })
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async getSubnets(page: number, pageSize: number) {
      try {
        return await subnetService.getSubnets({
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

    async getSubnetById() {
      try {
        const res = await subnetService.getSubnetById({ id: this.id })
        if (res.id !=0) {
          this.item = res
          this.chart = {
            labels: ['Занято', 'Свободно'],
            data: [{
              data: res.charts.map((item) => Number(item)),
              backgroundColor: ['#ff0000', '#1E252B']
            }]
          }
        }
      } catch (err) {
        if (err instanceof ConnectError) {
          console.log(err.message)
        }
        throw err
      }
    },

    async editVlanById(id: number, vlanId: number) {
      try {
        return await subnetService.editSubnetVlan({
          id: id,
          vlanId: vlanId,
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
        return await subnetService.editSubnetDescription({
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
    getItem: state => state.item,

    getChart: state => state.chart
  }
})
