<template>
	<div class="flex items-center gap-2">
		<HabitsEmojiPicker />
		<HabitsColorPicker />
	</div>

	<div class="mt-4">
		<Textarea
			placeholder="Enter details..."
			class="resize-none w-full h-52 border-none text-2xl"
			v-model="details"
		/>
	</div>

	<div class="mt-4 flex flex-col gap-2">
		<button
			class="w-full flex items-center justify-between gap-4 p-4 bg-card text-card-foreground rounded-xl"
		>
			<p class="text-base">Frequency</p>
			<div class="flex items-center gap-2">
				<p class="text-xs">Every day</p>
				<ChevronRight :size="14" />
			</div>
		</button>

		<button
			class="w-full flex items-center justify-between gap-4 p-4 bg-card text-card-foreground rounded-xl"
		>
			<p class="text-base">Repeats</p>
			<div class="flex items-center gap-2">
				<p class="text-xs">1 time per day</p>
				<ChevronRight :size="14" />
			</div>
		</button>

		<button
			class="w-full flex items-center justify-between gap-4 p-4 bg-card text-card-foreground rounded-xl"
		>
			<p class="text-base">Reminders</p>
			<div class="flex items-center gap-2">
				<p class="text-sm">On</p>
				<ChevronRight :size="14" />
			</div>
		</button>
	</div>

	<div class="mt-6">
		<Button
			@click="onCreate"
			class="w-full font-semibold"
		>
			Create new
		</Button>
	</div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Textarea } from '@/core/components/ui/textarea'
import HabitsColorPicker from '@/modules/habits/components/habits-color-picker.vue'
import HabitsEmojiPicker from '@/modules/habits/components/habits-emoji-picker.vue'

import { ChevronRight } from 'lucide-vue-next'
import { ref } from 'vue'

const emits = defineEmits({
	create: null
})

const details = ref('')

const onCreate = () => {
	emits("create")
}
</script>

<style scoped></style>
