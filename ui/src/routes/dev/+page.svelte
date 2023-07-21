<script lang="ts">
	import LikeButton from '$components/LikeButton.svelte'
	import PageTitle from '$components/text/PageTitle.svelte'

	let likeCount = 0
</script>

<PageTitle title="Ongoing Development" />
<div>
	<p>Liked {likeCount} times!</p>
	<div class="mt-2">
		<LikeButton bind:likeCount />
	</div>
</div>

<div class="border-base-200 mt-4 flex flex-wrap justify-center gap-2 rounded border p-3">
	<button class="btn">Base</button>
	<button class="btn btn-primary">Primary</button>
	<button class="btn btn-secondary">Secondary</button>
	<button class="btn btn-accent">Accent</button>
	<button class="btn btn-neutral">Neutral</button>
	<button class="btn btn-info">Info</button>
	<button class="btn btn-warning">Warning</button>
	<button class="btn btn-error">Error</button>
</div>
