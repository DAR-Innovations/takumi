<template>
	<form
		class="flex flex-col gap-6"
		@submit="onSubmit"
	>
		<FormField
			v-slot="{ componentField }"
			name="username"
		>
			<FormItem>
				<FormLabel>Username</FormLabel>
				<FormControl>
					<Input
						type="text"
						placeholder="shadcn"
						v-bind="componentField"
					/>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<FormField
			v-slot="{ componentField }"
			name="email"
		>
			<FormItem>
				<FormLabel>Email</FormLabel>
				<FormControl>
					<Input
						type="email"
						placeholder="takumi@gmail.com"
						v-bind="componentField"
					/>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<FormField
			v-slot="{ componentField }"
			name="firstName"
		>
			<FormItem>
				<FormLabel>First name</FormLabel>
				<FormControl>
					<Input
						type="text"
						placeholder="John"
						v-bind="componentField"
					/>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<FormField
			v-slot="{ componentField }"
			name="lastName"
		>
			<FormItem>
				<FormLabel>Last name</FormLabel>
				<FormControl>
					<Input
						type="text"
						placeholder="Snow"
						v-bind="componentField"
					/>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<FormField
			v-slot="{ componentField }"
			name="birthDate"
		>
			<FormItem>
				<FormLabel>Date of birth</FormLabel>
				<FormControl>
					<VueDatePicker
						placeholder="Select your birth date"
						v-bind="componentField"
						:enable-time-picker="false"
						class="bg-card text-card-foreground border border-border rounded-md"
					/>
				</FormControl>
				<FormMessage />
			</FormItem>
		</FormField>

		<FormField
			v-slot="{ componentField }"
			name="gender"
		>
			<FormItem>
				<FormLabel>Gender</FormLabel>
				<FormControl>
					<Select v-bind="componentField">
						<FormControl>
							<SelectTrigger>
								<SelectValue placeholder="Select your gender" />
							</SelectTrigger>
						</FormControl>
						<SelectContent>
							<SelectGroup>
								<template
									v-for="gender in genders"
									:key="gender.label"
								>
									<SelectItem :value="gender.value">{{ gender.label }}</SelectItem>
								</template>
							</SelectGroup>
						</SelectContent>
					</Select>
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
			Signup
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
import {
	Select,
	SelectContent,
	SelectGroup,
	SelectItem,
	SelectTrigger,
	SelectValue
} from '@/core/components/ui/select'
import { toast } from '@/core/components/ui/toast'
import { ROUTES } from '@/core/config/routes.config'
import { signupFormSchema } from '@/modules/auth/components/signup/signup-form-schema'
import type { SignupDTO } from '@/modules/auth/models/auth.dto'
import { Genders } from '@/modules/auth/models/genders.enum'
import { authService } from '@/modules/auth/service/auth.service'
import { useMutation } from '@tanstack/vue-query'
import VueDatePicker from '@vuepic/vue-datepicker'
import '@vuepic/vue-datepicker/dist/main.css'
import { Eye, EyeOff } from 'lucide-vue-next'
import { useForm } from 'vee-validate'
import { ref } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const {isPending, mutate} = useMutation({
		mutationFn: (dto: SignupDTO) => authService.signup(dto),
		onSuccess: () => {
			toast({title: 'You successfully signed up', variant: 'default'})
			router.push({name: ROUTES.HOME})
		},
		onError: () => {
			toast({title: 'Failed to sign up', variant: 'destructive'})
		},
})

const {meta, handleSubmit} = useForm({
  validationSchema: signupFormSchema,
})

const onSubmit = handleSubmit((formData) => {
  mutate(formData)
})

const isPasswordVisible = ref(false)

const changePasswordVisibility = () => {
	isPasswordVisible.value = !isPasswordVisible.value
}

const genders = [
	{label: "Female", value: Genders.FEMALE},
	{label: "Male", value: Genders.MALE},
	{label: "Other", value: Genders.OTHER}
]
</script>

<style scoped>
.dp__main :deep() {
	--dp-background-color: var(--background);
	--dp-text-color: var(--foreground);
	--dp-border-color: var(--border-input);
	--dp-border-color-focus: var(--primary);
	--dp-text-color: var(--foreground);
  --dp-primary-color: var(--primary);
  --dp-primary-text-color: var(--primary-foreground);
	--dp-border-radius: 6px;
	--dp-input-padding: 8px 12px 8px 12px;
	--dp-font-size: 0.875rem;
	--dp-font-family: inherit,
}
.dp__main :deep(.dp__theme_light) {
	--dp-background-color: var(--card);
	--dp-text-color: var(--card-foreground);
	--dp-border-color: var(--border);
	--dp-text-color: var(--foreground);
  --dp-primary-color: var(--primary);
  --dp-primary-text-color: var(--primary-foreground);
	--dp-menu-border-color:var(--border);
}
</style>
