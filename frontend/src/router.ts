import { createRouter, createWebHistory } from 'vue-router'

import { ROUTES_CONFIG } from './core/config/routes.config'

export const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: ROUTES_CONFIG,
})

export const DEFAULT_TITLE = 'Takumi | Habit Tracking'
const TITLE_TEMPLATE = (title: string) => `${title} | Takumi`

router.beforeEach((to, from, next) => {
	const metaTitle = to.meta.title ? TITLE_TEMPLATE(to.meta.title as string) : DEFAULT_TITLE
	document.title = metaTitle

	next()
})
