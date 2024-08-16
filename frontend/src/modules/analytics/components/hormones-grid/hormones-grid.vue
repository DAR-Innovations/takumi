<template>
	<div class="grid grid-cols-2 gap-2">
		<div
			v-for="hormone in hormones"
			:key="hormone.name"
			class="bg-card rounded-xl p-4"
		>
			<div class="flex items-center justify-center">
				<span
					class="w-10 h-10"
					:class="cn(hormonesVariants({variant: hormone.key}))"
				></span>
			</div>
			<p class="mt-5 font-semibold">{{ hormone.name }}</p>
			<p class="mt-1 text-xs">{{ hormone.value }}%</p>
			<div class="mt-2">
				<Progress
					v-model="hormone.value"
					class="bg-card"
					:indicator-class="cn(hormonesVariants({variant: hormone.key}))"
				>
				</Progress>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import Progress from '@/core/components/ui/progress/Progress.vue';
import { cn } from '@/core/utils/tailwind.utils';
import { hormonesVariants, type HormonesVariants } from '@/modules/analytics/components/hormones-grid/hormones-grid.styles';
import { ref } from 'vue';

interface Hormones {
	name: string
	key: HormonesVariants['variant'],
	logo: string
	value: number
}

const hormones = ref([
	{name: "Endorphin", key: "endorphin", logo: "@/core/assets/icons/endorphin.webp", value: 48},
	{name: "Dopamine", key: "dopamine",logo: "@/core/assets/icons/dopamine.webp", value: 25},
	{name: "Serotonin", key: "serotonin",logo: "@/core/assets/icons/serotonin.webp", value: 87},
	{name: "Oxytocin", key: "oxytocin",logo: "@/core/assets/icons/oxytocin.webp", value: 39}
] as Hormones[])
</script>

<style scoped></style>
