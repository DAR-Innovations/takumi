import { createRouter, createWebHistory } from 'vue-router'

import { ROUTES_CONFIG } from './core/config/routes.config'

export const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	scrollBehavior: function (to, _from, savedPosition) {
		if (savedPosition) {
			return savedPosition
		}
		if (to.hash) {
			return { el: to.hash, behavior: 'smooth' }
		} else {
			window.scrollTo(0, 0)
		}
	},
	routes: ROUTES_CONFIG,
})

export const DEFAULT_TITLE = 'Takumi | Habit Tracking'
const TITLE_TEMPLATE = (title: string) => `${title} | Takumi`

router.beforeEach((to, _from, next) => {
	const metaTitle = to.meta.title ? TITLE_TEMPLATE(to.meta.title as string) : DEFAULT_TITLE
	document.title = metaTitle

	next()
})