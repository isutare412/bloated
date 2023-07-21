<script lang="ts">
	export let src: string
	export let alt: string
	export let caption: string | undefined = undefined
</script>

<div class="not-prose mx-auto my-4">
	<img {alt} {src} class="mx-auto" />
	{#if caption}
		<p class="mt-4 text-center font-serif">{caption}</p>
	{/if}
</div>
