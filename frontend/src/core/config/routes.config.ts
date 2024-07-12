import HomePage from '@/modules/home/pages/HomePage.vue'

import type { RouteRecordRaw } from 'vue-router'

export enum APP_LAYOUTS {
	PRIMARY_LAYOUT = 'PRIMARY_LAYOUT',
}

type ROUTES_REGISTRY = 'home' | 'activities' | 'settings' | 'profile' | 'create'

type ExtendedRouteRecord = RouteRecordRaw & {
	meta: {
		layout?: APP_LAYOUTS
		title: string
		pageTitle?: string
	}
	name: ROUTES_REGISTRY
}

export const ROUTES_CONFIG: ExtendedRouteRecord[] = [
	{
		path: '/',
		name: 'home',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Home' },
		component: HomePage,
	},
	{
		path: '/activities',
		name: 'activities',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Activities' },
		component: HomePage,
	},
	{
		path: '/settings',
		name: 'settings',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Settings' },
		component: HomePage,
	},
	{
		path: '/profile',
		name: 'profile',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Profile' },
		component: HomePage,
	},
	{
		path: '/create',
		name: 'create',
		meta: { layout: APP_LAYOUTS.PRIMARY_LAYOUT, title: 'Create' },
		component: HomePage,
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
