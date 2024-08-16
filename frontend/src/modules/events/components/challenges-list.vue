<template>
	<div class="flex items-center gap-2 overflow-x-auto no-scrollbar">
		<div
			v-for="challenge in challenges"
			:key="challenge.title"
			class="rounded-xl bg-card w-64 p-5 flex-shrink-0"
		>
			<p class="font-semibold text-base truncate">{{ challenge.title }} {{ challenge.icon }}</p>

			<div class="flex items-center gap-2 mt-2">
				<i class="pi pi-users text-gray-500"></i>
				<p class="text-gray-500">{{ challenge.joined }} people joined</p>
			</div>

			<div class="flex items-center gap-2 mt-2">
				<i class="pi pi-clock text-gray-500"></i>
				<p class="text-gray-500">{{ getTimeLeft(today, challenge.startDate) }} left</p>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { getTimeLeft } from '@/core/utils/time.utils';
import { ref } from 'vue';

const today = new Date();

let nextDay = new Date(today);
nextDay.setDate(today.getDate() + 1);

let nextMonth = new Date(today);
nextMonth.setMonth(today.getMonth() + 1);

let threeDaysNext = new Date(today);
threeDaysNext.setDate(today.getDate() + 3);

let twoDaysNext = new Date(today);
twoDaysNext.setDate(today.getDate() + 2);

const challenges = ref([
  { title: "Best Runners", joined: 423, icon: "🏃", startDate: nextDay },
  { title: "Best Bikers", joined: 234, icon: "🚲", startDate: nextMonth },
  { title: "Half-marathon Runners", joined: 123, icon: "🏃", startDate: threeDaysNext },
  { title: "Long walkers", joined: 654, icon: "🚶‍♂️", startDate: twoDaysNext },
]);
</script>

<style scoped></style>
