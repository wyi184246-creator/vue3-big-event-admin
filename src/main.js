import { createApp } from 'vue'
import App from './App.vue'
import pinia from './stores/index'
import router from './router/index'
import './assets/main.scss'
import '@vueup/vue-quill/dist/vue-quill.snow.css'

const app = createApp(App)
app.use(pinia)
app.use(router)
app.mount('#app')
