
<template>
    <div class="card">
        <DataTable :value="products" tableStyle="min-width: 50rem">
            <Column v-for="col of columns" :key="col.field" :field="col.field" :header="col.header"></Column>
        </DataTable>
    </div>
</template>

<script setup>
    import { ref, onMounted } from 'vue';
    import { ProductService } from '@/service/ProductService';

    onMounted(() => {
        ProductService.getProductsMini().then((data) => (products.value = data));
    });

    const products = ref();
    const columns = [
        { field: 'code', header: 'Code' },
        { field: 'name', header: 'Name' },
        { field: 'category', header: 'Category' },
        { field: 'quantity', header: 'Quantity' }
    ];

</script>
