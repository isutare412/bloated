<script lang="ts">
	import { goto } from '$app/navigation'
	import { navigating } from '$app/stores'
	import LoadingDots from '$components/LoadingDots.svelte'
	import LabeledText from '$components/text/LabeledText.svelte'
	import PageTitle from '$components/text/PageTitle.svelte'
	import { pokemonVersionName } from '$lib/model/pokemon'
	import { addToast } from '$lib/toast/action'
	import type { PageData } from './$types'

	export let data: PageData
	$: ({ id, name, koreanName, flavorTexts, spriteUrl } = data)

	function navigatePrev() {
		goto(`/pokemons/${id - 1}`)
	}

	function navigateNext() {
		goto(`/pokemons/${id + 1}`)
	}

	function navigateTo(idStr: string) {
		const id = parseInt(idStr)
		if (isNaN(id)) {
			addToast(`pokemon ID ${idStr} is not valid`, 'error')
			return
		}

		goto(`/pokemons/${id}`)
	}
</script>

<div class="mb-4 flex items-center justify-between">
	<button
		on:click={navigatePrev}
		class="btn btn-link btn-sm text-base-content no-animation px-0 no-underline hover:no-underline"
		>〈<span class="underline">Prev</span></button
	>
	<input
		type="text"
		placeholder="ID"
		class="input input-bordered input-sm focus:border-primary w-full max-w-[4rem] text-center transition-colors focus:outline-none"
		value={id}
		on:keydown={(event) => {
			if (event.key !== 'Enter') return
			event.currentTarget.blur()
			navigateTo(event.currentTarget.value)
		}}
	/>
	<button
		on:click={navigateNext}
		class="btn btn-link btn-sm text-base-content no-animation px-0 no-underline hover:no-underline"
		><span class="underline">Next</span> 〉</button
	>
</div>
{#if $navigating !== null}
	<LoadingDots />
{:else}
	<PageTitle title={koreanName ?? name} />
	{#if spriteUrl}
		<img src={spriteUrl} alt="pokemon sprite" class="mx-auto w-1/2 max-w-xs" />
	{/if}
	<div class="mt-4 flex flex-col gap-y-2">
		<LabeledText label="ID" value={id.toString()} isCode />
		<LabeledText label="Name" value={name.toString()} />
	</div>
	{#if flavorTexts.length > 0}
		<table class="mt-4 table">
			<thead>
				<tr>
					<th class="text-base-content text-xs font-light">Version</th>
					<th class="text-base-content text-xs font-light">Flavor Text</th>
				</tr>
			</thead>
			<tbody>
				{#each flavorTexts as { version, text }}
					<tr>
						<td class="font-medium">{pokemonVersionName[version]}</td>
						<td>{text}</td>
					</tr>
				{/each}
			</tbody>
		</table>
	{/if}
{/if}
