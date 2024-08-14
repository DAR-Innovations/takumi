import { ROUTES } from './routes.config'

interface MenuItem {
	name: string
	routeName: string
	icon: string
}

export const sidebarItems: MenuItem[] = [
	{ name: 'Dashboard', routeName: ROUTES.HOME, icon: 'pi pi-home' },
	{ name: 'Analytics', routeName: ROUTES.ANALYTICS, icon: 'pi pi-chart-bar' },
	{ name: 'Settings', routeName: ROUTES.SETTINGS, icon: 'pi pi-cog' },
]
