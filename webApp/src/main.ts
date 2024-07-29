import { createApp } from 'vue'
import App from './App.vue'

import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import "./main.css"

const pinia = createPinia()
const app = createApp(App)

app.use(pinia)
app.use(ElementPlus)


app.mount('#app')

