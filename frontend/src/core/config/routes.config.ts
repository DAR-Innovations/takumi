import type { RouteRecordRaw } from 'vue-router'

export enum APP_LAYOUTS {
	PRIMARY_LAYOUT = 'PRIMARY_LAYOUT',
	GO_BACK_LAYOUT = 'GO_BACK_LAYOUT',
}

type ROUTES_REGISTRY =
	| 'home'
	| 'analytics'
	| 'events'
	| 'profile'
	| 'create'
	| 'profile-achievements'
	| 'notifications'

type ExtendedRouteRecord = RouteRecordRaw & {
	meta: {
		layout?: APP_LAYOUTS
		title: string
		pageTitle?: string
	}
	name: ROUTES_REGISTRY
	children?: ExtendedRouteRecord[]
}

export const ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/',
		name: 'home',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Home' },
		component: () => import('@/modules/home/pages/home-page.vue'),
	},
	{
		path: '/analytics',
		name: 'analytics',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Analytics' },
		component: () => import('@/modules/analytics/pages/analytics-page.vue'),
	},
	{
		path: '/notifications',
		name: 'notifications',
		meta: { layout: APP_LAYOUTS.GO_BACK_LAYOUT, title: 'Notifications' },
		component: () => import('@/modules/notifications/pages/notifications-page.vue'),
	},
	{
		path: '/events',
		name: 'events',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Events' },
		component: () => import('@/modules/events/pages/events-page.vue'),
},
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
	{
		path: '/create',
		name: 'create',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Create' },
		component: () => import('@/modules/home/pages/home-page.vue'),
	},
]

type RouteNames = (typeof ROUTES_CONFIG)[number]['name']

function transformRouteName(name: string | symbol): string {
	return name.toString().toUpperCase().replace(/-/g, '_')
}

type TransformNames<T extends string> = T extends `${infer F}-${infer R}`
	? `${Uppercase<F>}_${TransformNames<R>}`
	: Uppercase<T>

type TransformedRouteNames = {
	[K in RouteNames as TransformNames<K & string>]: K
}

export const ROUTES: TransformedRouteNames = {} as TransformedRouteNames

ROUTES_CONFIG.forEach(route => {
	if (route.name) {
		const transformedName = transformRouteName(route.name) as keyof TransformedRouteNames
		ROUTES[transformedName] = route.name as never
	}
})
