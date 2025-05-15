<template>
    <div class="border p-4 h-48">
      <h1 class="text-lg mb-2">{{ selectedStock.value }}</h1>
      <h2 class="text-md mb-2">{{ selectedStock.value }}</h2>
      <Chart type="line" :data="chartData" />

    </div>
  </template>


<script setup lang="ts">
    import { watch } from 'vue';
    import Chart from 'primevue/chart';
    import {}

    watch(props.selectedStock, async (stock) => {
        if (stock) {
            
            // const res = await fetch(`/api/stocks/${stock.id}/chart`);
            const res = await StockPrice;
            const data = await res.json();
            chartData.value = formatForChart(data);
        }
    });

    const props = defineProps<{
        selectedStock: RatingsHistoric;
    }>();



    function formatForChart(rawData: any) {
    // Example format
    return {
        labels: rawData.map((item: any) => item.date),
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
  