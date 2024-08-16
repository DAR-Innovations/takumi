<script setup lang="ts">
import PageTitle from '@/core/components/ui/page-title/page-title.vue'
import AppAdminSidebar from '@/core/layouts/admin/app-admin-sidebar.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentPageTitle = computed(() => route.meta.pageTitle ?? route.meta.title)
</script>

<template>
	<div class="flex w-full h-screen overflow-hidden">
		<aside class="flex flex-col h-full border-r border-border ">
			<AppAdminSidebar />
		</aside>

		<main class="relative flex-1 w-full h-full overflow-y-auto">
			<header class="w-full flex justify-between items-center p-6">
				<slot name="header">
					<PageTitle>{{ currentPageTitle }}</PageTitle>

					<div class="flex items-center gap-2">
						<slot name="right-side" />
					</div>
				</slot>
			</header>

			<div class="p-6 pt-2 h-full">
				<slot></slot>
			</div>
		</main>
	</div>
</template>

<style scoped></style>
