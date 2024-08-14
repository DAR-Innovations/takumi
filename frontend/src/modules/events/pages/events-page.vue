<template>
	<div>
		<page-section-label class="mb-4">Challenges</page-section-label>
		<challenges-list />
	</div>

	<div class="mt-8">
		<page-section-label class="mb-4">Habits Club</page-section-label>
		<habits-club />
	</div>

	<div class="mt-8">
		<page-section-label class="mb-4">Learning</page-section-label>
		<learnings-list />
	</div>
</template>

<script setup lang="ts">
import PageSectionLabel from '@/core/components/ui/page-section-label/PageSectionLabel.vue'
import ChallengesList from '@/modules/events/components/challenges-list.vue'
import HabitsClub from '@/modules/events/components/habits-club.vue'
import LearningsList from '@/modules/events/components/learnings-list.vue'
</script>

<style scoped></style>
