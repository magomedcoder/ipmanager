<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { Chart, registerables } from 'chart.js'
import type { ChartItem } from 'chart.js'

interface IChart {
  type: string
  label: string
  data: number[]
}

const props = defineProps<{
  labels: string[]
  data: IChart[]
}>()

const chartRef = ref<ChartItem>()
const chart = ref()

const init = () => {
  const ctx = chartRef.value instanceof HTMLCanvasElement
    ? chartRef.value.getContext('2d')
    : chartRef.value

  if (ctx){
    Chart.register(...registerables)
    chart.value = new Chart(ctx, {
      type: 'doughnut',
      data: {
        labels: props.labels,
        datasets: props.data
      },
      // options: {
      //   plugins: {
      //     legend: {
      //       display: false
      //     }
      //   }
      // }
    })
  }
}

onMounted(() => {
  init()
})
</script>

<template>
  <div class="rounded-lg shadow-md">
    <p class="text-center text-gray-600">
      <canvas
        ref="chartRef"
        width="200"
        height="100"
      />
    </p>
  </div>
</template>

<style scoped></style>
