<template>
	<div class="flex items-center gap-3 overflow-x-auto no-scrollbar sm:scrollable">
		<div
			v-for="item in calendarDays"
			:key="item.id"
			class="min-w-20"
			@click="onSelectDay(item.day)"
		>
			<div
				class="bg-[#242424] px-6 py-3 rounded-xl"
				:class='{
				"bg-primary text-primary-foreground": item.day === currentDay
			}'
			>
				<p class="font-bold text-center">{{ item.weekDay }}</p>
				<p class="mt-1 font-bold text-center">{{ item.day }}</p>
			</div>

			<div class="flex justify-center mt-2">
				<p>{{ item.mood ?? "Empty" }}</p>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const currentDay = ref(25)

const calendarDays = [
	{id: 1, weekDay: "Mon", day: 24, mood: "Sad"},
	{id: 2, weekDay: "Tue", day: 25, mood: "Angry"},
	{id: 3, weekDay: "Wed", day: 26, mood: "Angry"},
	{id: 4, weekDay: "Thu", day: 27, mood: "Neutral"},
	{id: 5, weekDay: "Fri", day: 28, mood: "Happy"},
	{id: 6, weekDay: "Sat", day: 29, mood: null},
	{id: 7, weekDay: "Sun", day: 30, mood: null}
]

const onSelectDay = (day: number) => {
	currentDay.value = day
}
</script>

<style scoped></style>
