import type { CurrentUser, LoginDTO, SignupDTO } from '@/modules/auth/models/auth.dto'
import axios from 'axios'

class AuthService {
	private readonly baseUrl = '/api/v1/auth'

	async login(dto: LoginDTO) {
		return axios.post(`${this.baseUrl}/login`, dto).then(res => res.data)
	}

	async signup(dto: SignupDTO) {
		return axios.post(`${this.baseUrl}/signup`, dto).then(res => res.data)
	}

	async getCurrentUser() {
		return axios.get<CurrentUser>(`${this.baseUrl}/current-user`)
	}
}

export const authService = new AuthService()
