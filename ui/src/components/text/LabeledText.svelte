<script lang="ts">
	export { clazz as class }
	let clazz: string = ''

	export let label: string
	export let value: string
	export let isCode: boolean = false
</script>

<div class={clazz}>
	<span class="ml-2 text-xs font-light">{label}</span>
	<div class="border-base-200 border-t">
		<span class="ml-2" class:font-mono={isCode}>{value}</span>
	</div>
</div>
