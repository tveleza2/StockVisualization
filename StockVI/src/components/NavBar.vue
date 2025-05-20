<template>
  <div class="flex justify-between items-center p-4 bg-[#C4E7D4]">
    <img src="@/assets/logo.png" class="w-1/30" />
    // ...existing code...
  <input
  v-model="search"
  placeholder="🔍 Search a Stock"
  class="border rounded px-2 py-1 w-1/3 text-center"
  @keyup.enter="handleSearch"><input>
    <div class="text-sm text-gray-600">{{ currentTime }} ({{ timeZoneName }})</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted, onUnmounted } from 'vue'
import { searchRatingsHistoric } from '@/services/SearchByStock'
import type { RatingsHistoric } from '@/ports/RatingHistoric'

const emit = defineEmits<{
  (e: 'search-results', results: RatingsHistoric[]): void
}>()

async function handleSearch() {
  if (!search.value.trim()) return
  try {
    const results = await searchRatingsHistoric(search.value)
    emit('search-results', results)
  } catch (err) {
    // Optionally handle error (e.g., show a message)
    emit('search-results', [])
  }
}



const search = ref('')
const currentTime = ref('')
const timeZoneName = ref('')

// Function to format the local time
function updateClock() {
  const now = new Date()
  currentTime.value = now.toLocaleTimeString([], {
    hour: '2-digit',
    minute: '2-digit',
    hour12: true
  })
}

// Get the user's local timezone (e.g., "Pacific Daylight Time")
function detectTimeZone() {
  const now = new Date()
  timeZoneName.value = Intl.DateTimeFormat(undefined, { timeZoneName: 'short' })
    .formatToParts(now)
    .find(part => part.type === 'timeZoneName')?.value || ''
}

let intervalId: number
onMounted(() => {
  detectTimeZone()
  updateClock()
  intervalId = window.setInterval(updateClock, 1000)
})

onUnmounted(() => {
  clearInterval(intervalId)
})
</script>
