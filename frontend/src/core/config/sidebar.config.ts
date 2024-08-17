import { ROUTES } from './routes.config'

interface MenuItem {
	name: string
	routeName: string
	icon: string
}

export const primarySidebarItems: MenuItem[] = [
	{ name: 'Create', routeName: ROUTES.HABITS_CREATE, icon: 'pi pi-plus' },
	{ name: 'Dashboard', routeName: ROUTES.HOME, icon: 'pi pi-home' },
	{ name: 'Analytics', routeName: ROUTES.ANALYTICS, icon: 'pi pi-chart-bar' },
	{ name: 'Events', routeName: ROUTES.EVENTS, icon: 'pi pi-compass' },
	{ name: 'Profile', routeName: ROUTES.PROFILE, icon: 'pi pi-user' },
	{ name: 'Admin', routeName: ROUTES.ADMIN_DASHBOARD, icon: 'pi pi-cog' },
]

export const adminSidebarItems: MenuItem[] = [
	{ name: 'Dashboard', routeName: ROUTES.ADMIN_DASHBOARD, icon: 'pi pi-th-large' },
	{ name: 'Users', routeName: ROUTES.ADMIN_USERS, icon: 'pi pi-users' },
	{ name: 'Challenges', routeName: ROUTES.ADMIN_CHALLENGES, icon: 'pi pi-compass' },
	{ name: 'Articles', routeName: ROUTES.ADMIN_ARTICLES, icon: 'pi pi-book' },
]
