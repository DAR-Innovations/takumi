import { ADMIN_ROUTES_CONFIG, type ADMIN_ROUTES_REGISTRY } from '@/core/routes/admin.routes'
import { AUTH_ROUTES_CONFIG, type AUTH_ROUTES_REGISTRY } from '@/core/routes/auth.routes'
import { LANDING_ROUTES_CONFIG, type LANDING_ROUTES_REGISTRY } from '@/core/routes/landing.routes'
import { PRIMARY_ROUTES_CONFIG, type PRIMARY_ROUTES_REGISTRY } from '@/core/routes/primary.routes'
import { PROFILE_ROUTES_CONFIG, type PROFILE_ROUTES_REGISTRY } from '@/core/routes/profile.routes'
import type { RouteRecordRaw } from 'vue-router'
import { ERRORS_ROUTES_CONFIG, type ERRORS_ROUTES_REGISTRY } from '../routes/errors.routes'
import type { APP_LAYOUTS } from '../types/routes.types'

type ROUTES_REGISTRY =
	| ADMIN_ROUTES_REGISTRY
	| PRIMARY_ROUTES_REGISTRY
	| PROFILE_ROUTES_REGISTRY
	| AUTH_ROUTES_REGISTRY
	| LANDING_ROUTES_REGISTRY
	| ERRORS_ROUTES_REGISTRY

export type ExtendedRouteRecord = RouteRecordRaw & {
	meta: {
		layout?: APP_LAYOUTS
		title: string
		pageTitle?: string
		requiresAuth?: boolean
	}
	name: ROUTES_REGISTRY
	children?: ExtendedRouteRecord[]
}

export const ROUTES_CONFIG: ExtendedRouteRecord[] = [
	...ADMIN_ROUTES_CONFIG,
	...PRIMARY_ROUTES_CONFIG,
	...PROFILE_ROUTES_CONFIG,
	...AUTH_ROUTES_CONFIG,
	...LANDING_ROUTES_CONFIG,

	// Always in the bottom
	...ERRORS_ROUTES_CONFIG,
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
