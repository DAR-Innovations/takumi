import './styles.css'

import { createPinia } from 'pinia'
import { createApp } from 'vue'

import { VueQueryPlugin } from '@tanstack/vue-query'
import App from './App.vue'
import directives from './core/directives'
import { router } from './router'

const app = createApp(App)

Object.entries(directives).forEach(([key, value]) => {
	app.directive(key, value)
})

app.use(VueQueryPlugin)
app.use(createPinia())
app.use(router)

app.mount('#app')
