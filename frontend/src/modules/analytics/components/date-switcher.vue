<template>
	<div class="flex items-center gap-2 overflow-hidden">
		<button
			v-for="range in ranges"
			:key="range.value"
			class="w-full p-3 rounded-xl"
			:class='{
				"bg-primary text-primary-foreground": selectedRangeValue === range.value,
				"bg-[#171717] text-foreground": selectedRangeValue !== range.value

			}'
			@click="setSelectedRangeValue(range.value)"
		>
			{{ range.label }}
		</button>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const ranges = [{label: "Day", value: "DAY"}, {label: "Week", value: "WEEK"}, {label: "Month", value: "MONTH"}]

const selectedRangeValue = ref(ranges[0].value)

const setSelectedRangeValue = (v: string) => {
	selectedRangeValue.value = v
}
</script>

<style scoped></style>
