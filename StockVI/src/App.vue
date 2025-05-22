<template>


  <div class="flex flex-col h-screen">
    <NavBar @search-results="handleSearchResults" />
    <div class="flex flex-1">
      <div :class="selectedStock ? 'w-3/4' : 'w-full'">
        <StockTable :ratingsData="ratingsData" @stock-selected="stock => selectedStock = stock"></StockTable>
      </div>
      
      <div v-if="selectedStock" class="w-1/4 p-4 space-y-4 relative">
        <button
          class="absolute top-4 right-7 text-gray-500 hover:text-black text-xl"
          @click="selectedStock = null"
          aria-label="Close panel"
          title="Close"
        >
          &times;
        </button>
        <StockChart :selectedStock="selectedStock"></StockChart>
        <BrokerRating :ratingsData="filteredRatingsData" :selectedStock="selectedStock"></BrokerRating>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
  import { computed, onMounted, ref } from 'vue'
  import NavBar from '@/components/NavBar.vue'
  import StockTable from '@/components/StockTable.vue'
  import StockChart from '@/components/StockChart.vue'
  import BrokerRating from '@/components/BrokerRating.vue'
  import type { RatingsHistoric } from './ports/RatingHistoric'
  import { GetRatingsHistoric } from './services/RatingsHistoric'

  const selectedStock = ref<RatingsHistoric | null>(null);
  const ratingsData = ref<RatingsHistoric[]>([])
  const filteredRatingsData = computed(() => {
    if (!selectedStock.value) return []
    return ratingsData.value.filter(r => r.stock === selectedStock.value?.stock)
  })


  function handleSearchResults(results: RatingsHistoric[]) {
    console.log(results)
    ratingsData.value = results
  }

  onMounted(async () => {
    ratingsData.value = await GetRatingsHistoric();
  });
</script>
