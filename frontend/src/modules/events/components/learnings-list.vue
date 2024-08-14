<template>
	<div class="flex items-center gap-4 overflow-x-auto no-scrollbar">
		<div
			v-for="learning in learnings"
			:key="learning.title"
			class="relative rounded-xl bg-[#171717] overflow-hidden min-w-64 max-w-64"
		>
			<img
				:src="learning.imageUrl"
				:alt="learning.title"
				class="w-full h-32 object-cover bg-[#242424]"
			/>
			<div class="p-4 flex items-center gap-3 z-10">
				<i class="pi pi-book"></i>
				<p class="font-medium truncate">{{ learning.title }}</p>
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">
import { ref } from 'vue'

const learnings = ref([
  { title: "Why should we drink water often?", imageUrl: "https://diy-ideas.oss-accelerate.aliyuncs.com/wp-content/uploads/2024/02/filtered-water-vs-distilled-water-1024x682.png!webp" },
  { title: "Benefits of regular walking", imageUrl: "https://experiencelife.lifetime.life/wp-content/uploads/2023/02/apr23-feat-born-to-walk.jpg" },
  { title: "What diet you must follow to lose weight", imageUrl: "https://totalwellnessandbariatrics.com/wp-content/uploads/want-to-lose-weight.jpg" },
  { title: "Running in the morning or in the evening?", imageUrl: "https://hips.hearstapps.com/hmg-prod/images/gettyimages-1341854189-6540f8266a2ef.jpg?crop=0.8888888888888888xw:1xh;center,top&resize=1200:*" },
])
</script>

<style scoped>
/* You may adjust the overlay color and opacity here if needed */
</style>
