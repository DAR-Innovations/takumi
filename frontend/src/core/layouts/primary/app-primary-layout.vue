<script setup lang="ts">
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
	<div class="flex w-full h-screen overflow-hidden">
		<aside class="sm:flex flex-col hidden p-6 border-r border-border w-full max-w-[250px] h-full">
			<AppPrimarySidebar />
		</aside>

		<main class="relative flex-1 w-full h-full overflow-y-auto">
			<header class="w-full flex justify-between items-center p-4 sm:p-7">
				<slot name="header">
					<PageTitle>{{ currentPageTitle }}</PageTitle>

					<div class="flex items-center gap-2">
						<slot name="right-side" />

						<router-link :to="{name:ROUTES.NOTIFICATIONS}">
							<button class="bg-[#242424] px-3 py-2 rounded-xl">
								<i class="text-xl pi pi-bell"></i>
							</button>
						</router-link>
					</div>
				</slot>
			</header>

			<div class="p-4 sm:p-7 mb-[300px] h-full">
				<slot></slot>
			</div>

			<div class="right-4 bottom-4 left-4 fixed sm:hidden bg-[#242424] px-8 py-4 rounded-3xl">
				<AppPrimaryTabs />
			</div>
		</main>
	</div>
</template>

<style scoped></style>
