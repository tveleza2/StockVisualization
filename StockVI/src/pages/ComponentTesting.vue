<template>
    <div class="flex flex-1">
      <div :class="selectedStock ? 'w-3/4' : 'w-full'">
        <StockTable @stock-selected="stock => selectedStock = stock"></StockTable>
      </div>
      <div v-if="selectedStock" class="w-1/4 p-4 space-y-4">
        <StockChart v-if="selectedStock" :selectedStock="selectedStock"></StockChart>
      </div>
    </div>
    
     
  </template>
  
  <script setup lang="ts">
    // import NavBar from '@/components/NavBar.vue'
    import StockTable from '@/components/StockTable.vue'
    import StockChart from '@/components/StockChart.vue'
    import { ref } from 'vue';
    import type { RatingsHistoric } from '@/ports/RatingHistoric';
    // import BrokerRating from '@/components/BrokerRating.vue'

    const selectedStock = ref<RatingsHistoric | null>(null);
  </script>
  