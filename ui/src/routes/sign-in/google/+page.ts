import { PUBLIC_GOOGLE_OAUTH_CLIENT_ID, PUBLIC_GOOGLE_OAUTH_ENDPOINT } from '$env/static/public'
import type { CsrfState } from '$lib/model/auth'
import { error } from '@sveltejs/kit'
import type { PageLoad } from './$types'

export const load = (async (event) => {
	const { url, fetch } = event
	const referer = url.searchParams.get('referer') ?? '/'

	const response = await fetch(`/api/auth/state?referer=${encodeURIComponent(referer)}`, {
		method: 'POST',
	})
	if (!response.ok) throw error(response.status, 'Failed to issue auth state')

	const state: CsrfState = await response.json()

	const querySet: Record<string, string> = {
		response_type: 'token id_token',
		client_id: PUBLIC_GOOGLE_OAUTH_CLIENT_ID,
		scope: 'openid profile email',
		redirect_uri: `${url.origin}/sign-in/google/callback`,
		state: state.value,
		nonce: crypto.randomUUID(),
	}

	const queries: string[] = []
	for (const key in querySet) {
		queries.push(`${key}=${encodeURIComponent(querySet[key])}`)
	}

	const googleOauthUri = `${PUBLIC_GOOGLE_OAUTH_ENDPOINT}?${queries.join('&')}`
	window.location.href = googleOauthUri
}) satisfies PageLoad
