import type { CreateTodoRequest, ListTodoReponse as ListTodosReponse, Todo } from '$lib/model/todo'
import { getBloatedApiBase } from '$lib/server/bloatedApi/base'
import { error } from '@sveltejs/kit'

export async function createTodo(
	request: CreateTodoRequest,
	customToken: string,
	options?: { fetch?: typeof fetch }
): Promise<Todo> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/todos`, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			Authorization: `Bearer ${customToken}`,
		},
		body: JSON.stringify(request),
	})
	if (!response.ok) {
		throw error(response.status, `Failed to create todo of user(${request.userId})`)
	}

	return response.json()
}

export async function listTodos(
	userId: string,
	customToken: string,
	options?: { fetch?: typeof fetch }
): Promise<ListTodosReponse> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/todos?userid=${userId}`, {
		headers: {
			Authorization: `Bearer ${customToken}`,
		},
	})
	if (!response.ok) {
		throw error(response.status, `Failed to get todos of user(${userId})`)
	}

	return response.json()
}

export async function deleteTodo(
	todoId: number,
	customToken: string,
	options?: { fetch?: typeof fetch }
): Promise<void> {
	const customFetch = options?.fetch ?? fetch

	const response = await customFetch(`${getBloatedApiBase()}/api/v1/todos/${todoId}`, {
		method: 'DELETE',
		headers: {
			Authorization: `Bearer ${customToken}`,
		},
	})
	if (!response.ok) {
		throw error(response.status, `Failed to delete todo(${todoId})`)
	}
}
