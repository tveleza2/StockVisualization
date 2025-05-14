<template>
    <div class="overflow-x-auto p-4">
      <table class="w-full table-auto border-collapse">
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
      </table>
    </div>
  </template>
  
  <script setup>
  const data = [
    { stock: 'AAPL', company: 'Apple Inc.', broker: 'JP Morgan', action: 'Buy', rating: '4.5' },
    // Add more items as needed
  ]
  </script>
  