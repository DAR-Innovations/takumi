<template>
	<div class="flex items-center gap-4 overflow-x-auto no-scrollbar">
		<div
			v-for="club in clubs"
			:key="club.name"
			class="rounded-xl bg-card min-w-fit p-5"
		>
			<span class="text-4xl">{{ club.icon }}</span>
			<p class="font-semibold mt-4">{{ club.name }}</p>
			<p class="text-gray-500 mt-1">{{club.members }} members</p>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

	const clubs = ref([
		{name: "Bikers", members: 423, icon: "🏍️"},
		{name: "Runners", members: 231, icon: "🏃"},
		{name: "Meditation", members: 23, icon: "🧘"},
		{name: "Swimming", members: 663, icon: "🏊"},
		{name: "Climbers", members: 54, icon: "⛰️"},
	])
</script>

<style scoped></style>
