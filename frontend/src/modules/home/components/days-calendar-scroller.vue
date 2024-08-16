<template>
	<div class="flex items-start gap-3 overflow-x-auto no-scrollbar sm:scrollable">
		<button
			v-for="item in calendarDays"
			:key="item.id"
			class="min-w-16"
			@click="onSelectDay(item.day)"
			:data-day="item.day"
		>
			<div
				class="bg-card flex items-center justify-center flex-col rounded-xl gap-1 py-3"
				:class="{
					'bg-primary text-primary-foreground': item.day === currentDay, 
					'bg-muted text-muted-foreground': item.day === selectedDay
				}"
			>
				<p class="font-bold">{{ item.weekDay }}</p>
				<p class="font-bold">{{ item.day }}</p>
			</div>
			<div class="flex justify-center mt-3">
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
					class="w-9 h-9 opacity-30 dark:opacity-100"
				/>
			</div>
		</button>
	</div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'

const currentDate = new Date();
const currentDay = currentDate.getDate()
const selectedDay = ref(currentDate.getDate());
const daysInMonth = new Date(currentDate.getFullYear(), currentDate.getMonth() + 1, 0).getDate();
const weekDays = ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat'];
const moods = ['😡', '😔', '🙂', '😐'];

const calendarDays = ref(
  Array.from({ length: daysInMonth }, (_, i) => {
    const date = new Date(currentDate.getFullYear(), currentDate.getMonth(), i + 1);
    return {
      id: i + 1,
      weekDay: weekDays[date.getDay()],
      day: i + 1,
      mood: i + 1 < selectedDay.value ? moods[Math.floor(Math.random() * moods.length)] : null,
    };
  })
);

const onSelectDay = (day: number) => {
  selectedDay.value = day;
};

onMounted(() => {
  const currentDayElement = document.querySelector(`[data-day="${selectedDay.value}"]`);
  if (currentDayElement) {
    currentDayElement.scrollIntoView({ behavior: 'smooth', block: 'center' });
  }
});
</script>

<style scoped></style>
