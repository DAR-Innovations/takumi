<script setup lang="ts">
import { Toaster } from '@/core/components/ui/toast'
import AppAdminLayout from '@/core/layouts/admin/app-admin-layout.vue'
import AppDefaultLayout from '@/core/layouts/default/app-default-layout.vue'
import AppGoBackLayout from '@/core/layouts/go-back/app-go-back-layout.vue'
import AppPrimaryLayout from '@/core/layouts/primary/app-primary-layout.vue'
import { APP_LAYOUTS } from '@/core/types/routes.types'
import { useColorMode } from '@vueuse/core'
import AppNoTabbarLayout from './core/layouts/no-tabbar/app-no-tabbar-layout.vue'

useColorMode()
</script>

<template>
	<Toaster />

	<AppPrimaryLayout v-if="$route.meta.layout === APP_LAYOUTS.PRIMARY_LAYOUT">
		<router-view />
	</AppPrimaryLayout>

	<AppGoBackLayout v-else-if="$route.meta.layout === APP_LAYOUTS.GO_BACK_LAYOUT">
		<router-view />
	</AppGoBackLayout>

	<AppAdminLayout v-else-if="$route.meta.layout === APP_LAYOUTS.ADMIN_LAYOUT">
		<router-view />
	</AppAdminLayout>

	<AppNoTabbarLayout v-else-if="$route.meta.layout === APP_LAYOUTS.NO_TABBAR">
		<router-view />
	</AppNoTabbarLayout>

	<AppDefaultLayout v-else>
		<router-view />
	</AppDefaultLayout>
</template>
