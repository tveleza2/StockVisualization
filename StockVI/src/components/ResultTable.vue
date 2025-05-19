<template>
    <div class="border p-4">
        <DataTable :value="ratings" tableStyle="min-width: 50rem">
            <Column :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }" v-for="col of columns" :key="col.field" :field="col.field" :header="col.header " ></Column>
            <!-- <Column field="stock" header="Stock" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="company" header="Company" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="broker" header="Broker" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="action" header="Action" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column>
            <Column field="rating" header="Rating" :headerStyle="{ fontWeight: 'bold', fontSize: '1.125rem' }"></Column> -->
        </DataTable>

    </div>
</template>


<script setup lang="ts">

    import DataTable from 'primevue/datatable';
    import Column from 'primevue/column';
    import ColumnGroup from 'primevue/columngroup';   // optional
    import Row from 'primevue/row';                   // optional


    import { watch,ref } from 'vue';
    import Chart from 'primevue/chart';
    import type { ChartData } from 'chart.js';
    import type { RatingsHistoric } from '@/ports/RatingHistoric';
    import type { StockPrice } from '@/ports/StockPrice';
    import {GetStockPrices} from '@/services/StockPrice';

    const chartData = ref<ChartData<'line'>>();
    const columns = [
        {field:"broker",header:"Broker"},
        {field:"rating",header:"Rating"}, 
        {field:"target",header:"Target"}
    ]

    const props = defineProps<{
        selectedStock: RatingsHistoric;
    }>();

    const ratings = ref<RatingsHistoric[]>([])

    
</script>
  