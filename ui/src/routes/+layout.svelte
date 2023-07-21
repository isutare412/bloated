<script lang="ts">
	import NavBar from '$components/NavBar.svelte'
	import '@fontsource/inter'
	import '../app.css'
</script>

<svelte:head>
	<title>Redshore</title>
</svelte:head>

<NavBar>
	<div class="mx-auto w-full max-w-3xl p-3">
		<slot />
	</div>
</NavBar>
