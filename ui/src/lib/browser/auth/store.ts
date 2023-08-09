import type { CustomTokenClaims } from '$lib/model/auth'
import { writable } from 'svelte/store'

interface AuthState {
	claims?: CustomTokenClaims
}

export const authState = writable<AuthState>({})
