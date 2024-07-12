import { ROUTES } from './routes.config'

interface MenuItem {
	name: string
	routeName: string
	icon: string
}

export const tabBarRoutes: MenuItem[] = [
	{ name: 'Dashboard', routeName: ROUTES.HOME, icon: 'pi pi-home' },
	{ name: 'Analytics', routeName: ROUTES.ACTIVITIES, icon: 'pi pi-chart-bar' },
	{ name: 'Create', routeName: ROUTES.CREATE, icon: 'pi pi-plus' },
	{ name: 'Settings', routeName: ROUTES.SETTINGS, icon: 'pi pi-calendar' },
	{ name: 'Profile', routeName: ROUTES.PROFILE, icon: 'pi pi-user' },
]
