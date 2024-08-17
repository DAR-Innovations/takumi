import type { LoginDTO, SignupDTO } from '@/modules/auth/models/auth.dto'
import axios from 'axios'

class AuthService {
	private readonly baseUrl = '/api/v1/auth'

	login(dto: LoginDTO) {
		return axios.post(`${this.baseUrl}/login`, dto)
	}

	signup(dto: SignupDTO) {
		return axios.post(`${this.baseUrl}/signup`, dto)
	}
}

export const authService = new AuthService()
