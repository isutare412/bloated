<script lang="ts">
	import { enhance } from '$app/forms'
	import TextInput from '$components/form/TextInput.svelte'
	import TrashBin from '$components/icons/TrashBin.svelte'
	import PageTitle from '$components/text/PageTitle.svelte'
	import type { PageData } from './$types'

	export let data: PageData
</script>

<PageTitle title="Todos" />
<form method="POST" action="?/create" use:enhance>
	<TextInput
		name="description"
		label="Todo"
		tooltip="Press enter to add the todo"
		placeholder="Task"
		required
	/>
</form>
<ul class="mt-4 space-y-2">
	{#each data.todos as todo (todo.id)}
		<li class="bg-base-200 flex w-full items-center gap-x-2 rounded p-3 px-4">
			<div class="flex-1">
				<span class="break-all">
					{todo.task}
				</span>
			</div>
			<div class="flex-none">
				<form method="POST" action="?/delete" use:enhance>
					<input type="hidden" name="todoId" value={todo.id} />
					<button class="btn btn-ghost btn-circle btn-sm">
						<div class="h-6 w-6 min-w-[1.5rem]">
							<TrashBin />
						</div>
					</button>
				</form>
			</div>
		</li>
	{/each}
</ul>
