<template>
	<div class="h-full flex flex-col items-center justify-center">
		<p class="text-center text-5xl md:text-7xl font-bold text-primary">Oops!</p>
		<p class="text-center text-lg md:text-xl mt-2 md:mt-5 text-muted-foreground">
			Something went wrong
		</p>
		<RouterLink :to="{name: ROUTES.LANDING}">
			<Button
				class="mt-8"
				size="lg"
			>
				Return to home
			</Button>
		</RouterLink>
	</div>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import { ROUTES } from '@/core/config/routes.config'
import { RouterLink } from 'vue-router'
</script>

<style scoped></style>
