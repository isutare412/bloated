<script lang="ts">
	import Avatar from '$components/Avatar.svelte'
	import Cross from '$components/icons/Cross.svelte'
	import Hamburger from '$components/icons/Hamburger.svelte'
	import { fade, fly } from 'svelte/transition'

	const title = 'Redshore'

	const pages = [
		{ name: 'Home', href: '/' },
		{ name: 'Dev', href: '/dev' },
		{ name: 'Pokemon', href: '/pokemons' },
		{ name: 'Todo', href: '/todos' },
		{ name: 'Roll', href: '/roll' },
	]

	let isDrawerOpen = false

	function openDrawer() {
		isDrawerOpen = true
	}

	function closeDrawer() {
		isDrawerOpen = false
	}
</script>

<nav class="border-base-300 border-b shadow md:mb-4">
	<div class="mx-auto flex h-16 max-w-3xl items-stretch justify-between p-3">
		<div class="flex items-center gap-x-2">
			<button on:click={openDrawer} class="btn btn-square btn-ghost md:hidden">
				<div class="h-7 w-7">
					<Hamburger />
				</div>
			</button>
			<a class="text-2xl" href="/">{title}</a>
		</div>
		<div class="flex items-center gap-x-3">
			<ul class="hidden items-center gap-x-3 md:flex">
				{#each pages.slice(1) as page (page.name)}
					<li><a href={page.href} class="link link-hover text-lg font-light">{page.name}</a></li>
				{/each}
			</ul>
			<Avatar />
		</div>
	</div>
</nav>

<slot />

{#if isDrawerOpen}
	<section>
		<button
			on:click={closeDrawer}
			transition:fade={{ duration: 150 }}
			class="fixed left-0 top-0 h-screen w-screen bg-black/40"
		/>
		<div class="fixed left-0 top-0">
			<div
				transition:fly={{ x: '-24rem', opacity: 100, duration: 300 }}
				class="bg-base-100 flex min-h-screen w-80 max-w-[100vw] flex-col overflow-y-scroll p-2"
			>
				<div class="ml-auto">
					<button on:click={closeDrawer} class="btn btn-circle btn-ghost btn-sm p-1"
						><div class="h-6 w-6"><Cross /></div></button
					>
				</div>
				<div class="menu px-0">
					<ul>
						{#each pages as page (page.name)}
							<li>
								<a
									href={page.href}
									on:click={closeDrawer}
									class="link link-hover text-lg font-light">{page.name}</a
								>
							</li>
						{/each}
					</ul>
				</div>
			</div>
		</div>
	</section>
{/if}
