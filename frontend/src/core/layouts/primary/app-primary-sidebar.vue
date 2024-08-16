<script setup lang="ts">

import { primarySidebarItems } from '@/core/config/sidebar.config';
import { computed } from 'vue';
import { useRoute } from 'vue-router';

const route = useRoute()
const currentRouteName = computed(() => route.name)

const isActiveRoute = (routeName: string) => {
	return currentRouteName.value?.toString().includes(routeName) ?? false
}
</script>

<template>
	<div class="h-full flex flex-col items-center no-scrollbar overflow-">
		<div>
			<button>
				<img
					src="@/core/assets/icons/logo.svg"
					alt="logo"
					class="w-10 h-10"
				/>
			</button>
		</div>

		<div
			class="h-full flex flex-col items-center justify-center flex-1 gap-2 overflow-y-auto no-scrollbar"
		>
			<template
				v-for="sidebarItem in primarySidebarItems"
				:key="sidebarItem.name"
			>
				<router-link
					v-slot="{ href, navigate }"
					:to="{name: sidebarItem.routeName}"
					custom
				>
					<a
						class="flex items-center justify-center cursor-pointer w-14 h-14 p-3 bg-card rounded-full"
						:class='{
								"text-primary": isActiveRoute(sidebarItem.routeName),
								"text-foreground": !isActiveRoute(sidebarItem.routeName)
							}'
						:href="href"
						@click="navigate"
					>
						<span
							class="text-xl"
							:class="sidebarItem.icon"
						/>
					</a>
				</router-link>
			</template>
		</div>
	</div>
</template>

<style lang="scss" scoped></style>
