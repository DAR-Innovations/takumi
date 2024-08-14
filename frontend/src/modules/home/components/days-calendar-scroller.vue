<template>
	<div class="flex items-start gap-3 overflow-x-auto no-scrollbar sm:scrollable">
		<div
			v-for="item in calendarDays"
			:key="item.id"
			class=" min-w-16"
			@click="onSelectDay(item.day)"
		>
			<div
				class="bg-[#242424] flex items-center justify-center flex-col rounded-xl gap-1 py-3"
				:class='{
				"bg-primary text-primary-foreground": item.day === currentDay
			}'
			>
				<p class="font-bold">{{ item.weekDay }}</p>
				<p class="font-bold">{{ item.day }}</p>
			</div>

			<div class="flex justify-center mt-3 ">
				<p
					v-if="item.mood"
					class="text-4xl"
				>
					{{ item.mood }}
				</p>

				<img
					v-else
					alt="empty emoji"
					src="@/core/assets/icons/empty_emoji.svg"
					class="w-9 h-9"
				/>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const currentDay = ref(25)

const calendarDays = [
	{id: 1, weekDay: "Mon", day: 24, mood: "😔"},
	{id: 2, weekDay: "Tue", day: 25, mood: "😡"},
	{id: 3, weekDay: "Wed", day: 26, mood: "😡"},
	{id: 4, weekDay: "Thu", day: 27, mood: "😐"},
	{id: 5, weekDay: "Fri", day: 28, mood: "🙂"},
	{id: 6, weekDay: "Sat", day: 29, mood: null},
	{id: 7, weekDay: "Sun", day: 30, mood: null}
]

const onSelectDay = (day: number) => {
	currentDay.value = day
}
</script>

<style scoped></style>
