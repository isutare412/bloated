import type { GetCustomClaimsResponse } from '$lib/model/auth'
import { verifyCustomToken } from '$lib/server/bloatedApi/token'
import { json, type RequestHandler } from '@sveltejs/kit'

export const GET = (async (event) => {
	const { cookies } = event
	const response: GetCustomClaimsResponse = {}

	const token = cookies.get('token')
	if (!token) return json(response)

	response.claims = await verifyCustomToken({ customToken: token })
	return json(response)
}) satisfies RequestHandler
