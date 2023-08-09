<script lang="ts">
	import { goto } from '$app/navigation'
	import { page } from '$app/stores'
	import LoadingDots from '$components/LoadingDots.svelte'
	import { clearAuthState } from '$lib/browser/auth/action'
	import { addToast } from '$lib/browser/toast/action'
	import { getErrorMessage } from '$lib/utils/errors'
	import { onMount } from 'svelte'

	onMount(async () => {
		try {
			await fetch('/api/auth/sign-out')
			clearAuthState()
		} catch (err) {
			addToast(getErrorMessage(err), 'error')
		}

		const referer = $page.url.searchParams.get('referer') ?? '/'
		goto(referer)
	})
</script>

<LoadingDots />
