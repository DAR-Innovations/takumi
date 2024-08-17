import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

export const loginFormSchema = toTypedSchema(
	z.object({
		login: z
			.string()
			.min(2, { message: 'Username or email must be at least 2 characters long' })
			.max(50, { message: 'Username or email must be no more than 50 characters long' }),

		password: z
			.string()
			.min(6, { message: 'Password must be at least 6 characters long' })
			.max(50, { message: 'Password must be no more than 50 characters long' })
			.regex(/^(?=.*[A-Z])(?=.*\d).{6,}$/, {
				message: 'Password must contain at least one uppercase letter and one number',
			}),
	}),
)
