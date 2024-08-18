import { createRouter, createWebHistory } from 'vue-router'

import { ROUTES, ROUTES_CONFIG } from './core/config/routes.config'
import { DEFAULT_TITLE, TITLE_TEMPLATE } from './core/constants/seo.constant'
import { authService } from './modules/auth/service/auth.service'
import { useAuthStore } from './modules/auth/stores/auth.store'

export const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	scrollBehavior() {
		return { top: 0 }
	},
	routes: ROUTES_CONFIG,
})

router.beforeEach(async (to, _from, next) => {
	const { isLoggedIn, setCurrentUser } = useAuthStore()

	if (to.meta?.requiresAuth && !isLoggedIn) {
		try {
			const currentUser = await authService.getCurrentUser()

			if (!currentUser) {
				return next({ name: ROUTES.LANDING })
			}

			setCurrentUser(currentUser)
		} catch (error) {
			return next({ name: ROUTES.INTERNAL_ERROR })
		}
	}

	const metaTitle = to.meta?.title ? TITLE_TEMPLATE(to.meta.title as string) : DEFAULT_TITLE
	document.title = metaTitle

	return next()
})
