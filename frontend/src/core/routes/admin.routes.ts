import { type ExtendedRouteRecord } from '@/core/config/routes.config'
import { APP_LAYOUTS } from '../types/routes.types'

export type ADMIN_ROUTES_REGISTRY =
	| 'admin-dashboard'
	| 'admin-users'
	| 'admin-challenges'
	| 'admin-articles'

export const ADMIN_ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/admin',
		name: 'admin-dashboard',
		meta: { layout: APP_LAYOUTS.ADMIN_LAYOUT, title: 'Admin Dashboard' },
		component: () => import('@/modules/admin/dashboard/pages/admin-dashboard-page.vue'),
	},
	{
		path: '/admin/users',
		name: 'admin-users',
		meta: { layout: APP_LAYOUTS.ADMIN_LAYOUT, title: 'Admin Users' },
		component: () => import('@/modules/admin/users/pages/admin-users-page.vue'),
	},
	{
		path: '/admin/challenges',
		name: 'admin-challenges',
		meta: { layout: APP_LAYOUTS.ADMIN_LAYOUT, title: 'Admin Challenges' },
		component: () => import('@/modules/admin/challenges/pages/admin-challenges-page.vue'),
	},
	{
		path: '/admin/articles',
		name: 'admin-articles',
		meta: { layout: APP_LAYOUTS.ADMIN_LAYOUT, title: 'Admin Challenges' },
		component: () => import('@/modules/admin/articles/pages/admin-articles-page.vue'),
	},
]
