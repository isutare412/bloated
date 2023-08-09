import { authState } from '$lib/browser/auth/store'
import type { GetCustomClaimsResponse } from '$lib/model/auth'

export async function updateAuthStateFromCookie() {
	const response = await fetch('/api/auth/claims')
	if (!response.ok) return

	const { claims }: GetCustomClaimsResponse = await response.json()
	if (!claims) return

	authState.update((state) => {
		state.claims = claims
		return state
	})
}

export function clearAuthState() {
	authState.update((state) => {
		state.claims = undefined
		return state
	})
}
