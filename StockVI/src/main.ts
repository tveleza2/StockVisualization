import './assets/main.css'

import { createApp } from 'vue'
import { createPinia } from 'pinia'
import App from './App.vue'
import PrimeVue from 'primevue/config';
import ComponentTesting from '@/pages/ComponentTesting.vue';

// const app = createApp(ComponentTesting)
const app = createApp(App)


app.use(createPinia())
app.use(PrimeVue)

app.mount('#app')
