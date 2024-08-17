import type { Genders } from '@/modules/auth/models/genders.enum'

export interface LoginDTO {
	login: string
	password: string
}

export interface SignupDTO {
	username: string
	firstName: string
	lastName: string
	gender: Genders
	birthDate: Date
	email: string
	password: string
}
