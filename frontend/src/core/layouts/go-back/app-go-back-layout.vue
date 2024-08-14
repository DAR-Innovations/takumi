<template>
	<AppPrimaryLayout>
		<template #header>
			<div class="w-full relative flex items-center gap-5">
				<button
					@click="$router.back()"
					class="border border-[#242424] px-3 py-2 rounded-xl"
				>
					<i class="text-xl pi pi-angle-left"></i>
				</button>
				<PageTitle>{{ currentPageTitle }}</PageTitle>
			</div>
		</template>

		<slot></slot>
	</AppPrimaryLayout>
</template>

<script setup lang="ts">
import PageTitle from '@/core/components/ui/page-title/page-title.vue'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AppPrimaryLayout from "../primary/app-primary-layout.vue"

const route = useRoute()
const currentPageTitle = computed(() => route.meta.pageTitle ?? route.meta.title)
</script>

<style scoped></style>
