import { type ExtendedRouteRecord } from '@/core/config/routes.config'

export type AUTH_ROUTES_REGISTRY = 'auth-signup' | 'auth-login'

export const AUTH_ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/auth/signup',
		name: 'auth-signup',
		meta: { title: 'Signup' },
		component: () => import('@/modules/auth/pages/signup-page.vue'),
	},
	{
		path: '/auth/login',
		name: 'auth-login',
		meta: { title: 'Login' },
		component: () => import('@/modules/auth/pages/login-page.vue'),
	},
]
