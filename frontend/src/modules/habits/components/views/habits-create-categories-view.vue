<template>
	<div class="flex flex-col gap-2">
		<template
			v-for="item in categories"
			:key="item.id"
		>
			<button
				@click="setLocalCategory(item.id)"
				:class="{'bg-primary': localCategoryId === item.id, 'bg-card': localCategoryId !== item.id}"
				class="flex justify-between items-center text-card-foreground p-4 rounded-xl"
			>
				<div class="flex items-center gap-3">
					<div class="p-2 rounded-full text-xl bg-muted w-10 h-10 flex items-center justify-center">
						{{ item.icon }}
					</div>
					<p
						class="font-semibold text-base"
						:class="{'text-primary-foreground': localCategoryId === item.id, 'text-card-foreground': localCategoryId !== item.id}"
					>
						{{ item.label }}
					</p>
				</div>

				<div class="p-2 rounded-full bg-muted w-fit text-xl">
					<Plus
						:size="20"
						class="text-primary text-bold"
					/>
				</div>
			</button>
		</template>
	</div>

	<div class="mt-6">
		<Button
			:disabled="!localCategoryId"
			@click="confirmCategorySelection"
			class="w-full font-semibold"
		>
			Done
		</Button>
	</div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { Plus } from 'lucide-vue-next'
import { ref } from 'vue'

const emits = defineEmits({
	select: null
})

const localCategoryId = ref<number | null>(null)

const categories = [
  { id: 1, label: "Health", icon: "🧘🏻‍♀️" },
  { id: 2, label: "Arts", icon: "🎨" },
  { id: 3, label: "Sport", icon: "🏀" },
  { id: 4, label: "Language", icon: "🔠" },
  { id: 5, label: "Mindfulness", icon: "🌱" },
  { id: 6, label: "Skills Development", icon: "🚀" },
]

const setLocalCategory = (categoryId: number) => {
	if (localCategoryId.value === categoryId) {
		localCategoryId.value = null
	} else {
		localCategoryId.value = categoryId
	}
}

const confirmCategorySelection = () => {
	emits("select", localCategoryId.value)
}
</script>

<style scoped></style>
