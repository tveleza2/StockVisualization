<template>
    <div class="flex flex-col h-screen">
      <NavBar/>
      <div class="flex flex-1">
        <div class="w-3/4">
          <StockTable />
        </div>
        <div class="w-1/4 p-4 space-y-4">
          <StockChart />
          <BrokerRating />
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
    import NavBar from './components/NavBar.vue'
    import StockTable from './components/StockTable.vue'
    import StockChart from './components/StockChart.vue'
    import BrokerRating from './components/BrokerRating.vue'
    
  
  </script>
  