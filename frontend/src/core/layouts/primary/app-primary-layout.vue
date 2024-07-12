<script setup lang="ts">
import AppPrimarySidebar from '@/core/layouts/primary/app-primary-sidebar.vue'
import AppPrimaryTabs from '@/core/layouts/primary/app-primary-tabs.vue'
import { computed } from 'vue'

import { useRoute } from 'vue-router'
const route = useRoute()
const currentPageTitle = computed(() => route.meta.pageTitle ?? route.meta.title)
</script>

<template>
	<div class="relative flex w-full h-screen overflow-hidden">
		<aside class="sm:flex flex-col hidden p-6 border-r border-border w-full max-w-[250px] h-full">
			<AppPrimarySidebar />
		</aside>

		<main class="relative flex-1 w-full h-full overflow-y-auto">
			<div class="flex justify-between items-center p-4 sm:p-7">
				<h1 class="font-medium text-4xl">{{currentPageTitle}}</h1>

				<div class="flex items-center gap-2">
					<button class="bg-[#242424] px-3 py-2 rounded-xl">
						<i class="text-xl pi pi-bell"></i>
					</button>
				</div>
			</div>

			<div class="mb-[160px] p-4 sm:p-7 h-full">
				<slot></slot>
			</div>
		</main>
	</div>

	<div class="block right-4 bottom-4 left-4 fixed sm:hidden bg-[#242424] px-8 py-4 rounded-3xl">
		<AppPrimaryTabs />
	</div>
</template>

<style scoped></style>
