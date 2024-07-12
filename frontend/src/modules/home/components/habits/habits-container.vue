<template>
	<div class="flex flex-col gap-2">
		<template
			v-for="habit in habits"
			:key="habit.id"
		>
			<habits-item
				:title="habit.title"
				:emoji="habit.emoji"
				:completed="habit.completed"
				@toggle-checkbox="onHabitCheck(habit.id)"
			/>
		</template>
	</div>
</template>

<script setup lang="ts">
import HabitsItem from '@/modules/home/components/habits/habits-item.vue'
import { ref, watchEffect } from 'vue'

const habits = ref([
	{id: 1, title: "Avoid sweets", emoji: "🍰", completed: false},
	{id: 2, title: "Keto diet", emoji: "🥗", completed: false},
	{id: 3, title: "Work out for 30 min", emoji: "🏋️", completed: true},
	{id: 4, title: "Breathe exercise", emoji: "😮‍💨", completed: false},
])

const onHabitCheck = (id: number) => {
  const foundIndex = habits.value.findIndex(item => item.id === id);

  if (foundIndex !== -1) {
		const previousCheckValue = habits.value[foundIndex].completed
    habits.value[foundIndex].completed = !previousCheckValue
  }
};

watchEffect( () => {
	console.log("Updated")
})
</script>

<style scoped></style>
