import { isIPBanned } from '$lib/server/blockIp'
import { error } from '@sveltejs/kit'
import type { LayoutServerLoad } from './$types'

export const load = (async ({ getClientAddress }) => {
	const ip = getClientAddress()
	console.log(`IP: ${ip}`)

	if (isIPBanned(ip)) {
		console.log(`Block banned IP: ${ip}`)
		throw error(403, 'Banned IP')
	}
}) satisfies LayoutServerLoad
