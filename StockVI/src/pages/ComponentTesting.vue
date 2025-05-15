<template>
    <StockTable @stock-selected="stock => selectedStock = stock"></StockTable>
     <StockChart v-if="selectedStock" :selectedStock="selectedStock"></StockChart>
  </template>
  
  <script setup>
    import NavBar from '@/components/NavBar.vue'
    import StockTable from '@/components/StockTable.vue'
    import StockChart from '@/components/StockChart.vue'
    import BrokerRating from '@/components/BrokerRating.vue'

    const selectedStock = ref<RatingsHistoric | null>(null);
  </script>
  