<template>
	<!-- <div class="flex items-center gap-2 overflow-hidden">
		<button
			v-for="range in ranges"
			:key="range.value"
			class="w-full p-3 rounded-xl"
			:class='{
				"bg-primary text-primary-foreground": selectedRangeValue === range.value,
				"bg-card text-foreground": selectedRangeValue !== range.value

			}'
			@click="setSelectedRangeValue(range.value)"
		>
			{{ range.label }}
		</button>
	</div> -->

	<Select>
		<SelectTrigger class="w-full">
			<SelectValue placeholder="Select a period" />
		</SelectTrigger>
		<SelectContent>
			<SelectGroup>
				<SelectLabel>Periods</SelectLabel>
				<template
					v-for="range in ranges"
					:key="range.label"
				>
					<SelectItem
						:value="range.value"
						@click="setSelectedRangeValue(range.value)"
					>
						{{ range.label }}
					</SelectItem>
				</template>
			</SelectGroup>
		</SelectContent>
	</Select>
</template>

<script setup lang="ts">
import {
	Select,
	SelectContent,
	SelectGroup,
	SelectItem,
	SelectLabel,
	SelectTrigger,
	SelectValue
} from '@/core/components/ui/select'
import { ref } from 'vue'

const ranges = [{label: "Day", value: "DAY"}, {label: "Week", value: "WEEK"}, {label: "Month", value: "MONTH"}]

const selectedRangeValue = ref(ranges[0].value)

const setSelectedRangeValue = (v: string) => {
	selectedRangeValue.value = v
}
</script>

<style scoped></style>
