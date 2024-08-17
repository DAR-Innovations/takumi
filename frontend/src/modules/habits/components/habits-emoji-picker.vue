<template>
	<button
		class="flex items-center gap-2 rounded-full px-4 py-2 bg-card text-card-foreground min-h-full"
		@click="openEmojiPicker"
	>
		<input
			ref="emojiInput"
			type="text"
			class="text-base bg-transparent border-none w-6"
			:value="selectedEmoji"
			@input="handleEmojiInput"
		/>
		<p>Icon</p>
	</button>
</template>

<script setup lang="ts">
import { ref, type Ref } from 'vue'

// Regular expression to match emojis
const emojiRegex = /^(?:[\p{Emoji}\p{Emoji_Modifier}\p{Emoji_Component}]+)$/u;

const selectedEmoji: Ref<string> = ref('😎');
const emojiInput: Ref<HTMLInputElement | null> = ref(null);

// Open emoji picker (if you have an actual picker implementation)
const openEmojiPicker = (): void => {
  emojiInput.value?.focus();
};

// Handle emoji input and validate it
const handleEmojiInput = (event: Event): void => {
  const target = event.target as HTMLInputElement;
  const newValue = target.value;

  // Ensure the new value is a single emoji character
  if (emojiRegex.test(newValue) || newValue === '') {
    selectedEmoji.value = newValue;
  } else {
    // If not valid, revert to the previous valid emoji
    target.value = selectedEmoji.value;
  }
};
</script>

<style scoped>
/* Add your custom styles here */
</style>
