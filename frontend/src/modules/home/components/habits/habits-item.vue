<template>
	<button
		class="flex justify-between items-center gap-4 bg-[#171717] p-4 rounded-xl"
		@click='emit("toggleCheckbox")'
	>
		<div class="flex items-center gap-3">
			<span class="bg-[#242424] p-1 rounded-full text-lg">{{ emoji }}</span>
			<p class="text-base">{{ title }}</p>
		</div>

		<Checkbox
			class="rounded-full w-6 h-6"
			:checked="completed"
			@update:checked="emit('toggleCheckbox')"
		/>
	</button>
</template>

<script setup lang="ts">
import { Checkbox } from '@/core/components/ui/checkbox'

defineProps({
	emoji: {type: String, required: true},
	title: {type: String, required: true},
	completed: {type: Boolean, required: true}
})

const emit = defineEmits({
	toggleCheckbox: null
})
</script>

<style scoped></style>
