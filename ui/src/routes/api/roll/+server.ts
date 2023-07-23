import { json } from '@sveltejs/kit'
import { random } from 'lodash-es'
import type { RequestHandler } from './$types'

export const GET = (() => {
	const number = random(1, 6)
	return json(number)
}) satisfies RequestHandler
