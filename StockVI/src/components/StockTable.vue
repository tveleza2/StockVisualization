<template>
    <div class="overflow-x-auto p-4">
        <DataTable :value="ratingsData" tableStyle="min-width: 50rem" @row-click="onRowClick">
            <Column sortable :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }" v-for="col of columns" :key="col.field" :field="col.field" :header="col.header " ></Column>
        </DataTable>
    </div>
  </template>
  
  <script setup lang="ts">
    import DataTable from 'primevue/datatable';
    import Column from 'primevue/column';
    import ColumnGroup from 'primevue/columngroup';   // optional
    import Row from 'primevue/row';                   // optional

    import { ref, onMounted } from 'vue';
    import type { RatingsHistoric } from '@/ports/RatingHistoric';
    import {GetRatingsHistoric} from '@/services/RatingsHistoric'
    // import { ProductService } from '@/service/ProductService';

    const props = defineProps<{
        ratingsData: RatingsHistoric[];
    }>();
    function onRowClick(event:any) {
        selectedStock.value = event.data.stock; // PrimeVue gives you the row object in event.data
        emit('stock-selected', event.data);
        console.log("Selected Stock",selectedStock.value);
    }

    const emit = defineEmits<{
        (e: 'stock-selected', stock: RatingsHistoric): void;
    }>();

    const columns = [
        {field:"stock",header:"Stock"},
        {field:"company",header:"Company"}, 
        {field:"broker",header:"Broker"},
        {field:"action",header:"Action"},
        {field:"rating",header:"Rating"},

    ]

    const selectedStock = ref(null);
  </script>
  