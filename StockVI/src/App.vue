<template>
  
  <div class="p-8 text-2xl text-green-600">
    If you see this, App.vue is rendering!
  </div>


  <!-- <div class="flex flex-col h-screen">
    <NavBar/>
    <div class="flex flex-1">
      <div :class="selectedStock ? 'w-3/4' : 'w-full'">
        <StockTable @stock-selected="stock => selectedStock = stock"></StockTable>
      </div>
      <div v-if="selectedStock" class="w-1/4 p-4 space-y-4">
        <StockChart v-if="selectedStock" :selectedStock="selectedStock"></StockChart>
        <BrokerRating v-if="selectedStock" :selectedStock="selectedStock"></BrokerRating>
      </div>
    </div>
  </div> -->
</template>

<script setup lang="ts">
  import { ref } from 'vue'
  import NavBar from '@/components/NavBar.vue'
  import StockTable from '@/components/StockTable.vue'
  import StockChart from '@/components/StockChart.vue'
  import BrokerRating from '@/components/BrokerRating.vue'
  import type { RatingsHistoric } from './ports/RatingHistoric'

  const selectedStock = ref<RatingsHistoric | null>(null);
</script>
