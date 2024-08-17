<template>
	<button
		class="bg-muted p-2 rounded-full"
		@click="$router.back()"
	>
		<ArrowLeft :size="20" />
	</button>

	<h1 class="mt-6 text-2xl">Log in to Takumi</h1>

	<div class="mt-8">
		<LoginForm />
	</div>
</template>

<script setup lang="ts">
import LoginForm from '@/modules/auth/components/login/login-form.vue'
import { ArrowLeft } from 'lucide-vue-next'
</script>

<style scoped></style>
