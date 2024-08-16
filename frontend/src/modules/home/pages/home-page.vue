<template>
	<days-calendar-scroller />

	<div class="mt-6">
		<goals-banner />
	</div>

	<div class="mt-6">
		<page-section-label class="mb-3">Daily check-in</page-section-label>
		<daily-check-in-container />
	</div>

	<div class="mt-6">
		<page-section-label class="mb-3">Habits</page-section-label>
		<habits-container />
	</div>
</template>

<script setup lang="ts">
import PageSectionLabel from '@/core/components/ui/page-section-label/PageSectionLabel.vue'
import DailyCheckInContainer from '../components/daily-check-in/daily-check-in-container.vue'
import DaysCalendarScroller from '../components/days-calendar-scroller.vue'
import GoalsBanner from '../components/goals-banner.vue'
import HabitsContainer from '../components/habits/habits-container.vue'
</script>

<style scoped></style>
