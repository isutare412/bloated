<script lang="ts">
	import { debounce } from '$lib/utils/debounce'
	import { createEventDispatcher } from 'svelte'

	export let likeCount = 0

	const dispatch = createEventDispatcher<{ like: { count: number } }>()

	const onLike = debounce(() => {
		likeCount++
		dispatch('like', { count: likeCount })
	}, 200)
</script>

<button on:click={onLike} class="btn">Like ({likeCount})</button>
