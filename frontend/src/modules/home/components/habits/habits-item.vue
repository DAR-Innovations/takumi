<template>
	<button
		class="flex justify-between items-center gap-4 bg-card p-4 rounded-xl"
		@click="onCheckBoxClick"
	>
		<div class="flex items-center gap-3">
			<span class="bg-card p-1 rounded-full text-lg">{{ emoji }}</span>
			<p class="text-base">{{ title }}</p>
		</div>

		<Checkbox
			class="rounded-full w-6 h-6"
			:checked="completed"
		/>
	</button>
</template>

<script setup lang="ts">
import { Checkbox } from '@/core/components/ui/checkbox'

defineProps({
  emoji: { type: String, required: true },
  title: { type: String, required: true },
  completed: { type: Boolean, required: true }
})

const emit = defineEmits({'toggleCheckbox': null})

const onCheckBoxClick = () => {
  emit('toggleCheckbox')
}
</script>

<style scoped></style>
