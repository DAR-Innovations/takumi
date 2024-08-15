<template>
	<div class="bg-card p-4 rounded-xl">
		<div class="flex items-center gap-3">
			<span class="bg-card p-1 rounded-full text-lg">{{ emoji }}</span>
			<p class="text-base">{{ title }}</p>
		</div>

		<div class="flex justify-center mt-4">
			<RadialProgress
				:diameter="140"
				:completed-steps="completedSteps"
				:total-steps="totalSteps"
				:stroke-width="12"
				:inner-stroke-width="12"
				stroke-linecap="square"
				inner-stroke-color="#353535"
				:start-color="progressColor"
				:stop-color="progressColor"
			>
				<p class="text-lg font-medium">
					<span class="text-2xl">{{ completedSteps }}</span> / {{ totalSteps }}
				</p>
				<p class="mt-1 text-gray-500">{{ progressLabel }}</p>
			</RadialProgress>
		</div>

		<div class="flex justify-center mt-5">
			<div class="grid grid-cols-2 bg-card rounded-full">
				<button
					class="flex justify-center items-center px-5 py-3 border-r border-r-[#464646]"
					@click='emit("subtract")'
				>
					<i class="text-xs pi pi-minus"></i>
				</button>
				<button
					class="flex justify-center items-center px-5 py-3"
					@click='emit("add")'
				>
					<i class="text-primary text-xs pi pi-plus"></i>
				</button>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import RadialProgress from "vue3-radial-progress";

defineProps({
	completedSteps: Number,
	totalSteps: Number,
	progressLabel:String,
	title: String,
	emoji: String,
	progressColor: String
})

const emit = defineEmits({
	add: null,
	subtract: null,
})
</script>

<style scoped></style>
