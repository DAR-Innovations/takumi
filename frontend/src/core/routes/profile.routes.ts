import { type ExtendedRouteRecord } from '@/core/config/routes.config'
import { APP_LAYOUTS } from '../types/routes.types'

export type PROFILE_ROUTES_REGISTRY = 'profile' | 'profile-achievements'

export const PROFILE_ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/profile',
		name: 'profile',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Profile' },
		component: () => import('@/modules/profile/pages/profile-page.vue'),
	},
	{
		path: '/profile/achievements',
		name: 'profile-achievements',
		meta: { layout: APP_LAYOUTS.GO_BACK_LAYOUT, title: 'Achievements' },
		component: () => import('@/modules/profile/pages/achievements-page.vue'),
	},
]
