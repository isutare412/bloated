import { createTodo, deleteTodo, getTodos } from '$lib/server/database'
import { error } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

const userIdKey = 'userid'

export const load = (({ cookies }) => {
	let id = cookies.get(userIdKey)
	if (id === undefined) {
		id = crypto.randomUUID()
		cookies.set(userIdKey, id, { path: '/', secure: false })
	}

	return {
		todos: getTodos(id),
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

		createTodo(userId, description)
	},

	delete: async ({ request, cookies }) => {
		const userId = cookies.get(userIdKey)
		if (!userId) {
			throw error(500, 'No userId found from cookie')
		}

		const data = await request.formData()
		const todoId = data.get('todoId') as string
		if (!todoId) throw error(400, 'No todo ID')

		deleteTodo(userId, todoId)
	},
} satisfies Actions
