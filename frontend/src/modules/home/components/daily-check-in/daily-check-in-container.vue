<template>
	<div class="gap-4 grid grid-cols-2">
		<daily-check-in-item
			:completed-steps="completedWaterSteps"
			:total-steps="totalWaterSteps"
			title="Water"
			emoji="💧"
			progress-label="glass"
			progress-color="var(--secondary)"
			@add="onWaterAdd"
			@subtract="onWaterSubtract"
		/>
		<daily-check-in-item
			:completed-steps="completedSleepSteps"
			:total-steps="totalSleepSteps"
			title="Sleep"
			emoji="😴"
			progress-label="hours"
			progress-color="var(--primary)"
			@add="onSleepAdd"
			@subtract="onSleepSubtract"
		/>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import DailyCheckInItem from './daily-check-in-item.vue'

const completedWaterSteps = ref(4)
const totalWaterSteps = ref(10)

const completedSleepSteps = ref(5)
const totalSleepSteps = ref(8)

const onWaterAdd = () => {
	if (completedWaterSteps.value < totalWaterSteps.value) {
		completedWaterSteps.value += 1
	}
}
const onWaterSubtract = () => {
	if (completedWaterSteps.value > 0) {
		completedWaterSteps.value -= 1
	}
}

const onSleepAdd = () => {
	if (completedSleepSteps.value < totalSleepSteps.value) {
		completedSleepSteps.value += 1
	}
}
const onSleepSubtract = () => {
	if (completedSleepSteps.value > 0) {
		completedSleepSteps.value -= 1
	}
}
</script>

<style scoped></style>
