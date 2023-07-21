<script lang="ts">
	import LikeButton from '$components/LikeButton.svelte'
	import PageTitle from '$components/text/PageTitle.svelte'

	let likeCount = 0
</script>

<PageTitle title="Develop Page" />
<div>
	<p>Liked {likeCount} times!</p>
	<div class="mt-2">
		<LikeButton bind:likeCount />
	</div>
</div>
