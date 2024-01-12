import "element-plus/dist/index.css"

import ElementPlus from "element-plus"

import './assets/main.css'

import { createApp } from 'vue'
import router from "./router";
import App from './App.vue'

const app = createApp(App)

app.use(router)
app.use(ElementPlus)



app.mount('#app')