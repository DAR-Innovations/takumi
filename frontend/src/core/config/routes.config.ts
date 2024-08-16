import { ADMIN_ROUTES_CONFIG, type ADMIN_ROUTES_REGISTRY } from '@/core/routes/admin.routes'
import { PRIMARY_ROUTES_CONFIG, type PRIMARY_ROUTES_REGISTRY } from '@/core/routes/primary.routes'
import { PROFILE_ROUTES_CONFIG, type PROFILE_ROUTES_REGISTRY } from '@/core/routes/profile.routes'
import type { RouteRecordRaw } from 'vue-router'
import type { APP_LAYOUTS } from '../types/routes.types'

type ROUTES_REGISTRY = ADMIN_ROUTES_REGISTRY | PRIMARY_ROUTES_REGISTRY | PROFILE_ROUTES_REGISTRY

export type ExtendedRouteRecord = RouteRecordRaw & {
	meta: {
		layout?: APP_LAYOUTS
		title: string
		pageTitle?: string
	}
	name: ROUTES_REGISTRY
	children?: ExtendedRouteRecord[]
}

export const ROUTES_CONFIG: ExtendedRouteRecord[] = [
	...ADMIN_ROUTES_CONFIG,
	...PRIMARY_ROUTES_CONFIG,
	...PROFILE_ROUTES_CONFIG,
]

export const ROUTES: TransformedRouteNames = {} as TransformedRouteNames

ROUTES_CONFIG.forEach(route => {
	if (route.name) {
		const transformedName = route.name
			.toString()
			.toUpperCase()
			.replace(/-/g, '_') as keyof TransformedRouteNames
		ROUTES[transformedName] = route.name as never
	}
})

type RouteNames = (typeof ROUTES_CONFIG)[number]['name']

type TransformNames<T extends string> = T extends `${infer F}-${infer R}`
	? `${Uppercase<F>}_${TransformNames<R>}`
	: Uppercase<T>

type TransformedRouteNames = {
	[K in RouteNames as TransformNames<K & string>]: K
}
