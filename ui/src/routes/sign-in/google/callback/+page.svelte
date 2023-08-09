<script lang="ts">
	import { goto } from '$app/navigation'
	import { page } from '$app/stores'
	import LoadingDots from '$components/LoadingDots.svelte'
	import { updateAuthStateFromCookie } from '$lib/browser/auth/action'
	import type { GoogleSignInRequest, GoogleSignInResponse } from '$lib/model/auth'
	import { addToast } from '$lib/toast/action'
	import { onMount } from 'svelte'

	onMount(async () => {
		let fragment = $page.url.hash
		if (!fragment.startsWith('#')) {
			addToast('No fragment found from Google OAuth URI', 'error')
			return
		}

		// Remove '#' for parsing.
		fragment = fragment.slice(1)

		const params = new Map<string, string>()
		fragment.split('&').forEach((param) => {
			const [key, value] = param.split('=', 2)
			if (key === undefined || value === undefined) return

			params.set(key, decodeURIComponent(value))
		})

		const state = params.get('state')
		const idToken = params.get('id_token')
		if (state === undefined || idToken === undefined) {
			addToast('Invalid fragment from Google OAuth URI', 'error')
			return
		}

		const requestBody = {
			state,
			idToken,
		} satisfies GoogleSignInRequest

		const response = await fetch('/api/auth/google', {
			method: 'POST',
			body: JSON.stringify(requestBody),
		})
		if (!response.ok) {
			addToast(`Failed to sign in with Google: ${response.statusText}`, 'error')
			return
		}

		const { referer }: GoogleSignInResponse = await response.json()
		await updateAuthStateFromCookie()
		goto(referer)
	})
</script>

<LoadingDots />
