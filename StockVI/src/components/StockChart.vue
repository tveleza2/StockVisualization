<template>
    <div class="border p-4">
      <h1 class="text-xl mb-2 text-center">{{ selectedStock.stock }}</h1>
      <h2 class="text-base mb-2 text-center">{{ selectedStock.company }}</h2>
      <Chart type="line" :data="chartData"/>

    </div>
  </template>


<script setup lang="ts">
    import { watch,ref } from 'vue';
    import Chart from 'primevue/chart';
    import type { ChartData } from 'chart.js';
    import type { RatingsHistoric } from '@/ports/RatingHistoric';
    import type { StockPrice } from '@/ports/StockPrice';
    import {GetStockPrices} from '@/services/StockPrice';

    const chartData = ref<ChartData<'line'>>();
    

    const props = defineProps<{
        selectedStock: RatingsHistoric;
    }>();

    watch(props.selectedStock, async (stock) => {
        if (stock) {
            
            // const res = await fetch(`/api/stocks/${stock.id}/chart`);
            // const data = await res.json();
            const data = await GetStockPrices();
            chartData.value = formatForChart(data);
        }
    },{ immediate: true });

    function formatForChart(rawData: any) {
    // Example format
    return {
        labels: rawData.map((item:StockPrice) =>
                item.time instanceof Date ? item.time.toISOString().split('T')[0] : item.time
                ),
        datasets: [
        {
            label: 'Price',
            data: rawData.map((item: any) => item.price),
            fill: false,
            borderColor: '#42A5F5',
        },
        ],
    };
    }
</script>
  