<script setup lang="ts">
import { APP_LAYOUTS } from '@/core/config/routes.config'
import AppDefaultLayout from '@/core/layouts/default/app-default-layout.vue'
import AppPrimaryLayout from '@/core/layouts/primary/app-primary-layout.vue'
</script>

<template>
	<AppPrimaryLayout v-if="$route.meta.layout === APP_LAYOUTS.PRIMARY_LAYOUT">
		<router-view />
	</AppPrimaryLayout>

	<AppDefaultLayout v-else>
		<router-view />
	</AppDefaultLayout>
</template>
