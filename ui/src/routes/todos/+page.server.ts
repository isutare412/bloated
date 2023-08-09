import { createTodo, deleteTodo, listTodos } from '$lib/server/bloatedApi/todo'
import { verifyCustomToken } from '$lib/server/bloatedApi/token'
import { error, redirect } from '@sveltejs/kit'
import type { Actions, PageServerLoad } from './$types'

export const load = (async (event) => {
	const { cookies, url, fetch } = event

	const token = cookies.get('token')
	if (!token) throw redirect(302, `/sign-in?referer=${url.pathname}`)

	const claims = await verifyCustomToken(token)

	return {
		todos: (await listTodos(claims.userId, token, { fetch })).todos,
	}
}) satisfies PageServerLoad

export const actions = {
	create: async (event) => {
		const { request, cookies, url, fetch } = event

		const token = cookies.get('token')
		if (!token) throw redirect(302, `/sign-in?referer=${url.pathname}`)

		const claims = await verifyCustomToken(token)

		const data = await request.formData()
		const description = data.get('description') as string
		if (!description) throw error(400, 'Todo description should not be empty')

		await createTodo({ userId: claims.userId, task: description }, token, { fetch })
	},

	delete: async (event) => {
		const { request, cookies, url, fetch } = event

		const token = cookies.get('token')
		if (!token) throw redirect(302, `/sign-in?referer=${url.pathname}`)

		await verifyCustomToken(token)

		const data = await request.formData()
		const todoIdString = data.get('todoId') as string
		const todoId = Number(todoIdString)
		if (!todoId) throw error(400, 'Invalid todo ID')

		deleteTodo(todoId, token, { fetch })
	},
} satisfies Actions
