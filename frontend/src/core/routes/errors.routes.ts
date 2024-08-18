import { type ExtendedRouteRecord } from '@/core/config/routes.config'

export type ERRORS_ROUTES_REGISTRY = 'not-found' | 'internal-error'

export const ERRORS_ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/:pathMatch(.*)*',
		name: 'not-found',
		meta: { title: 'Not found' },
		component: () => import('@/modules/errors/pages/not-found-page.vue'),
	},
	{
		path: '/:pathMatch(.*)*',
		name: 'internal-error',
		meta: { title: 'Internal server error' },
		component: () => import('@/modules/errors/pages/internal-error-page.vue'),
	},
]
