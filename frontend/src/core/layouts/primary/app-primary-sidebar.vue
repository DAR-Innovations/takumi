<script setup lang="ts">

import { sidebarItems } from '@/core/config/sidebar.config'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentRouteName = computed(() => route.name)

const isActiveRoute = (routeName: string) => {
	return currentRouteName.value?.toString().includes(routeName) ?? false
}
</script>

<template>
	<div class="flex flex-col flex-1 gap-6 overflow-y-auto no-scrollbar">
		<div
			v-for="sidebarItem in sidebarItems"
			:key="sidebarItem.name"
		>
			<router-link
				v-slot="{ href, navigate }"
				:to="{name: sidebarItem.routeName}"
				custom
			>
				<a
					ripple
					class="flex items-center cursor-pointer"
					:class='{
								"text-primary": isActiveRoute(sidebarItem.routeName),
								"text-foreground": !isActiveRoute(sidebarItem.routeName)
							}'
					:href="href"
					@click="navigate"
				>
					<span :class="sidebarItem.icon" />
					<span class="ml-3">{{ sidebarItem.name }}</span>
				</a>
			</router-link>
		</div>
	</div>
</template>

<style lang="scss" scoped></style>
