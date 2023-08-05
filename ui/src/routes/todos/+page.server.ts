import { createTodo, deleteTodo, listTodos } from '$lib/server/bloatedApi/todo'
import { error } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

const userIdKey = 'userid'

export const load = (async ({ cookies }) => {
	let id = cookies.get(userIdKey)
	if (id === undefined) {
		id = crypto.randomUUID()
		cookies.set(userIdKey, id, { path: '/', secure: false, maxAge: 31_536_000 })
	}

	return {
		todos: (await listTodos(id)).todos,
	}
}) satisfies PageServerLoad

export const actions = {
	create: async ({ request, cookies }) => {
		const userId = cookies.get(userIdKey)
		if (!userId) {
			throw error(500, 'No userId found from cookie')
		}

		const data = await request.formData()
		const description = data.get('description') as string
		if (!description) throw error(400, 'Todo description should not be empty')

		await createTodo({ userId, task: description })
	},

	delete: async ({ request, cookies }) => {
		const userId = cookies.get(userIdKey)
		if (!userId) {
			throw error(500, 'No userId found from cookie')
		}

		const data = await request.formData()
		const todoIdString = data.get('todoId') as string
		const todoId = Number(todoIdString)
		if (!todoId) throw error(400, 'Invalid todo ID')

		deleteTodo(todoId)
	},
} satisfies Actions
