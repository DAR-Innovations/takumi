<template>
	<div class="w-full h-full sm:flex sm:items-center sm:justify-center">
		<div class="w-full max-w-lg">
			<button
				class="bg-muted p-2 rounded-full"
				@click="$router.back()"
			>
				<ArrowLeft :size="20" />
			</button>

			<h1 class="mt-6 text-2xl">Sign up to Takumi</h1>

			<div class="mt-8">
				<SignupForm />
			</div>
		</div>
	</div>
</template>

<script setup lang="ts">

import SignupForm from '@/modules/auth/components/signup/signup-form.vue'
import { ArrowLeft } from 'lucide-vue-next'
</script>

<style scoped></style>
