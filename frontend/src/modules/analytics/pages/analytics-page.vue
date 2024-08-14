<template>
	<date-switcher />

	<div class="mt-4">
		<section-switcher
			:selected-section="selectedSection"
			@select="onSelectSection"
		/>
	</div>

	<div class="mt-8">
		<component :is="currentPageComponent" />
	</div>
</template>

<script setup lang="ts">
import { AnalyticsSections } from '@/modules/analytics/models/analytics-sections.type'
import { computed, defineAsyncComponent, ref } from 'vue'
import DateSwitcher from "../components/date-switcher.vue"
import SectionSwitcher from "../components/section-switcher.vue"

const componentsMap = {
	[AnalyticsSections.HORMONES_LEVEL]: defineAsyncComponent(() => import('@/modules/analytics/pages/analytics-hormones-page.vue')),
	[AnalyticsSections.MENTAL_HEALTH]: defineAsyncComponent(() => import('@/modules/analytics/pages/analytics-mental-page.vue')),
	[AnalyticsSections.PRODUCTIVITY]: defineAsyncComponent(() => import('@/modules/analytics/pages/analytics-productivity-page.vue')),
}

const selectedSection = ref(AnalyticsSections.HORMONES_LEVEL)

const onSelectSection = (v: AnalyticsSections) => {
	selectedSection.value = v
}

const currentPageComponent = computed(() => componentsMap[selectedSection.value])
</script>

<style scoped></style>
