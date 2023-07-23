<script lang="ts">
	import { delay, round } from 'lodash-es'
	import { onDestroy } from 'svelte'

	let roll: number
	let rollCount: { [key: number]: number } = {
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
	}
	let totalRoll = 0

	let delayTimer: number | undefined
	let keepRollTimer: ReturnType<typeof setInterval> | undefined

	async function rollDice() {
		const response = await fetch('/api/roll')
		roll = await response.json()
		rollCount[roll]++
		totalRoll++
	}

	function keepRollDice() {
		if (delayTimer || keepRollTimer) return

		delayTimer = delay(() => {
			keepRollTimer = setInterval(() => {
				rollDice()
			}, 10)
		}, 200)
	}

	function stopRollDice() {
		clearTimeout(delayTimer)
		clearInterval(keepRollTimer)
		delayTimer = undefined
		keepRollTimer = undefined
	}

	onDestroy(() => {
		clearTimeout(delayTimer)
		clearInterval(keepRollTimer)
	})
</script>

<button
	on:click={rollDice}
	on:mousedown={keepRollDice}
	on:mouseup={stopRollDice}
	class="btn btn-primary">ROLL</button
>
<div class="mt-2">
	{#if roll === undefined}
		<span>Waiting for a roll...</span>
	{:else}
		<span>You've got '<span class="text-primary">{roll}</span>'</span>
	{/if}
</div>
{#if totalRoll > 0}
	<table class="table table-fixed">
		<thead>
			<tr>
				<th>Face</th>
				<th>Count</th>
				<th>Ratio</th>
			</tr>
		</thead>
		<tbody class="font-mono">
			{#each Object.entries(rollCount) as [face, count] (face)}
				<tr>
					<td>{face}</td>
					<td>{count}</td>
					<td>{round(count / totalRoll, 4)}</td>
				</tr>
			{/each}
		</tbody>
	</table>
{/if}
