<template>
	<div class="flex items-center gap-5 bg-[#171717] p-4 rounded-xl">
		<RadialProgress
			:diameter="65"
			:completed-steps="completedSteps"
			:total-steps="totalSteps"
			:stroke-width="4"
			:inner-stroke-width="4"
			stroke-linecap="square"
			inner-stroke-color="#353535"
			start-color="var(--primary)"
			stop-color="var(--primary)"
		>
			<p class="">25%</p>
		</RadialProgress>

		<div>
			<p class="text-base text-medium">Your daily goals almost done! 🔥</p>
			<p class="mt-1 text-gray-400">{{ completedSteps }} of {{ totalSteps }} completed</p>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import RadialProgress from "vue3-radial-progress"

const completedSteps = ref(4)
const totalSteps = ref(10)
</script>

<style scoped></style>
