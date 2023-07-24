<script lang="ts">
	import { navigating } from '$app/stores'
	import LoadingDots from '$components/LoadingDots.svelte'
	import PageTitle from '$components/text/PageTitle.svelte'
	import type { PageData } from './$types'

	export let data: PageData
</script>

{#if $navigating}
	<LoadingDots />
{:else}
<PageTitle title="Pokemons" />
<ul class="list-inside list-disc space-y-1">
	{#each data.defaultPokemons as pokemon}
		<li><a href={`/pokemons/${pokemon.id}`} class="link">{pokemon.name}</a></li>
	{/each}
</ul>
{/if}
