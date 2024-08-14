import { ROUTES } from './routes.config'

interface MenuItem {
	name: string
	routeName: string
	icon: string
}

export const tabBarRoutes: MenuItem[] = [
	{ name: 'Dashboard', routeName: ROUTES.HOME, icon: 'pi pi-home' },
	{ name: 'Analytics', routeName: ROUTES.ANALYTICS, icon: 'pi pi-chart-bar' },
	{ name: 'Create', routeName: ROUTES.CREATE, icon: 'pi pi-plus' },
	{ name: 'Events', routeName: ROUTES.EVENTS, icon: 'pi pi-compass' },
	{ name: 'Profile', routeName: ROUTES.PROFILE, icon: 'pi pi-user' },
]
