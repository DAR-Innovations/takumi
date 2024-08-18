import { defineStore } from 'pinia'
import { computed, ref, type Ref } from 'vue'
import type { CurrentUser } from '../models/auth.dto'

export const useAuthStore = defineStore('auth', () => {
	const currentUser: Ref<CurrentUser | null> = ref(null)
	const isLoggedIn = computed(() => Boolean(currentUser.value))

	const setCurrentUser = (user: CurrentUser | null) => {
		currentUser.value = user
	}

	return { currentUser, isLoggedIn, setCurrentUser }
})
