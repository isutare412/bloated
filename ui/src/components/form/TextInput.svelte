<script lang="ts">
	import Info from '$components/icons/Info.svelte'

	export let name: string = ''
	export let label: string = ''
	export let tooltip: string = ''
	export let required: boolean = false
	export let placeholder = ''

	export let value: string = ''

	$: needValue = required && !value
</script>

<label class="flex flex-col">
	{#if label}
		<div class="ml-2 flex items-center">
			<span class="text-sm"
				>{label}
				{#if required}
					<span class="text-warning">*</span>
				{/if}
			</span>
			{#if tooltip}
				<div class="tooltip tooltip-bottom my-auto ml-1 h-4 w-4" data-tip={tooltip}><Info /></div>
			{/if}
		</div>
	{/if}
	<input
		bind:value
		type="text"
		{name}
		{placeholder}
		{required}
		autocomplete="off"
		class="input input-bordered focus:border-primary w-full rounded-md transition-colors focus:outline-none"
		class:border-warning={needValue}
	/>
</label>
