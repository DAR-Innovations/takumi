<script setup lang="ts">
import ThemeToggler from '@/core/components/themes/theme-toggler.vue'
import PageTitle from '@/core/components/ui/page-title/page-title.vue'
import { ROUTES } from '@/core/config/routes.config'
import AppPrimarySidebar from '@/core/layouts/primary/app-primary-sidebar.vue'
import AppPrimaryTabs from '@/core/layouts/primary/app-primary-tabs.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentPageTitle = computed(() => route.meta.pageTitle ?? route.meta.title)
</script>

<template>
	<div class="w-full h-screen sm:flex sm:container sm:items-center sm:justify-center">
		<div
			class="flex w-full h-full sm:h-3/4 sm:border sm:border-border sm:rounded-3xl sm:backdrop-blur-sm	sm:p-2"
		>
			<aside class="sm:flex flex-col hidden p-10 pr-4 py-6 w-fit h-full">
				<AppPrimarySidebar />
			</aside>

			<main class="relative flex-1 w-full h-full overflow-y-auto">
				<header class="w-full flex justify-between items-center p-4 sm:p-6">
					<slot name="header">
						<PageTitle>{{ currentPageTitle }}</PageTitle>

						<div class="flex items-center gap-2">
							<slot name="right-side" />

							<router-link :to="{name:ROUTES.NOTIFICATIONS}">
								<button class="bg-card w-11 h-11 p-2 rounded-lg flex items-center justify-center">
									<i class="text-lg pi pi-bell"></i>
								</button>
							</router-link>

							<button class="bg-card w-11 h-11 p-2 rounded-lg flex items-center justify-center">
								<ThemeToggler />
							</button>
						</div>
					</slot>
				</header>

				<div class="relative p-4 sm:p-6 mb-[300px] h-full">
					<slot></slot>
				</div>

				<slot name="tabbar">
					<div
						class="right-4 bottom-4 left-4 fixed sm:hidden bg-card px-8 py-4 rounded-3xl border border-border"
					>
						<AppPrimaryTabs />
					</div>
				</slot>
			</main>
		</div>
	</div>
</template>

<style scoped></style>
