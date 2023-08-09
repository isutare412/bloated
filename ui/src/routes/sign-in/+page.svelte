<script lang="ts">
	import { goto } from '$app/navigation'
	import GoogleG from '$components/icons/GoogleG.svelte'
	import type { PageData } from './$types'

	export let data: PageData

	function signInGoogle() {
		goto(`/sign-in/google?referer=${data.redirectPath}`)
	}
</script>

<div class="mx-auto w-fit">
	<button on:click={signInGoogle} class="btn btn-outline flex w-60 gap-x-3">
		<div class="w-[24px]">
			<GoogleG />
		</div>
		<span>Sign in with Google</span>
	</button>
</div>
