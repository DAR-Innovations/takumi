import { type ExtendedRouteRecord } from '@/core/config/routes.config'
import { APP_LAYOUTS } from '../types/routes.types'

export type PRIMARY_ROUTES_REGISTRY =
	| 'home'
	| 'analytics'
	| 'events'
	| 'habits-create'
	| 'notifications'

export const PRIMARY_ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/',
		name: 'home',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Home', requiresAuth: true },
		component: () => import('@/modules/home/pages/home-page.vue'),
	},
	{
		path: '/analytics',
		name: 'analytics',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Analytics', requiresAuth: true },
		component: () => import('@/modules/analytics/pages/analytics-page.vue'),
	},
	{
		path: '/notifications',
		name: 'notifications',
		meta: { layout: APP_LAYOUTS.GO_BACK_LAYOUT, title: 'Notifications', requiresAuth: true },
		component: () => import('@/modules/notifications/pages/notifications-page.vue'),
	},
	{
		path: '/events',
		name: 'events',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Events', requiresAuth: true },
		component: () => import('@/modules/events/pages/events-page.vue'),
	},

	{
		path: '/habits/create',
		name: 'habits-create',
		meta: {
			layout: APP_LAYOUTS.NO_TABBAR,
			title: 'Create habit',
			pageTitle: 'Add new habit',
			requiresAuth: true,
		},
		component: () => import('@/modules/habits/pages/habits-create-page.vue'),
	},
]
