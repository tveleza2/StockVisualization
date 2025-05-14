<template>
    <div class="flex justify-between items-center p-4 bg-[#C4E7D4]">
      <img src="@/assets/logo.svg" class="w-1/20"></img>
      <input v-model="search" placeholder="🔍 Search" class="border rounded px-2 py-1 w-1/3 text-center" />
      <div class="text-sm text-gray-600">11:30 PT</div>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  const search = ref('')
  </script>
  