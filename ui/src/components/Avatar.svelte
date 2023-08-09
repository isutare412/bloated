<script lang="ts">
	import { page } from '$app/stores'
	import userImage from '$assets/user.png'
	import { authState } from '$lib/browser/auth/store'

	$: claims = $authState.claims
	$: photoUrl = claims?.picture ?? userImage
	$: currentPath = $page.url.pathname
</script>

<div class="dropdown dropdown-end">
	<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
	<!-- svelte-ignore a11y-label-has-associated-control -->
	<label tabindex="0" class="btn btn-ghost btn-circle avatar">
		<div class="h-10 w-10 rounded-full">
			<img src={photoUrl} alt="profile" />
		</div>
	</label>
	<!-- svelte-ignore a11y-no-noninteractive-tabindex -->
	<ul
		tabindex="0"
		class="menu menu-sm dropdown-content bg-base-100 rounded-box z-[1] mt-3 w-52 p-2 shadow"
	>
		{#if claims}
			<li><a href={`/sign-out?referer=${currentPath}`}>Sign Out</a></li>
		{:else}
			<li><a href={`/sign-in?referer=${currentPath}`}>Sign In</a></li>
		{/if}
	</ul>
</div>
