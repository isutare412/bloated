import type { Todo, UserId } from '$lib/model/todo'
import { error } from '@sveltejs/kit'

const db = new Map<UserId, Todo[]>()

export function createTodo(id: UserId, description: string) {
	const todos = getInitializedTodos(id)
	if (todos.length >= 100) {
		throw error(400, 'Cannot make more than 100 todos')
	}

	todos.push({
		id: crypto.randomUUID(),
		description,
		done: false,
	})
}

export function getTodos(id: UserId) {
	return getInitializedTodos(id)
}

export function deleteTodo(id: UserId, todoId: string): boolean {
	const todos = getInitializedTodos(id)
	const index = todos.findIndex((todo) => todo.id === todoId)
	if (index === -1) {
		return false
	}

	todos.splice(index, 1)
	return true
}

function getInitializedTodos(id: UserId): Todo[] {
	let todos = db.get(id)
	if (todos !== undefined) {
		return todos
	}

	todos = [
		{
			id: crypto.randomUUID(),
			description: 'Catch all pokemons',
			done: false,
		},
	]
	db.set(id, todos)
	return todos
}
