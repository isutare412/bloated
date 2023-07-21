<script lang="ts">
	import NavBar from '$components/NavBar.svelte'
	import '@fontsource-variable/inter'
	import '../app.css'
</script>

<svelte:head>
	<title>Redshore</title>
</svelte:head>

<NavBar>
	<section class="mx-auto w-full max-w-3xl p-3">
		<slot />
	</section>
</NavBar>
