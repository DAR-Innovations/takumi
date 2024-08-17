<template>
	<div>
		<img
			src="@/core/assets/backgrounds/landing.jpg"
			alt="landing"
			class="w-full h-96 object-cover"
		/>
	</div>
	<div class="mt-6">
		<p class="text-center text-2xl font-semibold">Health. Strong. Flexible.</p>
		<p class="text-center text-2xl font-semibold">Be the best!</p>
		<p class="text-center text-muted-foreground text-sm mt-4">
			The neuroscience of happiness in your pocket!
		</p>
	</div>

	<div class="mt-6">
		<router-link :to="{name: ROUTES.AUTH_SIGNUP}">
			<button class="w-full bg-primary text-primary-foreground p-4 rounded-xl font-semibold">
				I'm new here
			</button>
		</router-link>
		<button class="w-full mt-6 text-primary font-semibold">
			<router-link :to="{name: ROUTES.AUTH_LOGIN}">I'm already have an account</router-link>
		</button>
	</div>
</template>

<script setup lang="ts">
import { ROUTES } from '@/core/config/routes.config'
import { RouterLink } from 'vue-router'
</script>

<style scoped></style>
