import type { CsrfState } from '$lib/model/auth'
import ExpiryMap from 'expiry-map'

const EXPIRE_MS = 180_000
const volatileStates = new ExpiryMap<string, CsrfState>(EXPIRE_MS)

export function createCsrfState(referer: string): CsrfState {
	const state = {
		value: crypto.randomUUID(),
		referer,
		createdAt: new Date(),
	} satisfies CsrfState

	volatileStates.set(state.value, state)
	return state
}

export function getCsrfSate(val: string): CsrfState | undefined {
	return volatileStates.get(val)
}
