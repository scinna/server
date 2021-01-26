import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from "@vue/cli-service/lib/options";

const token = localStorage.getItem('scinna-token')
if (token)
    axios.defaults.headers.common['Authorization'] = "Bearer " + token

createApp(App).use(store).use(router).mount('#app')
