<template>
	<days-calendar-scroller />
	<div class="mt-8">
		<goals-banner />
	</div>

	<div class="mt-5">
		<h3 class="mb-4 text-xl">Daily check-in</h3>
		<daily-check-in-container />
	</div>

	<div class="mt-5">
		<h3 class="mb-4 text-xl">Habits</h3>
		<habits-container />
	</div>
</template>

<script setup lang="ts">
import DailyCheckInContainer from '../components/daily-check-in/daily-check-in-container.vue'
import DaysCalendarScroller from '../components/days-calendar-scroller.vue'
import GoalsBanner from '../components/goals-banner.vue'
import HabitsContainer from '../components/habits/habits-container.vue'
</script>

<style scoped></style>
