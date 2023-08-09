import { createCsrfState } from '$lib/server/auth/csrfState'
import { json, type RequestHandler } from '@sveltejs/kit'

export const POST = ((event) => {
	const { url, getClientAddress } = event
	const ip = getClientAddress()

	const referer = url.searchParams.get('referer') ?? '/'
	const csrfState = createCsrfState(referer)

	console.log(`Created CSRF state: referer(${csrfState.referer}) ip(${ip})`)

	return json(csrfState)
}) satisfies RequestHandler
