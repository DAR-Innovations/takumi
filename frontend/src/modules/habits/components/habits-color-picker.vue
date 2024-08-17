<template>
	<button
		class="flex items-center gap-2 rounded-full px-4 py-2 bg-card text-card-foreground min-h-full"
		@click="openColorPicker"
	>
		<input
			ref="colorInput"
			type="color"
			class="w-6 h-6 p-0 border-none bg-transparent"
			:value="selectedColor"
			@input="handleColorInput"
		/>

		<p>Color</p>
	</button>
</template>

<script setup lang="ts">
import { ref, type Ref } from 'vue'

const selectedColor: Ref<string> = ref('#4f46e5');
const colorInput: Ref<HTMLInputElement | null> = ref(null);

const openColorPicker = (): void => {
  colorInput.value?.focus();
};

const handleColorInput = (event: Event): void => {
  const target = event.target as HTMLInputElement;
  selectedColor.value = target.value;
};
</script>

<style scoped></style>
