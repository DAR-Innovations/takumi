<template>
	<AppPrimaryLayout>
		<template #header>
			<div class="w-full relative flex items-center justify-center">
				<button
					@click="$router.back()"
					class="absolute left-0 top-0 h-full flex items-center justify-center"
				>
					<ChevronLeft :size="28" />
				</button>
				<PageTitle>{{ currentPageTitle }}</PageTitle>
			</div>
		</template>

		<slot></slot>
	</AppPrimaryLayout>
</template>

<script setup lang="ts">
import PageTitle from '@/core/components/ui/page-title/page-title.vue'
import { ChevronLeft } from 'lucide-vue-next'
import { computed } from 'vue'
import { useRoute } from 'vue-router'
import AppPrimaryLayout from "../primary/app-primary-layout.vue"

const route = useRoute()
const currentPageTitle = computed(() => route.meta.pageTitle ?? route.meta.title)
</script>

<style scoped></style>
