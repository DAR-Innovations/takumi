<template>
	<div class="flex items-center gap-2 overflow-x-auto no-scrollbar">
		<button
			v-for="section in sections"
			:key="section.value"
			class="min-w-fit px-5 py-3 rounded-xl"
			:class='{
				"bg-secondary text-secondary-foreground": selectedSection === section.value,
				"bg-[#171717] text-white": selectedSection !== section.value

			}'
			@click="setSelectedSection(section.value)"
		>
			{{ section.label }}
		</button>
	</div>
</template>

<script setup lang="ts">
import { AnalyticsSections } from '@/modules/analytics/models/analytics-sections.type'

defineProps<{
	selectedSection: AnalyticsSections
}>()

const emit = defineEmits<{
  (event: 'select', value: AnalyticsSections): void
}>()

const sections = [
	{label: "Hormones Level", value: AnalyticsSections.HORMONES_LEVEL},
	{label: "Mental Health", value: AnalyticsSections.MENTAL_HEALTH},
	{label: "Productivity", value: AnalyticsSections.PRODUCTIVITY}
]

const setSelectedSection = (v: AnalyticsSections) => {
	emit("select", v)
}
</script>

<style scoped></style>
