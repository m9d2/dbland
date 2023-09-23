import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import App from './App.vue'
import router from './plugins/router'
import ContextMenu from '@imengyu/vue3-context-menu'
import i18n from './plugins/i18n'
import './style/index.scss'

const app = createApp(App)

app.use(createPinia())
app.use(i18n)
app.use(router)
app.use(ElementPlus, {
    size: 'extra-small'
})
app.use(ContextMenu)

app.mount('#app')