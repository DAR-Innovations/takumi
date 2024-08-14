<template>
	<div class="flex items-center gap-5">
		<Avatar class="w-16 h-16">
			<AvatarImage
				:src="profile.imageUrl"
				:alt="profile.name"
			/>
			<AvatarFallback>{{profile.name.charAt(0)}}</AvatarFallback>
		</Avatar>

		<div>
			<p class="font-semibold text-xl">{{profile.name}}</p>
			<div class="flex items-center gap-2 mt-1">
				<img
					src="@/core/assets/icons/kaiteki-primary.svg"
					alt="kaiteki coin"
					class="w-4 h-4"
				/>
				<p class="text-primary">{{ profile.coins }}</p>
			</div>
		</div>
	</div>

	<div class="flex flex-col mt-8">
		<router-link
			v-for="section in sections"
			:key="section.name"
			:to="{name: section.routeName}"
			class="w-full flex items-center gap-4 justify-between p-4"
		>
			<div class="flex items-center gap-4">
				<i
					class="text-lg"
					:class="section.icon"
				/>
				<p class="text-base">{{ section.name }}</p>
			</div>

			<i class="pi pi-angle-right text-gray-300 text-lg"></i>
		</router-link>
	</div>
</template>

<script setup lang="ts">
import Avatar from '@/core/components/ui/avatar/Avatar.vue'
import AvatarFallback from '@/core/components/ui/avatar/AvatarFallback.vue'
import AvatarImage from '@/core/components/ui/avatar/AvatarImage.vue'
import { ROUTES } from '@/core/config/routes.config'
import { ref } from 'vue'

const profile = ref({
	name: "Diar Begisbayev",
	imageUrl: "https://www.perfocal.com/blog/content/images/2021/01/Perfocal_17-11-2019_TYWFAQ_100_standard-3.jpg",
	coins: "2321"
})

const sections = ref([
	{name: "Achievements", routeName: ROUTES.PROFILE_ACHIEVEMENTS, icon: "pi pi-check-square"},
	{name: "Friends", routeName: ROUTES.PROFILE_ACHIEVEMENTS, icon: "pi pi-users"},
	{name: "Joined events", routeName: ROUTES.PROFILE_ACHIEVEMENTS, icon: "pi pi-compass"},
	{name: "Settings", routeName: ROUTES.PROFILE_ACHIEVEMENTS, icon: "pi pi-cog"}
])
</script>

<style scoped></style>
