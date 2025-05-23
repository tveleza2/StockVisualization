<template>
    <div class="overflow-x-auto p-4 text-center">
        <DataTable :value="ratingsData" tableStyle="min-width: 50rem" @row-click="onRowClick" :rowClass="rowClass" >
            <Column field="stock" header="Stock" sortable :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem'}">
                <template #body="slotProps">
                <span>
                    <span v-if="slotProps.data.stock === bestStockId"><i class="pi pi-star text-yellow-500 mr-1"></i></span>
                    {{ slotProps.data.stock }}
                </span>
                </template>
            </Column>
            <Column sortable :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }" v-for="col of columns" :key="col.field" :field="col.field" :header="col.header " class="text-center"></Column>
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
    import { GetStocks } from '@/services/GetStocks';
    // import { ProductService } from '@/service/ProductService';

    const props = defineProps<{
        ratingsData: RatingsHistoric[];
    }>();
    function onRowClick(event:any) {
        selectedStock.value = event.data.stock; // PrimeVue gives you the row object in event.data
        emit('stock-selected', event.data);
        console.log("Selected Stock",selectedStock.value);
    }
    function rowClass(rowData: RatingsHistoric) {
        const stockId = rowData.stock;
        return stockId === bestStockId.value
        ? 'font-semibold text-center'
        : 'text-center';
    }
    const bestStockId = ref<string | null>(null);

    onMounted(async () => {
        try {
            const stocks = await GetStocks();
            const bestStock = stocks.reduce((max, stock) =>
            stock.score > max.score ? stock : max
            );
            bestStockId.value = bestStock.id;
        } catch (error) {
            console.error('Error fetching best stock:', error);
        }
    });

    

    const emit = defineEmits<{
        (e: 'stock-selected', stock: RatingsHistoric): void;
    }>();

    const columns = [
        {field:"company",header:"Company"}, 
        {field:"broker",header:"Broker"},
        {field:"action",header:"Action"},
        {field:"rating",header:"Rating"},

    ]
    
    const selectedStock = ref(null);
  </script>
  