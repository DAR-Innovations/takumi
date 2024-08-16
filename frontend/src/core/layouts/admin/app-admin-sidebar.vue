<script setup lang="ts">

import { ROUTES } from '@/core/config/routes.config'
import { adminSidebarItems } from '@/core/config/sidebar.config'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentRouteName = computed(() => route.name)

const isActiveRoute = (routeName: string) => {
	return currentRouteName.value?.toString().includes(routeName) ?? false
}
</script>

<template>
	<div class="h-full flex flex-col overflow-hidden">
		<button class="w-full flex items-center  gap-3 border-b border-border p-6">
			<img
				src="@/core/assets/icons/logo.svg"
				alt="logo"
				class="w-8 h-8"
			/>

			<p class="text-xl">Takumi</p>
		</button>

		<div class="h-full flex flex-col gap-3 mt-8 overflow-y-auto no-scrollbar p-6 pt-0">
			<template
				v-for="sidebarItem in adminSidebarItems"
				:key="sidebarItem.name"
			>
				<router-link
					v-slot="{ href, navigate }"
					:to="{name: sidebarItem.routeName}"
					custom
				>
					<a
						class="flex cursor-pointer "
						:class='{
								"text-primary": isActiveRoute(sidebarItem.routeName),
								"text-foreground": !isActiveRoute(sidebarItem.routeName)
							}'
						:href="href"
						@click="navigate"
					>
						<div class="flex items-center gap-4">
							<span
								class="text-lg"
								:class="sidebarItem.icon"
							/>

							<p class="text-sm">{{ sidebarItem.name }}</p>
						</div>
					</a>
				</router-link>
			</template>
		</div>

		<div class="flex-1 p-6">
			<router-link
				:to="{name: ROUTES.HOME}"
				class="w-full flex gap-4 cursor-pointer bg-card items-center justify-center px-4 py-2 rounded-lg text-card-foreground"
			>
				<span class="text-lg pi pi-angle-left" />

				<p class="text-sm">Return to App</p>
			</router-link>
		</div>
	</div>
</template>

<style lang="scss" scoped></style>
