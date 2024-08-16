<template>
	<div class="grid grid-cols-3 gap-2 overflow-x-auto no-scrollbar">
		<button
			v-for="section in sections"
			:key="section.value"
			class="min-w-fit px-5 py-3 rounded-xl"
			:class='{
				"bg-primary text-primary-foreground": selectedSection === section.value,
				"bg-card text-card-foreground": selectedSection !== section.value

			}'
			@click="setSelectedSection(section.value)"
		>
			{{ section.label }}
		</button>
	</div>
</template>

<script setup lang="ts">
import { AnalyticsSections } from '@/modules/analytics/models/analytics-sections.type';

defineProps<{
	selectedSection: AnalyticsSections
}>()

const emit = defineEmits<{
  (event: 'select', value: AnalyticsSections): void
}>()

const sections = [
	{label: "Hormones", value: AnalyticsSections.HORMONES_LEVEL},
	{label: "Mental", value: AnalyticsSections.MENTAL_HEALTH},
	{label: "Productivity", value: AnalyticsSections.PRODUCTIVITY}
]

const setSelectedSection = (v: AnalyticsSections) => {
	emit("select", v)
}
</script>

<style scoped></style>
