<script lang="ts">
	import Footer from '$components/Footer.svelte'
	import NavBar from '$components/NavBar.svelte'
	import ToastContainer from '$components/toast/ToastContainer.svelte'
	import { updateAuthStateFromCookie } from '$lib/browser/auth/action'
	import '@fontsource-variable/inter'
	import { onMount } from 'svelte'
	import '../app.css'

	onMount(() => {
		updateAuthStateFromCookie()
	})
</script>

<svelte:head>
	<title>Redshore</title>
</svelte:head>

<div class="flex min-h-screen flex-col">
	<NavBar>
		<section
			class="md:border-base-300 mx-auto mb-6 w-full max-w-3xl p-6 md:rounded-xl md:border md:shadow-md"
		>
			<slot />
		</section>
		<Footer />
		<ToastContainer />
	</NavBar>
</div>
