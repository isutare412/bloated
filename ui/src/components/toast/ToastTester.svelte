<script lang="ts">
	import TextInput from '$components/form/TextInput.svelte'
	import { addToast } from '$lib/browser/toast/action'
	import type { ToastLevel } from '$lib/model/toast'

	let message: string = 'Burn some toasts :)'

	function createToast(level: ToastLevel) {
		if (!message) return
		addToast(message, level)
	}
</script>

<div class="my-4">
	<TextInput bind:value={message} label="Message" placeholder="Toast message" />
	<div class="mt-2 flex flex-wrap gap-2">
		<button on:click={() => createToast('info')} class="btn btn-info">Info</button>
		<button on:click={() => createToast('success')} class="btn btn-success">Success</button>
		<button on:click={() => createToast('warning')} class="btn btn-warning">Warning</button>
		<button on:click={() => createToast('error')} class="btn btn-error">Error</button>
	</div>
</div>
