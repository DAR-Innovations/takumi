import { Genders } from '@/modules/auth/models/genders.enum'
import { toTypedSchema } from '@vee-validate/zod'
import * as z from 'zod'

export const signupFormSchema = toTypedSchema(
	z.object({
		username: z
			.string()
			.min(2, { message: 'Username must be at least 2 characters long' })
			.max(50, { message: 'Username must be no more than 50 characters long' })
			.regex(/^[a-zA-Z0-9_]+$/, {
				message: 'Username can only contain letters, numbers, and underscores',
			}),

		email: z
			.string()
			.email({ message: 'Invalid email address' })
			.min(5, { message: 'Email must be at least 5 characters long' })
			.max(50, { message: 'Email must be no more than 50 characters long' }),

		firstName: z
			.string()
			.min(2, { message: 'First name must be at least 2 characters long' })
			.max(50, { message: 'First name must be no more than 50 characters long' }),

		lastName: z
			.string()
			.min(2, { message: 'Last name must be at least 2 characters long' })
			.max(50, { message: 'Last name must be no more than 50 characters long' }),

		password: z
			.string()
			.min(6, { message: 'Password must be at least 6 characters long' })
			.max(50, { message: 'Password must be no more than 50 characters long' })
			.regex(/^(?=.*[A-Z])(?=.*\d).{6,}$/, {
				message: 'Password must contain at least one uppercase letter and one number',
			}),

		gender: z.nativeEnum(Genders, {
			errorMap: () => ({ message: 'Please select a valid gender' }),
		}),

		birthDate: z.date({
			required_error: 'Birth date is required',
			invalid_type_error: 'Invalid date format',
		}),
	}),
)
