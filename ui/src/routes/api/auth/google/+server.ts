import type { GoogleSignInRequest, GoogleSignInResponse } from '$lib/model/auth'
import { getCsrfSate } from '$lib/server/auth/csrfState'
import { createCustomTokenFromGoogleIdToken } from '$lib/server/bloatedApi/token'
import { error, json, type RequestHandler } from '@sveltejs/kit'

export const POST = (async (event) => {
	const { request, cookies, fetch } = event

	const requestBody: GoogleSignInRequest = await request.json()

	const state = getCsrfSate(requestBody.state)
	if (!state) {
		console.log(`Unknown Google CSRF state: ${requestBody.state}`)
		throw error(404, 'OAuth state not found')
	}

	const { customToken } = await createCustomTokenFromGoogleIdToken(
		{ token: requestBody.idToken },
		{ fetch }
	)

	cookies.set('token', customToken, {
		path: '/',
		secure: false,
		maxAge: 2_505_600,
		httpOnly: true,
	})

	console.log(`Google sign-in success; Redirect to referer(${state.referer})`)

	const responseBody = {
		referer: state.referer,
	} satisfies GoogleSignInResponse

	return json(responseBody)
}) satisfies RequestHandler
