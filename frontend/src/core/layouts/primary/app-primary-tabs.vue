<script setup lang="ts">
import { ROUTES } from '@/core/config/routes.config'
import { tabBarRoutes } from '@/core/config/tabbar.config'
import { computed } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const currentRouteName = computed(() => route.name)

const isActiveRoute = (routeName: string) => {
	return currentRouteName.value?.toString().includes(routeName) ?? false
}
</script>

<template>
	<div class="flex flex-1 justify-between items-center gap-4">
		<div
			v-for="item in tabBarRoutes"
			:key="item.name"
		>
			<router-link
				v-slot="{ href, navigate }"
				:to="{name: item.routeName}"
				custom
			>
				<a
					class="flex items-center text-[22px]"
					:class='{
								"text-primary": isActiveRoute(item.routeName),
								"text-[#8f8f8f]": !isActiveRoute(item.routeName),
								"p-2 rounded-xl bg-primary !text-md !text-primary-foreground": item.routeName === ROUTES.CREATE
							}'
					:href="href"
					@click="navigate"
				>
					<span :class="item.icon"></span>
				</a>
			</router-link>
		</div>
	</div>
</template>

<style lang="scss" scoped></style>
