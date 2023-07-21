<script lang="ts">
	export let src: string
	export let alt: string
	export let caption: string | undefined = undefined
</script>

<figure>
	<img {alt} {src} class="mx-auto" />
	{#if caption}
		<figcaption class="text-center font-serif">{caption}</figcaption>
	{/if}
</figure>
