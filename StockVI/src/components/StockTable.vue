<template>
    <div class="overflow-x-auto p-4">
        <DataTable :value="ratings" tableStyle="min-width: 50rem" @row-click="onRowClick">
            <Column field="stock" header="Stock" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="company" header="Company" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="broker" header="Broker" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="action" header="Action" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="rating" header="Rating" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
        </DataTable>


      <!-- <table class="w-full table-auto border-collapse">
        <thead>
          <tr class="bg-gray-200">
            <th class="p-2 border">Stock</th>
            <th class="p-2 border">Company</th>
            <th class="p-2 border">Broker</th>
            <th class="p-2 border">Action</th>
            <th class="p-2 border">Rating</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="(item, index) in data" :key="index" class="text-center">
            <td class="border p-2">{{ item.stock }}</td>
            <td class="border p-2">{{ item.company }}</td>
            <td class="border p-2">{{ item.broker }}</td>
            <td class="border p-2">{{ item.action }}</td>
            <td class="border p-2">{{ item.rating }}</td>
          </tr>
        </tbody>
      </table> -->
    </div>
  </template>
  
  <script setup lang="ts">
    import DataTable from 'primevue/datatable';
    import Column from 'primevue/column';
    import ColumnGroup from 'primevue/columngroup';   // optional
    import Row from 'primevue/row';                   // optional

    import { ref, onMounted } from 'vue';

    import {GetRatingsHistoric} from '@/services/RatingsHistoric'
    // import { ProductService } from '@/service/ProductService';

    function onRowClick(event) {
        selectedStock.value = event.data.stock; // PrimeVue gives you the row object in event.data
        console.log("Selected Stock",selectedStock.value);
    }


    const ratings = ref<RatingsHistoric[]>([])
    const selectedStock = ref(null);

    

    onMounted(() => {ratings.value = GetRatingsHistoric()});
  </script>
  