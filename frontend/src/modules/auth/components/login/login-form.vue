<template>
	<form
		class="flex flex-col gap-6"
		@submit="onSubmit"
	>
		<FormField
			v-slot="{ componentField }"
			name="login"
		>
			<FormItem>
				<FormLabel>Username / Email</FormLabel>
				<FormControl>
					<Input
						type="text"
						placeholder="johnsnow@takumi.com"
						v-bind="componentField"
					/>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<FormField
			v-slot="{ componentField }"
			name="password"
		>
			<FormItem>
				<FormLabel>Password</FormLabel>
				<FormControl>
					<div class="relative">
						<Input
							:type="isPasswordVisible ? 'text' : 'password'"
							placeholder="Password"
							v-bind="componentField"
						>
						</Input>
						<button
							type="button"
							class="absolute right-3 top-[10px]"
							@click="changePasswordVisibility"
						>
							<Eye
								v-if="!isPasswordVisible"
								class="opacity-50"
								:size="20"
							/>
							<EyeOff
								v-if="isPasswordVisible"
								class="opacity-50"
								:size="20"
							/>
						</button>
					</div>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<Button
			class="w-full"
			type="submit"
			@click="onSubmit"
			:disabled="!meta.valid ?? isPending"
		>
			Login
		</Button>
	</form>
</template>

<script setup lang="ts">
import { Button } from '@/core/components/ui/button'
import {
	FormControl,
	FormField,
	FormItem,
	FormLabel,
	FormMessage
} from '@/core/components/ui/form'
import { Input } from '@/core/components/ui/input'
import { useToast } from '@/core/components/ui/toast'
import { ROUTES } from '@/core/config/routes.config'
import { loginFormSchema } from '@/modules/auth/components/login/login-form-schema'
import type { LoginDTO } from '@/modules/auth/models/auth.dto'
import { authService } from '@/modules/auth/service/auth.service'
import { useMutation } from '@tanstack/vue-query'
import '@vuepic/vue-datepicker/dist/main.css'
import { Eye, EyeOff } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const { toast } = useToast()

const {isPending, mutate} = useMutation({
		mutationFn: (dto: LoginDTO) => authService.login(dto),
		onSuccess: () => {
			toast({title: 'You successfully logged in', variant: 'default'})
			router.push({name: ROUTES.HOME})
		},
		onError: () => {
			toast({title: 'Failed to log in', variant: 'destructive'})
		},
})

const {meta, handleSubmit} = useForm({
  validationSchema: loginFormSchema,
})

const onSubmit = handleSubmit((formData) => {
  mutate(formData)
})

const isPasswordVisible = ref(false)

const changePasswordVisibility = () => {
	isPasswordVisible.value = !isPasswordVisible.value
}
</script>

<style scoped></style>
