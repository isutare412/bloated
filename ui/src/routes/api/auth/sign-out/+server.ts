import type { RequestHandler } from '@sveltejs/kit'

export const GET = ((event) => {
	const { cookies } = event

	cookies.delete('token', { path: '/' })
	return new Response(null)
}) satisfies RequestHandler
