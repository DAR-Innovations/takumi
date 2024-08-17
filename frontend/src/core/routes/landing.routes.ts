import { type ExtendedRouteRecord } from '@/core/config/routes.config'
import LandingPage from '@/modules/landing/pages/landing-page.vue'

export type LANDING_ROUTES_REGISTRY = 'landing'

export const LANDING_ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/landing',
		name: 'landing',
		meta: { title: 'Welcome' },
		component: LandingPage,
	},
]
